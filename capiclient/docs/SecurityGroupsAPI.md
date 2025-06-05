# \SecurityGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3SecurityGroupsGet**](SecurityGroupsAPI.md#V3SecurityGroupsGet) | **Get** /v3/security_groups | List security groups
[**V3SecurityGroupsGuidDelete**](SecurityGroupsAPI.md#V3SecurityGroupsGuidDelete) | **Delete** /v3/security_groups/{guid} | Delete a security group
[**V3SecurityGroupsGuidGet**](SecurityGroupsAPI.md#V3SecurityGroupsGuidGet) | **Get** /v3/security_groups/{guid} | Get a security group
[**V3SecurityGroupsGuidPatch**](SecurityGroupsAPI.md#V3SecurityGroupsGuidPatch) | **Patch** /v3/security_groups/{guid} | Update a security group
[**V3SecurityGroupsGuidRelationshipsRunningSpacesPost**](SecurityGroupsAPI.md#V3SecurityGroupsGuidRelationshipsRunningSpacesPost) | **Post** /v3/security_groups/{guid}/relationships/running_spaces | Bind running security group to spaces
[**V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete**](SecurityGroupsAPI.md#V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete) | **Delete** /v3/security_groups/{guid}/relationships/running_spaces/{space_guid} | Unbind running security group from a space
[**V3SecurityGroupsGuidRelationshipsStagingSpacesPost**](SecurityGroupsAPI.md#V3SecurityGroupsGuidRelationshipsStagingSpacesPost) | **Post** /v3/security_groups/{guid}/relationships/staging_spaces | Bind staging security group to spaces
[**V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete**](SecurityGroupsAPI.md#V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete) | **Delete** /v3/security_groups/{guid}/relationships/staging_spaces/{space_guid} | Unbind staging security group from a space
[**V3SecurityGroupsPost**](SecurityGroupsAPI.md#V3SecurityGroupsPost) | **Post** /v3/security_groups | Create a security group
[**V3SpacesGuidRunningSecurityGroupsGet**](SecurityGroupsAPI.md#V3SpacesGuidRunningSecurityGroupsGet) | **Get** /v3/spaces/{guid}/running_security_groups | List running security groups for a space
[**V3SpacesGuidStagingSecurityGroupsGet**](SecurityGroupsAPI.md#V3SpacesGuidStagingSecurityGroupsGet) | **Get** /v3/spaces/{guid}/staging_security_groups | List staging security groups for a space



## V3SecurityGroupsGet

> V3SecurityGroupsGet200Response V3SecurityGroupsGet(ctx).Guids(guids).Names(names).GloballyEnabledRunning(globallyEnabledRunning).GloballyEnabledStaging(globallyEnabledStaging).RunningSpaceGuids(runningSpaceGuids).StagingSpaceGuids(stagingSpaceGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List security groups



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
	guids := "1234,5678" // string | Comma-delimited list of security group guids to filter by (optional)
	names := "my-security-group,another-security-group" // string | Comma-delimited list of security group names to filter by (optional)
	globallyEnabledRunning := true // bool | If true, only include security groups enabled globally for running apps (optional)
	globallyEnabledStaging := true // bool | If true, only include security groups enabled globally for staging apps (optional)
	runningSpaceGuids := "runningSpaceGuids_example" // string | Comma-delimited list of space guids to filter by for running security groups (optional)
	stagingSpaceGuids := "stagingSpaceGuids_example" // string | Comma-delimited list of space guids to filter by for staging security groups (optional)
	labelSelector := "environment==production,tier!=backend" // string | Filter by label selector (optional)
	page := int32(56) // int32 | Page to display (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results per page (optional) (default to 50)
	orderBy := "orderBy_example" // string | Value to sort by (optional) (default to "created_at")
	createdAts := "createdAts_example" // string | Filter by creation time (optional)
	updatedAts := "updatedAts_example" // string | Filter by update time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGet(context.Background()).Guids(guids).Names(names).GloballyEnabledRunning(globallyEnabledRunning).GloballyEnabledStaging(globallyEnabledStaging).RunningSpaceGuids(runningSpaceGuids).StagingSpaceGuids(stagingSpaceGuids).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGet`: V3SecurityGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** | Comma-delimited list of security group guids to filter by | 
 **names** | **string** | Comma-delimited list of security group names to filter by | 
 **globallyEnabledRunning** | **bool** | If true, only include security groups enabled globally for running apps | 
 **globallyEnabledStaging** | **bool** | If true, only include security groups enabled globally for staging apps | 
 **runningSpaceGuids** | **string** | Comma-delimited list of space guids to filter by for running security groups | 
 **stagingSpaceGuids** | **string** | Comma-delimited list of space guids to filter by for staging security groups | 
 **labelSelector** | **string** | Filter by label selector | 
 **page** | **int32** | Page to display | [default to 1]
 **perPage** | **int32** | Number of results per page | [default to 50]
 **orderBy** | **string** | Value to sort by | [default to &quot;created_at&quot;]
 **createdAts** | **string** | Filter by creation time | 
 **updatedAts** | **string** | Filter by update time | 

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


## V3SecurityGroupsGuidDelete

> DeleteOrganizationQuota202Response V3SecurityGroupsGuidDelete(ctx, guid).Execute()

Delete a security group



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidDelete`: DeleteOrganizationQuota202Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGuidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidDeleteRequest struct via the builder pattern


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


## V3SecurityGroupsGuidGet

> SecurityGroup V3SecurityGroupsGuidGet(ctx, guid).Execute()

Get a security group



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidGet`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SecurityGroupsGuidPatch

> SecurityGroup V3SecurityGroupsGuidPatch(ctx, guid).SecurityGroupUpdate(securityGroupUpdate).Execute()

Update a security group



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID
	securityGroupUpdate := *openapiclient.NewSecurityGroupUpdate() // SecurityGroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidPatch(context.Background(), guid).SecurityGroupUpdate(securityGroupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidPatch`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityGroupUpdate** | [**SecurityGroupUpdate**](SecurityGroupUpdate.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SecurityGroupsGuidRelationshipsRunningSpacesPost

> ApplyOrganizationQuotaToOrganizations201Response V3SecurityGroupsGuidRelationshipsRunningSpacesPost(ctx, guid).V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest(v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest).Execute()

Bind running security group to spaces



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID
	v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest := *openapiclient.NewV3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest([]openapiclient.V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequestDataInner{*openapiclient.NewV3SecurityGroupsGuidRelationshipsRunningSpacesPostRequestDataInner("Guid_example")}) // V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsRunningSpacesPost(context.Background(), guid).V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest(v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsRunningSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidRelationshipsRunningSpacesPost`: ApplyOrganizationQuotaToOrganizations201Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsRunningSpacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest** | [**V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest**](V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest.md) |  | 

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


## V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete

> V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete(ctx, guid, spaceGuid).Execute()

Unbind running security group from a space



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID
	spaceGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID to unbind

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete(context.Background(), guid, spaceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 
**spaceGuid** | **string** | The space GUID to unbind | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidRelationshipsRunningSpacesSpaceGuidDeleteRequest struct via the builder pattern


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


## V3SecurityGroupsGuidRelationshipsStagingSpacesPost

> ApplyOrganizationQuotaToOrganizations201Response V3SecurityGroupsGuidRelationshipsStagingSpacesPost(ctx, guid).V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest(v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest).Execute()

Bind staging security group to spaces



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID
	v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest := *openapiclient.NewV3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest([]openapiclient.V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequestDataInner{*openapiclient.NewV3SecurityGroupsGuidRelationshipsRunningSpacesPostRequestDataInner("Guid_example")}) // V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsStagingSpacesPost(context.Background(), guid).V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest(v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsStagingSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidRelationshipsStagingSpacesPost`: ApplyOrganizationQuotaToOrganizations201Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsStagingSpacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidRelationshipsStagingSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest** | [**V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest**](V3SecurityGroupsGuidRelationshipsRunningSpacesPostRequest.md) |  | 

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


## V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete

> V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete(ctx, guid, spaceGuid).Execute()

Unbind staging security group from a space



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The security group GUID
	spaceGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The space GUID to unbind

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete(context.Background(), guid, spaceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The security group GUID | 
**spaceGuid** | **string** | The space GUID to unbind | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGuidRelationshipsStagingSpacesSpaceGuidDeleteRequest struct via the builder pattern


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


## V3SecurityGroupsPost

> SecurityGroup V3SecurityGroupsPost(ctx).SecurityGroupCreate(securityGroupCreate).Execute()

Create a security group



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
	securityGroupCreate := *openapiclient.NewSecurityGroupCreate("my-security-group") // SecurityGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupsAPI.V3SecurityGroupsPost(context.Background()).SecurityGroupCreate(securityGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SecurityGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsPost`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SecurityGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityGroupCreate** | [**SecurityGroupCreate**](SecurityGroupCreate.md) |  | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.SecurityGroupsAPI.V3SpacesGuidRunningSecurityGroupsGet(context.Background(), guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SpacesGuidRunningSecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidRunningSecurityGroupsGet`: V3SecurityGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SpacesGuidRunningSecurityGroupsGet`: %v\n", resp)
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
	resp, r, err := apiClient.SecurityGroupsAPI.V3SpacesGuidStagingSecurityGroupsGet(context.Background(), guid).Guids(guids).Names(names).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsAPI.V3SpacesGuidStagingSecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidStagingSecurityGroupsGet`: V3SecurityGroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsAPI.V3SpacesGuidStagingSecurityGroupsGet`: %v\n", resp)
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

