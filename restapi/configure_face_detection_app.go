// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/regeda/faced/handlers"
	"github.com/regeda/faced/internal/pigo"
	"github.com/regeda/faced/restapi/operations"
)

//go:generate swagger generate server --target ../../faced --name FaceDetectionApp --spec ../swagger.yml --principal interface{}

var pigoConfig = struct {
	CascadeFile string `long:"pigo-cascade-file" description:"Cascade binary file"`
	PuplocFile  string `long:"pigo-puploc-file" description:"Pupil localization cascade file"`
	FlplocDir   string `long:"pigo-flploc-dir" description:"The facial landmark points base directory"`
}{}

var httpClientConfig = struct {
	Timeout time.Duration `long:"http-client-timeout" description:"HTTP client timeout" default:"5s"`
}{}

func configureFlags(api *operations.FaceDetectionAppAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Pigo Options",
			Options:          &pigoConfig,
		},
		{
			ShortDescription: "HTTP Client Options",
			Options:          &httpClientConfig,
		}}
}

func configureAPI(api *operations.FaceDetectionAppAPI) http.Handler {
	detector, err := pigo.New(pigoConfig.CascadeFile, pigoConfig.PuplocFile, pigoConfig.FlplocDir)
	if err != nil {
		log.Fatal(err)
	}

	api.ServeError = errors.ServeError

	api.Logger = log.Printf

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetFacesHandler = handlers.NewFaces(detector,
		handlers.FacesWithHTTPClient(&http.Client{
			Timeout: httpClientConfig.Timeout,
		}))

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
