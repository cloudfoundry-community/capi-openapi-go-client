# \InstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3ProcessesGuidInstancesIndexDelete**](InstancesAPI.md#V3ProcessesGuidInstancesIndexDelete) | **Delete** /v3/processes/{guid}/instances/{index} | Terminate a process instance



## V3ProcessesGuidInstancesIndexDelete

> V3ProcessesGuidInstancesIndexDelete(ctx, guid, index).Execute()

Terminate a process instance



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The process GUID
	index := int32(56) // int32 | The instance index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstancesAPI.V3ProcessesGuidInstancesIndexDelete(context.Background(), guid, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.V3ProcessesGuidInstancesIndexDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The process GUID | 
**index** | **int32** | The instance index | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidInstancesIndexDeleteRequest struct via the builder pattern


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

