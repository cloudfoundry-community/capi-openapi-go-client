# \MonitoringAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ProcessesGuidStatsGet**](MonitoringAPI.md#V3ProcessesGuidStatsGet) | **Get** /v3/processes/{guid}/stats | Get stats for a process



## V3ProcessesGuidStatsGet

> V3ProcessesGuidStatsGet200Response V3ProcessesGuidStatsGet(ctx, guid).Execute()

Get stats for a process



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitoringAPI.V3ProcessesGuidStatsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAPI.V3ProcessesGuidStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidStatsGet`: V3ProcessesGuidStatsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MonitoringAPI.V3ProcessesGuidStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3ProcessesGuidStatsGet200Response**](V3ProcessesGuidStatsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

