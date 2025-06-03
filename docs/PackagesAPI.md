# \PackagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidPackagesGet**](PackagesAPI.md#V3AppsGuidPackagesGet) | **Get** /v3/apps/{guid}/packages | List packages for an app
[**V3PackagesGet**](PackagesAPI.md#V3PackagesGet) | **Get** /v3/packages | List packages
[**V3PackagesGuidCopyPost**](PackagesAPI.md#V3PackagesGuidCopyPost) | **Post** /v3/packages/{guid}/copy | Copy a package
[**V3PackagesGuidDelete**](PackagesAPI.md#V3PackagesGuidDelete) | **Delete** /v3/packages/{guid} | Delete a package
[**V3PackagesGuidDownloadGet**](PackagesAPI.md#V3PackagesGuidDownloadGet) | **Get** /v3/packages/{guid}/download | Download package bits
[**V3PackagesGuidGet**](PackagesAPI.md#V3PackagesGuidGet) | **Get** /v3/packages/{guid} | Get a package
[**V3PackagesGuidPatch**](PackagesAPI.md#V3PackagesGuidPatch) | **Patch** /v3/packages/{guid} | Update a package
[**V3PackagesGuidUploadPost**](PackagesAPI.md#V3PackagesGuidUploadPost) | **Post** /v3/packages/{guid}/upload | Upload package bits
[**V3PackagesPost**](PackagesAPI.md#V3PackagesPost) | **Post** /v3/packages | Create a package



## V3AppsGuidPackagesGet

> V3AppsGuidPackagesGet200Response V3AppsGuidPackagesGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Types(types).Execute()

List packages for an app



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
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field to sort results by (optional)
	states := "states_example" // string | Filter by package states (comma-separated) (optional)
	types := "types_example" // string | Filter by package types (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3AppsGuidPackagesGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3AppsGuidPackagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPackagesGet`: V3AppsGuidPackagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3AppsGuidPackagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidPackagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **states** | **string** | Filter by package states (comma-separated) | 
 **types** | **string** | Filter by package types (comma-separated) | 

### Return type

[**V3AppsGuidPackagesGet200Response**](V3AppsGuidPackagesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGet

> V3AppsGuidPackagesGet200Response V3PackagesGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Guids(guids).States(states).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List packages



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
	orderBy := "orderBy_example" // string | Field to sort results by (optional)
	guids := "guids_example" // string | Filter by package GUIDs (comma-separated) (optional)
	states := "READY,PROCESSING_UPLOAD" // string | Filter by package states (comma-separated) (optional)
	types := "bits,docker,cnb" // string | Filter by package types (comma-separated) (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	spaceGuids := "spaceGuids_example" // string | Filter by space GUIDs (comma-separated) (optional)
	organizationGuids := "organizationGuids_example" // string | Filter by organization GUIDs (comma-separated) (optional)
	labelSelector := "environment=production,tier!=backend" // string | Filter by labels using label selector syntax (optional)
	createdAts := "2020-01-01T00:00:00Z,2020-12-31T23:59:59Z" // string | Filter by creation timestamp range (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamp range (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Guids(guids).States(states).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGet`: V3AppsGuidPackagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **guids** | **string** | Filter by package GUIDs (comma-separated) | 
 **states** | **string** | Filter by package states (comma-separated) | 
 **types** | **string** | Filter by package types (comma-separated) | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 
 **spaceGuids** | **string** | Filter by space GUIDs (comma-separated) | 
 **organizationGuids** | **string** | Filter by organization GUIDs (comma-separated) | 
 **labelSelector** | **string** | Filter by labels using label selector syntax | 
 **createdAts** | **string** | Filter by creation timestamp range | 
 **updatedAts** | **string** | Filter by update timestamp range | 

### Return type

[**V3AppsGuidPackagesGet200Response**](V3AppsGuidPackagesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidCopyPost

> Package V3PackagesGuidCopyPost(ctx, guid).V3PackagesGuidCopyPostRequest(v3PackagesGuidCopyPostRequest).Execute()

Copy a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The source package GUID
	v3PackagesGuidCopyPostRequest := *openapiclient.NewV3PackagesGuidCopyPostRequest(*openapiclient.NewV3PackagesGuidCopyPostRequestRelationships(*openapiclient.NewV3PackagesGuidCopyPostRequestRelationshipsApp(*openapiclient.NewV3PackagesGuidCopyPostRequestRelationshipsAppData("Guid_example")))) // V3PackagesGuidCopyPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGuidCopyPost(context.Background(), guid).V3PackagesGuidCopyPostRequest(v3PackagesGuidCopyPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidCopyPost`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGuidCopyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The source package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3PackagesGuidCopyPostRequest** | [**V3PackagesGuidCopyPostRequest**](V3PackagesGuidCopyPostRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidDelete

> V3PackagesGuidDelete(ctx, guid).Execute()

Delete a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The package GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackagesAPI.V3PackagesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidDeleteRequest struct via the builder pattern


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


## V3PackagesGuidDownloadGet

> *os.File V3PackagesGuidDownloadGet(ctx, guid).Execute()

Download package bits



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The package GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGuidDownloadGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGuidDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidGet

> Package V3PackagesGuidGet(ctx, guid).Execute()

Get a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The package GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidGet`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Package**](Package.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidPatch

> Package V3PackagesGuidPatch(ctx, guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()

Update a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The package GUID
	v3PackagesGuidPatchRequest := *openapiclient.NewV3PackagesGuidPatchRequest() // V3PackagesGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGuidPatch(context.Background(), guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidPatch`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3PackagesGuidPatchRequest** | [**V3PackagesGuidPatchRequest**](V3PackagesGuidPatchRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidUploadPost

> Package V3PackagesGuidUploadPost(ctx, guid).Bits(bits).Resources(resources).Execute()

Upload package bits



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The package GUID
	bits := os.NewFile(1234, "some_file") // *os.File | Zip file containing application source code
	resources := "resources_example" // string | Fingerprints of already-uploaded bits for resource matching (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesGuidUploadPost(context.Background(), guid).Bits(bits).Resources(resources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesGuidUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidUploadPost`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesGuidUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The package GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bits** | ***os.File** | Zip file containing application source code | 
 **resources** | **string** | Fingerprints of already-uploaded bits for resource matching | 

### Return type

[**Package**](Package.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesPost

> Package V3PackagesPost(ctx).V3PackagesPostRequest(v3PackagesPostRequest).Execute()

Create a package



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
	v3PackagesPostRequest := *openapiclient.NewV3PackagesPostRequest(*openapiclient.NewV3PackagesPostRequestRelationships(*openapiclient.NewV3AppsPostRequestRelationshipsSpace(*openapiclient.NewV3AppsPostRequestRelationshipsSpaceData("Guid_example"))), "Type_example") // V3PackagesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackagesAPI.V3PackagesPost(context.Background()).V3PackagesPostRequest(v3PackagesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackagesAPI.V3PackagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesPost`: Package
	fmt.Fprintf(os.Stdout, "Response from `PackagesAPI.V3PackagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3PackagesPostRequest** | [**V3PackagesPostRequest**](V3PackagesPostRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

