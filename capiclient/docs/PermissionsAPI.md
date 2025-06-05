# \PermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppPermissions**](PermissionsAPI.md#GetAppPermissions) | **Get** /v3/apps/{guid}/permissions | Get app permissions



## GetAppPermissions

> GetAppPermissions200Response GetAppPermissions(ctx, guid).Execute()

Get app permissions



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
	resp, r, err := apiClient.PermissionsAPI.GetAppPermissions(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetAppPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppPermissions`: GetAppPermissions200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetAppPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppPermissions200Response**](GetAppPermissions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

