# \ScalingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ProcessesGuidActionsScalePost**](ScalingAPI.md#V3ProcessesGuidActionsScalePost) | **Post** /v3/processes/{guid}/actions/scale | Scale a process



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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	v3ProcessesGuidActionsScalePostRequest := *openapiclient.NewV3ProcessesGuidActionsScalePostRequest() // V3ProcessesGuidActionsScalePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingAPI.V3ProcessesGuidActionsScalePost(context.Background(), guid).V3ProcessesGuidActionsScalePostRequest(v3ProcessesGuidActionsScalePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingAPI.V3ProcessesGuidActionsScalePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidActionsScalePost`: Process
	fmt.Fprintf(os.Stdout, "Response from `ScalingAPI.V3ProcessesGuidActionsScalePost`: %v\n", resp)
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

