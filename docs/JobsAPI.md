# \JobsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3JobsGet**](JobsAPI.md#V3JobsGet) | **Get** /v3/jobs | List jobs
[**V3JobsGuidGet**](JobsAPI.md#V3JobsGuidGet) | **Get** /v3/jobs/{guid} | Get a job



## V3JobsGet

> V3JobsGet200Response V3JobsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Guids(guids).Execute()

List jobs



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
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "-created_at")
	states := "PROCESSING,COMPLETE" // string | Comma-delimited list of job states to filter by (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)
	createdAts := "created_ats[gte]=2024-01-01T00:00:00Z" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)
	guids := "guids_example" // string | Comma-delimited list of job GUIDs to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.V3JobsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Guids(guids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.V3JobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3JobsGet`: V3JobsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.V3JobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3JobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;-created_at&quot;]
 **states** | **string** | Comma-delimited list of job states to filter by | 
 **labelSelector** | **string** | Filter by label selector | 
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 
 **guids** | **string** | Comma-delimited list of job GUIDs to filter by | 

### Return type

[**V3JobsGet200Response**](V3JobsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3JobsGuidGet

> Job V3JobsGuidGet(ctx, guid).Execute()

Get a job



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The job GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.V3JobsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.V3JobsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3JobsGuidGet`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.V3JobsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The job GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3JobsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

