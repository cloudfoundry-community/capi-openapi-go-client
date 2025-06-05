# \FeatureFlagsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3FeatureFlagsGet**](FeatureFlagsAPI.md#V3FeatureFlagsGet) | **Get** /v3/feature_flags | List feature flags
[**V3FeatureFlagsNameGet**](FeatureFlagsAPI.md#V3FeatureFlagsNameGet) | **Get** /v3/feature_flags/{name} | Get a feature flag
[**V3FeatureFlagsNamePatch**](FeatureFlagsAPI.md#V3FeatureFlagsNamePatch) | **Patch** /v3/feature_flags/{name} | Update a feature flag



## V3FeatureFlagsGet

> V3FeatureFlagsGet200Response V3FeatureFlagsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).UpdatedAts(updatedAts).Names(names).LabelSelector(labelSelector).Execute()

List feature flags



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
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "name")
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)
	names := "user_org_creation,private_domain_creation" // string | Comma-delimited list of feature flag names to filter by (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsAPI.V3FeatureFlagsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).UpdatedAts(updatedAts).Names(names).LabelSelector(labelSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.V3FeatureFlagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsGet`: V3FeatureFlagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsAPI.V3FeatureFlagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;name&quot;]
 **updatedAts** | **string** | Filter by update time | 
 **names** | **string** | Comma-delimited list of feature flag names to filter by | 
 **labelSelector** | **string** | Filter by label selector | 

### Return type

[**V3FeatureFlagsGet200Response**](V3FeatureFlagsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3FeatureFlagsNameGet

> FeatureFlag V3FeatureFlagsNameGet(ctx, name).Execute()

Get a feature flag



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
	name := "user_org_creation" // string | The name of the feature flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsAPI.V3FeatureFlagsNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.V3FeatureFlagsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsNameGet`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsAPI.V3FeatureFlagsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3FeatureFlagsNamePatch

> FeatureFlag V3FeatureFlagsNamePatch(ctx, name).FeatureFlagUpdate(featureFlagUpdate).Execute()

Update a feature flag



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
	name := "user_org_creation" // string | The name of the feature flag
	featureFlagUpdate := *openapiclient.NewFeatureFlagUpdate() // FeatureFlagUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsAPI.V3FeatureFlagsNamePatch(context.Background(), name).FeatureFlagUpdate(featureFlagUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.V3FeatureFlagsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsNamePatch`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsAPI.V3FeatureFlagsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureFlagUpdate** | [**FeatureFlagUpdate**](FeatureFlagUpdate.md) |  | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

