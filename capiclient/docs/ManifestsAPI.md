# \ManifestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidActionsApplyManifestPost**](ManifestsAPI.md#V3AppsGuidActionsApplyManifestPost) | **Post** /v3/apps/{guid}/actions/apply_manifest | Apply manifest to an app
[**V3AppsGuidManifestGet**](ManifestsAPI.md#V3AppsGuidManifestGet) | **Get** /v3/apps/{guid}/manifest | Get app manifest
[**V3SpacesGuidManifestDiffPost**](ManifestsAPI.md#V3SpacesGuidManifestDiffPost) | **Post** /v3/spaces/{guid}/manifest_diff | Generate manifest diff for a space
[**V3SpacesGuidManifestPost**](ManifestsAPI.md#V3SpacesGuidManifestPost) | **Post** /v3/spaces/{guid}/manifest | Apply manifest to a space



## V3AppsGuidActionsApplyManifestPost

> V3AppsGuidActionsApplyManifestPost202Response V3AppsGuidActionsApplyManifestPost(ctx, guid).Body(body).Execute()

Apply manifest to an app



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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestsAPI.V3AppsGuidActionsApplyManifestPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestsAPI.V3AppsGuidActionsApplyManifestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsApplyManifestPost`: V3AppsGuidActionsApplyManifestPost202Response
	fmt.Fprintf(os.Stdout, "Response from `ManifestsAPI.V3AppsGuidActionsApplyManifestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsApplyManifestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**V3AppsGuidActionsApplyManifestPost202Response**](V3AppsGuidActionsApplyManifestPost202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidManifestGet

> string V3AppsGuidManifestGet(ctx, guid).Execute()

Get app manifest



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestsAPI.V3AppsGuidManifestGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestsAPI.V3AppsGuidManifestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidManifestGet`: string
	fmt.Fprintf(os.Stdout, "Response from `ManifestsAPI.V3AppsGuidManifestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidManifestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidManifestDiffPost

> ManifestDiff V3SpacesGuidManifestDiffPost(ctx, guid).Body(body).Execute()

Generate manifest diff for a space



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestsAPI.V3SpacesGuidManifestDiffPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestsAPI.V3SpacesGuidManifestDiffPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidManifestDiffPost`: ManifestDiff
	fmt.Fprintf(os.Stdout, "Response from `ManifestsAPI.V3SpacesGuidManifestDiffPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidManifestDiffPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**ManifestDiff**](ManifestDiff.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidManifestPost

> V3AppsGuidActionsApplyManifestPost202Response V3SpacesGuidManifestPost(ctx, guid).Body(body).Execute()

Apply manifest to a space



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestsAPI.V3SpacesGuidManifestPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestsAPI.V3SpacesGuidManifestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidManifestPost`: V3AppsGuidActionsApplyManifestPost202Response
	fmt.Fprintf(os.Stdout, "Response from `ManifestsAPI.V3SpacesGuidManifestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidManifestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**V3AppsGuidActionsApplyManifestPost202Response**](V3AppsGuidActionsApplyManifestPost202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

