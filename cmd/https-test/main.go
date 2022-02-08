package main

import (
	"crypto/tls"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jessevdk/go-flags"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/loads"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"

	"wwwin-github.cisco.com/eti/swagger-authentication-test/client/client"
	clientop "wwwin-github.cisco.com/eti/swagger-authentication-test/client/client/operations"
	clientmo "wwwin-github.cisco.com/eti/swagger-authentication-test/client/models"
	"wwwin-github.cisco.com/eti/swagger-authentication-test/server/models"
	"wwwin-github.cisco.com/eti/swagger-authentication-test/server/restapi"
	serverop "wwwin-github.cisco.com/eti/swagger-authentication-test/server/restapi/operations"
)

func initLogs() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableTimestamp:       false,
		DisableSorting:         true,
		DisableLevelTruncation: true,
		QuoteEmptyFields:       true,
	})

	log.SetReportCaller(true)

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func runClient(*cli.Context) error {
	var err error

	httpClient := newClient()

	apiKeyAuth := httptransport.APIKeyAuth("X-API-Key", "header", viper.GetString("API_KEY"))

	for i := 0; i < viper.GetInt("CLIENT_CALLS_NUMBER"); i++ {
		_, err = httpClient.Operations.GetTest(clientop.NewGetTestParams(), apiKeyAuth)
		if err != nil {
			log.Errorf("GetTest failed: %v", err)
		} else {
			log.Info("GetTest succeeded")
		}

		_, err = httpClient.Operations.PostTest(clientop.NewPostTestParams().WithBody(&clientmo.Body{Body: "test"}), apiKeyAuth)
		if err != nil {
			log.Errorf("PostTest failed: %v", err)
		} else {
			log.Info("PostTest succeeded")
		}

		time.Sleep(viper.GetDuration("SLEEP_BETWEEN_CALLS"))
	}

	return nil
}

func newClient() *client.HTTPs {
	host := viper.GetString("HOST")

	var transport *httptransport.Runtime
	if viper.GetBool("DISABLE_TLS") {
		transport = httptransport.New(host, client.DefaultBasePath, []string{"http"})
	} else if viper.GetBool("INSECURE_SKIP_VERIFY") {
		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport = httptransport.NewWithClient(host, client.DefaultBasePath, []string{"https"},
			&http.Client{Transport: customTransport})
	} else {
		transport = httptransport.New(host, client.DefaultBasePath, []string{"https"})
	}

	return client.New(transport, strfmt.Default)
}

func runServer(*cli.Context) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := serverop.NewHTTPsAPI(swaggerSpec)
	api.APIKeyAuthAuth = func(token string) (interface{}, error) {
		log.Infof("APIKeyAuthAuth: Got token (%v)", token)
		if token != "allow" {
			return "forbidden", nil
		}
		return "ok", nil
	}

	api.GetTestHandler = serverop.GetTestHandlerFunc(func(params serverop.GetTestParams, principal interface{}) middleware.Responder {
		log.Infof("GetTestHandlerFunc: Got principal (%v)", principal)
		if principal != "ok" {
			return serverop.NewGetTestForbidden()
		}
		return serverop.NewGetTestOK().WithPayload(&models.OK{ID: 1})
	})
	api.PostTestHandler = serverop.PostTestHandlerFunc(func(params serverop.PostTestParams, principal interface{}) middleware.Responder {
		log.Infof("PostTestHandlerFunc: Got principal (%v)", principal)
		if principal != "ok" {
			return serverop.NewPostTestForbidden()
		}
		return serverop.NewPostTestCreated().WithPayload(&models.OK{ID: 1})
	})

	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	server.ConfigureAPI()
	server.EnabledListeners = []string{"http", "https"}
	// The port to listen on for insecure connections
	server.Port = 8080
	// The port to listen on for secure connections
	server.TLSPort = 8433
	//// The certificate authority file to be used with mutual tls auth
	//server.TLSCACertificate = flags.Filename(viper.GetString("TLS_CA_CERTIFICATE"))
	// The certificate to use for secure connections
	server.TLSCertificate = flags.Filename(viper.GetString("TLS_CERTIFICATE"))
	// The private key to use for secure connections
	server.TLSCertificateKey = flags.Filename(viper.GetString("TLS_PRIVATE_KEY"))

	if err := server.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("HOST", "server:8433")
	viper.SetDefault("DISABLE_TLS", false)
	viper.SetDefault("INSECURE_SKIP_VERIFY", true)
	viper.SetDefault("API_KEY", "api-key")
	viper.SetDefault("SLEEP_BETWEEN_CALLS", 30*time.Second)
	viper.SetDefault("CLIENT_CALLS_NUMBER", 100) // app will run ~30 min
	//viper.SetDefault("TLS_CA_CERTIFICATE", "/etc/certs/ca.crt")
	viper.SetDefault("TLS_CERTIFICATE", "/etc/certs/tls.crt")
	viper.SetDefault("TLS_PRIVATE_KEY", "/etc/certs/tls.key")

	initLogs()

	app := cli.NewApp()
	app.Name = "HTTPs Test"
	app.Version = "0.1.0"
	app.Usage = ""
	app.UsageText = ""
	runClientCommand := cli.Command{
		Name:   "run-client",
		Usage:  "Starts https client test",
		Action: runClient,
		Flags:  []cli.Flag{},
	}
	runClientCommand.UsageText = runClientCommand.Name

	runServerCommand := cli.Command{
		Name:   "run-server",
		Usage:  "Starts https server test",
		Action: runServer,
		Flags:  []cli.Flag{},
	}
	runServerCommand.UsageText = runServerCommand.Name

	app.Commands = []cli.Command{
		runClientCommand,
		runServerCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
