# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AdminActionsClearBuildpackCachePost**](AdminAPI.md#V3AdminActionsClearBuildpackCachePost) | **Post** /v3/admin/actions/clear_buildpack_cache | Clear buildpack cache



## V3AdminActionsClearBuildpackCachePost

> Job V3AdminActionsClearBuildpackCachePost(ctx).Execute()

Clear buildpack cache



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.V3AdminActionsClearBuildpackCachePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.V3AdminActionsClearBuildpackCachePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AdminActionsClearBuildpackCachePost`: Job
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.V3AdminActionsClearBuildpackCachePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3AdminActionsClearBuildpackCachePostRequest struct via the builder pattern


### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

