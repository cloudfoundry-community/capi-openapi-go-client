# \ServicePlanVisibilityAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ServicePlansGuidVisibilityGet**](ServicePlanVisibilityAPI.md#V3ServicePlansGuidVisibilityGet) | **Get** /v3/service_plans/{guid}/visibility | Get a service plan visibility
[**V3ServicePlansGuidVisibilityOrganizationGuidDelete**](ServicePlanVisibilityAPI.md#V3ServicePlansGuidVisibilityOrganizationGuidDelete) | **Delete** /v3/service_plans/{guid}/visibility/{organization_guid} | Remove organization from a service plan visibility
[**V3ServicePlansGuidVisibilityPatch**](ServicePlanVisibilityAPI.md#V3ServicePlansGuidVisibilityPatch) | **Patch** /v3/service_plans/{guid}/visibility | Update a service plan visibility
[**V3ServicePlansGuidVisibilityPost**](ServicePlanVisibilityAPI.md#V3ServicePlansGuidVisibilityPost) | **Post** /v3/service_plans/{guid}/visibility | Apply a service plan visibility



## V3ServicePlansGuidVisibilityGet

> ServicePlanVisibility V3ServicePlansGuidVisibilityGet(ctx, guid).Execute()

Get a service plan visibility



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The service plan GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityGet`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityOrganizationGuidDelete

> V3ServicePlansGuidVisibilityOrganizationGuidDelete(ctx, guid, organizationGuid).Execute()

Remove organization from a service plan visibility



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The service plan GUID
	organizationGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization GUID to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityOrganizationGuidDelete(context.Background(), guid, organizationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityOrganizationGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 
**organizationGuid** | **string** | The organization GUID to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityOrganizationGuidDeleteRequest struct via the builder pattern


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


## V3ServicePlansGuidVisibilityPatch

> ServicePlanVisibility V3ServicePlansGuidVisibilityPatch(ctx, guid).ServicePlanVisibilityUpdate(servicePlanVisibilityUpdate).Execute()

Update a service plan visibility



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The service plan GUID
	servicePlanVisibilityUpdate := *openapiclient.NewServicePlanVisibilityUpdate("organization") // ServicePlanVisibilityUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPatch(context.Background(), guid).ServicePlanVisibilityUpdate(servicePlanVisibilityUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityPatch`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePlanVisibilityUpdate** | [**ServicePlanVisibilityUpdate**](ServicePlanVisibilityUpdate.md) |  | 

### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityPost

> ServicePlanVisibility V3ServicePlansGuidVisibilityPost(ctx, guid).ServicePlanVisibilityApply(servicePlanVisibilityApply).Execute()

Apply a service plan visibility



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The service plan GUID
	servicePlanVisibilityApply := *openapiclient.NewServicePlanVisibilityApply("organization") // ServicePlanVisibilityApply | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPost(context.Background(), guid).ServicePlanVisibilityApply(servicePlanVisibilityApply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityPost`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `ServicePlanVisibilityAPI.V3ServicePlansGuidVisibilityPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePlanVisibilityApply** | [**ServicePlanVisibilityApply**](ServicePlanVisibilityApply.md) |  | 

### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

