# \RouteMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3RoutesGuidDestinationsDestinationGuidDelete**](RouteMappingsAPI.md#V3RoutesGuidDestinationsDestinationGuidDelete) | **Delete** /v3/routes/{guid}/destinations/{destination_guid} | Remove destination from a route
[**V3RoutesGuidDestinationsGet**](RouteMappingsAPI.md#V3RoutesGuidDestinationsGet) | **Get** /v3/routes/{guid}/destinations | List destinations for a route
[**V3RoutesGuidDestinationsPatch**](RouteMappingsAPI.md#V3RoutesGuidDestinationsPatch) | **Patch** /v3/routes/{guid}/destinations | Update destinations for a route
[**V3RoutesGuidDestinationsPost**](RouteMappingsAPI.md#V3RoutesGuidDestinationsPost) | **Post** /v3/routes/{guid}/destinations | Add destinations to a route



## V3RoutesGuidDestinationsDestinationGuidDelete

> V3RoutesGuidDestinationsDestinationGuidDelete(ctx, guid, destinationGuid).Execute()

Remove destination from a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	destinationGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The destination GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RouteMappingsAPI.V3RoutesGuidDestinationsDestinationGuidDelete(context.Background(), guid, destinationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteMappingsAPI.V3RoutesGuidDestinationsDestinationGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 
**destinationGuid** | **string** | The destination GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsDestinationGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDestinationsGet

> V3RoutesGuidDestinationsGet200Response V3RoutesGuidDestinationsGet(ctx, guid).Page(page).PerPage(perPage).AppGuids(appGuids).Execute()

List destinations for a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteMappingsAPI.V3RoutesGuidDestinationsGet(context.Background(), guid).Page(page).PerPage(perPage).AppGuids(appGuids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteMappingsAPI.V3RoutesGuidDestinationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsGet`: V3RoutesGuidDestinationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RouteMappingsAPI.V3RoutesGuidDestinationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 

### Return type

[**V3RoutesGuidDestinationsGet200Response**](V3RoutesGuidDestinationsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDestinationsPatch

> V3RoutesGuidDestinationsPost200Response V3RoutesGuidDestinationsPatch(ctx, guid).V3RoutesGuidDestinationsPatchRequest(v3RoutesGuidDestinationsPatchRequest).Execute()

Update destinations for a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidDestinationsPatchRequest := *openapiclient.NewV3RoutesGuidDestinationsPatchRequest([]openapiclient.V3RoutesGuidDestinationsPatchRequestDestinationsInner{*openapiclient.NewV3RoutesGuidDestinationsPatchRequestDestinationsInner(*openapiclient.NewV3RoutesGuidDestinationsPatchRequestDestinationsInnerApp("Guid_example"))}) // V3RoutesGuidDestinationsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteMappingsAPI.V3RoutesGuidDestinationsPatch(context.Background(), guid).V3RoutesGuidDestinationsPatchRequest(v3RoutesGuidDestinationsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteMappingsAPI.V3RoutesGuidDestinationsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsPatch`: V3RoutesGuidDestinationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RouteMappingsAPI.V3RoutesGuidDestinationsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidDestinationsPatchRequest** | [**V3RoutesGuidDestinationsPatchRequest**](V3RoutesGuidDestinationsPatchRequest.md) |  | 

### Return type

[**V3RoutesGuidDestinationsPost200Response**](V3RoutesGuidDestinationsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDestinationsPost

> V3RoutesGuidDestinationsPost200Response V3RoutesGuidDestinationsPost(ctx, guid).V3RoutesGuidDestinationsPostRequest(v3RoutesGuidDestinationsPostRequest).Execute()

Add destinations to a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidDestinationsPostRequest := *openapiclient.NewV3RoutesGuidDestinationsPostRequest([]openapiclient.V3RoutesGuidDestinationsPostRequestDestinationsInner{*openapiclient.NewV3RoutesGuidDestinationsPostRequestDestinationsInner(*openapiclient.NewV3RoutesGuidDestinationsPostRequestDestinationsInnerApp("Guid_example"))}) // V3RoutesGuidDestinationsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteMappingsAPI.V3RoutesGuidDestinationsPost(context.Background(), guid).V3RoutesGuidDestinationsPostRequest(v3RoutesGuidDestinationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteMappingsAPI.V3RoutesGuidDestinationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsPost`: V3RoutesGuidDestinationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RouteMappingsAPI.V3RoutesGuidDestinationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidDestinationsPostRequest** | [**V3RoutesGuidDestinationsPostRequest**](V3RoutesGuidDestinationsPostRequest.md) |  | 

### Return type

[**V3RoutesGuidDestinationsPost200Response**](V3RoutesGuidDestinationsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

