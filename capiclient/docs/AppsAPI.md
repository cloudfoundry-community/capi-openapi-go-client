# \AppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearAppBuildpackCache**](AppsAPI.md#ClearAppBuildpackCache) | **Post** /v3/apps/{guid}/actions/clear_buildpack_cache | Clear buildpack cache
[**CreateApp**](AppsAPI.md#CreateApp) | **Post** /v3/apps | Create an app
[**DeleteApp**](AppsAPI.md#DeleteApp) | **Delete** /v3/apps/{guid} | Delete an app
[**GetApp**](AppsAPI.md#GetApp) | **Get** /v3/apps/{guid} | Get an app
[**GetAppEnvironmentVariables**](AppsAPI.md#GetAppEnvironmentVariables) | **Get** /v3/apps/{guid}/environment_variables | Get app environment variables
[**GetAppPermissions**](AppsAPI.md#GetAppPermissions) | **Get** /v3/apps/{guid}/permissions | Get app permissions
[**GetAppSshEnabled**](AppsAPI.md#GetAppSshEnabled) | **Get** /v3/apps/{guid}/ssh_enabled | Get SSH enabled
[**ListApps**](AppsAPI.md#ListApps) | **Get** /v3/apps | List apps
[**RestartApp**](AppsAPI.md#RestartApp) | **Post** /v3/apps/{guid}/actions/restart | Restart an app
[**StartApp**](AppsAPI.md#StartApp) | **Post** /v3/apps/{guid}/actions/start | Start an app
[**StopApp**](AppsAPI.md#StopApp) | **Post** /v3/apps/{guid}/actions/stop | Stop an app
[**UpdateApp**](AppsAPI.md#UpdateApp) | **Patch** /v3/apps/{guid} | Update an app
[**UpdateAppEnvironmentVariables**](AppsAPI.md#UpdateAppEnvironmentVariables) | **Patch** /v3/apps/{guid}/environment_variables | Update app environment variables
[**V3AppsGuidActionsApplyManifestPost**](AppsAPI.md#V3AppsGuidActionsApplyManifestPost) | **Post** /v3/apps/{guid}/actions/apply_manifest | Apply manifest to an app
[**V3AppsGuidBuildsGet**](AppsAPI.md#V3AppsGuidBuildsGet) | **Get** /v3/apps/{guid}/builds | List builds for an app
[**V3AppsGuidFeaturesGet**](AppsAPI.md#V3AppsGuidFeaturesGet) | **Get** /v3/apps/{guid}/features | List app features
[**V3AppsGuidFeaturesNameGet**](AppsAPI.md#V3AppsGuidFeaturesNameGet) | **Get** /v3/apps/{guid}/features/{name} | Get an app feature
[**V3AppsGuidFeaturesNamePatch**](AppsAPI.md#V3AppsGuidFeaturesNamePatch) | **Patch** /v3/apps/{guid}/features/{name} | Update an app feature
[**V3AppsGuidManifestDiffPost**](AppsAPI.md#V3AppsGuidManifestDiffPost) | **Post** /v3/apps/{guid}/manifest_diff | Generate manifest diff (experimental)
[**V3AppsGuidManifestGet**](AppsAPI.md#V3AppsGuidManifestGet) | **Get** /v3/apps/{guid}/manifest | Get app manifest
[**V3AppsGuidPackagesGet**](AppsAPI.md#V3AppsGuidPackagesGet) | **Get** /v3/apps/{guid}/packages | List packages for an app
[**V3AppsGuidProcessesGet**](AppsAPI.md#V3AppsGuidProcessesGet) | **Get** /v3/apps/{guid}/processes | List processes for an app
[**V3AppsGuidProcessesTypeGet**](AppsAPI.md#V3AppsGuidProcessesTypeGet) | **Get** /v3/apps/{guid}/processes/{type} | Get a process by type for an app
[**V3AppsGuidRelationshipsCurrentDropletGet**](AppsAPI.md#V3AppsGuidRelationshipsCurrentDropletGet) | **Get** /v3/apps/{guid}/relationships/current_droplet | Get current droplet relationship
[**V3AppsGuidRelationshipsCurrentDropletPatch**](AppsAPI.md#V3AppsGuidRelationshipsCurrentDropletPatch) | **Patch** /v3/apps/{guid}/relationships/current_droplet | Update current droplet relationship



## ClearAppBuildpackCache

> ClearAppBuildpackCache(ctx, guid).Execute()

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
	r, err := apiClient.AppsAPI.ClearAppBuildpackCache(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ClearAppBuildpackCache``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClearAppBuildpackCacheRequest struct via the builder pattern


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


## CreateApp

> App CreateApp(ctx).CreateAppRequest(createAppRequest).Execute()

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
	createAppRequest := *openapiclient.NewCreateAppRequest("Name_example", *openapiclient.NewCreateAppRequestRelationships(*openapiclient.NewCreateAppRequestRelationshipsSpace(*openapiclient.NewCreateAppRequestRelationshipsSpaceData("Guid_example")))) // CreateAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.CreateApp(context.Background()).CreateAppRequest(createAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.CreateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAppRequest** | [**CreateAppRequest**](CreateAppRequest.md) |  | 

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


## DeleteApp

> DeleteApp(ctx, guid).Execute()

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
	r, err := apiClient.AppsAPI.DeleteApp(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


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


## GetApp

> App GetApp(ctx, guid).Include(include).Execute()

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
	resp, r, err := apiClient.AppsAPI.GetApp(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


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


## GetAppEnvironmentVariables

> GetAppEnvironmentVariables200Response GetAppEnvironmentVariables(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.GetAppEnvironmentVariables(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppEnvironmentVariables`: GetAppEnvironmentVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppEnvironmentVariables200Response**](GetAppEnvironmentVariables200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppPermissions

> GetAppPermissions200Response GetAppPermissions(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.GetAppPermissions(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppPermissions`: GetAppPermissions200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppPermissions200Response**](GetAppPermissions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppSshEnabled

> GetAppSshEnabled200Response GetAppSshEnabled(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.GetAppSshEnabled(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetAppSshEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppSshEnabled`: GetAppSshEnabled200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetAppSshEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSshEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAppSshEnabled200Response**](GetAppSshEnabled200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> ListApps200Response ListApps(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).States(states).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()

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
	resp, r, err := apiClient.AppsAPI.ListApps(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).States(states).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.ListApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApps`: ListApps200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


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

[**ListApps200Response**](ListApps200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartApp

> App RestartApp(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.RestartApp(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.RestartApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.RestartApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartAppRequest struct via the builder pattern


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


## StartApp

> App StartApp(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.StartApp(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.StartApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.StartApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartAppRequest struct via the builder pattern


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


## StopApp

> App StopApp(ctx, guid).Execute()

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
	resp, r, err := apiClient.AppsAPI.StopApp(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.StopApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.StopApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAppRequest struct via the builder pattern


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


## UpdateApp

> App UpdateApp(ctx, guid).UpdateAppRequest(updateAppRequest).Execute()

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
	updateAppRequest := *openapiclient.NewUpdateAppRequest() // UpdateAppRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateApp(context.Background(), guid).UpdateAppRequest(updateAppRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApp`: App
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppRequest** | [**UpdateAppRequest**](UpdateAppRequest.md) |  | 

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


## UpdateAppEnvironmentVariables

> UpdateAppEnvironmentVariables200Response UpdateAppEnvironmentVariables(ctx, guid).UpdateAppEnvironmentVariablesRequest(updateAppEnvironmentVariablesRequest).Execute()

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
	updateAppEnvironmentVariablesRequest := *openapiclient.NewUpdateAppEnvironmentVariablesRequest() // UpdateAppEnvironmentVariablesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.UpdateAppEnvironmentVariables(context.Background(), guid).UpdateAppEnvironmentVariablesRequest(updateAppEnvironmentVariablesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.UpdateAppEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppEnvironmentVariables`: UpdateAppEnvironmentVariables200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.UpdateAppEnvironmentVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAppEnvironmentVariablesRequest** | [**UpdateAppEnvironmentVariablesRequest**](UpdateAppEnvironmentVariablesRequest.md) |  | 

### Return type

[**UpdateAppEnvironmentVariables200Response**](UpdateAppEnvironmentVariables200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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

