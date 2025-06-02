# \SpaceFeaturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3SpacesGuidFeaturesGet**](SpaceFeaturesAPI.md#V3SpacesGuidFeaturesGet) | **Get** /v3/spaces/{guid}/features | List space features
[**V3SpacesGuidFeaturesNameGet**](SpaceFeaturesAPI.md#V3SpacesGuidFeaturesNameGet) | **Get** /v3/spaces/{guid}/features/{name} | Get a space feature
[**V3SpacesGuidFeaturesNamePatch**](SpaceFeaturesAPI.md#V3SpacesGuidFeaturesNamePatch) | **Patch** /v3/spaces/{guid}/features/{name} | Update a space feature



## V3SpacesGuidFeaturesGet

> V3SpacesGuidFeaturesGet200Response V3SpacesGuidFeaturesGet(ctx, guid).Execute()

List space features



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceFeaturesAPI.V3SpacesGuidFeaturesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceFeaturesAPI.V3SpacesGuidFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesGet`: V3SpacesGuidFeaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceFeaturesAPI.V3SpacesGuidFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3SpacesGuidFeaturesGet200Response**](V3SpacesGuidFeaturesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidFeaturesNameGet

> SpaceFeature V3SpacesGuidFeaturesNameGet(ctx, guid, name).Execute()

Get a space feature



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID
	name := "name_example" // string | The feature name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceFeaturesAPI.V3SpacesGuidFeaturesNameGet(context.Background(), guid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceFeaturesAPI.V3SpacesGuidFeaturesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesNameGet`: SpaceFeature
	fmt.Fprintf(os.Stdout, "Response from `SpaceFeaturesAPI.V3SpacesGuidFeaturesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SpaceFeature**](SpaceFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidFeaturesNamePatch

> SpaceFeature V3SpacesGuidFeaturesNamePatch(ctx, guid, name).SpaceFeatureUpdate(spaceFeatureUpdate).Execute()

Update a space feature



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID
	name := "name_example" // string | The feature name
	spaceFeatureUpdate := *openapiclient.NewSpaceFeatureUpdate(true) // SpaceFeatureUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceFeaturesAPI.V3SpacesGuidFeaturesNamePatch(context.Background(), guid, name).SpaceFeatureUpdate(spaceFeatureUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceFeaturesAPI.V3SpacesGuidFeaturesNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesNamePatch`: SpaceFeature
	fmt.Fprintf(os.Stdout, "Response from `SpaceFeaturesAPI.V3SpacesGuidFeaturesNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spaceFeatureUpdate** | [**SpaceFeatureUpdate**](SpaceFeatureUpdate.md) |  | 

### Return type

[**SpaceFeature**](SpaceFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

