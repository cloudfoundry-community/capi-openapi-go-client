# \ProcessesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidProcessesGet**](ProcessesAPI.md#V3AppsGuidProcessesGet) | **Get** /v3/apps/{guid}/processes | List processes for an app
[**V3AppsGuidProcessesTypeGet**](ProcessesAPI.md#V3AppsGuidProcessesTypeGet) | **Get** /v3/apps/{guid}/processes/{type} | Get a process by type for an app
[**V3ProcessesGet**](ProcessesAPI.md#V3ProcessesGet) | **Get** /v3/processes | List processes
[**V3ProcessesGuidActionsScalePost**](ProcessesAPI.md#V3ProcessesGuidActionsScalePost) | **Post** /v3/processes/{guid}/actions/scale | Scale a process
[**V3ProcessesGuidGet**](ProcessesAPI.md#V3ProcessesGuidGet) | **Get** /v3/processes/{guid} | Get a process
[**V3ProcessesGuidInstancesIndexDelete**](ProcessesAPI.md#V3ProcessesGuidInstancesIndexDelete) | **Delete** /v3/processes/{guid}/instances/{index} | Terminate a process instance
[**V3ProcessesGuidPatch**](ProcessesAPI.md#V3ProcessesGuidPatch) | **Patch** /v3/processes/{guid} | Update a process
[**V3ProcessesGuidSidecarsGet**](ProcessesAPI.md#V3ProcessesGuidSidecarsGet) | **Get** /v3/processes/{guid}/sidecars | List sidecars for a process
[**V3ProcessesGuidStatsGet**](ProcessesAPI.md#V3ProcessesGuidStatsGet) | **Get** /v3/processes/{guid}/stats | Get stats for a process
[**V3SidecarsGuidProcessesGet**](ProcessesAPI.md#V3SidecarsGuidProcessesGet) | **Get** /v3/sidecars/{guid}/processes | List processes for a sidecar



## V3AppsGuidProcessesGet

> V3AppsGuidProcessesGet200Response V3AppsGuidProcessesGet(ctx, guid).Page(page).PerPage(perPage).Types(types).Execute()

List processes for an app



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
	types := "types_example" // string | Filter by process types (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3AppsGuidProcessesGet(context.Background(), guid).Page(page).PerPage(perPage).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3AppsGuidProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidProcessesGet`: V3AppsGuidProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3AppsGuidProcessesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **types** | **string** | Filter by process types (comma-separated) | 

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


## V3AppsGuidProcessesTypeGet

> Process V3AppsGuidProcessesTypeGet(ctx, guid, type_).Execute()

Get a process by type for an app



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
	type_ := "web" // string | The process type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3AppsGuidProcessesTypeGet(context.Background(), guid, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3AppsGuidProcessesTypeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidProcessesTypeGet`: Process
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3AppsGuidProcessesTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**type_** | **string** | The process type | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidProcessesTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Process**](Process.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGet

> V3AppsGuidProcessesGet200Response V3ProcessesGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Guids(guids).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List processes



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
	guids := "guids_example" // string | Filter by process GUIDs (comma-separated) (optional)
	types := "web,worker" // string | Filter by process types (comma-separated) (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	spaceGuids := "spaceGuids_example" // string | Filter by space GUIDs (comma-separated) (optional)
	organizationGuids := "organizationGuids_example" // string | Filter by organization GUIDs (comma-separated) (optional)
	labelSelector := "environment=production,tier!=backend" // string | Filter by labels using label selector syntax (optional)
	createdAts := "createdAts_example" // string | Filter by creation timestamp range (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamp range (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Guids(guids).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGet`: V3AppsGuidProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **guids** | **string** | Filter by process GUIDs (comma-separated) | 
 **types** | **string** | Filter by process types (comma-separated) | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 
 **spaceGuids** | **string** | Filter by space GUIDs (comma-separated) | 
 **organizationGuids** | **string** | Filter by organization GUIDs (comma-separated) | 
 **labelSelector** | **string** | Filter by labels using label selector syntax | 
 **createdAts** | **string** | Filter by creation timestamp range | 
 **updatedAts** | **string** | Filter by update timestamp range | 

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


## V3ProcessesGuidActionsScalePost

> Process V3ProcessesGuidActionsScalePost(ctx, guid).V3ProcessesGuidActionsScalePostRequest(v3ProcessesGuidActionsScalePostRequest).Execute()

Scale a process



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	v3ProcessesGuidActionsScalePostRequest := *openapiclient.NewV3ProcessesGuidActionsScalePostRequest() // V3ProcessesGuidActionsScalePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGuidActionsScalePost(context.Background(), guid).V3ProcessesGuidActionsScalePostRequest(v3ProcessesGuidActionsScalePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidActionsScalePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidActionsScalePost`: Process
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGuidActionsScalePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidActionsScalePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ProcessesGuidActionsScalePostRequest** | [**V3ProcessesGuidActionsScalePostRequest**](V3ProcessesGuidActionsScalePostRequest.md) |  | 

### Return type

[**Process**](Process.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidGet

> Process V3ProcessesGuidGet(ctx, guid).Execute()

Get a process



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidGet`: Process
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Process**](Process.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidInstancesIndexDelete

> V3ProcessesGuidInstancesIndexDelete(ctx, guid, index).Execute()

Terminate a process instance



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	index := int32(56) // int32 | The instance index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcessesAPI.V3ProcessesGuidInstancesIndexDelete(context.Background(), guid, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidInstancesIndexDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 
**index** | **int32** | The instance index | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidInstancesIndexDeleteRequest struct via the builder pattern


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


## V3ProcessesGuidPatch

> Process V3ProcessesGuidPatch(ctx, guid).V3ProcessesGuidPatchRequest(v3ProcessesGuidPatchRequest).Execute()

Update a process



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	v3ProcessesGuidPatchRequest := *openapiclient.NewV3ProcessesGuidPatchRequest() // V3ProcessesGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGuidPatch(context.Background(), guid).V3ProcessesGuidPatchRequest(v3ProcessesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidPatch`: Process
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ProcessesGuidPatchRequest** | [**V3ProcessesGuidPatchRequest**](V3ProcessesGuidPatchRequest.md) |  | 

### Return type

[**Process**](Process.md)

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
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
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
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGuidSidecarsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidSidecarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidSidecarsGet`: V3AppsGuidSidecarsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGuidSidecarsGet`: %v\n", resp)
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


## V3ProcessesGuidStatsGet

> V3ProcessesGuidStatsGet200Response V3ProcessesGuidStatsGet(ctx, guid).Execute()

Get stats for a process



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3ProcessesGuidStatsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3ProcessesGuidStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidStatsGet`: V3ProcessesGuidStatsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3ProcessesGuidStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3ProcessesGuidStatsGet200Response**](V3ProcessesGuidStatsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The sidecar GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcessesAPI.V3SidecarsGuidProcessesGet(context.Background(), guid).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcessesAPI.V3SidecarsGuidProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidProcessesGet`: V3AppsGuidProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ProcessesAPI.V3SidecarsGuidProcessesGet`: %v\n", resp)
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

