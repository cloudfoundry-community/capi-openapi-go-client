# \DeploymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3DeploymentsGet**](DeploymentsAPI.md#V3DeploymentsGet) | **Get** /v3/deployments | List deployments
[**V3DeploymentsGuidActionsCancelPost**](DeploymentsAPI.md#V3DeploymentsGuidActionsCancelPost) | **Post** /v3/deployments/{guid}/actions/cancel | Cancel a deployment
[**V3DeploymentsGuidActionsContinuePost**](DeploymentsAPI.md#V3DeploymentsGuidActionsContinuePost) | **Post** /v3/deployments/{guid}/actions/continue | Continue a deployment
[**V3DeploymentsGuidGet**](DeploymentsAPI.md#V3DeploymentsGuidGet) | **Get** /v3/deployments/{guid} | Get a deployment
[**V3DeploymentsGuidPatch**](DeploymentsAPI.md#V3DeploymentsGuidPatch) | **Patch** /v3/deployments/{guid} | Update a deployment
[**V3DeploymentsPost**](DeploymentsAPI.md#V3DeploymentsPost) | **Post** /v3/deployments | Create a deployment



## V3DeploymentsGet

> V3DeploymentsGet200Response V3DeploymentsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).AppGuids(appGuids).StatusReasons(statusReasons).StatusValues(statusValues).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List deployments



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
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field by which to order results (optional)
	states := "states_example" // string | Filter by deployment states (comma-separated) (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	statusReasons := "statusReasons_example" // string | Filter by status reasons (comma-separated) (optional)
	statusValues := "statusValues_example" // string | Filter by status values (comma-separated) (optional)
	labelSelector := "labelSelector_example" // string | Filter by label selector (optional)
	createdAts := "createdAts_example" // string | Filter by creation timestamps (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).AppGuids(appGuids).StatusReasons(statusReasons).StatusValues(statusValues).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsGet`: V3DeploymentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field by which to order results | 
 **states** | **string** | Filter by deployment states (comma-separated) | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 
 **statusReasons** | **string** | Filter by status reasons (comma-separated) | 
 **statusValues** | **string** | Filter by status values (comma-separated) | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation timestamps | 
 **updatedAts** | **string** | Filter by update timestamps | 

### Return type

[**V3DeploymentsGet200Response**](V3DeploymentsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidActionsCancelPost

> Deployment V3DeploymentsGuidActionsCancelPost(ctx, guid).Execute()

Cancel a deployment



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsGuidActionsCancelPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsGuidActionsCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsGuidActionsCancelPost`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsGuidActionsCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidActionsCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidActionsContinuePost

> Deployment V3DeploymentsGuidActionsContinuePost(ctx, guid).Execute()

Continue a deployment



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsGuidActionsContinuePost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsGuidActionsContinuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsGuidActionsContinuePost`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsGuidActionsContinuePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidActionsContinuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidGet

> Deployment V3DeploymentsGuidGet(ctx, guid).Execute()

Get a deployment



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsGuidGet`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deployment**](Deployment.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidPatch

> Deployment V3DeploymentsGuidPatch(ctx, guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()

Update a deployment



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the deployment
	v3DeploymentsGuidPatchRequest := *openapiclient.NewV3DeploymentsGuidPatchRequest() // V3DeploymentsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsGuidPatch(context.Background(), guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsGuidPatch`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DeploymentsGuidPatchRequest** | [**V3DeploymentsGuidPatchRequest**](V3DeploymentsGuidPatchRequest.md) |  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsPost

> Deployment V3DeploymentsPost(ctx).V3DeploymentsPostRequest(v3DeploymentsPostRequest).Execute()

Create a deployment



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
	v3DeploymentsPostRequest := *openapiclient.NewV3DeploymentsPostRequest(*openapiclient.NewV3DeploymentsPostRequestRelationships(*openapiclient.NewV3DeploymentsPostRequestRelationshipsApp(*openapiclient.NewV3DeploymentsPostRequestRelationshipsAppData("Guid_example")))) // V3DeploymentsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.V3DeploymentsPost(context.Background()).V3DeploymentsPostRequest(v3DeploymentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.V3DeploymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DeploymentsPost`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.V3DeploymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3DeploymentsPostRequest** | [**V3DeploymentsPostRequest**](V3DeploymentsPostRequest.md) |  | 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

