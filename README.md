### Generate certificates

First generate a self-signed rsa key and certificate that the server can use for TLS.

```sh
$ make keys KEY=/tmp/ca.key CERT=/tmp/ca.crt
```

### Create a https server & client application running in a kubernetes cluster

Create a TLS secret.

```sh
$ kubectl create secret tls https-test-certs --key /tmp/ca.key --cert /tmp/ca.crt
```

Deploy client & server.

```sh
$ kubectl apply -f deploy/deploy.yaml
```
