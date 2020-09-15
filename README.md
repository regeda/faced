# Face Detection Web API

Face Detection Web API is written on Go. The detection approach is powered by [github.com/esimov/pigo](https://github.com/esimov/pigo) library. The library gives good results out of the box and it's free of charge :)

### Installation
The API spec is implemented in [go-swagger 2.0](https://goswagger.io/).

Install go-swagger with brew (visit https://goswagger.io/install.html for more options):
```sh
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```
Get the app:
```sh
$ go get -u github.com/regeda/faced
```

### Development

*API Spec*

Run `go generate` to get "go" code if `swagger.yml` changed.

Run the web server:
```sh
$ ./run.sh
```
or
```sh
$ go run cmd/face-detection-app-server/main.go \
  --port=5555 \
  --pigo-cascade-file=internal/pigo/cascade/facefinder \
  --pigo-puploc-file=internal/pigo/cascade/puploc \
  --pigo-flploc-dir=internal/pigo/cascade/lps
```
> Run `go run cmd/face-detection-app-server/main.go --help` for more options

Then navigate to http://127.0.0.1:5555/docs in your preferred browser.

*Tests*

Put your commands in `generate.go` file if you require mocks for tests. Then run `go generate`.

### Deployment

Face Detection Web API is a cloud-enabled web service. No 3d-party APIs required. You can pack the application into the Docker image and spin up containers in the Kubernetes cluster.

### Optimizations

To reduce CPU utilization of the server, the web service can be run behind a caching reverse proxy (aka Squid, Varnish, Nginx). Because of the HTTP handler for faces detection is implemented as HTTP GET method and the detection algorithm is based on a static data model.
