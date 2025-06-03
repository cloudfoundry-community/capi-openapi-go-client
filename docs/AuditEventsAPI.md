# \AuditEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AuditEventsGet**](AuditEventsAPI.md#V3AuditEventsGet) | **Get** /v3/audit_events | List audit events
[**V3AuditEventsGuidGet**](AuditEventsAPI.md#V3AuditEventsGuidGet) | **Get** /v3/audit_events/{guid} | Retrieve an audit event



## V3AuditEventsGet

> AuditEventList V3AuditEventsGet(ctx).Types(types).TargetGuids(targetGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List audit events

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
	types := []string{"Inner_example"} // []string | Comma-delimited list of event types to filter by (optional)
	targetGuids := []string{"Inner_example"} // []string | Comma-delimited list of target guids to filter by. Also supports filtering by exclusion. (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	createdAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by (optional)
	updatedAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditEventsAPI.V3AuditEventsGet(context.Background()).Types(types).TargetGuids(targetGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsAPI.V3AuditEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AuditEventsGet`: AuditEventList
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsAPI.V3AuditEventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AuditEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]string** | Comma-delimited list of event types to filter by | 
 **targetGuids** | **[]string** | Comma-delimited list of target guids to filter by. Also supports filtering by exclusion. | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **createdAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by | 
 **updatedAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by | 

### Return type

[**AuditEventList**](AuditEventList.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AuditEventsGuidGet

> AuditEvent V3AuditEventsGuidGet(ctx, guid).Execute()

Retrieve an audit event

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
	resp, r, err := apiClient.AuditEventsAPI.V3AuditEventsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditEventsAPI.V3AuditEventsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AuditEventsGuidGet`: AuditEvent
	fmt.Fprintf(os.Stdout, "Response from `AuditEventsAPI.V3AuditEventsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AuditEventsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditEvent**](AuditEvent.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

