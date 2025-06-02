# \RoutesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3RoutesGet**](RoutesAPI.md#V3RoutesGet) | **Get** /v3/routes | List routes
[**V3RoutesGuidDelete**](RoutesAPI.md#V3RoutesGuidDelete) | **Delete** /v3/routes/{guid} | Delete a route
[**V3RoutesGuidDestinationsDestinationGuidDelete**](RoutesAPI.md#V3RoutesGuidDestinationsDestinationGuidDelete) | **Delete** /v3/routes/{guid}/destinations/{destination_guid} | Remove destination from a route
[**V3RoutesGuidDestinationsGet**](RoutesAPI.md#V3RoutesGuidDestinationsGet) | **Get** /v3/routes/{guid}/destinations | List destinations for a route
[**V3RoutesGuidDestinationsPatch**](RoutesAPI.md#V3RoutesGuidDestinationsPatch) | **Patch** /v3/routes/{guid}/destinations | Update destinations for a route
[**V3RoutesGuidDestinationsPost**](RoutesAPI.md#V3RoutesGuidDestinationsPost) | **Post** /v3/routes/{guid}/destinations | Add destinations to a route
[**V3RoutesGuidGet**](RoutesAPI.md#V3RoutesGuidGet) | **Get** /v3/routes/{guid} | Get a route
[**V3RoutesGuidPatch**](RoutesAPI.md#V3RoutesGuidPatch) | **Patch** /v3/routes/{guid} | Update a route
[**V3RoutesGuidRelationshipsSharedSpacesGet**](RoutesAPI.md#V3RoutesGuidRelationshipsSharedSpacesGet) | **Get** /v3/routes/{guid}/relationships/shared_spaces | List shared spaces for a route (experimental)
[**V3RoutesGuidRelationshipsSpacePatch**](RoutesAPI.md#V3RoutesGuidRelationshipsSpacePatch) | **Patch** /v3/routes/{guid}/relationships/space | Share a route to another space (experimental)
[**V3RoutesGuidTransferOwnerPost**](RoutesAPI.md#V3RoutesGuidTransferOwnerPost) | **Post** /v3/routes/{guid}/transfer_owner | Transfer route ownership
[**V3RoutesPost**](RoutesAPI.md#V3RoutesPost) | **Post** /v3/routes | Create a route



## V3RoutesGet

> V3RoutesGet200Response V3RoutesGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).AppGuids(appGuids).DomainGuids(domainGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Hosts(hosts).Paths(paths).Ports(ports).ServiceInstanceGuids(serviceInstanceGuids).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()

List routes



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
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field to sort results by (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	domainGuids := "domainGuids_example" // string | Filter by domain GUIDs (comma-separated) (optional)
	spaceGuids := "spaceGuids_example" // string | Filter by space GUIDs (comma-separated) (optional)
	organizationGuids := "organizationGuids_example" // string | Filter by organization GUIDs (comma-separated) (optional)
	hosts := "app,api" // string | Filter by hostnames (comma-separated) (optional)
	paths := "/api/v1,/api/v2" // string | Filter by paths (comma-separated) (optional)
	ports := "8080,9000" // string | Filter by ports (comma-separated) (optional)
	serviceInstanceGuids := "serviceInstanceGuids_example" // string | Filter by service instance GUIDs (comma-separated) (optional)
	labelSelector := "environment=production,tier!=backend" // string | Filter by labels using label selector syntax (optional)
	include := "include_example" // string | Include related resources (optional)
	createdAts := "created_ats[gte]=2020-01-01T00:00:00Z" // string | Filter by creation timestamp. Supports multiple formats: - Range: created_ats=2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: created_ats[gt]=2020-01-01T00:00:00Z - Greater than or equal: created_ats[gte]=2020-01-01T00:00:00Z - Less than: created_ats[lt]=2020-12-31T23:59:59Z - Less than or equal: created_ats[lte]=2020-12-31T23:59:59Z  (optional)
	updatedAts := "updated_ats[lt]=2020-12-31T23:59:59Z" // string | Filter by update timestamp. Supports multiple formats: - Range: updated_ats=2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: updated_ats[gt]=2020-01-01T00:00:00Z - Greater than or equal: updated_ats[gte]=2020-01-01T00:00:00Z - Less than: updated_ats[lt]=2020-12-31T23:59:59Z - Less than or equal: updated_ats[lte]=2020-12-31T23:59:59Z  (optional)
	fields := "fields[routes]=guid,host,path" // string | Fields to include in the response. Use dot notation for nested fields. Example: fields[routes]=guid,host,path or fields[domain]=name,guid  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).AppGuids(appGuids).DomainGuids(domainGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Hosts(hosts).Paths(paths).Ports(ports).ServiceInstanceGuids(serviceInstanceGuids).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGet`: V3RoutesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 
 **domainGuids** | **string** | Filter by domain GUIDs (comma-separated) | 
 **spaceGuids** | **string** | Filter by space GUIDs (comma-separated) | 
 **organizationGuids** | **string** | Filter by organization GUIDs (comma-separated) | 
 **hosts** | **string** | Filter by hostnames (comma-separated) | 
 **paths** | **string** | Filter by paths (comma-separated) | 
 **ports** | **string** | Filter by ports (comma-separated) | 
 **serviceInstanceGuids** | **string** | Filter by service instance GUIDs (comma-separated) | 
 **labelSelector** | **string** | Filter by labels using label selector syntax | 
 **include** | **string** | Include related resources | 
 **createdAts** | **string** | Filter by creation timestamp. Supports multiple formats: - Range: created_ats&#x3D;2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: created_ats[gt]&#x3D;2020-01-01T00:00:00Z - Greater than or equal: created_ats[gte]&#x3D;2020-01-01T00:00:00Z - Less than: created_ats[lt]&#x3D;2020-12-31T23:59:59Z - Less than or equal: created_ats[lte]&#x3D;2020-12-31T23:59:59Z  | 
 **updatedAts** | **string** | Filter by update timestamp. Supports multiple formats: - Range: updated_ats&#x3D;2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: updated_ats[gt]&#x3D;2020-01-01T00:00:00Z - Greater than or equal: updated_ats[gte]&#x3D;2020-01-01T00:00:00Z - Less than: updated_ats[lt]&#x3D;2020-12-31T23:59:59Z - Less than or equal: updated_ats[lte]&#x3D;2020-12-31T23:59:59Z  | 
 **fields** | **string** | Fields to include in the response. Use dot notation for nested fields. Example: fields[routes]&#x3D;guid,host,path or fields[domain]&#x3D;name,guid  | 

### Return type

[**V3RoutesGet200Response**](V3RoutesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDelete

> V3RoutesGuidDelete(ctx, guid).Execute()

Delete a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutesAPI.V3RoutesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDeleteRequest struct via the builder pattern


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


## V3RoutesGuidDestinationsDestinationGuidDelete

> V3RoutesGuidDestinationsDestinationGuidDelete(ctx, guid, destinationGuid).Execute()

Remove destination from a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	destinationGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The destination GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoutesAPI.V3RoutesGuidDestinationsDestinationGuidDelete(context.Background(), guid, destinationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidDestinationsDestinationGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 
**destinationGuid** | **string** | The destination GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsDestinationGuidDeleteRequest struct via the builder pattern


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


## V3RoutesGuidDestinationsGet

> V3RoutesGuidDestinationsGet200Response V3RoutesGuidDestinationsGet(ctx, guid).Page(page).PerPage(perPage).AppGuids(appGuids).Execute()

List destinations for a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	appGuids := "appGuids_example" // string | Filter by app GUIDs (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidDestinationsGet(context.Background(), guid).Page(page).PerPage(perPage).AppGuids(appGuids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidDestinationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsGet`: V3RoutesGuidDestinationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidDestinationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **appGuids** | **string** | Filter by app GUIDs (comma-separated) | 

### Return type

[**V3RoutesGuidDestinationsGet200Response**](V3RoutesGuidDestinationsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDestinationsPatch

> V3RoutesGuidDestinationsPost200Response V3RoutesGuidDestinationsPatch(ctx, guid).V3RoutesGuidDestinationsPatchRequest(v3RoutesGuidDestinationsPatchRequest).Execute()

Update destinations for a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidDestinationsPatchRequest := *openapiclient.NewV3RoutesGuidDestinationsPatchRequest([]openapiclient.V3RoutesGuidDestinationsPatchRequestDestinationsInner{*openapiclient.NewV3RoutesGuidDestinationsPatchRequestDestinationsInner(*openapiclient.NewV3RoutesGuidDestinationsPatchRequestDestinationsInnerApp("Guid_example"))}) // V3RoutesGuidDestinationsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidDestinationsPatch(context.Background(), guid).V3RoutesGuidDestinationsPatchRequest(v3RoutesGuidDestinationsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidDestinationsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsPatch`: V3RoutesGuidDestinationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidDestinationsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidDestinationsPatchRequest** | [**V3RoutesGuidDestinationsPatchRequest**](V3RoutesGuidDestinationsPatchRequest.md) |  | 

### Return type

[**V3RoutesGuidDestinationsPost200Response**](V3RoutesGuidDestinationsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidDestinationsPost

> V3RoutesGuidDestinationsPost200Response V3RoutesGuidDestinationsPost(ctx, guid).V3RoutesGuidDestinationsPostRequest(v3RoutesGuidDestinationsPostRequest).Execute()

Add destinations to a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidDestinationsPostRequest := *openapiclient.NewV3RoutesGuidDestinationsPostRequest([]openapiclient.V3RoutesGuidDestinationsPostRequestDestinationsInner{*openapiclient.NewV3RoutesGuidDestinationsPostRequestDestinationsInner(*openapiclient.NewV3RoutesGuidDestinationsPostRequestDestinationsInnerApp("Guid_example"))}) // V3RoutesGuidDestinationsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidDestinationsPost(context.Background(), guid).V3RoutesGuidDestinationsPostRequest(v3RoutesGuidDestinationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidDestinationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidDestinationsPost`: V3RoutesGuidDestinationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidDestinationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidDestinationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidDestinationsPostRequest** | [**V3RoutesGuidDestinationsPostRequest**](V3RoutesGuidDestinationsPostRequest.md) |  | 

### Return type

[**V3RoutesGuidDestinationsPost200Response**](V3RoutesGuidDestinationsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidGet

> Route V3RoutesGuidGet(ctx, guid).Include(include).Execute()

Get a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	include := "include_example" // string | Include related resources (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidGet`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Include related resources | 

### Return type

[**Route**](Route.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidPatch

> Route V3RoutesGuidPatch(ctx, guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()

Update a route



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3PackagesGuidPatchRequest := *openapiclient.NewV3PackagesGuidPatchRequest() // V3PackagesGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidPatch(context.Background(), guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidPatch`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3PackagesGuidPatchRequest** | [**V3PackagesGuidPatchRequest**](V3PackagesGuidPatchRequest.md) |  | 

### Return type

[**Route**](Route.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidRelationshipsSharedSpacesGet

> V3RoutesGuidRelationshipsSharedSpacesGet200Response V3RoutesGuidRelationshipsSharedSpacesGet(ctx, guid).Execute()

List shared spaces for a route (experimental)



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidRelationshipsSharedSpacesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidRelationshipsSharedSpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidRelationshipsSharedSpacesGet`: V3RoutesGuidRelationshipsSharedSpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidRelationshipsSharedSpacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidRelationshipsSharedSpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3RoutesGuidRelationshipsSharedSpacesGet200Response**](V3RoutesGuidRelationshipsSharedSpacesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidRelationshipsSpacePatch

> V3RoutesGuidRelationshipsSpacePatch200Response V3RoutesGuidRelationshipsSpacePatch(ctx, guid).V3RoutesGuidRelationshipsSpacePatchRequest(v3RoutesGuidRelationshipsSpacePatchRequest).Execute()

Share a route to another space (experimental)



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidRelationshipsSpacePatchRequest := *openapiclient.NewV3RoutesGuidRelationshipsSpacePatchRequest([]openapiclient.V3RoutesGuidRelationshipsSpacePatchRequestDataInner{*openapiclient.NewV3RoutesGuidRelationshipsSpacePatchRequestDataInner("Guid_example")}) // V3RoutesGuidRelationshipsSpacePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidRelationshipsSpacePatch(context.Background(), guid).V3RoutesGuidRelationshipsSpacePatchRequest(v3RoutesGuidRelationshipsSpacePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidRelationshipsSpacePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidRelationshipsSpacePatch`: V3RoutesGuidRelationshipsSpacePatch200Response
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidRelationshipsSpacePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidRelationshipsSpacePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidRelationshipsSpacePatchRequest** | [**V3RoutesGuidRelationshipsSpacePatchRequest**](V3RoutesGuidRelationshipsSpacePatchRequest.md) |  | 

### Return type

[**V3RoutesGuidRelationshipsSpacePatch200Response**](V3RoutesGuidRelationshipsSpacePatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGuidTransferOwnerPost

> Route V3RoutesGuidTransferOwnerPost(ctx, guid).V3RoutesGuidTransferOwnerPostRequest(v3RoutesGuidTransferOwnerPostRequest).Execute()

Transfer route ownership



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The route GUID
	v3RoutesGuidTransferOwnerPostRequest := *openapiclient.NewV3RoutesGuidTransferOwnerPostRequest(*openapiclient.NewV3AppsPostRequestRelationships(*openapiclient.NewV3AppsPostRequestRelationshipsSpace(*openapiclient.NewV3AppsPostRequestRelationshipsSpaceData("Guid_example")))) // V3RoutesGuidTransferOwnerPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesGuidTransferOwnerPost(context.Background(), guid).V3RoutesGuidTransferOwnerPostRequest(v3RoutesGuidTransferOwnerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesGuidTransferOwnerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidTransferOwnerPost`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesGuidTransferOwnerPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The route GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGuidTransferOwnerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3RoutesGuidTransferOwnerPostRequest** | [**V3RoutesGuidTransferOwnerPostRequest**](V3RoutesGuidTransferOwnerPostRequest.md) |  | 

### Return type

[**Route**](Route.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesPost

> Route V3RoutesPost(ctx).V3RoutesPostRequest(v3RoutesPostRequest).Execute()

Create a route



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
	v3RoutesPostRequest := *openapiclient.NewV3RoutesPostRequest(*openapiclient.NewV3RoutesPostRequestRelationships(*openapiclient.NewV3AppsPostRequestRelationshipsSpace(*openapiclient.NewV3AppsPostRequestRelationshipsSpaceData("Guid_example")), *openapiclient.NewV3AppsPostRequestRelationshipsSpace(*openapiclient.NewV3AppsPostRequestRelationshipsSpaceData("Guid_example")))) // V3RoutesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.V3RoutesPost(context.Background()).V3RoutesPostRequest(v3RoutesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.V3RoutesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesPost`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.V3RoutesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3RoutesPostRequest** | [**V3RoutesPostRequest**](V3RoutesPostRequest.md) |  | 

### Return type

[**Route**](Route.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

