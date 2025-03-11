package docs

import (
	_ "embed"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

const (
	OpenAPIDocsEndpoint        = "docs"
	OpenAPIRapiEndpoint        = "rapi"
	OpenAPISwaggerUIEndpoint   = "swaggerui"
	OpenAPIswaggerJSONEndpoint = "swagger.yaml"
	OpenAPIV2Prefix            = "/openapi/v2/"
)

//go:embed swagger.yaml
var specJSONBytes []byte

// ServeOpenAPISpec returns the OpenAPI specification file
func ServeOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(specJSONBytes)
}

// NewOpenAPIUIHandler returns a handler that serves OpenAPI specs via UI
func NewOpenAPIUIHandler() http.Handler {
	redocOpts := middleware.RedocOpts{
		BasePath: OpenAPIV2Prefix,
		SpecURL:  OpenAPIswaggerJSONEndpoint,
	}
	RapiDocOpts := middleware.RapiDocOpts{
		BasePath: OpenAPIV2Prefix,
		SpecURL:  OpenAPIswaggerJSONEndpoint,
		Path:     OpenAPIRapiEndpoint,
	}
	opts := middleware.SwaggerUIOpts{
		BasePath: OpenAPIV2Prefix,
		SpecURL:  OpenAPIswaggerJSONEndpoint,
		Path:     OpenAPISwaggerUIEndpoint,
	}

	var openAPISpecHandler http.Handler = http.HandlerFunc(ServeOpenAPISpec)

	openAPIUIHandler := middleware.Redoc(redocOpts, openAPISpecHandler)
	openAPIUIHandler = middleware.RapiDoc(RapiDocOpts, openAPIUIHandler)
	openAPIUIHandler = middleware.SwaggerUI(opts, openAPIUIHandler)

	return openAPIUIHandler
}
