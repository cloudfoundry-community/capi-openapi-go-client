# \DomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3DomainsGet**](DomainsAPI.md#V3DomainsGet) | **Get** /v3/domains | List domains
[**V3DomainsGuidDelete**](DomainsAPI.md#V3DomainsGuidDelete) | **Delete** /v3/domains/{guid} | Delete a domain
[**V3DomainsGuidGet**](DomainsAPI.md#V3DomainsGuidGet) | **Get** /v3/domains/{guid} | Get a domain
[**V3DomainsGuidPatch**](DomainsAPI.md#V3DomainsGuidPatch) | **Patch** /v3/domains/{guid} | Update a domain
[**V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete**](DomainsAPI.md#V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete) | **Delete** /v3/domains/{guid}/relationships/shared_organizations/{org_guid} | Unshare a domain
[**V3DomainsGuidRelationshipsSharedOrganizationsPost**](DomainsAPI.md#V3DomainsGuidRelationshipsSharedOrganizationsPost) | **Post** /v3/domains/{guid}/relationships/shared_organizations | Share a domain
[**V3DomainsPost**](DomainsAPI.md#V3DomainsPost) | **Post** /v3/domains | Create a domain



## V3DomainsGet

> V3DomainsGet200Response V3DomainsGet(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List domains



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
	guids := []string{"Inner_example"} // []string | Comma-delimited list of guids to filter by (optional)
	names := []string{"Inner_example"} // []string | Comma-delimited list of domain names to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of owning organization guids to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	labelSelector := "labelSelector_example" // string | List of label selector requirements (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V3DomainsGet(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DomainsGet`: V3DomainsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V3DomainsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of guids to filter by | 
 **names** | **[]string** | Comma-delimited list of domain names to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of owning organization guids to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **labelSelector** | **string** | List of label selector requirements | 
 **createdAts** | **string** | Timestamp to filter by | 
 **updatedAts** | **string** | Timestamp to filter by | 

### Return type

[**V3DomainsGet200Response**](V3DomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsGuidDelete

> V3DomainsGuidDelete(ctx, guid).Execute()

Delete a domain



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.V3DomainsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsGuidGet

> Domain V3DomainsGuidGet(ctx, guid).Execute()

Get a domain



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V3DomainsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DomainsGuidGet`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V3DomainsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsGuidPatch

> Domain V3DomainsGuidPatch(ctx, guid).V3DomainsGuidPatchRequest(v3DomainsGuidPatchRequest).Execute()

Update a domain



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	v3DomainsGuidPatchRequest := *openapiclient.NewV3DomainsGuidPatchRequest() // V3DomainsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V3DomainsGuidPatch(context.Background(), guid).V3DomainsGuidPatchRequest(v3DomainsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DomainsGuidPatch`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V3DomainsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DomainsGuidPatchRequest** | [**V3DomainsGuidPatchRequest**](V3DomainsGuidPatchRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete

> V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete(ctx, guid, orgGuid).Execute()

Unshare a domain



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orgGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete(context.Background(), guid, orgGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 
**orgGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGuidRelationshipsSharedOrganizationsOrgGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsGuidRelationshipsSharedOrganizationsPost

> V3DomainsGuidRelationshipsSharedOrganizationsPost200Response V3DomainsGuidRelationshipsSharedOrganizationsPost(ctx, guid).V3DomainsGuidRelationshipsSharedOrganizationsPostRequest(v3DomainsGuidRelationshipsSharedOrganizationsPostRequest).Execute()

Share a domain



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	v3DomainsGuidRelationshipsSharedOrganizationsPostRequest := *openapiclient.NewV3DomainsGuidRelationshipsSharedOrganizationsPostRequest() // V3DomainsGuidRelationshipsSharedOrganizationsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V3DomainsGuidRelationshipsSharedOrganizationsPost(context.Background(), guid).V3DomainsGuidRelationshipsSharedOrganizationsPostRequest(v3DomainsGuidRelationshipsSharedOrganizationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsGuidRelationshipsSharedOrganizationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DomainsGuidRelationshipsSharedOrganizationsPost`: V3DomainsGuidRelationshipsSharedOrganizationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V3DomainsGuidRelationshipsSharedOrganizationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsGuidRelationshipsSharedOrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DomainsGuidRelationshipsSharedOrganizationsPostRequest** | [**V3DomainsGuidRelationshipsSharedOrganizationsPostRequest**](V3DomainsGuidRelationshipsSharedOrganizationsPostRequest.md) |  | 

### Return type

[**V3DomainsGuidRelationshipsSharedOrganizationsPost200Response**](V3DomainsGuidRelationshipsSharedOrganizationsPost200Response.md)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DomainsPost

> Domain V3DomainsPost(ctx).V3DomainsPostRequest(v3DomainsPostRequest).Execute()

Create a domain



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
	v3DomainsPostRequest := *openapiclient.NewV3DomainsPostRequest() // V3DomainsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.V3DomainsPost(context.Background()).V3DomainsPostRequest(v3DomainsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.V3DomainsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DomainsPost`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.V3DomainsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3DomainsPostRequest** | [**V3DomainsPostRequest**](V3DomainsPostRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[OrgManager](../README.md#OrgManager)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

