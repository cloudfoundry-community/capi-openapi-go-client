# \SidecarsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidSidecarsGet**](SidecarsAPI.md#V3AppsGuidSidecarsGet) | **Get** /v3/apps/{guid}/sidecars | List sidecars for an app
[**V3AppsGuidSidecarsPost**](SidecarsAPI.md#V3AppsGuidSidecarsPost) | **Post** /v3/apps/{guid}/sidecars | Create a sidecar for an app
[**V3ProcessesGuidSidecarsGet**](SidecarsAPI.md#V3ProcessesGuidSidecarsGet) | **Get** /v3/processes/{guid}/sidecars | List sidecars for a process
[**V3SidecarsGuidDelete**](SidecarsAPI.md#V3SidecarsGuidDelete) | **Delete** /v3/sidecars/{guid} | Delete a sidecar
[**V3SidecarsGuidGet**](SidecarsAPI.md#V3SidecarsGuidGet) | **Get** /v3/sidecars/{guid} | Get a sidecar
[**V3SidecarsGuidPatch**](SidecarsAPI.md#V3SidecarsGuidPatch) | **Patch** /v3/sidecars/{guid} | Update a sidecar
[**V3SidecarsGuidProcessesGet**](SidecarsAPI.md#V3SidecarsGuidProcessesGet) | **Get** /v3/sidecars/{guid}/processes | List processes for a sidecar



## V3AppsGuidSidecarsGet

> V3AppsGuidSidecarsGet200Response V3AppsGuidSidecarsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List sidecars for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Sort order for results (optional) (default to "created_at")
	names := "names_example" // string | Filter by sidecar names (comma-separated) (optional)
	labelSelector := "labelSelector_example" // string | Filter by label selector (optional)
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3AppsGuidSidecarsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3AppsGuidSidecarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSidecarsGet`: V3AppsGuidSidecarsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3AppsGuidSidecarsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSidecarsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Sort order for results | [default to &quot;created_at&quot;]
 **names** | **string** | Filter by sidecar names (comma-separated) | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 

### Return type

[**V3AppsGuidSidecarsGet200Response**](V3AppsGuidSidecarsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidSidecarsPost

> Sidecar V3AppsGuidSidecarsPost(ctx, guid).SidecarCreate(sidecarCreate).Execute()

Create a sidecar for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	sidecarCreate := *openapiclient.NewSidecarCreate("bundle exec rackup config.ru -p $PORT", "auth-proxy", []string{"web"}) // SidecarCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3AppsGuidSidecarsPost(context.Background(), guid).SidecarCreate(sidecarCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3AppsGuidSidecarsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSidecarsPost`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3AppsGuidSidecarsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSidecarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sidecarCreate** | [**SidecarCreate**](SidecarCreate.md) |  | 

### Return type

[**Sidecar**](Sidecar.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidSidecarsGet

> V3AppsGuidSidecarsGet200Response V3ProcessesGuidSidecarsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List sidecars for a process



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	page := int32(56) // int32 | Page number for pagination (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Sort order for results (optional) (default to "created_at")
	names := "names_example" // string | Filter by sidecar names (comma-separated) (optional)
	labelSelector := "labelSelector_example" // string | Filter by label selector (optional)
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3ProcessesGuidSidecarsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3ProcessesGuidSidecarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidSidecarsGet`: V3AppsGuidSidecarsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3ProcessesGuidSidecarsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidSidecarsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Sort order for results | [default to &quot;created_at&quot;]
 **names** | **string** | Filter by sidecar names (comma-separated) | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 

### Return type

[**V3AppsGuidSidecarsGet200Response**](V3AppsGuidSidecarsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidDelete

> V3SidecarsGuidDelete(ctx, guid).Execute()

Delete a sidecar



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sidecar GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SidecarsAPI.V3SidecarsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3SidecarsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The sidecar GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidDeleteRequest struct via the builder pattern


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


## V3SidecarsGuidGet

> Sidecar V3SidecarsGuidGet(ctx, guid).Execute()

Get a sidecar



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sidecar GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3SidecarsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3SidecarsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidGet`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3SidecarsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The sidecar GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sidecar**](Sidecar.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidPatch

> Sidecar V3SidecarsGuidPatch(ctx, guid).SidecarUpdate(sidecarUpdate).Execute()

Update a sidecar



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sidecar GUID
	sidecarUpdate := *openapiclient.NewSidecarUpdate() // SidecarUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3SidecarsGuidPatch(context.Background(), guid).SidecarUpdate(sidecarUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3SidecarsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidPatch`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3SidecarsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The sidecar GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sidecarUpdate** | [**SidecarUpdate**](SidecarUpdate.md) |  | 

### Return type

[**Sidecar**](Sidecar.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidProcessesGet

> V3AppsGuidProcessesGet200Response V3SidecarsGuidProcessesGet(ctx, guid).Page(page).PerPage(perPage).Execute()

List processes for a sidecar



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sidecar GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SidecarsAPI.V3SidecarsGuidProcessesGet(context.Background(), guid).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SidecarsAPI.V3SidecarsGuidProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidProcessesGet`: V3AppsGuidProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SidecarsAPI.V3SidecarsGuidProcessesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The sidecar GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 

### Return type

[**V3AppsGuidProcessesGet200Response**](V3AppsGuidProcessesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

