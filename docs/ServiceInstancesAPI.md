# \ServiceInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ServiceInstancesGuidRelationshipsSharedSpacesGet**](ServiceInstancesAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesGet) | **Get** /v3/service_instances/{guid}/relationships/shared_spaces | List shared spaces relationship (experimental)
[**V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet**](ServiceInstancesAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet) | **Get** /v3/service_instances/{guid}/relationships/shared_spaces/usage_summary | Get usage summary in shared spaces (experimental)



## V3ServiceInstancesGuidRelationshipsSharedSpacesGet

> V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response V3ServiceInstancesGuidRelationshipsSharedSpacesGet(ctx, guid).Execute()

List shared spaces relationship (experimental)



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
	guid := "guid_example" // string | GUID of the service instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidRelationshipsSharedSpacesGet`: V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidRelationshipsSharedSpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response**](V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet

> UsageSummary V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet(ctx, guid).Execute()

Get usage summary in shared spaces (experimental)



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
	guid := "guid_example" // string | GUID of the service instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageSummary**](UsageSummary.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

