# \BuildpacksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3BuildpacksGet**](BuildpacksAPI.md#V3BuildpacksGet) | **Get** /v3/buildpacks | List buildpacks
[**V3BuildpacksGuidDelete**](BuildpacksAPI.md#V3BuildpacksGuidDelete) | **Delete** /v3/buildpacks/{guid} | Delete a buildpack
[**V3BuildpacksGuidGet**](BuildpacksAPI.md#V3BuildpacksGuidGet) | **Get** /v3/buildpacks/{guid} | Get a buildpack
[**V3BuildpacksGuidPatch**](BuildpacksAPI.md#V3BuildpacksGuidPatch) | **Patch** /v3/buildpacks/{guid} | Update a buildpack
[**V3BuildpacksGuidUploadPost**](BuildpacksAPI.md#V3BuildpacksGuidUploadPost) | **Post** /v3/buildpacks/{guid}/upload | Upload buildpack bits
[**V3BuildpacksPost**](BuildpacksAPI.md#V3BuildpacksPost) | **Post** /v3/buildpacks | Create a buildpack



## V3BuildpacksGet

> V3BuildpacksGet200Response V3BuildpacksGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Stacks(stacks).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List buildpacks



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
	names := "names_example" // string | Filter by buildpack names (comma-separated) (optional)
	stacks := "stacks_example" // string | Filter by stack names (comma-separated) (optional)
	labelSelector := "labelSelector_example" // string | Filter by label selector (optional)
	createdAts := "createdAts_example" // string | Filter by creation timestamps (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildpacksAPI.V3BuildpacksGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Stacks(stacks).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildpacksGet`: V3BuildpacksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BuildpacksAPI.V3BuildpacksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field by which to order results | 
 **names** | **string** | Filter by buildpack names (comma-separated) | 
 **stacks** | **string** | Filter by stack names (comma-separated) | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation timestamps | 
 **updatedAts** | **string** | Filter by update timestamps | 

### Return type

[**V3BuildpacksGet200Response**](V3BuildpacksGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGuidDelete

> V3BuildpacksGuidDelete(ctx, guid).Execute()

Delete a buildpack



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the buildpack

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildpacksAPI.V3BuildpacksGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidDeleteRequest struct via the builder pattern


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


## V3BuildpacksGuidGet

> Buildpack V3BuildpacksGuidGet(ctx, guid).Execute()

Get a buildpack



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the buildpack

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildpacksAPI.V3BuildpacksGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildpacksGuidGet`: Buildpack
	fmt.Fprintf(os.Stdout, "Response from `BuildpacksAPI.V3BuildpacksGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Buildpack**](Buildpack.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGuidPatch

> Buildpack V3BuildpacksGuidPatch(ctx, guid).V3BuildpacksGuidPatchRequest(v3BuildpacksGuidPatchRequest).Execute()

Update a buildpack



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the buildpack
	v3BuildpacksGuidPatchRequest := *openapiclient.NewV3BuildpacksGuidPatchRequest() // V3BuildpacksGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildpacksAPI.V3BuildpacksGuidPatch(context.Background(), guid).V3BuildpacksGuidPatchRequest(v3BuildpacksGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildpacksGuidPatch`: Buildpack
	fmt.Fprintf(os.Stdout, "Response from `BuildpacksAPI.V3BuildpacksGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3BuildpacksGuidPatchRequest** | [**V3BuildpacksGuidPatchRequest**](V3BuildpacksGuidPatchRequest.md) |  | 

### Return type

[**Buildpack**](Buildpack.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGuidUploadPost

> Job V3BuildpacksGuidUploadPost(ctx, guid).V3BuildpacksGuidUploadPostRequest(v3BuildpacksGuidUploadPostRequest).Execute()

Upload buildpack bits



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the buildpack
	v3BuildpacksGuidUploadPostRequest := *openapiclient.NewV3BuildpacksGuidUploadPostRequest() // V3BuildpacksGuidUploadPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildpacksAPI.V3BuildpacksGuidUploadPost(context.Background(), guid).V3BuildpacksGuidUploadPostRequest(v3BuildpacksGuidUploadPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksGuidUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildpacksGuidUploadPost`: Job
	fmt.Fprintf(os.Stdout, "Response from `BuildpacksAPI.V3BuildpacksGuidUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3BuildpacksGuidUploadPostRequest** | [**V3BuildpacksGuidUploadPostRequest**](V3BuildpacksGuidUploadPostRequest.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksPost

> Buildpack V3BuildpacksPost(ctx).V3BuildpacksPostRequest(v3BuildpacksPostRequest).Execute()

Create a buildpack



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
	v3BuildpacksPostRequest := *openapiclient.NewV3BuildpacksPostRequest("nodejs_buildpack") // V3BuildpacksPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildpacksAPI.V3BuildpacksPost(context.Background()).V3BuildpacksPostRequest(v3BuildpacksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildpacksAPI.V3BuildpacksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildpacksPost`: Buildpack
	fmt.Fprintf(os.Stdout, "Response from `BuildpacksAPI.V3BuildpacksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3BuildpacksPostRequest** | [**V3BuildpacksPostRequest**](V3BuildpacksPostRequest.md) |  | 

### Return type

[**Buildpack**](Buildpack.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

