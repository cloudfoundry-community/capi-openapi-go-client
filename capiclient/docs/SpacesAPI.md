# \SpacesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3SpacesGuidManifestDiffPost**](SpacesAPI.md#V3SpacesGuidManifestDiffPost) | **Post** /v3/spaces/{guid}/manifest_diff | Generate manifest diff for a space
[**V3SpacesGuidManifestPost**](SpacesAPI.md#V3SpacesGuidManifestPost) | **Post** /v3/spaces/{guid}/manifest | Apply manifest to a space
[**V3SpacesGuidRunningSecurityGroupsGet**](SpacesAPI.md#V3SpacesGuidRunningSecurityGroupsGet) | **Get** /v3/spaces/{guid}/running_security_groups | List running security groups for a space
[**V3SpacesGuidStagingSecurityGroupsGet**](SpacesAPI.md#V3SpacesGuidStagingSecurityGroupsGet) | **Get** /v3/spaces/{guid}/staging_security_groups | List staging security groups for a space



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
	resp, r, err := apiClient.SpacesAPI.V3SpacesGuidManifestDiffPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.V3SpacesGuidManifestDiffPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidManifestDiffPost`: ManifestDiff
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.V3SpacesGuidManifestDiffPost`: %v\n", resp)
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
	resp, r, err := apiClient.SpacesAPI.V3SpacesGuidManifestPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.V3SpacesGuidManifestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidManifestPost`: V3AppsGuidActionsApplyManifestPost202Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.V3SpacesGuidManifestPost`: %v\n", resp)
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


## V3SpacesGuidRunningSecurityGroupsGet

> V3SecurityGroupsGet200Response V3SpacesGuidRunningSecurityGroupsGet(ctx, guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()

List running security groups for a space



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
	guids := "guids_example" // string | Comma-delimited list of security group guids to filter by (optional)
	names := "names_example" // string | Comma-delimited list of security group names to filter by (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.V3SpacesGuidRunningSecurityGroupsGet(context.Background(), guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.V3SpacesGuidRunningSecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidRunningSecurityGroupsGet`: V3SecurityGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.V3SpacesGuidRunningSecurityGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidRunningSecurityGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **string** | Comma-delimited list of security group guids to filter by | 
 **names** | **string** | Comma-delimited list of security group names to filter by | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]

### Return type

[**V3SecurityGroupsGet200Response**](V3SecurityGroupsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidStagingSecurityGroupsGet

> V3SecurityGroupsGet200Response V3SpacesGuidStagingSecurityGroupsGet(ctx, guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()

List staging security groups for a space



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
	guids := "guids_example" // string | Comma-delimited list of security group guids to filter by (optional)
	names := "names_example" // string | Comma-delimited list of security group names to filter by (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.V3SpacesGuidStagingSecurityGroupsGet(context.Background(), guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.V3SpacesGuidStagingSecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidStagingSecurityGroupsGet`: V3SecurityGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.V3SpacesGuidStagingSecurityGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidStagingSecurityGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **string** | Comma-delimited list of security group guids to filter by | 
 **names** | **string** | Comma-delimited list of security group names to filter by | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]

### Return type

[**V3SecurityGroupsGet200Response**](V3SecurityGroupsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

