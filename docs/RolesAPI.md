# \RolesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3RolesGet**](RolesAPI.md#V3RolesGet) | **Get** /v3/roles | List roles
[**V3RolesGuidDelete**](RolesAPI.md#V3RolesGuidDelete) | **Delete** /v3/roles/{guid} | Delete a role
[**V3RolesGuidGet**](RolesAPI.md#V3RolesGuidGet) | **Get** /v3/roles/{guid} | Get a role
[**V3RolesPost**](RolesAPI.md#V3RolesPost) | **Post** /v3/roles | Create a role



## V3RolesGet

> V3RolesGet200Response V3RolesGet(ctx).Guids(guids).Types(types).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).UserGuids(userGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List roles



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
	guids := "1234,5678" // string | Comma-delimited list of role guids to filter by (optional)
	types := "organization_manager,space_developer" // string | Comma-delimited list of role types to filter by (optional)
	spaceGuids := "spaceGuids_example" // string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := "organizationGuids_example" // string | Comma-delimited list of organization guids to filter by (optional)
	userGuids := "userGuids_example" // string | Comma-delimited list of user guids to filter by (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")
	include := "user,organization" // string | Optionally include related resources in the response (optional)
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.V3RolesGet(context.Background()).Guids(guids).Types(types).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).UserGuids(userGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.V3RolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesGet`: V3RolesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.V3RolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** | Comma-delimited list of role guids to filter by | 
 **types** | **string** | Comma-delimited list of role types to filter by | 
 **spaceGuids** | **string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **string** | Comma-delimited list of organization guids to filter by | 
 **userGuids** | **string** | Comma-delimited list of user guids to filter by | 
 **labelSelector** | **string** | Filter by label selector | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]
 **include** | **string** | Optionally include related resources in the response | 
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 

### Return type

[**V3RolesGet200Response**](V3RolesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesGuidDelete

> DeleteOrganizationQuota202Response V3RolesGuidDelete(ctx, guid).Execute()

Delete a role



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The role GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.V3RolesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.V3RolesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesGuidDelete`: DeleteOrganizationQuota202Response
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.V3RolesGuidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The role GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOrganizationQuota202Response**](DeleteOrganizationQuota202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesGuidGet

> Role V3RolesGuidGet(ctx, guid).Include(include).Execute()

Get a role



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The role GUID
	include := "user,organization" // string | Optionally include related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.V3RolesGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.V3RolesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesGuidGet`: Role
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.V3RolesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The role GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Optionally include related resources in the response | 

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesPost

> Role V3RolesPost(ctx).RoleCreate(roleCreate).Execute()

Create a role



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
	roleCreate := *openapiclient.NewRoleCreate(*openapiclient.NewRoleCreateRelationships(*openapiclient.NewRoleCreateRelationshipsUser(openapiclient.RoleCreate_relationships_user_data{RoleCreateRelationshipsUserDataOneOf: openapiclient.NewRoleCreateRelationshipsUserDataOneOf("Guid_example")})), "space_developer") // RoleCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.V3RolesPost(context.Background()).RoleCreate(roleCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.V3RolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesPost`: Role
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.V3RolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCreate** | [**RoleCreate**](RoleCreate.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

