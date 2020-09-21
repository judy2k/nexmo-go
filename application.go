package vonage

import (
	"context"
	"encoding/json"
	"runtime"

	"github.com/antihax/optional"
	"github.com/vonage/vonage-go-sdk/application"
)

// ApplicationClient for working with the Application API
type ApplicationClient struct {
	Config    *application.Configuration
	apiKey    string
	apiSecret string
}

// NewApplicationClient Creates a new Application Client, supplying an Auth to work with
func NewApplicationClient(Auth Auth) *ApplicationClient {
	client := new(ApplicationClient)
	creds := Auth.GetCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]

	client.Config = application.NewConfiguration()
	client.Config.UserAgent = "vonage-go/0.15-dev Go/" + runtime.Version()
	return client
}

// ApplicationErrorResponse respresents error responses
type ApplicationErrorResponse struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

type GetApplicationsOpts struct {
	PageSize int32
	Page     int32
}

// List your Applications
func (client *ApplicationClient) GetApplications(opts GetApplicationsOpts) (application.ApplicationResponseCollection, ApplicationErrorResponse, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	AppOpts := application.ListApplicationOpts{}

	if opts.Page != 0 {
		AppOpts.Page = optional.NewInt32(opts.Page)
	}

	if opts.PageSize != 0 {
		AppOpts.PageSize = optional.NewInt32(opts.PageSize)
	}

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	result, _, err := applicationClient.DefaultApi.ListApplication(ctx, &AppOpts)
	if err != nil {
		e, ok := err.(application.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp ApplicationErrorResponse
			json.Unmarshal(data, &errResp)
			return application.ApplicationResponseCollection{}, errResp, err
		}

		// this catches other error types
		return result, ApplicationErrorResponse{}, err
	}
	return result, ApplicationErrorResponse{}, nil
}

// GetApplication returns one application, by app ID
func (client *ApplicationClient) GetApplication(app_id string) (application.ApplicationResponse, ApplicationErrorResponse, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	result, _, err := applicationClient.DefaultApi.GetApplication(ctx, app_id)
	if err != nil {
		e, ok := err.(application.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp ApplicationErrorResponse
			json.Unmarshal(data, &errResp)
			return application.ApplicationResponse{}, errResp, err
		}
		return result, ApplicationErrorResponse{}, err
	}

	return result, ApplicationErrorResponse{}, nil
}

type ApplicationUrl struct {
	Address    string `json:"address,omitempty"`
	HttpMethod string `json:"http_method,omitempty"`
}

type ApplicationMessagesWebhooks struct {
	InboundUrl ApplicationUrl `json:"inbound_url,omitempty"`
	StatusUrl  ApplicationUrl `json:"status_url,omitempty"`
}

type ApplicationMessages struct {
	Webhooks ApplicationMessagesWebhooks `json:"webhooks,omitempty"`
}

type ApplicationKeys struct {
	PublicKey string `json:"public_key,omitempty"`
}

type ApplicationCapabilities struct {
	//Voice    interface{}
	// Rtc      interface{}
	Messages ApplicationMessages `json:"messages,omitempty"`
	// Vbc      interface{}
}

type CreateApplicationOpts struct {
	Keys         ApplicationKeys         `json:"keys,omitempty"`
	Capabilities ApplicationCapabilities `json:"capabilities,omitempty"`
}

type CreateApplicationRequestOpts struct {
	Name         string                   `json:"name,omitempty"`
	Keys         *ApplicationKeys         `json:"keys,omitempty"`
	Capabilities *ApplicationCapabilities `json:"capabilities,omitempty"`
}

// CreateApplication creates a new application
func (client *ApplicationClient) CreateApplication(name string, opts CreateApplicationOpts) (application.ApplicationResponse, ApplicationErrorResponse, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	AppOpts := CreateApplicationRequestOpts{}
	AppOpts.Name = name
	AppOpts.Capabilities = &opts.Capabilities
	createOpts := application.CreateApplicationOpts{Opts: optional.NewInterface(AppOpts)}

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	result, _, err := applicationClient.DefaultApi.CreateApplication(ctx, &createOpts)
	if err != nil {
		e, ok := err.(application.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp ApplicationErrorResponse
			json.Unmarshal(data, &errResp)
			return application.ApplicationResponse{}, errResp, err
		}
		return result, ApplicationErrorResponse{}, err
	}

	return result, ApplicationErrorResponse{}, nil
}

// Delete application deletes an application
func (client *ApplicationClient) DeleteApplication(app_id string) (bool, ApplicationErrorResponse, error) {
	// create the client
	applicationClient := application.NewAPIClient(client.Config)

	ctx := context.WithValue(context.Background(), application.ContextBasicAuth, application.BasicAuth{
		UserName: client.apiKey,
		Password: client.apiSecret,
	})

	_, err := applicationClient.DefaultApi.DeleteApplication(ctx, app_id)
	if err != nil {
		e, ok := err.(application.GenericOpenAPIError)
		if ok {
			data := e.Body()

			var errResp ApplicationErrorResponse
			json.Unmarshal(data, &errResp)
			return false, errResp, err
		}
		return false, ApplicationErrorResponse{}, err
	}

	return true, ApplicationErrorResponse{}, nil
}
