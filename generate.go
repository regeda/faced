package faced

//go:generate swagger generate server -f swagger.yml

//go:generate mockgen -destination handlers/mock_http_test.go -package handlers_test net/http RoundTripper
//go:generate mockgen -destination handlers/mock_detector_test.go -package handlers_test github.com/regeda/faced/handlers Detector
//go:generate mockgen -destination handlers/mock_readcloser_test.go -package handlers_test io ReadCloser
