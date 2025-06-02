# \OrganizationQuotasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyOrganizationQuotaToOrganizations**](OrganizationQuotasAPI.md#ApplyOrganizationQuotaToOrganizations) | **Post** /v3/organization_quotas/{guid}/relationships/organizations | Apply an organization quota to organizations
[**CreateOrganizationQuota**](OrganizationQuotasAPI.md#CreateOrganizationQuota) | **Post** /v3/organization_quotas | Create an organization quota
[**DeleteOrganizationQuota**](OrganizationQuotasAPI.md#DeleteOrganizationQuota) | **Delete** /v3/organization_quotas/{guid} | Delete an organization quota
[**GetOrganizationQuota**](OrganizationQuotasAPI.md#GetOrganizationQuota) | **Get** /v3/organization_quotas/{guid} | Get an organization quota
[**ListOrganizationQuotas**](OrganizationQuotasAPI.md#ListOrganizationQuotas) | **Get** /v3/organization_quotas | List organization quotas
[**UpdateOrganizationQuota**](OrganizationQuotasAPI.md#UpdateOrganizationQuota) | **Patch** /v3/organization_quotas/{guid} | Update an organization quota



## ApplyOrganizationQuotaToOrganizations

> ApplyOrganizationQuotaToOrganizations201Response ApplyOrganizationQuotaToOrganizations(ctx, guid).ApplyOrganizationQuotaToOrganizationsRequest(applyOrganizationQuotaToOrganizationsRequest).Execute()

Apply an organization quota to organizations



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization quota GUID
	applyOrganizationQuotaToOrganizationsRequest := *openapiclient.NewApplyOrganizationQuotaToOrganizationsRequest([]openapiclient.ApplyOrganizationQuotaToOrganizationsRequestDataInner{*openapiclient.NewApplyOrganizationQuotaToOrganizationsRequestDataInner("Guid_example")}) // ApplyOrganizationQuotaToOrganizationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.ApplyOrganizationQuotaToOrganizations(context.Background(), guid).ApplyOrganizationQuotaToOrganizationsRequest(applyOrganizationQuotaToOrganizationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.ApplyOrganizationQuotaToOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyOrganizationQuotaToOrganizations`: ApplyOrganizationQuotaToOrganizations201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.ApplyOrganizationQuotaToOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The organization quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyOrganizationQuotaToOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyOrganizationQuotaToOrganizationsRequest** | [**ApplyOrganizationQuotaToOrganizationsRequest**](ApplyOrganizationQuotaToOrganizationsRequest.md) |  | 

### Return type

[**ApplyOrganizationQuotaToOrganizations201Response**](ApplyOrganizationQuotaToOrganizations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationQuota

> OrganizationQuota CreateOrganizationQuota(ctx).OrganizationQuotaCreate(organizationQuotaCreate).Execute()

Create an organization quota



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
	organizationQuotaCreate := *openapiclient.NewOrganizationQuotaCreate("default-quota") // OrganizationQuotaCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.CreateOrganizationQuota(context.Background()).OrganizationQuotaCreate(organizationQuotaCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.CreateOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.CreateOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationQuotaCreate** | [**OrganizationQuotaCreate**](OrganizationQuotaCreate.md) |  | 

### Return type

[**OrganizationQuota**](OrganizationQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationQuota

> DeleteOrganizationQuota202Response DeleteOrganizationQuota(ctx, guid).Execute()

Delete an organization quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization quota GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.DeleteOrganizationQuota(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.DeleteOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationQuota`: DeleteOrganizationQuota202Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.DeleteOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The organization quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationQuotaRequest struct via the builder pattern


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


## GetOrganizationQuota

> OrganizationQuota GetOrganizationQuota(ctx, guid).Include(include).Execute()

Get an organization quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization quota GUID
	include := "include_example" // string | Optionally include related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.GetOrganizationQuota(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.GetOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.GetOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The organization quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Optionally include related resources in the response | 

### Return type

[**OrganizationQuota**](OrganizationQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationQuotas

> ListOrganizationQuotas200Response ListOrganizationQuotas(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Execute()

List organization quotas



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
	guids := "1234,5678" // string | Comma-delimited list of organization quota guids to filter by (optional)
	names := "small,medium,large" // string | Comma-delimited list of organization quota names to filter by (optional)
	organizationGuids := "organizationGuids_example" // string | Comma-delimited list of organization guids to filter by (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)
	include := "organizations" // string | Optionally include related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.ListOrganizationQuotas(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.ListOrganizationQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationQuotas`: ListOrganizationQuotas200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.ListOrganizationQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** | Comma-delimited list of organization quota guids to filter by | 
 **names** | **string** | Comma-delimited list of organization quota names to filter by | 
 **organizationGuids** | **string** | Comma-delimited list of organization guids to filter by | 
 **labelSelector** | **string** | Filter by label selector | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 
 **include** | **string** | Optionally include related resources in the response | 

### Return type

[**ListOrganizationQuotas200Response**](ListOrganizationQuotas200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationQuota

> OrganizationQuota UpdateOrganizationQuota(ctx, guid).OrganizationQuotaUpdate(organizationQuotaUpdate).Execute()

Update an organization quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The organization quota GUID
	organizationQuotaUpdate := *openapiclient.NewOrganizationQuotaUpdate() // OrganizationQuotaUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationQuotasAPI.UpdateOrganizationQuota(context.Background(), guid).OrganizationQuotaUpdate(organizationQuotaUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationQuotasAPI.UpdateOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `OrganizationQuotasAPI.UpdateOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The organization quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationQuotaUpdate** | [**OrganizationQuotaUpdate**](OrganizationQuotaUpdate.md) |  | 

### Return type

[**OrganizationQuota**](OrganizationQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

