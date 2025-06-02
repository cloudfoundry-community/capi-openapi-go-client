# \AppFeaturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidFeaturesGet**](AppFeaturesAPI.md#V3AppsGuidFeaturesGet) | **Get** /v3/apps/{guid}/features | List app features
[**V3AppsGuidFeaturesNameGet**](AppFeaturesAPI.md#V3AppsGuidFeaturesNameGet) | **Get** /v3/apps/{guid}/features/{name} | Get an app feature
[**V3AppsGuidFeaturesNamePatch**](AppFeaturesAPI.md#V3AppsGuidFeaturesNamePatch) | **Patch** /v3/apps/{guid}/features/{name} | Update an app feature



## V3AppsGuidFeaturesGet

> V3AppsGuidFeaturesGet200Response V3AppsGuidFeaturesGet(ctx, guid).Execute()

List app features



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFeaturesAPI.V3AppsGuidFeaturesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFeaturesAPI.V3AppsGuidFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesGet`: V3AppsGuidFeaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppFeaturesAPI.V3AppsGuidFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidFeaturesGet200Response**](V3AppsGuidFeaturesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNameGet

> AppFeature V3AppsGuidFeaturesNameGet(ctx, guid, name).Execute()

Get an app feature



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
	name := "name_example" // string | The feature name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFeaturesAPI.V3AppsGuidFeaturesNameGet(context.Background(), guid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFeaturesAPI.V3AppsGuidFeaturesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesNameGet`: AppFeature
	fmt.Fprintf(os.Stdout, "Response from `AppFeaturesAPI.V3AppsGuidFeaturesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppFeature**](AppFeature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNamePatch

> AppFeature V3AppsGuidFeaturesNamePatch(ctx, guid, name).AppFeatureUpdate(appFeatureUpdate).Execute()

Update an app feature



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
	name := "name_example" // string | The feature name
	appFeatureUpdate := *openapiclient.NewAppFeatureUpdate(true) // AppFeatureUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppFeaturesAPI.V3AppsGuidFeaturesNamePatch(context.Background(), guid, name).AppFeatureUpdate(appFeatureUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppFeaturesAPI.V3AppsGuidFeaturesNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesNamePatch`: AppFeature
	fmt.Fprintf(os.Stdout, "Response from `AppFeaturesAPI.V3AppsGuidFeaturesNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appFeatureUpdate** | [**AppFeatureUpdate**](AppFeatureUpdate.md) |  | 

### Return type

[**AppFeature**](AppFeature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

