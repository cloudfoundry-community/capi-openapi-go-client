# \SshAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGuidSshEnabledGet**](SshAPI.md#V3AppsGuidSshEnabledGet) | **Get** /v3/apps/{guid}/ssh_enabled | Get SSH enabled



## V3AppsGuidSshEnabledGet

> V3AppsGuidSshEnabledGet200Response V3AppsGuidSshEnabledGet(ctx, guid).Execute()

Get SSH enabled



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
	resp, r, err := apiClient.SshAPI.V3AppsGuidSshEnabledGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshAPI.V3AppsGuidSshEnabledGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSshEnabledGet`: V3AppsGuidSshEnabledGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SshAPI.V3AppsGuidSshEnabledGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSshEnabledGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidSshEnabledGet200Response**](V3AppsGuidSshEnabledGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

