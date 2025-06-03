# \AppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3AppsGet**](AppsAPI.md#V3AppsGet) | **Get** /v3/apps | List apps
[**V3AppsGuidActionsApplyManifestPost**](AppsAPI.md#V3AppsGuidActionsApplyManifestPost) | **Post** /v3/apps/{guid}/actions/apply_manifest | Apply manifest to an app
[**V3AppsGuidActionsClearBuildpackCachePost**](AppsAPI.md#V3AppsGuidActionsClearBuildpackCachePost) | **Post** /v3/apps/{guid}/actions/clear_buildpack_cache | Clear buildpack cache
[**V3AppsGuidActionsRestartPost**](AppsAPI.md#V3AppsGuidActionsRestartPost) | **Post** /v3/apps/{guid}/actions/restart | Restart an app
[**V3AppsGuidActionsStartPost**](AppsAPI.md#V3AppsGuidActionsStartPost) | **Post** /v3/apps/{guid}/actions/start | Start an app
[**V3AppsGuidActionsStopPost**](AppsAPI.md#V3AppsGuidActionsStopPost) | **Post** /v3/apps/{guid}/actions/stop | Stop an app
[**V3AppsGuidBuildsGet**](AppsAPI.md#V3AppsGuidBuildsGet) | **Get** /v3/apps/{guid}/builds | List builds for an app
[**V3AppsGuidDelete**](AppsAPI.md#V3AppsGuidDelete) | **Delete** /v3/apps/{guid} | Delete an app
[**V3AppsGuidEnvironmentVariablesGet**](AppsAPI.md#V3AppsGuidEnvironmentVariablesGet) | **Get** /v3/apps/{guid}/environment_variables | Get app environment variables
[**V3AppsGuidEnvironmentVariablesPatch**](AppsAPI.md#V3AppsGuidEnvironmentVariablesPatch) | **Patch** /v3/apps/{guid}/environment_variables | Update app environment variables
[**V3AppsGuidFeaturesGet**](AppsAPI.md#V3AppsGuidFeaturesGet) | **Get** /v3/apps/{guid}/features | List app features
[**V3AppsGuidFeaturesNameGet**](AppsAPI.md#V3AppsGuidFeaturesNameGet) | **Get** /v3/apps/{guid}/features/{name} | Get an app feature
[**V3AppsGuidFeaturesNamePatch**](AppsAPI.md#V3AppsGuidFeaturesNamePatch) | **Patch** /v3/apps/{guid}/features/{name} | Update an app feature
[**V3AppsGuidGet**](AppsAPI.md#V3AppsGuidGet) | **Get** /v3/apps/{guid} | Get an app
[**V3AppsGuidManifestDiffPost**](AppsAPI.md#V3AppsGuidManifestDiffPost) | **Post** /v3/apps/{guid}/manifest_diff | Generate manifest diff (experimental)
[**V3AppsGuidManifestGet**](AppsAPI.md#V3AppsGuidManifestGet) | **Get** /v3/apps/{guid}/manifest | Get app manifest
[**V3AppsGuidPackagesGet**](AppsAPI.md#V3AppsGuidPackagesGet) | **Get** /v3/apps/{guid}/packages | List packages for an app
[**V3AppsGuidPatch**](AppsAPI.md#V3AppsGuidPatch) | **Patch** /v3/apps/{guid} | Update an app
[**V3AppsGuidPermissionsGet**](AppsAPI.md#V3AppsGuidPermissionsGet) | **Get** /v3/apps/{guid}/permissions | Get app permissions
[**V3AppsGuidProcessesGet**](AppsAPI.md#V3AppsGuidProcessesGet) | **Get** /v3/apps/{guid}/processes | List processes for an app
[**V3AppsGuidProcessesTypeGet**](AppsAPI.md#V3AppsGuidProcessesTypeGet) | **Get** /v3/apps/{guid}/processes/{type} | Get a process by type for an app
[**V3AppsGuidRelationshipsCurrentDropletGet**](AppsAPI.md#V3AppsGuidRelationshipsCurrentDropletGet) | **Get** /v3/apps/{guid}/relationships/current_droplet | Get current droplet relationship
[**V3AppsGuidRelationshipsCurrentDropletPatch**](AppsAPI.md#V3AppsGuidRelationshipsCurrentDropletPatch) | **Patch** /v3/apps/{guid}/relationships/current_droplet | Update current droplet relationship
[**V3AppsGuidSshEnabledGet**](AppsAPI.md#V3AppsGuidSshEnabledGet) | **Get** /v3/apps/{guid}/ssh_enabled | Get SSH enabled
[**V3AppsPost**](AppsAPI.md#V3AppsPost) | **Post** /v3/apps | Create an app



## V3AppsGet

> V3AppsGet200Response V3AppsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).States(states).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()

List apps



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
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field to sort results by (optional)
	names := "app1,app2" // string | Filter by app names (comma-separated) (optional)
	guids := "guids_example" // string | Filter by app GUIDs (comma-separated) (optional)
	organizationGuids := "organizationGuids_example" // string | Filter by organization GUIDs (comma-separated) (optional)
	spaceGuids := "spaceGuids_example" // string | Filter by space GUIDs (comma-separated) (optional)
	stacks := "stacks_example" // string | Filter by stack names (comma-separated) (optional)
	states := "STARTED,STOPPED" // string | Filter by app states (comma-separated) (optional)
	include := "include_example" // string | Include related resources (optional)
	lifecycleType := "lifecycleType_example" // string | Filter by lifecycle type (optional)
	labelSelector := "environment=production,tier!=backend" // string | Filter by labels using label selector syntax (optional)
	createdAts := "created_ats[gte]=2020-01-01T00:00:00Z" // string | Filter by creation timestamp. Supports multiple formats: - Range: created_ats=2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: created_ats[gt]=2020-01-01T00:00:00Z - Greater than or equal: created_ats[gte]=2020-01-01T00:00:00Z - Less than: created_ats[lt]=2020-12-31T23:59:59Z - Less than or equal: created_ats[lte]=2020-12-31T23:59:59Z  (optional)
	updatedAts := "updatedAts_example" // string | Filter by update timestamp. Supports multiple formats: - Range: updated_ats=2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: updated_ats[gt]=2020-01-01T00:00:00Z - Greater than or equal: updated_ats[gte]=2020-01-01T00:00:00Z - Less than: updated_ats[lt]=2020-12-31T23:59:59Z - Less than or equal: updated_ats[lte]=2020-12-31T23:59:59Z  (optional)
	fields := "fields[apps]=name,guid,state" // string | Fields to include in the response. Use dot notation for nested fields. Example: fields[apps]=name,guid,state or fields[space]=name,guid  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).States(states).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGet`: V3AppsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **names** | **string** | Filter by app names (comma-separated) | 
 **guids** | **string** | Filter by app GUIDs (comma-separated) | 
 **organizationGuids** | **string** | Filter by organization GUIDs (comma-separated) | 
 **spaceGuids** | **string** | Filter by space GUIDs (comma-separated) | 
 **stacks** | **string** | Filter by stack names (comma-separated) | 
 **states** | **string** | Filter by app states (comma-separated) | 
 **include** | **string** | Include related resources | 
 **lifecycleType** | **string** | Filter by lifecycle type | 
 **labelSelector** | **string** | Filter by labels using label selector syntax | 
 **createdAts** | **string** | Filter by creation timestamp. Supports multiple formats: - Range: created_ats&#x3D;2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: created_ats[gt]&#x3D;2020-01-01T00:00:00Z - Greater than or equal: created_ats[gte]&#x3D;2020-01-01T00:00:00Z - Less than: created_ats[lt]&#x3D;2020-12-31T23:59:59Z - Less than or equal: created_ats[lte]&#x3D;2020-12-31T23:59:59Z  | 
 **updatedAts** | **string** | Filter by update timestamp. Supports multiple formats: - Range: updated_ats&#x3D;2020-01-01T00:00:00Z,2020-12-31T23:59:59Z - Greater than: updated_ats[gt]&#x3D;2020-01-01T00:00:00Z - Greater than or equal: updated_ats[gte]&#x3D;2020-01-01T00:00:00Z - Less than: updated_ats[lt]&#x3D;2020-12-31T23:59:59Z - Less than or equal: updated_ats[lte]&#x3D;2020-12-31T23:59:59Z  | 
 **fields** | **string** | Fields to include in the response. Use dot notation for nested fields. Example: fields[apps]&#x3D;name,guid,state or fields[space]&#x3D;name,guid  | 

### Return type

[**V3AppsGet200Response**](V3AppsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsApplyManifestPost

> V3AppsGuidActionsApplyManifestPost202Response V3AppsGuidActionsApplyManifestPost(ctx, guid).Body(body).Execute()

Apply manifest to an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidActionsApplyManifestPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidActionsApplyManifestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsApplyManifestPost`: V3AppsGuidActionsApplyManifestPost202Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidActionsApplyManifestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsApplyManifestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**V3AppsGuidActionsApplyManifestPost202Response**](V3AppsGuidActionsApplyManifestPost202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsClearBuildpackCachePost

> V3AppsGuidActionsClearBuildpackCachePost(ctx, guid).Execute()

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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.V3AppsGuidActionsClearBuildpackCachePost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidActionsClearBuildpackCachePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsClearBuildpackCachePostRequest struct via the builder pattern


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


## V3AppsGuidActionsRestartPost

> App V3AppsGuidActionsRestartPost(ctx, guid).Execute()

Restart an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidActionsRestartPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidActionsRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsRestartPost`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidActionsRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsStartPost

> App V3AppsGuidActionsStartPost(ctx, guid).Execute()

Start an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidActionsStartPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidActionsStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsStartPost`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidActionsStartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsStopPost

> App V3AppsGuidActionsStopPost(ctx, guid).Execute()

Stop an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidActionsStopPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidActionsStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsStopPost`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidActionsStopPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidBuildsGet

> V3AppsGuidBuildsGet200Response V3AppsGuidBuildsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Execute()

List builds for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the app
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field by which to order results (optional)
	states := "states_example" // string | Filter by build states (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidBuildsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidBuildsGet`: V3AppsGuidBuildsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidBuildsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field by which to order results | 
 **states** | **string** | Filter by build states (comma-separated) | 

### Return type

[**V3AppsGuidBuildsGet200Response**](V3AppsGuidBuildsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidDelete

> V3AppsGuidDelete(ctx, guid).Execute()

Delete an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.V3AppsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidDeleteRequest struct via the builder pattern


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


## V3AppsGuidEnvironmentVariablesGet

> V3AppsGuidEnvironmentVariablesGet200Response V3AppsGuidEnvironmentVariablesGet(ctx, guid).Execute()

Get app environment variables



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidEnvironmentVariablesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidEnvironmentVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesGet`: V3AppsGuidEnvironmentVariablesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidEnvironmentVariablesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvironmentVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidEnvironmentVariablesGet200Response**](V3AppsGuidEnvironmentVariablesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidEnvironmentVariablesPatch

> V3AppsGuidEnvironmentVariablesPatch200Response V3AppsGuidEnvironmentVariablesPatch(ctx, guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()

Update app environment variables



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	v3AppsGuidEnvironmentVariablesPatchRequest := *openapiclient.NewV3AppsGuidEnvironmentVariablesPatchRequest() // V3AppsGuidEnvironmentVariablesPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidEnvironmentVariablesPatch(context.Background(), guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidEnvironmentVariablesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesPatch`: V3AppsGuidEnvironmentVariablesPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidEnvironmentVariablesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvironmentVariablesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidEnvironmentVariablesPatchRequest** | [**V3AppsGuidEnvironmentVariablesPatchRequest**](V3AppsGuidEnvironmentVariablesPatchRequest.md) |  | 

### Return type

[**V3AppsGuidEnvironmentVariablesPatch200Response**](V3AppsGuidEnvironmentVariablesPatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesGet

> V3AppsGuidFeaturesGet200Response V3AppsGuidFeaturesGet(ctx, guid).Execute()

List app features



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidFeaturesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesGet`: V3AppsGuidFeaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidFeaturesGet200Response**](V3AppsGuidFeaturesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNameGet

> AppFeature V3AppsGuidFeaturesNameGet(ctx, guid, name).Execute()

Get an app feature



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	name := "name_example" // string | The feature name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidFeaturesNameGet(context.Background(), guid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidFeaturesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesNameGet`: AppFeature
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidFeaturesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppFeature**](AppFeature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNamePatch

> AppFeature V3AppsGuidFeaturesNamePatch(ctx, guid, name).AppFeatureUpdate(appFeatureUpdate).Execute()

Update an app feature



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	name := "name_example" // string | The feature name
	appFeatureUpdate := *openapiclient.NewAppFeatureUpdate(true) // AppFeatureUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidFeaturesNamePatch(context.Background(), guid, name).AppFeatureUpdate(appFeatureUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidFeaturesNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidFeaturesNamePatch`: AppFeature
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidFeaturesNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**name** | **string** | The feature name | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appFeatureUpdate** | [**AppFeatureUpdate**](AppFeatureUpdate.md) |  | 

### Return type

[**AppFeature**](AppFeature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidGet

> App V3AppsGuidGet(ctx, guid).Include(include).Execute()

Get an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	include := "include_example" // string | Include related resources (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidGet`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Include related resources | 

### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidManifestDiffPost

> V3AppsGuidManifestDiffPost200Response V3AppsGuidManifestDiffPost(ctx, guid).Body(body).Execute()

Generate manifest diff (experimental)



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidManifestDiffPost(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidManifestDiffPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidManifestDiffPost`: V3AppsGuidManifestDiffPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidManifestDiffPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidManifestDiffPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**V3AppsGuidManifestDiffPost200Response**](V3AppsGuidManifestDiffPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-yaml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidManifestGet

> string V3AppsGuidManifestGet(ctx, guid).Execute()

Get app manifest



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidManifestGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidManifestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidManifestGet`: string
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidManifestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidManifestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidPackagesGet

> V3AppsGuidPackagesGet200Response V3AppsGuidPackagesGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Types(types).Execute()

List packages for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Field to sort results by (optional)
	states := "states_example" // string | Filter by package states (comma-separated) (optional)
	types := "types_example" // string | Filter by package types (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidPackagesGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).States(states).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidPackagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPackagesGet`: V3AppsGuidPackagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidPackagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidPackagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Field to sort results by | 
 **states** | **string** | Filter by package states (comma-separated) | 
 **types** | **string** | Filter by package types (comma-separated) | 

### Return type

[**V3AppsGuidPackagesGet200Response**](V3AppsGuidPackagesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidPatch

> App V3AppsGuidPatch(ctx, guid).V3AppsGuidPatchRequest(v3AppsGuidPatchRequest).Execute()

Update an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	v3AppsGuidPatchRequest := *openapiclient.NewV3AppsGuidPatchRequest() // V3AppsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidPatch(context.Background(), guid).V3AppsGuidPatchRequest(v3AppsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPatch`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidPatchRequest** | [**V3AppsGuidPatchRequest**](V3AppsGuidPatchRequest.md) |  | 

### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidPermissionsGet

> V3AppsGuidPermissionsGet200Response V3AppsGuidPermissionsGet(ctx, guid).Execute()

Get app permissions



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidPermissionsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPermissionsGet`: V3AppsGuidPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidPermissionsGet200Response**](V3AppsGuidPermissionsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidProcessesGet

> V3AppsGuidProcessesGet200Response V3AppsGuidProcessesGet(ctx, guid).Page(page).PerPage(perPage).Types(types).Execute()

List processes for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	page := int32(56) // int32 | Page number for pagination (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	types := "types_example" // string | Filter by process types (comma-separated) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidProcessesGet(context.Background(), guid).Page(page).PerPage(perPage).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidProcessesGet`: V3AppsGuidProcessesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidProcessesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number for pagination | 
 **perPage** | **int32** | Number of results per page | 
 **types** | **string** | Filter by process types (comma-separated) | 

### Return type

[**V3AppsGuidProcessesGet200Response**](V3AppsGuidProcessesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidProcessesTypeGet

> Process V3AppsGuidProcessesTypeGet(ctx, guid, type_).Execute()

Get a process by type for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	type_ := "web" // string | The process type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidProcessesTypeGet(context.Background(), guid, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidProcessesTypeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidProcessesTypeGet`: Process
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidProcessesTypeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 
**type_** | **string** | The process type | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidProcessesTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Process**](Process.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRelationshipsCurrentDropletGet

> V3AppsGuidRelationshipsCurrentDropletGet200Response V3AppsGuidRelationshipsCurrentDropletGet(ctx, guid).Execute()

Get current droplet relationship



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidRelationshipsCurrentDropletGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidRelationshipsCurrentDropletGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRelationshipsCurrentDropletGet`: V3AppsGuidRelationshipsCurrentDropletGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidRelationshipsCurrentDropletGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidRelationshipsCurrentDropletGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidRelationshipsCurrentDropletGet200Response**](V3AppsGuidRelationshipsCurrentDropletGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRelationshipsCurrentDropletPatch

> V3AppsGuidRelationshipsCurrentDropletPatch200Response V3AppsGuidRelationshipsCurrentDropletPatch(ctx, guid).V3AppsGuidRelationshipsCurrentDropletPatchRequest(v3AppsGuidRelationshipsCurrentDropletPatchRequest).Execute()

Update current droplet relationship



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID
	v3AppsGuidRelationshipsCurrentDropletPatchRequest := *openapiclient.NewV3AppsGuidRelationshipsCurrentDropletPatchRequest() // V3AppsGuidRelationshipsCurrentDropletPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidRelationshipsCurrentDropletPatch(context.Background(), guid).V3AppsGuidRelationshipsCurrentDropletPatchRequest(v3AppsGuidRelationshipsCurrentDropletPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidRelationshipsCurrentDropletPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRelationshipsCurrentDropletPatch`: V3AppsGuidRelationshipsCurrentDropletPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidRelationshipsCurrentDropletPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidRelationshipsCurrentDropletPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidRelationshipsCurrentDropletPatchRequest** | [**V3AppsGuidRelationshipsCurrentDropletPatchRequest**](V3AppsGuidRelationshipsCurrentDropletPatchRequest.md) |  | 

### Return type

[**V3AppsGuidRelationshipsCurrentDropletPatch200Response**](V3AppsGuidRelationshipsCurrentDropletPatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidSshEnabledGet

> V3AppsGuidSshEnabledGet200Response V3AppsGuidSshEnabledGet(ctx, guid).Execute()

Get SSH enabled



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsGuidSshEnabledGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsGuidSshEnabledGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSshEnabledGet`: V3AppsGuidSshEnabledGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsGuidSshEnabledGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSshEnabledGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidSshEnabledGet200Response**](V3AppsGuidSshEnabledGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsPost

> App V3AppsPost(ctx).V3AppsPostRequest(v3AppsPostRequest).Execute()

Create an app



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
	v3AppsPostRequest := *openapiclient.NewV3AppsPostRequest("Name_example", *openapiclient.NewV3AppsPostRequestRelationships(*openapiclient.NewV3AppsPostRequestRelationshipsSpace(*openapiclient.NewV3AppsPostRequestRelationshipsSpaceData("Guid_example")))) // V3AppsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.V3AppsPost(context.Background()).V3AppsPostRequest(v3AppsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.V3AppsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsPost`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.V3AppsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3AppsPostRequest** | [**V3AppsPostRequest**](V3AppsPostRequest.md) |  | 

### Return type

[**App**](App.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

