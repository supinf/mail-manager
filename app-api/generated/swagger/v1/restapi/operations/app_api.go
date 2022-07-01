// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/admin"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/basic"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/history"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/suppression"
)

// NewAppAPI creates a new App instance
func NewAppAPI(spec *loads.Document) *AppAPI {
	return &AppAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		SuppressionDeleteSuppressionsHandler: suppression.DeleteSuppressionsHandlerFunc(func(params suppression.DeleteSuppressionsParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation suppression.DeleteSuppressions has not yet been implemented")
		}),
		HistoryGetMailsHistoriesHandler: history.GetMailsHistoriesHandlerFunc(func(params history.GetMailsHistoriesParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation history.GetMailsHistories has not yet been implemented")
		}),
		SuppressionGetSuppressionsHandler: suppression.GetSuppressionsHandlerFunc(func(params suppression.GetSuppressionsParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation suppression.GetSuppressions has not yet been implemented")
		}),
		SuppressionGetSuppressionsMailHandler: suppression.GetSuppressionsMailHandlerFunc(func(params suppression.GetSuppressionsMailParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation suppression.GetSuppressionsMail has not yet been implemented")
		}),
		AdminPatchAdminUsersEnabledHandler: admin.PatchAdminUsersEnabledHandlerFunc(func(params admin.PatchAdminUsersEnabledParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation admin.PatchAdminUsersEnabled has not yet been implemented")
		}),
		AdminPostAdminUsagePlanHandler: admin.PostAdminUsagePlanHandlerFunc(func(params admin.PostAdminUsagePlanParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation admin.PostAdminUsagePlan has not yet been implemented")
		}),
		AdminPostAdminUsersHandler: admin.PostAdminUsersHandlerFunc(func(params admin.PostAdminUsersParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation admin.PostAdminUsers has not yet been implemented")
		}),
		BasicPostBulkMailsHandler: basic.PostBulkMailsHandlerFunc(func(params basic.PostBulkMailsParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation basic.PostBulkMails has not yet been implemented")
		}),
		BasicPostMailsHandler: basic.PostMailsHandlerFunc(func(params basic.PostMailsParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation basic.PostMails has not yet been implemented")
		}),
		SuppressionPostSuppressionsHandler: suppression.PostSuppressionsHandlerFunc(func(params suppression.PostSuppressionsParams, principal *auth.Session) middleware.Responder {
			return middleware.NotImplemented("operation suppression.PostSuppressions has not yet been implemented")
		}),

		// Applies when the "x-api-key" header is set
		APIKeyAuth: func(token string) (*auth.Session, error) {
			return nil, errors.NotImplemented("api key auth (api_key) x-api-key from header param [x-api-key] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*AppAPI SUPINF MAIL API 仕様
 */
type AppAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-api-key provided in the header
	APIKeyAuth func(string) (*auth.Session, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// SuppressionDeleteSuppressionsHandler sets the operation handler for the delete suppressions operation
	SuppressionDeleteSuppressionsHandler suppression.DeleteSuppressionsHandler
	// HistoryGetMailsHistoriesHandler sets the operation handler for the get mails histories operation
	HistoryGetMailsHistoriesHandler history.GetMailsHistoriesHandler
	// SuppressionGetSuppressionsHandler sets the operation handler for the get suppressions operation
	SuppressionGetSuppressionsHandler suppression.GetSuppressionsHandler
	// SuppressionGetSuppressionsMailHandler sets the operation handler for the get suppressions mail operation
	SuppressionGetSuppressionsMailHandler suppression.GetSuppressionsMailHandler
	// AdminPatchAdminUsersEnabledHandler sets the operation handler for the patch admin users enabled operation
	AdminPatchAdminUsersEnabledHandler admin.PatchAdminUsersEnabledHandler
	// AdminPostAdminUsagePlanHandler sets the operation handler for the post admin usage plan operation
	AdminPostAdminUsagePlanHandler admin.PostAdminUsagePlanHandler
	// AdminPostAdminUsersHandler sets the operation handler for the post admin users operation
	AdminPostAdminUsersHandler admin.PostAdminUsersHandler
	// BasicPostBulkMailsHandler sets the operation handler for the post bulk mails operation
	BasicPostBulkMailsHandler basic.PostBulkMailsHandler
	// BasicPostMailsHandler sets the operation handler for the post mails operation
	BasicPostMailsHandler basic.PostMailsHandler
	// SuppressionPostSuppressionsHandler sets the operation handler for the post suppressions operation
	SuppressionPostSuppressionsHandler suppression.PostSuppressionsHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *AppAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *AppAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *AppAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AppAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AppAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AppAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AppAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AppAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AppAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AppAPI
func (o *AppAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XAPIKeyAuth")
	}

	if o.SuppressionDeleteSuppressionsHandler == nil {
		unregistered = append(unregistered, "suppression.DeleteSuppressionsHandler")
	}
	if o.HistoryGetMailsHistoriesHandler == nil {
		unregistered = append(unregistered, "history.GetMailsHistoriesHandler")
	}
	if o.SuppressionGetSuppressionsHandler == nil {
		unregistered = append(unregistered, "suppression.GetSuppressionsHandler")
	}
	if o.SuppressionGetSuppressionsMailHandler == nil {
		unregistered = append(unregistered, "suppression.GetSuppressionsMailHandler")
	}
	if o.AdminPatchAdminUsersEnabledHandler == nil {
		unregistered = append(unregistered, "admin.PatchAdminUsersEnabledHandler")
	}
	if o.AdminPostAdminUsagePlanHandler == nil {
		unregistered = append(unregistered, "admin.PostAdminUsagePlanHandler")
	}
	if o.AdminPostAdminUsersHandler == nil {
		unregistered = append(unregistered, "admin.PostAdminUsersHandler")
	}
	if o.BasicPostBulkMailsHandler == nil {
		unregistered = append(unregistered, "basic.PostBulkMailsHandler")
	}
	if o.BasicPostMailsHandler == nil {
		unregistered = append(unregistered, "basic.PostMailsHandler")
	}
	if o.SuppressionPostSuppressionsHandler == nil {
		unregistered = append(unregistered, "suppression.PostSuppressionsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AppAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AppAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *AppAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *AppAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *AppAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AppAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the app API
func (o *AppAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AppAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/suppressions"] = suppression.NewDeleteSuppressions(o.context, o.SuppressionDeleteSuppressionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/mails/histories"] = history.NewGetMailsHistories(o.context, o.HistoryGetMailsHistoriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/suppressions"] = suppression.NewGetSuppressions(o.context, o.SuppressionGetSuppressionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/suppressions/{mail}"] = suppression.NewGetSuppressionsMail(o.context, o.SuppressionGetSuppressionsMailHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/admin/users/enabled"] = admin.NewPatchAdminUsersEnabled(o.context, o.AdminPatchAdminUsersEnabledHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/admin/usage_plan"] = admin.NewPostAdminUsagePlan(o.context, o.AdminPostAdminUsagePlanHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/admin/users"] = admin.NewPostAdminUsers(o.context, o.AdminPostAdminUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/bulk/mails"] = basic.NewPostBulkMails(o.context, o.BasicPostBulkMailsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mails"] = basic.NewPostMails(o.context, o.BasicPostMailsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/suppressions"] = suppression.NewPostSuppressions(o.context, o.SuppressionPostSuppressionsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AppAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AppAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AppAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AppAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *AppAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
