# \AppUsageEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost**](AppUsageEventsAPI.md#V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost) | **Post** /v3/app_usage_events/actions/destructively_purge_all_and_reseed | Purge and seed app usage events
[**V3AppUsageEventsGet**](AppUsageEventsAPI.md#V3AppUsageEventsGet) | **Get** /v3/app_usage_events | List app usage events
[**V3AppUsageEventsGuidGet**](AppUsageEventsAPI.md#V3AppUsageEventsGuidGet) | **Get** /v3/app_usage_events/{guid} | Retrieve an app usage event



## V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost

> map[string]interface{} V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost(ctx).Execute()

Purge and seed app usage events

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppUsageEventsAPI.V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppUsageEventsAPI.V3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppUsageEventsActionsDestructivelyPurgeAllAndReseedPostRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppUsageEventsGet

> AppUsageEventList V3AppUsageEventsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).AfterGuid(afterGuid).Guids(guids).CreatedAts(createdAts).Execute()

List app usage events

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	afterGuid := "afterGuid_example" // string | Filters out events before and including the event with the given guid (optional)
	guids := []string{"Inner_example"} // []string | Comma-delimited list of usage event guids to filter by (optional)
	createdAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).AfterGuid(afterGuid).Guids(guids).CreatedAts(createdAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppUsageEventsAPI.V3AppUsageEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppUsageEventsGet`: AppUsageEventList
	fmt.Fprintf(os.Stdout, "Response from `AppUsageEventsAPI.V3AppUsageEventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AppUsageEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **afterGuid** | **string** | Filters out events before and including the event with the given guid | 
 **guids** | **[]string** | Comma-delimited list of usage event guids to filter by | 
 **createdAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by | 

### Return type

[**AppUsageEventList**](AppUsageEventList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppUsageEventsGuidGet

> AppUsageEvent V3AppUsageEventsGuidGet(ctx, guid).Execute()

Retrieve an app usage event

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
	guid := "guid_example" // string | Unique identifier for the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppUsageEventsAPI.V3AppUsageEventsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppUsageEventsAPI.V3AppUsageEventsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppUsageEventsGuidGet`: AppUsageEvent
	fmt.Fprintf(os.Stdout, "Response from `AppUsageEventsAPI.V3AppUsageEventsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppUsageEventsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUsageEvent**](AppUsageEvent.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

