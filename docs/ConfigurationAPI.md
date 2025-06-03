# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidEnvironmentVariablesGet**](ConfigurationAPI.md#V3AppsGuidEnvironmentVariablesGet) | **Get** /v3/apps/{guid}/environment_variables | Get app environment variables
[**V3AppsGuidEnvironmentVariablesPatch**](ConfigurationAPI.md#V3AppsGuidEnvironmentVariablesPatch) | **Patch** /v3/apps/{guid}/environment_variables | Update app environment variables



## V3AppsGuidEnvironmentVariablesGet

> V3AppsGuidEnvironmentVariablesGet200Response V3AppsGuidEnvironmentVariablesGet(ctx, guid).Execute()

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
	resp, r, err := apiClient.ConfigurationAPI.V3AppsGuidEnvironmentVariablesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.V3AppsGuidEnvironmentVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesGet`: V3AppsGuidEnvironmentVariablesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.V3AppsGuidEnvironmentVariablesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvironmentVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidEnvironmentVariablesGet200Response**](V3AppsGuidEnvironmentVariablesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidEnvironmentVariablesPatch

> V3AppsGuidEnvironmentVariablesPatch200Response V3AppsGuidEnvironmentVariablesPatch(ctx, guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()

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
	v3AppsGuidEnvironmentVariablesPatchRequest := *openapiclient.NewV3AppsGuidEnvironmentVariablesPatchRequest() // V3AppsGuidEnvironmentVariablesPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.V3AppsGuidEnvironmentVariablesPatch(context.Background(), guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.V3AppsGuidEnvironmentVariablesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesPatch`: V3AppsGuidEnvironmentVariablesPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.V3AppsGuidEnvironmentVariablesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvironmentVariablesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidEnvironmentVariablesPatchRequest** | [**V3AppsGuidEnvironmentVariablesPatchRequest**](V3AppsGuidEnvironmentVariablesPatchRequest.md) |  | 

### Return type

[**V3AppsGuidEnvironmentVariablesPatch200Response**](V3AppsGuidEnvironmentVariablesPatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

