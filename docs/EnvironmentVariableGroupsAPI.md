# \EnvironmentVariableGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3EnvironmentVariableGroupsNameGet**](EnvironmentVariableGroupsAPI.md#V3EnvironmentVariableGroupsNameGet) | **Get** /v3/environment_variable_groups/{name} | Get an environment variable group
[**V3EnvironmentVariableGroupsNamePatch**](EnvironmentVariableGroupsAPI.md#V3EnvironmentVariableGroupsNamePatch) | **Patch** /v3/environment_variable_groups/{name} | Update environment variable group



## V3EnvironmentVariableGroupsNameGet

> EnvironmentVariableGroup V3EnvironmentVariableGroupsNameGet(ctx, name).Execute()

Get an environment variable group



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
	name := "name_example" // string | The name of the environment variable group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnvironmentVariableGroupsNameGet`: EnvironmentVariableGroup
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the environment variable group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnvironmentVariableGroupsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableGroup**](EnvironmentVariableGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnvironmentVariableGroupsNamePatch

> EnvironmentVariableGroup V3EnvironmentVariableGroupsNamePatch(ctx, name).EnvironmentVariableGroupUpdate(environmentVariableGroupUpdate).Execute()

Update environment variable group



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
	name := "name_example" // string | The name of the environment variable group
	environmentVariableGroupUpdate := *openapiclient.NewEnvironmentVariableGroupUpdate() // EnvironmentVariableGroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNamePatch(context.Background(), name).EnvironmentVariableGroupUpdate(environmentVariableGroupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnvironmentVariableGroupsNamePatch`: EnvironmentVariableGroup
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentVariableGroupsAPI.V3EnvironmentVariableGroupsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the environment variable group | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnvironmentVariableGroupsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentVariableGroupUpdate** | [**EnvironmentVariableGroupUpdate**](EnvironmentVariableGroupUpdate.md) |  | 

### Return type

[**EnvironmentVariableGroup**](EnvironmentVariableGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

