# \SpaceQuotasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3SpaceQuotasGet**](SpaceQuotasAPI.md#V3SpaceQuotasGet) | **Get** /v3/space_quotas | List space quotas
[**V3SpaceQuotasGuidDelete**](SpaceQuotasAPI.md#V3SpaceQuotasGuidDelete) | **Delete** /v3/space_quotas/{guid} | Delete a space quota
[**V3SpaceQuotasGuidGet**](SpaceQuotasAPI.md#V3SpaceQuotasGuidGet) | **Get** /v3/space_quotas/{guid} | Get a space quota
[**V3SpaceQuotasGuidPatch**](SpaceQuotasAPI.md#V3SpaceQuotasGuidPatch) | **Patch** /v3/space_quotas/{guid} | Update a space quota
[**V3SpaceQuotasGuidRelationshipsSpacesPost**](SpaceQuotasAPI.md#V3SpaceQuotasGuidRelationshipsSpacesPost) | **Post** /v3/space_quotas/{guid}/relationships/spaces | Apply a space quota to spaces
[**V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete**](SpaceQuotasAPI.md#V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete) | **Delete** /v3/space_quotas/{guid}/relationships/spaces/{space_guid} | Remove a space quota from a space
[**V3SpaceQuotasPost**](SpaceQuotasAPI.md#V3SpaceQuotasPost) | **Post** /v3/space_quotas | Create a space quota



## V3SpaceQuotasGet

> V3SpaceQuotasGet200Response V3SpaceQuotasGet(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Execute()

List space quotas



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
	guids := "1234,5678" // string | Comma-delimited list of space quota guids to filter by (optional)
	names := "small,medium,large" // string | Comma-delimited list of space quota names to filter by (optional)
	organizationGuids := "organizationGuids_example" // string | Comma-delimited list of organization guids to filter by (optional)
	spaceGuids := "spaceGuids_example" // string | Comma-delimited list of space guids to filter by (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)
	include := "organization" // string | Optionally include related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGet(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGet`: V3SpaceQuotasGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** | Comma-delimited list of space quota guids to filter by | 
 **names** | **string** | Comma-delimited list of space quota names to filter by | 
 **organizationGuids** | **string** | Comma-delimited list of organization guids to filter by | 
 **spaceGuids** | **string** | Comma-delimited list of space guids to filter by | 
 **labelSelector** | **string** | Filter by label selector | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 
 **include** | **string** | Optionally include related resources in the response | 

### Return type

[**V3SpaceQuotasGet200Response**](V3SpaceQuotasGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidDelete

> DeleteOrganizationQuota202Response V3SpaceQuotasGuidDelete(ctx, guid).Execute()

Delete a space quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space quota GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidDelete`: DeleteOrganizationQuota202Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasGuidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidDeleteRequest struct via the builder pattern


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


## V3SpaceQuotasGuidGet

> SpaceQuota V3SpaceQuotasGuidGet(ctx, guid).Include(include).Execute()

Get a space quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space quota GUID
	include := "include_example" // string | Optionally include related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidGet`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Optionally include related resources in the response | 

### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidPatch

> SpaceQuota V3SpaceQuotasGuidPatch(ctx, guid).SpaceQuotaUpdate(spaceQuotaUpdate).Execute()

Update a space quota



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space quota GUID
	spaceQuotaUpdate := *openapiclient.NewSpaceQuotaUpdate() // SpaceQuotaUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGuidPatch(context.Background(), guid).SpaceQuotaUpdate(spaceQuotaUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidPatch`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spaceQuotaUpdate** | [**SpaceQuotaUpdate**](SpaceQuotaUpdate.md) |  | 

### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidRelationshipsSpacesPost

> ApplyOrganizationQuotaToOrganizations201Response V3SpaceQuotasGuidRelationshipsSpacesPost(ctx, guid).V3SpaceQuotasGuidRelationshipsSpacesPostRequest(v3SpaceQuotasGuidRelationshipsSpacesPostRequest).Execute()

Apply a space quota to spaces



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space quota GUID
	v3SpaceQuotasGuidRelationshipsSpacesPostRequest := *openapiclient.NewV3SpaceQuotasGuidRelationshipsSpacesPostRequest([]openapiclient.V3SpaceQuotasGuidRelationshipsSpacesPostRequestDataInner{*openapiclient.NewV3SpaceQuotasGuidRelationshipsSpacesPostRequestDataInner("Guid_example")}) // V3SpaceQuotasGuidRelationshipsSpacesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGuidRelationshipsSpacesPost(context.Background(), guid).V3SpaceQuotasGuidRelationshipsSpacesPostRequest(v3SpaceQuotasGuidRelationshipsSpacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGuidRelationshipsSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidRelationshipsSpacesPost`: ApplyOrganizationQuotaToOrganizations201Response
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasGuidRelationshipsSpacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space quota GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidRelationshipsSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SpaceQuotasGuidRelationshipsSpacesPostRequest** | [**V3SpaceQuotasGuidRelationshipsSpacesPostRequest**](V3SpaceQuotasGuidRelationshipsSpacesPostRequest.md) |  | 

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


## V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete

> V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete(ctx, guid, spaceGuid).Execute()

Remove a space quota from a space



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space quota GUID
	spaceGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID to remove the quota from

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete(context.Background(), guid, spaceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The space quota GUID | 
**spaceGuid** | **string** | The space GUID to remove the quota from | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidRelationshipsSpacesSpaceGuidDeleteRequest struct via the builder pattern


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


## V3SpaceQuotasPost

> SpaceQuota V3SpaceQuotasPost(ctx).SpaceQuotaCreate(spaceQuotaCreate).Execute()

Create a space quota



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
	spaceQuotaCreate := *openapiclient.NewSpaceQuotaCreate("small-space-quota", *openapiclient.NewSpaceQuotaCreateRelationships(*openapiclient.NewSpaceQuotaRelationshipsOrganization(*openapiclient.NewSpaceQuotaRelationshipsOrganizationData("Guid_example")))) // SpaceQuotaCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceQuotasAPI.V3SpaceQuotasPost(context.Background()).SpaceQuotaCreate(spaceQuotaCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceQuotasAPI.V3SpaceQuotasPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasPost`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `SpaceQuotasAPI.V3SpaceQuotasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceQuotaCreate** | [**SpaceQuotaCreate**](SpaceQuotaCreate.md) |  | 

### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

