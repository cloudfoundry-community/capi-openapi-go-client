# \BuildsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidBuildsGet**](BuildsAPI.md#V3AppsGuidBuildsGet) | **Get** /v3/apps/{guid}/builds | List builds for an app
[**V3BuildsGet**](BuildsAPI.md#V3BuildsGet) | **Get** /v3/builds | List builds
[**V3BuildsGuidGet**](BuildsAPI.md#V3BuildsGuidGet) | **Get** /v3/builds/{guid} | Get a build
[**V3BuildsGuidPatch**](BuildsAPI.md#V3BuildsGuidPatch) | **Patch** /v3/builds/{guid} | Update a build
[**V3BuildsPost**](BuildsAPI.md#V3BuildsPost) | **Post** /v3/builds | Create a build



## V3AppsGuidBuildsGet

> V3AppsGuidBuildsGet200Response V3AppsGuidBuildsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Execute()

List builds for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the app
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field by which to order results (optional)
	states := "states_example" // string | Filter by build states (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.V3AppsGuidBuildsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.V3AppsGuidBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidBuildsGet`: V3AppsGuidBuildsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.V3AppsGuidBuildsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field by which to order results | 
 **states** | **string** | Filter by build states (comma-separated) | 

### Return type

[**V3AppsGuidBuildsGet200Response**](V3AppsGuidBuildsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsGet

> V3AppsGuidBuildsGet200Response V3BuildsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).AppGuids(appGuids).PackageGuids(packageGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List builds



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
	states := "states_example" // string | Filter by build states (comma-separated) (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	packageGuids := "packageGuids_example" // string | Filter by package GUIDs (comma-separated) (optional)
	labelSelector := "labelSelector_example" // string | Filter by label selector (optional)
	createdAts := "createdAts_example" // string | Filter by creation timestamps (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.V3BuildsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).AppGuids(appGuids).PackageGuids(packageGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.V3BuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildsGet`: V3AppsGuidBuildsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.V3BuildsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field by which to order results | 
 **states** | **string** | Filter by build states (comma-separated) | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 
 **packageGuids** | **string** | Filter by package GUIDs (comma-separated) | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation timestamps | 
 **updatedAts** | **string** | Filter by update timestamps | 

### Return type

[**V3AppsGuidBuildsGet200Response**](V3AppsGuidBuildsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsGuidGet

> Build V3BuildsGuidGet(ctx, guid).Execute()

Get a build



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.V3BuildsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.V3BuildsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildsGuidGet`: Build
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.V3BuildsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the build | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Build**](Build.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsGuidPatch

> Build V3BuildsGuidPatch(ctx, guid).V3BuildsGuidPatchRequest(v3BuildsGuidPatchRequest).Execute()

Update a build



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the build
	v3BuildsGuidPatchRequest := *openapiclient.NewV3BuildsGuidPatchRequest() // V3BuildsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.V3BuildsGuidPatch(context.Background(), guid).V3BuildsGuidPatchRequest(v3BuildsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.V3BuildsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildsGuidPatch`: Build
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.V3BuildsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the build | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3BuildsGuidPatchRequest** | [**V3BuildsGuidPatchRequest**](V3BuildsGuidPatchRequest.md) |  | 

### Return type

[**Build**](Build.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsPost

> Build V3BuildsPost(ctx).V3BuildsPostRequest(v3BuildsPostRequest).Execute()

Create a build



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
	v3BuildsPostRequest := *openapiclient.NewV3BuildsPostRequest(*openapiclient.NewV3BuildsPostRequestPackage("Guid_example")) // V3BuildsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.V3BuildsPost(context.Background()).V3BuildsPostRequest(v3BuildsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.V3BuildsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3BuildsPost`: Build
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.V3BuildsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3BuildsPostRequest** | [**V3BuildsPostRequest**](V3BuildsPostRequest.md) |  | 

### Return type

[**Build**](Build.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

