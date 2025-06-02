# \RouteSharingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3RoutesGuidRelationshipsSharedSpacesGet**](RouteSharingAPI.md#V3RoutesGuidRelationshipsSharedSpacesGet) | **Get** /v3/routes/{guid}/relationships/shared_spaces | List shared spaces for a route (experimental)
[**V3RoutesGuidRelationshipsSpacePatch**](RouteSharingAPI.md#V3RoutesGuidRelationshipsSpacePatch) | **Patch** /v3/routes/{guid}/relationships/space | Share a route to another space (experimental)



## V3RoutesGuidRelationshipsSharedSpacesGet

> V3RoutesGuidRelationshipsSharedSpacesGet200Response V3RoutesGuidRelationshipsSharedSpacesGet(ctx, guid).Execute()

List shared spaces for a route (experimental)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteSharingAPI.V3RoutesGuidRelationshipsSharedSpacesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteSharingAPI.V3RoutesGuidRelationshipsSharedSpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidRelationshipsSharedSpacesGet`: V3RoutesGuidRelationshipsSharedSpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RouteSharingAPI.V3RoutesGuidRelationshipsSharedSpacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidRelationshipsSharedSpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3RoutesGuidRelationshipsSharedSpacesGet200Response**](V3RoutesGuidRelationshipsSharedSpacesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidRelationshipsSpacePatch

> V3RoutesGuidRelationshipsSpacePatch200Response V3RoutesGuidRelationshipsSpacePatch(ctx, guid).V3RoutesGuidRelationshipsSpacePatchRequest(v3RoutesGuidRelationshipsSpacePatchRequest).Execute()

Share a route to another space (experimental)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidRelationshipsSpacePatchRequest := *openapiclient.NewV3RoutesGuidRelationshipsSpacePatchRequest([]openapiclient.V3RoutesGuidRelationshipsSpacePatchRequestDataInner{*openapiclient.NewV3RoutesGuidRelationshipsSpacePatchRequestDataInner("Guid_example")}) // V3RoutesGuidRelationshipsSpacePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RouteSharingAPI.V3RoutesGuidRelationshipsSpacePatch(context.Background(), guid).V3RoutesGuidRelationshipsSpacePatchRequest(v3RoutesGuidRelationshipsSpacePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RouteSharingAPI.V3RoutesGuidRelationshipsSpacePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidRelationshipsSpacePatch`: V3RoutesGuidRelationshipsSpacePatch200Response
	fmt.Fprintf(os.Stdout, "Response from `RouteSharingAPI.V3RoutesGuidRelationshipsSpacePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidRelationshipsSpacePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidRelationshipsSpacePatchRequest** | [**V3RoutesGuidRelationshipsSpacePatchRequest**](V3RoutesGuidRelationshipsSpacePatchRequest.md) |  | 

### Return type

[**V3RoutesGuidRelationshipsSpacePatch200Response**](V3RoutesGuidRelationshipsSpacePatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

