# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppEnvironmentVariables**](ConfigurationAPI.md#GetAppEnvironmentVariables) | **Get** /v3/apps/{guid}/environment_variables | Get app environment variables
[**UpdateAppEnvironmentVariables**](ConfigurationAPI.md#UpdateAppEnvironmentVariables) | **Patch** /v3/apps/{guid}/environment_variables | Update app environment variables



## GetAppEnvironmentVariables

> GetAppEnvironmentVariables200Response GetAppEnvironmentVariables(ctx, guid).Execute()

Get app environment variables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetAppEnvironmentVariables(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetAppEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentVariables`: GetAppEnvironmentVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetAppEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppEnvironmentVariables200Response**](GetAppEnvironmentVariables200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppEnvironmentVariables

> UpdateAppEnvironmentVariables200Response UpdateAppEnvironmentVariables(ctx, guid).UpdateAppEnvironmentVariablesRequest(updateAppEnvironmentVariablesRequest).Execute()

Update app environment variables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	updateAppEnvironmentVariablesRequest := *openapiclient.NewUpdateAppEnvironmentVariablesRequest() // UpdateAppEnvironmentVariablesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.UpdateAppEnvironmentVariables(context.Background(), guid).UpdateAppEnvironmentVariablesRequest(updateAppEnvironmentVariablesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateAppEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironmentVariables`: UpdateAppEnvironmentVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.UpdateAppEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppEnvironmentVariablesRequest** | [**UpdateAppEnvironmentVariablesRequest**](UpdateAppEnvironmentVariablesRequest.md) |  | 

### Return type

[**UpdateAppEnvironmentVariables200Response**](UpdateAppEnvironmentVariables200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

