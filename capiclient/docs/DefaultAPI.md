# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignDefaultIsolationSegment**](DefaultAPI.md#AssignDefaultIsolationSegment) | **Patch** /v3/organizations/{guid}/relationships/default_isolation_segment | Assign Default Isolation Segment
[**CreateOrganization**](DefaultAPI.md#CreateOrganization) | **Post** /v3/organizations | Create an Organization
[**DeleteOrganization**](DefaultAPI.md#DeleteOrganization) | **Delete** /v3/organizations/{guid} | Delete an Organization
[**GetDefaultDomain**](DefaultAPI.md#GetDefaultDomain) | **Get** /v3/organizations/{guid}/domains/default | Get Default Domain
[**GetDefaultIsolationSegment**](DefaultAPI.md#GetDefaultIsolationSegment) | **Get** /v3/organizations/{guid}/relationships/default_isolation_segment | Get Default Isolation Segment
[**GetOrganization**](DefaultAPI.md#GetOrganization) | **Get** /v3/organizations/{guid} | Get an Organization
[**GetUsageSummary**](DefaultAPI.md#GetUsageSummary) | **Get** /v3/organizations/{guid}/usage_summary | Get Usage Summary
[**ListOrganizations**](DefaultAPI.md#ListOrganizations) | **Get** /v3/organizations | List Organizations
[**ListUsersForOrganization**](DefaultAPI.md#ListUsersForOrganization) | **Get** /v3/organizations/{guid}/users | List Users for an Organization
[**RootGet**](DefaultAPI.md#RootGet) | **Get** / | Global API Root
[**UpdateOrganization**](DefaultAPI.md#UpdateOrganization) | **Patch** /v3/organizations/{guid} | Update an Organization
[**V3AppsGuidDropletsGet**](DefaultAPI.md#V3AppsGuidDropletsGet) | **Get** /v3/apps/{guid}/droplets | List droplets for an app
[**V3AppsGuidRevisionsDeployedGet**](DefaultAPI.md#V3AppsGuidRevisionsDeployedGet) | **Get** /v3/apps/{guid}/revisions/deployed | List deployed revisions for an app
[**V3AppsGuidRevisionsGet**](DefaultAPI.md#V3AppsGuidRevisionsGet) | **Get** /v3/apps/{guid}/revisions | List revisions for an app
[**V3AppsGuidTasksPost**](DefaultAPI.md#V3AppsGuidTasksPost) | **Post** /v3/apps/{guid}/tasks | Create a task
[**V3DropletsGet**](DefaultAPI.md#V3DropletsGet) | **Get** /v3/droplets | List droplets
[**V3DropletsGuidDelete**](DefaultAPI.md#V3DropletsGuidDelete) | **Delete** /v3/droplets/{guid} | Delete a droplet
[**V3DropletsGuidGet**](DefaultAPI.md#V3DropletsGuidGet) | **Get** /v3/droplets/{guid} | Get a droplet
[**V3DropletsGuidPatch**](DefaultAPI.md#V3DropletsGuidPatch) | **Patch** /v3/droplets/{guid} | Update a droplet
[**V3DropletsPost**](DefaultAPI.md#V3DropletsPost) | **Post** /v3/droplets | Create or copy a droplet
[**V3Get**](DefaultAPI.md#V3Get) | **Get** /v3 | V3 API Root
[**V3InfoGet**](DefaultAPI.md#V3InfoGet) | **Get** /v3/info | Get platform info
[**V3InfoUsageSummaryGet**](DefaultAPI.md#V3InfoUsageSummaryGet) | **Get** /v3/info/usage_summary | Get platform usage summary
[**V3IsolationSegmentsGet**](DefaultAPI.md#V3IsolationSegmentsGet) | **Get** /v3/isolation_segments | List isolation segments
[**V3IsolationSegmentsGuidDelete**](DefaultAPI.md#V3IsolationSegmentsGuidDelete) | **Delete** /v3/isolation_segments/{guid} | Delete an isolation segment
[**V3IsolationSegmentsGuidGet**](DefaultAPI.md#V3IsolationSegmentsGuidGet) | **Get** /v3/isolation_segments/{guid} | Get an isolation segment
[**V3IsolationSegmentsGuidPatch**](DefaultAPI.md#V3IsolationSegmentsGuidPatch) | **Patch** /v3/isolation_segments/{guid} | Update an isolation segment
[**V3IsolationSegmentsGuidRelationshipsOrganizationsGet**](DefaultAPI.md#V3IsolationSegmentsGuidRelationshipsOrganizationsGet) | **Get** /v3/isolation_segments/{guid}/relationships/organizations | List organizations relationship
[**V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete**](DefaultAPI.md#V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete) | **Delete** /v3/isolation_segments/{guid}/relationships/organizations/{org_guid} | Revoke entitlement to isolation segment for an organization
[**V3IsolationSegmentsGuidRelationshipsOrganizationsPost**](DefaultAPI.md#V3IsolationSegmentsGuidRelationshipsOrganizationsPost) | **Post** /v3/isolation_segments/{guid}/relationships/organizations | Entitle organizations for an isolation segment
[**V3IsolationSegmentsGuidRelationshipsSpacesGet**](DefaultAPI.md#V3IsolationSegmentsGuidRelationshipsSpacesGet) | **Get** /v3/isolation_segments/{guid}/relationships/spaces | List spaces relationship
[**V3IsolationSegmentsPost**](DefaultAPI.md#V3IsolationSegmentsPost) | **Post** /v3/isolation_segments | Create an isolation segment
[**V3PackagesGuidDropletsGet**](DefaultAPI.md#V3PackagesGuidDropletsGet) | **Get** /v3/packages/{guid}/droplets | List droplets for a package
[**V3ResourceMatchesPost**](DefaultAPI.md#V3ResourceMatchesPost) | **Post** /v3/resource_matches | Create a resource match
[**V3RevisionsGuidEnvironmentVariablesGet**](DefaultAPI.md#V3RevisionsGuidEnvironmentVariablesGet) | **Get** /v3/revisions/{guid}/environment_variables | Get environment variables for a revision
[**V3RevisionsGuidGet**](DefaultAPI.md#V3RevisionsGuidGet) | **Get** /v3/revisions/{guid} | Get a revision
[**V3RevisionsGuidPatch**](DefaultAPI.md#V3RevisionsGuidPatch) | **Patch** /v3/revisions/{guid} | Update a revision
[**V3ServiceBrokersGet**](DefaultAPI.md#V3ServiceBrokersGet) | **Get** /v3/service_brokers | List service brokers
[**V3ServiceBrokersGuidCatalogPost**](DefaultAPI.md#V3ServiceBrokersGuidCatalogPost) | **Post** /v3/service_brokers/{guid}/catalog | Sync service broker catalog
[**V3ServiceBrokersGuidDelete**](DefaultAPI.md#V3ServiceBrokersGuidDelete) | **Delete** /v3/service_brokers/{guid} | Delete a service broker
[**V3ServiceBrokersGuidGet**](DefaultAPI.md#V3ServiceBrokersGuidGet) | **Get** /v3/service_brokers/{guid} | Get a service broker
[**V3ServiceBrokersGuidJobsSynchronizationGet**](DefaultAPI.md#V3ServiceBrokersGuidJobsSynchronizationGet) | **Get** /v3/service_brokers/{guid}/jobs/synchronization | Get broker synchronization job status
[**V3ServiceBrokersGuidPatch**](DefaultAPI.md#V3ServiceBrokersGuidPatch) | **Patch** /v3/service_brokers/{guid} | Update a service broker
[**V3ServiceBrokersPost**](DefaultAPI.md#V3ServiceBrokersPost) | **Post** /v3/service_brokers | Create a service broker
[**V3ServiceCredentialBindingsGuidDelete**](DefaultAPI.md#V3ServiceCredentialBindingsGuidDelete) | **Delete** /v3/service_credential_bindings/{guid} | Delete a service credential binding
[**V3ServiceCredentialBindingsGuidDetailsGet**](DefaultAPI.md#V3ServiceCredentialBindingsGuidDetailsGet) | **Get** /v3/service_credential_bindings/{guid}/details | Get a service credential binding details
[**V3ServiceCredentialBindingsGuidGet**](DefaultAPI.md#V3ServiceCredentialBindingsGuidGet) | **Get** /v3/service_credential_bindings/{guid} | Get a service credential binding
[**V3ServiceCredentialBindingsGuidParametersGet**](DefaultAPI.md#V3ServiceCredentialBindingsGuidParametersGet) | **Get** /v3/service_credential_bindings/{guid}/parameters | Get parameters for a service credential binding
[**V3ServiceCredentialBindingsGuidPatch**](DefaultAPI.md#V3ServiceCredentialBindingsGuidPatch) | **Patch** /v3/service_credential_bindings/{guid} | Update a service credential binding
[**V3ServiceCredentialBindingsPost**](DefaultAPI.md#V3ServiceCredentialBindingsPost) | **Post** /v3/service_credential_bindings | Create a service credential binding
[**V3ServiceInstancesGet**](DefaultAPI.md#V3ServiceInstancesGet) | **Get** /v3/service_instances | Retrieve service instances
[**V3ServiceInstancesGuidCredentialsGet**](DefaultAPI.md#V3ServiceInstancesGuidCredentialsGet) | **Get** /v3/service_instances/{guid}/credentials | Get credentials for a user-provided service instance
[**V3ServiceInstancesGuidDelete**](DefaultAPI.md#V3ServiceInstancesGuidDelete) | **Delete** /v3/service_instances/{guid} | Delete a service instance
[**V3ServiceInstancesGuidPatch**](DefaultAPI.md#V3ServiceInstancesGuidPatch) | **Patch** /v3/service_instances/{guid} | Update a service instance
[**V3ServiceInstancesGuidRelationshipsSharedSpacesPost**](DefaultAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesPost) | **Post** /v3/service_instances/{guid}/relationships/shared_spaces | Share a service instance to other spaces
[**V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete**](DefaultAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete) | **Delete** /v3/service_instances/{guid}/relationships/shared_spaces/{space_guid} | Unshare a service instance from another space
[**V3ServiceOfferingsGet**](DefaultAPI.md#V3ServiceOfferingsGet) | **Get** /v3/service_offerings | List service offerings
[**V3ServiceOfferingsGuidDelete**](DefaultAPI.md#V3ServiceOfferingsGuidDelete) | **Delete** /v3/service_offerings/{guid} | Delete a service offering
[**V3ServiceOfferingsGuidGet**](DefaultAPI.md#V3ServiceOfferingsGuidGet) | **Get** /v3/service_offerings/{guid} | Get a service offering
[**V3ServiceOfferingsGuidPatch**](DefaultAPI.md#V3ServiceOfferingsGuidPatch) | **Patch** /v3/service_offerings/{guid} | Update a service offering
[**V3ServiceOfferingsPost**](DefaultAPI.md#V3ServiceOfferingsPost) | **Post** /v3/service_offerings | Create a service offering
[**V3ServicePlansGet**](DefaultAPI.md#V3ServicePlansGet) | **Get** /v3/service_plans | List service plans
[**V3ServicePlansGuidDelete**](DefaultAPI.md#V3ServicePlansGuidDelete) | **Delete** /v3/service_plans/{guid} | Delete a service plan
[**V3ServicePlansGuidGet**](DefaultAPI.md#V3ServicePlansGuidGet) | **Get** /v3/service_plans/{guid} | Get a service plan
[**V3ServicePlansGuidPatch**](DefaultAPI.md#V3ServicePlansGuidPatch) | **Patch** /v3/service_plans/{guid} | Update a service plan
[**V3ServicePlansPost**](DefaultAPI.md#V3ServicePlansPost) | **Post** /v3/service_plans | Create a service plan
[**V3ServiceRouteBindingsGet**](DefaultAPI.md#V3ServiceRouteBindingsGet) | **Get** /v3/service_route_bindings | List service route bindings
[**V3ServiceRouteBindingsGuidDelete**](DefaultAPI.md#V3ServiceRouteBindingsGuidDelete) | **Delete** /v3/service_route_bindings/{guid} | Delete a service route binding
[**V3ServiceRouteBindingsGuidGet**](DefaultAPI.md#V3ServiceRouteBindingsGuidGet) | **Get** /v3/service_route_bindings/{guid} | Get a service route binding
[**V3ServiceRouteBindingsGuidParametersGet**](DefaultAPI.md#V3ServiceRouteBindingsGuidParametersGet) | **Get** /v3/service_route_bindings/{guid}/parameters | Get parameters for a route binding
[**V3ServiceRouteBindingsGuidPatch**](DefaultAPI.md#V3ServiceRouteBindingsGuidPatch) | **Patch** /v3/service_route_bindings/{guid} | Update a service route binding
[**V3ServiceRouteBindingsPost**](DefaultAPI.md#V3ServiceRouteBindingsPost) | **Post** /v3/service_route_bindings | Create a service route binding
[**V3ServiceUsageEventsGet**](DefaultAPI.md#V3ServiceUsageEventsGet) | **Get** /v3/service_usage_events | List service usage events
[**V3ServiceUsageEventsGuidGet**](DefaultAPI.md#V3ServiceUsageEventsGuidGet) | **Get** /v3/service_usage_events/{guid} | Get a service usage event
[**V3ServiceUsageEventsPost**](DefaultAPI.md#V3ServiceUsageEventsPost) | **Post** /v3/service_usage_events | Purge and seed service usage events
[**V3SpacesGet**](DefaultAPI.md#V3SpacesGet) | **Get** /v3/spaces | List spaces
[**V3SpacesGuidDelete**](DefaultAPI.md#V3SpacesGuidDelete) | **Delete** /v3/spaces/{guid} | Delete a space
[**V3SpacesGuidGet**](DefaultAPI.md#V3SpacesGuidGet) | **Get** /v3/spaces/{guid} | Get a space
[**V3SpacesGuidPatch**](DefaultAPI.md#V3SpacesGuidPatch) | **Patch** /v3/spaces/{guid} | Update a space
[**V3SpacesGuidRelationshipsIsolationSegmentGet**](DefaultAPI.md#V3SpacesGuidRelationshipsIsolationSegmentGet) | **Get** /v3/spaces/{guid}/relationships/isolation_segment | Get assigned isolation segment
[**V3SpacesGuidRelationshipsIsolationSegmentPatch**](DefaultAPI.md#V3SpacesGuidRelationshipsIsolationSegmentPatch) | **Patch** /v3/spaces/{guid}/relationships/isolation_segment | Manage isolation segment
[**V3SpacesGuidUsersGet**](DefaultAPI.md#V3SpacesGuidUsersGet) | **Get** /v3/spaces/{guid}/users | List users for a space
[**V3SpacesPost**](DefaultAPI.md#V3SpacesPost) | **Post** /v3/spaces | Create a space
[**V3StacksGet**](DefaultAPI.md#V3StacksGet) | **Get** /v3/stacks | List all stacks
[**V3StacksGuidAppsGet**](DefaultAPI.md#V3StacksGuidAppsGet) | **Get** /v3/stacks/{guid}/apps | List apps on a stack
[**V3StacksGuidDelete**](DefaultAPI.md#V3StacksGuidDelete) | **Delete** /v3/stacks/{guid} | Delete a stack
[**V3StacksGuidGet**](DefaultAPI.md#V3StacksGuidGet) | **Get** /v3/stacks/{guid} | Get a stack by GUID
[**V3StacksGuidPatch**](DefaultAPI.md#V3StacksGuidPatch) | **Patch** /v3/stacks/{guid} | Update a stack
[**V3StacksPost**](DefaultAPI.md#V3StacksPost) | **Post** /v3/stacks | Create a stack
[**V3TasksGet**](DefaultAPI.md#V3TasksGet) | **Get** /v3/tasks | List all tasks
[**V3TasksGuidGet**](DefaultAPI.md#V3TasksGuidGet) | **Get** /v3/tasks/{guid} | Get a task
[**V3TasksGuidPatch**](DefaultAPI.md#V3TasksGuidPatch) | **Patch** /v3/tasks/{guid} | Update a task
[**V3TasksGuidPost**](DefaultAPI.md#V3TasksGuidPost) | **Post** /v3/tasks/{guid} | Cancel a task
[**V3UsersGet**](DefaultAPI.md#V3UsersGet) | **Get** /v3/users | List users
[**V3UsersGuidDelete**](DefaultAPI.md#V3UsersGuidDelete) | **Delete** /v3/users/{guid} | Delete a user
[**V3UsersGuidGet**](DefaultAPI.md#V3UsersGuidGet) | **Get** /v3/users/{guid} | Get a user
[**V3UsersGuidPatch**](DefaultAPI.md#V3UsersGuidPatch) | **Patch** /v3/users/{guid} | Update a user
[**V3UsersPost**](DefaultAPI.md#V3UsersPost) | **Post** /v3/users | Create a user



## AssignDefaultIsolationSegment

> map[string]interface{} AssignDefaultIsolationSegment(ctx, guid).AssignDefaultIsolationSegmentRequest(assignDefaultIsolationSegmentRequest).Execute()

Assign Default Isolation Segment

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
	guid := "guid_example" // string | 
	assignDefaultIsolationSegmentRequest := *openapiclient.NewAssignDefaultIsolationSegmentRequest() // AssignDefaultIsolationSegmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssignDefaultIsolationSegment(context.Background(), guid).AssignDefaultIsolationSegmentRequest(assignDefaultIsolationSegmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AssignDefaultIsolationSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignDefaultIsolationSegment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AssignDefaultIsolationSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignDefaultIsolationSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignDefaultIsolationSegmentRequest** | [**AssignDefaultIsolationSegmentRequest**](AssignDefaultIsolationSegmentRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> Organization CreateOrganization(ctx).CreateOrganizationRequest(createOrganizationRequest).Execute()

Create an Organization

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
	createOrganizationRequest := *openapiclient.NewCreateOrganizationRequest("Name_example") // CreateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateOrganization(context.Background()).CreateOrganizationRequest(createOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> map[string]interface{} DeleteOrganization(ctx, guid).Execute()

Delete an Organization

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteOrganization(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganization`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultDomain

> map[string]interface{} GetDefaultDomain(ctx, guid).Execute()

Get Default Domain

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDefaultDomain(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDefaultDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultDomain`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDefaultDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultIsolationSegment

> map[string]interface{} GetDefaultIsolationSegment(ctx, guid).Execute()

Get Default Isolation Segment

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetDefaultIsolationSegment(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetDefaultIsolationSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultIsolationSegment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetDefaultIsolationSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultIsolationSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx, guid).Execute()

Get an Organization

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOrganization(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> map[string]interface{} GetUsageSummary(ctx, guid).Execute()

Get Usage Summary

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUsageSummary(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageSummary`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUsageSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> []Organization ListOrganizations(ctx).Names(names).Guids(guids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List Organizations

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
	names := "names_example" // string |  (optional)
	guids := "guids_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListOrganizations(context.Background()).Names(names).Guids(guids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizations`: []Organization
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **string** |  | 
 **guids** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersForOrganization

> []User ListUsersForOrganization(ctx, guid).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List Users for an Organization

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
	guid := "guid_example" // string | 
	guids := "guids_example" // string |  (optional)
	usernames := "usernames_example" // string |  (optional)
	partialUsernames := "partialUsernames_example" // string |  (optional)
	origins := "origins_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListUsersForOrganization(context.Background(), guid).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListUsersForOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsersForOrganization`: []User
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListUsersForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **string** |  | 
 **usernames** | **string** |  | 
 **partialUsernames** | **string** |  | 
 **origins** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RootGet

> Get200Response RootGet(ctx).Execute()

Global API Root



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RootGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RootGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RootGet`: Get200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RootGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRootGetRequest struct via the builder pattern


### Return type

[**Get200Response**](Get200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> Organization UpdateOrganization(ctx, guid).UpdateOrganizationRequest(updateOrganizationRequest).Execute()

Update an Organization

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
	guid := "guid_example" // string | 
	updateOrganizationRequest := *openapiclient.NewUpdateOrganizationRequest() // UpdateOrganizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateOrganization(context.Background(), guid).UpdateOrganizationRequest(updateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationRequest** | [**UpdateOrganizationRequest**](UpdateOrganizationRequest.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidDropletsGet

> V3AppsGuidDropletsGet200Response V3AppsGuidDropletsGet(ctx, guid).Guids(guids).States(states).Current(current).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()

List droplets for an app



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The guid of the app
	guids := []string{"Inner_example"} // []string | Comma-delimited list of droplet guids to filter by (optional)
	states := []string{"Inner_example"} // []string | Comma-delimited list of droplet states to filter by (optional)
	current := true // bool | If true, only include the droplet currently assigned to the app (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidDropletsGet(context.Background(), guid).Guids(guids).States(states).Current(current).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidDropletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidDropletsGet`: V3AppsGuidDropletsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidDropletsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidDropletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **[]string** | Comma-delimited list of droplet guids to filter by | 
 **states** | **[]string** | Comma-delimited list of droplet states to filter by | 
 **current** | **bool** | If true, only include the droplet currently assigned to the app | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 

### Return type

[**V3AppsGuidDropletsGet200Response**](V3AppsGuidDropletsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRevisionsDeployedGet

> RevisionsList V3AppsGuidRevisionsDeployedGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()

List deployed revisions for an app



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
	guid := "guid_example" // string | Unique identifier for the app
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidRevisionsDeployedGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidRevisionsDeployedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRevisionsDeployedGet`: RevisionsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidRevisionsDeployedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidRevisionsDeployedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. | 

### Return type

[**RevisionsList**](RevisionsList.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRevisionsGet

> RevisionsList V3AppsGuidRevisionsGet(ctx, guid).Versions(versions).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List revisions for an app



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
	guid := "guid_example" // string | Unique identifier for the app
	versions := []string{"Inner_example"} // []string | Comma-delimited list of revision versions to filter by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by; supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by; supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidRevisionsGet(context.Background(), guid).Versions(versions).LabelSelector(labelSelector).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidRevisionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRevisionsGet`: RevisionsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidRevisionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidRevisionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versions** | **[]string** | Comma-delimited list of revision versions to filter by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. | 
 **createdAts** | **string** | Timestamp to filter by; supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by; supports filtering with relational operators | 

### Return type

[**RevisionsList**](RevisionsList.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidTasksPost

> Task V3AppsGuidTasksPost(ctx, guid).V3AppsGuidTasksPostRequest(v3AppsGuidTasksPostRequest).Execute()

Create a task



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
	guid := "guid_example" // string | 
	v3AppsGuidTasksPostRequest := *openapiclient.NewV3AppsGuidTasksPostRequest() // V3AppsGuidTasksPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidTasksPost(context.Background(), guid).V3AppsGuidTasksPostRequest(v3AppsGuidTasksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidTasksPost`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidTasksPostRequest** | [**V3AppsGuidTasksPostRequest**](V3AppsGuidTasksPostRequest.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DropletsGet

> V3AppsGuidDropletsGet200Response V3DropletsGet(ctx).Guids(guids).States(states).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List droplets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guids := []string{"Inner_example"} // []string | Comma-delimited list of droplet guids to filter by (optional)
	states := []string{"Inner_example"} // []string | Comma-delimited list of droplet states to filter by (optional)
	appGuids := []string{"Inner_example"} // []string | Comma-delimited list of app guids to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := time.Now() // time.Time | Timestamp to filter by (optional)
	updatedAts := time.Now() // time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3DropletsGet(context.Background()).Guids(guids).States(states).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DropletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DropletsGet`: V3AppsGuidDropletsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3DropletsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DropletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of droplet guids to filter by | 
 **states** | **[]string** | Comma-delimited list of droplet states to filter by | 
 **appGuids** | **[]string** | Comma-delimited list of app guids to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | **time.Time** | Timestamp to filter by | 
 **updatedAts** | **time.Time** | Timestamp to filter by | 

### Return type

[**V3AppsGuidDropletsGet200Response**](V3AppsGuidDropletsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DropletsGuidDelete

> V3DropletsGuidDelete(ctx, guid).Execute()

Delete a droplet



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The guid of the droplet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3DropletsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DropletsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the droplet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DropletsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DropletsGuidGet

> Droplet V3DropletsGuidGet(ctx, guid).Execute()

Get a droplet



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The guid of the droplet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3DropletsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DropletsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DropletsGuidGet`: Droplet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3DropletsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the droplet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DropletsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Droplet**](Droplet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DropletsGuidPatch

> Droplet V3DropletsGuidPatch(ctx, guid).V3DropletsGuidPatchRequest(v3DropletsGuidPatchRequest).Execute()

Update a droplet



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The guid of the droplet
	v3DropletsGuidPatchRequest := *openapiclient.NewV3DropletsGuidPatchRequest() // V3DropletsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3DropletsGuidPatch(context.Background(), guid).V3DropletsGuidPatchRequest(v3DropletsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DropletsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DropletsGuidPatch`: Droplet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3DropletsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the droplet | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DropletsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DropletsGuidPatchRequest** | [**V3DropletsGuidPatchRequest**](V3DropletsGuidPatchRequest.md) |  | 

### Return type

[**Droplet**](Droplet.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DropletsPost

> Droplet V3DropletsPost(ctx).V3DropletsPostRequest(v3DropletsPostRequest).SourceGuid(sourceGuid).Execute()

Create or copy a droplet



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
	v3DropletsPostRequest := *openapiclient.NewV3DropletsPostRequest(*openapiclient.NewV3DropletsPostRequestRelationships()) // V3DropletsPostRequest | 
	sourceGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source guid of the droplet to be copied (for copy operation) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3DropletsPost(context.Background()).V3DropletsPostRequest(v3DropletsPostRequest).SourceGuid(sourceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DropletsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3DropletsPost`: Droplet
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3DropletsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DropletsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3DropletsPostRequest** | [**V3DropletsPostRequest**](V3DropletsPostRequest.md) |  | 
 **sourceGuid** | **string** | Source guid of the droplet to be copied (for copy operation) | 

### Return type

[**Droplet**](Droplet.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3Get

> V3Get200Response V3Get(ctx).Execute()

V3 API Root



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3Get`: V3Get200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetRequest struct via the builder pattern


### Return type

[**V3Get200Response**](V3Get200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3InfoGet

> PlatformInfo V3InfoGet(ctx).Execute()

Get platform info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3InfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3InfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3InfoGet`: PlatformInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3InfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3InfoGetRequest struct via the builder pattern


### Return type

[**PlatformInfo**](PlatformInfo.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3InfoUsageSummaryGet

> UsageSummary V3InfoUsageSummaryGet(ctx).Execute()

Get platform usage summary



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3InfoUsageSummaryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3InfoUsageSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3InfoUsageSummaryGet`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3InfoUsageSummaryGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3InfoUsageSummaryGetRequest struct via the builder pattern


### Return type

[**UsageSummary**](UsageSummary.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGet

> V3IsolationSegmentsGet200Response V3IsolationSegmentsGet(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List isolation segments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guids := []string{"Inner_example"} // []string | Comma-delimited list of isolation segment guids to filter by (optional)
	names := []string{"Inner_example"} // []string | Comma-delimited list of isolation segment names to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by; defaults to ascending. Prepend with - to sort descending (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := time.Now() // time.Time | Timestamp to filter by (optional)
	updatedAts := time.Now() // time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGet(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGet`: V3IsolationSegmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of isolation segment guids to filter by | 
 **names** | **[]string** | Comma-delimited list of isolation segment names to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by; defaults to ascending. Prepend with - to sort descending | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | **time.Time** | Timestamp to filter by | 
 **updatedAts** | **time.Time** | Timestamp to filter by | 

### Return type

[**V3IsolationSegmentsGet200Response**](V3IsolationSegmentsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidDelete

> V3IsolationSegmentsGuidDelete(ctx, guid).Execute()

Delete an isolation segment



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
	guid := "guid_example" // string | The guid of the isolation segment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidGet

> IsolationSegment V3IsolationSegmentsGuidGet(ctx, guid).Execute()

Get an isolation segment



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
	guid := "guid_example" // string | The guid of the isolation segment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGuidGet`: IsolationSegment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsolationSegment**](IsolationSegment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidPatch

> IsolationSegment V3IsolationSegmentsGuidPatch(ctx, guid).IsolationSegment(isolationSegment).Execute()

Update an isolation segment



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
	guid := "guid_example" // string | The guid of the isolation segment
	isolationSegment := *openapiclient.NewIsolationSegment() // IsolationSegment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidPatch(context.Background(), guid).IsolationSegment(isolationSegment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGuidPatch`: IsolationSegment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isolationSegment** | [**IsolationSegment**](IsolationSegment.md) |  | 

### Return type

[**IsolationSegment**](IsolationSegment.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidRelationshipsOrganizationsGet

> V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response V3IsolationSegmentsGuidRelationshipsOrganizationsGet(ctx, guid).Execute()

List organizations relationship



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
	guid := "guid_example" // string | The guid of the isolation segment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGuidRelationshipsOrganizationsGet`: V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidRelationshipsOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete

> V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete(ctx, guid, orgGuid).Execute()

Revoke entitlement to isolation segment for an organization



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
	guid := "guid_example" // string | The guid of the isolation segment
	orgGuid := "orgGuid_example" // string | The guid of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete(context.Background(), guid, orgGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 
**orgGuid** | **string** | The guid of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidRelationshipsOrganizationsOrgGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidRelationshipsOrganizationsPost

> V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response V3IsolationSegmentsGuidRelationshipsOrganizationsPost(ctx, guid).V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest(v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest).Execute()

Entitle organizations for an isolation segment



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
	guid := "guid_example" // string | The guid of the isolation segment
	v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest := *openapiclient.NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest() // V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsPost(context.Background(), guid).V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest(v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGuidRelationshipsOrganizationsPost`: V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGuidRelationshipsOrganizationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest** | [**V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest**](V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest.md) |  | 

### Return type

[**V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response**](V3IsolationSegmentsGuidRelationshipsOrganizationsGet200Response.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsGuidRelationshipsSpacesGet

> V3IsolationSegmentsGuidRelationshipsSpacesGet200Response V3IsolationSegmentsGuidRelationshipsSpacesGet(ctx, guid).Execute()

List spaces relationship



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
	guid := "guid_example" // string | The guid of the isolation segment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsGuidRelationshipsSpacesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsGuidRelationshipsSpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsGuidRelationshipsSpacesGet`: V3IsolationSegmentsGuidRelationshipsSpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsGuidRelationshipsSpacesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the isolation segment | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsGuidRelationshipsSpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3IsolationSegmentsGuidRelationshipsSpacesGet200Response**](V3IsolationSegmentsGuidRelationshipsSpacesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3IsolationSegmentsPost

> IsolationSegment V3IsolationSegmentsPost(ctx).IsolationSegment(isolationSegment).Execute()

Create an isolation segment



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
	isolationSegment := *openapiclient.NewIsolationSegment() // IsolationSegment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3IsolationSegmentsPost(context.Background()).IsolationSegment(isolationSegment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3IsolationSegmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3IsolationSegmentsPost`: IsolationSegment
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3IsolationSegmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3IsolationSegmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isolationSegment** | [**IsolationSegment**](IsolationSegment.md) |  | 

### Return type

[**IsolationSegment**](IsolationSegment.md)

### Authorization

[Admin](../README.md#Admin)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidDropletsGet

> V3AppsGuidDropletsGet200Response V3PackagesGuidDropletsGet(ctx, guid).Guids(guids).States(states).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()

List droplets for a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The guid of the package
	guids := []string{"Inner_example"} // []string | Comma-delimited list of droplet guids to filter by (optional)
	states := []string{"Inner_example"} // []string | Comma-delimited list of droplet states to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGuidDropletsGet(context.Background(), guid).Guids(guids).States(states).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGuidDropletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidDropletsGet`: V3AppsGuidDropletsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGuidDropletsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The guid of the package | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidDropletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **[]string** | Comma-delimited list of droplet guids to filter by | 
 **states** | **[]string** | Comma-delimited list of droplet states to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 

### Return type

[**V3AppsGuidDropletsGet200Response**](V3AppsGuidDropletsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ResourceMatchesPost

> ResourceMatchResponse V3ResourceMatchesPost(ctx).ResourceMatchRequest(resourceMatchRequest).Execute()

Create a resource match



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
	resourceMatchRequest := *openapiclient.NewResourceMatchRequest() // ResourceMatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ResourceMatchesPost(context.Background()).ResourceMatchRequest(resourceMatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ResourceMatchesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ResourceMatchesPost`: ResourceMatchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ResourceMatchesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ResourceMatchesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceMatchRequest** | [**ResourceMatchRequest**](ResourceMatchRequest.md) |  | 

### Return type

[**ResourceMatchResponse**](ResourceMatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RevisionsGuidEnvironmentVariablesGet

> EnvironmentVariables V3RevisionsGuidEnvironmentVariablesGet(ctx, guid).Execute()

Get environment variables for a revision



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
	guid := "guid_example" // string | Unique identifier for the revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RevisionsGuidEnvironmentVariablesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RevisionsGuidEnvironmentVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RevisionsGuidEnvironmentVariablesGet`: EnvironmentVariables
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RevisionsGuidEnvironmentVariablesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RevisionsGuidEnvironmentVariablesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariables**](EnvironmentVariables.md)

### Authorization

[AdminReadOnly](../README.md#AdminReadOnly), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RevisionsGuidGet

> Revision V3RevisionsGuidGet(ctx, guid).Execute()

Get a revision



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
	guid := "guid_example" // string | Unique identifier for the revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RevisionsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RevisionsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RevisionsGuidGet`: Revision
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RevisionsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RevisionsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Revision**](Revision.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RevisionsGuidPatch

> Revision V3RevisionsGuidPatch(ctx, guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()

Update a revision



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
	guid := "guid_example" // string | Unique identifier for the revision
	v3DeploymentsGuidPatchRequest := *openapiclient.NewV3DeploymentsGuidPatchRequest() // V3DeploymentsGuidPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RevisionsGuidPatch(context.Background(), guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RevisionsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RevisionsGuidPatch`: Revision
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RevisionsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RevisionsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DeploymentsGuidPatchRequest** | [**V3DeploymentsGuidPatchRequest**](V3DeploymentsGuidPatchRequest.md) |  | 

### Return type

[**Revision**](Revision.md)

### Authorization

[OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGet

> V3ServiceBrokersGet200Response V3ServiceBrokersGet(ctx).Names(names).SpaceGuids(spaceGuids).OrderBy(orderBy).Page(page).PerPage(perPage).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List service brokers



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
	names := []string{"Inner_example"} // []string | Comma-delimited list of service broker names to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space GUIDs to filter by (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGet(context.Background()).Names(names).SpaceGuids(spaceGuids).OrderBy(orderBy).Page(page).PerPage(perPage).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersGet`: V3ServiceBrokersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Comma-delimited list of service broker names to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space GUIDs to filter by | 
 **orderBy** | **string** | Value to sort by | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3ServiceBrokersGet200Response**](V3ServiceBrokersGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidCatalogPost

> V3ServiceBrokersGuidCatalogPost(ctx, guid).Execute()

Sync service broker catalog



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidCatalogPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGuidCatalogPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidCatalogPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidDelete

> V3ServiceBrokersGuidDelete(ctx, guid).Execute()

Delete a service broker

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidGet

> ServiceBroker V3ServiceBrokersGuidGet(ctx, guid).Execute()

Get a service broker

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersGuidGet`: ServiceBroker
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [AdminReadOnly](../README.md#AdminReadOnly), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidJobsSynchronizationGet

> LastOperation V3ServiceBrokersGuidJobsSynchronizationGet(ctx, guid).Execute()

Get broker synchronization job status



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidJobsSynchronizationGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGuidJobsSynchronizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersGuidJobsSynchronizationGet`: LastOperation
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersGuidJobsSynchronizationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidJobsSynchronizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LastOperation**](LastOperation.md)

### Authorization

[AdminReadOnly](../README.md#AdminReadOnly), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidPatch

> ServiceBroker V3ServiceBrokersGuidPatch(ctx, guid).V3ServiceBrokersGuidPatchRequest(v3ServiceBrokersGuidPatchRequest).Execute()

Update a service broker



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
	guid := "guid_example" // string | 
	v3ServiceBrokersGuidPatchRequest := *openapiclient.NewV3ServiceBrokersGuidPatchRequest() // V3ServiceBrokersGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidPatch(context.Background(), guid).V3ServiceBrokersGuidPatchRequest(v3ServiceBrokersGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersGuidPatch`: ServiceBroker
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ServiceBrokersGuidPatchRequest** | [**V3ServiceBrokersGuidPatchRequest**](V3ServiceBrokersGuidPatchRequest.md) |  | 

### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

[Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersPost

> ServiceBroker V3ServiceBrokersPost(ctx).V3ServiceBrokersPostRequest(v3ServiceBrokersPostRequest).Execute()

Create a service broker



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
	v3ServiceBrokersPostRequest := *openapiclient.NewV3ServiceBrokersPostRequest(*openapiclient.NewAuthentication(*openapiclient.NewAuthenticationCredentials("Password_example", "Username_example"), "Type_example"), "Name_example", "Url_example") // V3ServiceBrokersPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersPost(context.Background()).V3ServiceBrokersPostRequest(v3ServiceBrokersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersPost`: ServiceBroker
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3ServiceBrokersPostRequest** | [**V3ServiceBrokersPostRequest**](V3ServiceBrokersPostRequest.md) |  | 

### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

[Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidDelete

> V3ServiceCredentialBindingsGuidDelete(ctx, guid).Execute()

Delete a service credential binding



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidDetailsGet

> V3ServiceCredentialBindingsGuidDetailsGet200Response V3ServiceCredentialBindingsGuidDetailsGet(ctx, guid).Execute()

Get a service credential binding details



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidDetailsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsGuidDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceCredentialBindingsGuidDetailsGet`: V3ServiceCredentialBindingsGuidDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceCredentialBindingsGuidDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsGuidDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3ServiceCredentialBindingsGuidDetailsGet200Response**](V3ServiceCredentialBindingsGuidDetailsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidGet

> ServiceCredentialBinding V3ServiceCredentialBindingsGuidGet(ctx, guid).Include(include).Execute()

Get a service credential binding



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
	guid := "guid_example" // string | 
	include := []string{"Include_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceCredentialBindingsGuidGet`: ServiceCredentialBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceCredentialBindingsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** |  | 

### Return type

[**ServiceCredentialBinding**](ServiceCredentialBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidParametersGet

> map[string]interface{} V3ServiceCredentialBindingsGuidParametersGet(ctx, guid).Execute()

Get parameters for a service credential binding



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidParametersGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsGuidParametersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceCredentialBindingsGuidParametersGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceCredentialBindingsGuidParametersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsGuidParametersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidPatch

> ServiceCredentialBinding V3ServiceCredentialBindingsGuidPatch(ctx, guid).V3DropletsGuidPatchRequest(v3DropletsGuidPatchRequest).Execute()

Update a service credential binding



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
	guid := "guid_example" // string | 
	v3DropletsGuidPatchRequest := *openapiclient.NewV3DropletsGuidPatchRequest() // V3DropletsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidPatch(context.Background(), guid).V3DropletsGuidPatchRequest(v3DropletsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceCredentialBindingsGuidPatch`: ServiceCredentialBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceCredentialBindingsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DropletsGuidPatchRequest** | [**V3DropletsGuidPatchRequest**](V3DropletsGuidPatchRequest.md) |  | 

### Return type

[**ServiceCredentialBinding**](ServiceCredentialBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsPost

> ServiceCredentialBinding V3ServiceCredentialBindingsPost(ctx).V3ServiceCredentialBindingsPostRequest(v3ServiceCredentialBindingsPostRequest).Execute()

Create a service credential binding



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
	v3ServiceCredentialBindingsPostRequest := *openapiclient.NewV3ServiceCredentialBindingsPostRequest(*openapiclient.NewV3ServiceCredentialBindingsPostRequestRelationships(*openapiclient.NewAssignDefaultIsolationSegmentRequest()), "Type_example") // V3ServiceCredentialBindingsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsPost(context.Background()).V3ServiceCredentialBindingsPostRequest(v3ServiceCredentialBindingsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceCredentialBindingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceCredentialBindingsPost`: ServiceCredentialBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceCredentialBindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceCredentialBindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3ServiceCredentialBindingsPostRequest** | [**V3ServiceCredentialBindingsPostRequest**](V3ServiceCredentialBindingsPostRequest.md) |  | 

### Return type

[**ServiceCredentialBinding**](ServiceCredentialBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGet

> V3ServiceInstancesGet200Response V3ServiceInstancesGet(ctx).Names(names).Guids(guids).Type_(type_).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).ServicePlanGuids(servicePlanGuids).ServicePlanNames(servicePlanNames).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

Retrieve service instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	names := []string{"Inner_example"} // []string | Comma-delimited list of service instance names to filter by (optional)
	guids := []string{"Inner_example"} // []string | Comma-delimited list of service instance guids to filter by (optional)
	type_ := "type__example" // string | Filter by type (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	servicePlanGuids := []string{"Inner_example"} // []string | Comma-delimited list of service plan guids to filter by (optional)
	servicePlanNames := []string{"Inner_example"} // []string | Comma-delimited list of service plan names to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by (optional)
	updatedAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGet(context.Background()).Names(names).Guids(guids).Type_(type_).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).ServicePlanGuids(servicePlanGuids).ServicePlanNames(servicePlanNames).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGet`: V3ServiceInstancesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceInstancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Comma-delimited list of service instance names to filter by | 
 **guids** | **[]string** | Comma-delimited list of service instance guids to filter by | 
 **type_** | **string** | Filter by type | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **servicePlanGuids** | **[]string** | Comma-delimited list of service plan guids to filter by | 
 **servicePlanNames** | **[]string** | Comma-delimited list of service plan names to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by | 
 **updatedAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by | 

### Return type

[**V3ServiceInstancesGet200Response**](V3ServiceInstancesGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidCredentialsGet

> map[string]string V3ServiceInstancesGuidCredentialsGet(ctx, guid).Execute()

Get credentials for a user-provided service instance



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
	guid := "guid_example" // string | GUID of the service instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidCredentialsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidCredentialsGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceInstancesGuidCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidDelete

> V3ServiceInstancesGuidDelete(ctx, guid).Purge(purge).Execute()

Delete a service instance



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
	guid := "guid_example" // string | GUID of the service instance
	purge := true // bool | If true, deletes the service instance and all associated resources without any interaction with the service broker (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidDelete(context.Background(), guid).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If true, deletes the service instance and all associated resources without any interaction with the service broker | 

### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidPatch

> ServiceInstance V3ServiceInstancesGuidPatch(ctx, guid).V3ServiceInstancesGuidPatchRequest(v3ServiceInstancesGuidPatchRequest).Execute()

Update a service instance



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
	guid := "guid_example" // string | GUID of the service instance
	v3ServiceInstancesGuidPatchRequest := openapiclient._v3_service_instances__guid__patch_request{ManagedServiceInstanceUpdate: openapiclient.NewManagedServiceInstanceUpdate()} // V3ServiceInstancesGuidPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidPatch(context.Background(), guid).V3ServiceInstancesGuidPatchRequest(v3ServiceInstancesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidPatch`: ServiceInstance
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceInstancesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ServiceInstancesGuidPatchRequest** | [**V3ServiceInstancesGuidPatchRequest**](V3ServiceInstancesGuidPatchRequest.md) |  | 

### Return type

[**ServiceInstance**](ServiceInstance.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidRelationshipsSharedSpacesPost

> V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response V3ServiceInstancesGuidRelationshipsSharedSpacesPost(ctx, guid).V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response(v3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response).Execute()

Share a service instance to other spaces



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
	guid := "guid_example" // string | GUID of the service instance
	v3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response := *openapiclient.NewV3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response() // V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesPost(context.Background(), guid).V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response(v3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidRelationshipsSharedSpacesPost`: V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response** | [**V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response**](V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response.md) |  | 

### Return type

[**V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response**](V3ServiceInstancesGuidRelationshipsSharedSpacesGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete

> V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete(ctx, guid, spaceGuid).Execute()

Unshare a service instance from another space



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
	guid := "guid_example" // string | GUID of the service instance
	spaceGuid := "spaceGuid_example" // string | GUID of the space

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete(context.Background(), guid, spaceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 
**spaceGuid** | **string** | GUID of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceOfferingsGet

> ServiceOfferingList V3ServiceOfferingsGet(ctx).Names(names).Available(available).ServiceBrokerGuids(serviceBrokerGuids).ServiceBrokerNames(serviceBrokerNames).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).OrderBy(orderBy).Page(page).PerPage(perPage).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List service offerings



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
	names := []string{"Inner_example"} // []string | Comma-delimited list of names to filter by (optional)
	available := true // bool | Filter by the available property; valid values are true or false (optional)
	serviceBrokerGuids := []string{"Inner_example"} // []string | Comma-delimited list of service broker GUIDs to filter by (optional)
	serviceBrokerNames := []string{"Inner_example"} // []string | Comma-delimited list of service broker names to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space GUIDs to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization GUIDs to filter by (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceOfferingsGet(context.Background()).Names(names).Available(available).ServiceBrokerGuids(serviceBrokerGuids).ServiceBrokerNames(serviceBrokerNames).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).LabelSelector(labelSelector).OrderBy(orderBy).Page(page).PerPage(perPage).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceOfferingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceOfferingsGet`: ServiceOfferingList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceOfferingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceOfferingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Comma-delimited list of names to filter by | 
 **available** | **bool** | Filter by the available property; valid values are true or false | 
 **serviceBrokerGuids** | **[]string** | Comma-delimited list of service broker GUIDs to filter by | 
 **serviceBrokerNames** | **[]string** | Comma-delimited list of service broker names to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space GUIDs to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization GUIDs to filter by | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**ServiceOfferingList**](ServiceOfferingList.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceOfferingsGuidDelete

> V3ServiceOfferingsGuidDelete(ctx, guid).Purge(purge).Execute()

Delete a service offering



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
	guid := "guid_example" // string | The service offering GUID
	purge := true // bool | If true, any service plans, instances, and bindings associated with this service offering will also be deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceOfferingsGuidDelete(context.Background(), guid).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceOfferingsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service offering GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceOfferingsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If true, any service plans, instances, and bindings associated with this service offering will also be deleted | 

### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceOfferingsGuidGet

> ServiceOffering V3ServiceOfferingsGuidGet(ctx, guid).Execute()

Get a service offering



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
	guid := "guid_example" // string | The service offering GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceOfferingsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceOfferingsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceOfferingsGuidGet`: ServiceOffering
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceOfferingsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service offering GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceOfferingsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceOfferingsGuidPatch

> ServiceOffering V3ServiceOfferingsGuidPatch(ctx, guid).ServiceOfferingUpdate(serviceOfferingUpdate).Execute()

Update a service offering



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
	guid := "guid_example" // string | The service offering GUID
	serviceOfferingUpdate := *openapiclient.NewServiceOfferingUpdate() // ServiceOfferingUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceOfferingsGuidPatch(context.Background(), guid).ServiceOfferingUpdate(serviceOfferingUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceOfferingsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceOfferingsGuidPatch`: ServiceOffering
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceOfferingsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service offering GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceOfferingsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceOfferingUpdate** | [**ServiceOfferingUpdate**](ServiceOfferingUpdate.md) |  | 

### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceOfferingsPost

> ServiceOffering V3ServiceOfferingsPost(ctx).ServiceOfferingCreate(serviceOfferingCreate).Execute()

Create a service offering



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
	serviceOfferingCreate := *openapiclient.NewServiceOfferingCreate() // ServiceOfferingCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceOfferingsPost(context.Background()).ServiceOfferingCreate(serviceOfferingCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceOfferingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceOfferingsPost`: ServiceOffering
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceOfferingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceOfferingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceOfferingCreate** | [**ServiceOfferingCreate**](ServiceOfferingCreate.md) |  | 

### Return type

[**ServiceOffering**](ServiceOffering.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGet

> ServicePlanList V3ServicePlansGet(ctx).Names(names).Available(available).BrokerCatalogIds(brokerCatalogIds).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).ServiceBrokerGuids(serviceBrokerGuids).ServiceBrokerNames(serviceBrokerNames).ServiceOfferingGuids(serviceOfferingGuids).ServiceOfferingNames(serviceOfferingNames).ServiceInstanceGuids(serviceInstanceGuids).Include(include).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Fields(fields).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List service plans



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
	names := []string{"Inner_example"} // []string | Comma-delimited list of names to filter by (optional)
	available := true // bool | Filter by the available property; valid values are true or false (optional)
	brokerCatalogIds := []string{"Inner_example"} // []string | Comma-delimited list of IDs provided by the service broker for the service plan to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space GUIDs to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization GUIDs to filter by (optional)
	serviceBrokerGuids := []string{"Inner_example"} // []string | Comma-delimited list of service broker GUIDs to filter by (optional)
	serviceBrokerNames := []string{"Inner_example"} // []string | Comma-delimited list of service broker names to filter by (optional)
	serviceOfferingGuids := []string{"Inner_example"} // []string | Comma-delimited list of service Offering GUIDs to filter by (optional)
	serviceOfferingNames := []string{"Inner_example"} // []string | Comma-delimited list of service Offering names to filter by (optional)
	serviceInstanceGuids := []string{"Inner_example"} // []string | Comma-delimited list of service Instance GUIDs to filter by (optional)
	include := []string{"Inner_example"} // []string | Optionally include a list of related resources in the response; valid values are space.organization and service_offering (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	fields := "fields_example" // string | Allowed values for fields (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGet(context.Background()).Names(names).Available(available).BrokerCatalogIds(brokerCatalogIds).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).ServiceBrokerGuids(serviceBrokerGuids).ServiceBrokerNames(serviceBrokerNames).ServiceOfferingGuids(serviceOfferingGuids).ServiceOfferingNames(serviceOfferingNames).ServiceInstanceGuids(serviceInstanceGuids).Include(include).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Fields(fields).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGet`: ServicePlanList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Comma-delimited list of names to filter by | 
 **available** | **bool** | Filter by the available property; valid values are true or false | 
 **brokerCatalogIds** | **[]string** | Comma-delimited list of IDs provided by the service broker for the service plan to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space GUIDs to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization GUIDs to filter by | 
 **serviceBrokerGuids** | **[]string** | Comma-delimited list of service broker GUIDs to filter by | 
 **serviceBrokerNames** | **[]string** | Comma-delimited list of service broker names to filter by | 
 **serviceOfferingGuids** | **[]string** | Comma-delimited list of service Offering GUIDs to filter by | 
 **serviceOfferingNames** | **[]string** | Comma-delimited list of service Offering names to filter by | 
 **serviceInstanceGuids** | **[]string** | Comma-delimited list of service Instance GUIDs to filter by | 
 **include** | **[]string** | Optionally include a list of related resources in the response; valid values are space.organization and service_offering | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **fields** | **string** | Allowed values for fields | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**ServicePlanList**](ServicePlanList.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidDelete

> V3ServicePlansGuidDelete(ctx, guid).Purge(purge).Execute()

Delete a service plan



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
	guid := "guid_example" // string | The service plan GUID
	purge := true // bool | If true, any service plans, instances, and bindings associated with this service plan will also be deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServicePlansGuidDelete(context.Background(), guid).Purge(purge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **bool** | If true, any service plans, instances, and bindings associated with this service plan will also be deleted | 

### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidGet

> ServicePlan V3ServicePlansGuidGet(ctx, guid).Execute()

Get a service plan



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
	guid := "guid_example" // string | The service plan GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidGet`: ServicePlan
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicePlan**](ServicePlan.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidPatch

> ServicePlan V3ServicePlansGuidPatch(ctx, guid).ServicePlanUpdate(servicePlanUpdate).Execute()

Update a service plan



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
	guid := "guid_example" // string | The service plan GUID
	servicePlanUpdate := *openapiclient.NewServicePlanUpdate() // ServicePlanUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGuidPatch(context.Background(), guid).ServicePlanUpdate(servicePlanUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidPatch`: ServicePlan
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePlanUpdate** | [**ServicePlanUpdate**](ServicePlanUpdate.md) |  | 

### Return type

[**ServicePlan**](ServicePlan.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansPost

> ServicePlan V3ServicePlansPost(ctx).ServicePlanCreate(servicePlanCreate).Execute()

Create a service plan



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
	servicePlanCreate := *openapiclient.NewServicePlanCreate() // ServicePlanCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansPost(context.Background()).ServicePlanCreate(servicePlanCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansPost`: ServicePlan
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servicePlanCreate** | [**ServicePlanCreate**](ServicePlanCreate.md) |  | 

### Return type

[**ServicePlan**](ServicePlan.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsGet

> V3ServiceRouteBindingsGet200Response V3ServiceRouteBindingsGet(ctx).RouteGuids(routeGuids).ServiceInstanceGuids(serviceInstanceGuids).ServiceInstanceNames(serviceInstanceNames).LabelSelector(labelSelector).Guids(guids).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()

List service route bindings



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
	routeGuids := []string{"Inner_example"} // []string |  (optional)
	serviceInstanceGuids := []string{"Inner_example"} // []string |  (optional)
	serviceInstanceNames := []string{"Inner_example"} // []string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	guids := []string{"Inner_example"} // []string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)
	include := []string{"Inner_example"} // []string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsGet(context.Background()).RouteGuids(routeGuids).ServiceInstanceGuids(serviceInstanceGuids).ServiceInstanceNames(serviceInstanceNames).LabelSelector(labelSelector).Guids(guids).CreatedAts(createdAts).UpdatedAts(updatedAts).Include(include).Page(page).PerPage(perPage).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceRouteBindingsGet`: V3ServiceRouteBindingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceRouteBindingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeGuids** | **[]string** |  | 
 **serviceInstanceGuids** | **[]string** |  | 
 **serviceInstanceNames** | **[]string** |  | 
 **labelSelector** | **string** |  | 
 **guids** | **[]string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 
 **include** | **[]string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 

### Return type

[**V3ServiceRouteBindingsGet200Response**](V3ServiceRouteBindingsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsGuidDelete

> V3ServiceRouteBindingsGuidDelete(ctx, guid).Execute()

Delete a service route binding



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsGuidGet

> ServiceRouteBinding V3ServiceRouteBindingsGuidGet(ctx, guid).Include(include).Execute()

Get a service route binding



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
	guid := "guid_example" // string | 
	include := []string{"Include_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceRouteBindingsGuidGet`: ServiceRouteBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceRouteBindingsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** |  | 

### Return type

[**ServiceRouteBinding**](ServiceRouteBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsGuidParametersGet

> map[string]string V3ServiceRouteBindingsGuidParametersGet(ctx, guid).Execute()

Get parameters for a route binding



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsGuidParametersGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsGuidParametersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceRouteBindingsGuidParametersGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceRouteBindingsGuidParametersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsGuidParametersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsGuidPatch

> ServiceRouteBinding V3ServiceRouteBindingsGuidPatch(ctx, guid).V3ServiceRouteBindingsGuidPatchRequest(v3ServiceRouteBindingsGuidPatchRequest).Execute()

Update a service route binding



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
	guid := "guid_example" // string | 
	v3ServiceRouteBindingsGuidPatchRequest := *openapiclient.NewV3ServiceRouteBindingsGuidPatchRequest() // V3ServiceRouteBindingsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsGuidPatch(context.Background(), guid).V3ServiceRouteBindingsGuidPatchRequest(v3ServiceRouteBindingsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceRouteBindingsGuidPatch`: ServiceRouteBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceRouteBindingsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3ServiceRouteBindingsGuidPatchRequest** | [**V3ServiceRouteBindingsGuidPatchRequest**](V3ServiceRouteBindingsGuidPatchRequest.md) |  | 

### Return type

[**ServiceRouteBinding**](ServiceRouteBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceRouteBindingsPost

> ServiceRouteBinding V3ServiceRouteBindingsPost(ctx).V3ServiceRouteBindingsPostRequest(v3ServiceRouteBindingsPostRequest).Execute()

Create a service route binding



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
	v3ServiceRouteBindingsPostRequest := *openapiclient.NewV3ServiceRouteBindingsPostRequest(*openapiclient.NewServiceRouteBindingRelationships()) // V3ServiceRouteBindingsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceRouteBindingsPost(context.Background()).V3ServiceRouteBindingsPostRequest(v3ServiceRouteBindingsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceRouteBindingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceRouteBindingsPost`: ServiceRouteBinding
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceRouteBindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceRouteBindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3ServiceRouteBindingsPostRequest** | [**V3ServiceRouteBindingsPostRequest**](V3ServiceRouteBindingsPostRequest.md) |  | 

### Return type

[**ServiceRouteBinding**](ServiceRouteBinding.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceUsageEventsGet

> V3ServiceUsageEventsGet200Response V3ServiceUsageEventsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).AfterGuid(afterGuid).Guids(guids).ServiceInstanceTypes(serviceInstanceTypes).ServiceOfferingGuids(serviceOfferingGuids).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List service usage events



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
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	afterGuid := "afterGuid_example" // string |  (optional)
	guids := []string{"Inner_example"} // []string |  (optional)
	serviceInstanceTypes := []string{"ServiceInstanceTypes_example"} // []string |  (optional)
	serviceOfferingGuids := []string{"Inner_example"} // []string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceUsageEventsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).AfterGuid(afterGuid).Guids(guids).ServiceInstanceTypes(serviceInstanceTypes).ServiceOfferingGuids(serviceOfferingGuids).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceUsageEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceUsageEventsGet`: V3ServiceUsageEventsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceUsageEventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceUsageEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **afterGuid** | **string** |  | 
 **guids** | **[]string** |  | 
 **serviceInstanceTypes** | **[]string** |  | 
 **serviceOfferingGuids** | **[]string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3ServiceUsageEventsGet200Response**](V3ServiceUsageEventsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceUsageEventsGuidGet

> ServiceUsageEvent V3ServiceUsageEventsGuidGet(ctx, guid).Execute()

Get a service usage event



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceUsageEventsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceUsageEventsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceUsageEventsGuidGet`: ServiceUsageEvent
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceUsageEventsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceUsageEventsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceUsageEvent**](ServiceUsageEvent.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceUsageEventsPost

> V3ServiceUsageEventsPost(ctx).Execute()

Purge and seed service usage events



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServiceUsageEventsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceUsageEventsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceUsageEventsPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGet

> V3SpacesGet200Response V3SpacesGet(ctx).Names(names).Guids(guids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List spaces



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
	names := "names_example" // string |  (optional)
	guids := "guids_example" // string |  (optional)
	organizationGuids := "organizationGuids_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	include := "include_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGet(context.Background()).Names(names).Guids(guids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGet`: V3SpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **string** |  | 
 **guids** | **string** |  | 
 **organizationGuids** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **include** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3SpacesGet200Response**](V3SpacesGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidDelete

> V3SpacesGuidDelete(ctx, guid).Execute()

Delete a space



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3SpacesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3SpacesGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidGet

> Space V3SpacesGuidGet(ctx, guid).Execute()

Get a space



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidGet`: Space
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Space**](Space.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidPatch

> Space V3SpacesGuidPatch(ctx, guid).V3SpacesGuidPatchRequest(v3SpacesGuidPatchRequest).Execute()

Update a space



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
	guid := "guid_example" // string | 
	v3SpacesGuidPatchRequest := *openapiclient.NewV3SpacesGuidPatchRequest() // V3SpacesGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidPatch(context.Background(), guid).V3SpacesGuidPatchRequest(v3SpacesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidPatch`: Space
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SpacesGuidPatchRequest** | [**V3SpacesGuidPatchRequest**](V3SpacesGuidPatchRequest.md) |  | 

### Return type

[**Space**](Space.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidRelationshipsIsolationSegmentGet

> AssignDefaultIsolationSegmentRequest V3SpacesGuidRelationshipsIsolationSegmentGet(ctx, guid).Execute()

Get assigned isolation segment



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidRelationshipsIsolationSegmentGet`: AssignDefaultIsolationSegmentRequest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidRelationshipsIsolationSegmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssignDefaultIsolationSegmentRequest**](AssignDefaultIsolationSegmentRequest.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidRelationshipsIsolationSegmentPatch

> AssignDefaultIsolationSegmentRequest V3SpacesGuidRelationshipsIsolationSegmentPatch(ctx, guid).AssignDefaultIsolationSegmentRequest(assignDefaultIsolationSegmentRequest).Execute()

Manage isolation segment



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
	guid := "guid_example" // string | 
	assignDefaultIsolationSegmentRequest := *openapiclient.NewAssignDefaultIsolationSegmentRequest() // AssignDefaultIsolationSegmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentPatch(context.Background(), guid).AssignDefaultIsolationSegmentRequest(assignDefaultIsolationSegmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidRelationshipsIsolationSegmentPatch`: AssignDefaultIsolationSegmentRequest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidRelationshipsIsolationSegmentPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignDefaultIsolationSegmentRequest** | [**AssignDefaultIsolationSegmentRequest**](AssignDefaultIsolationSegmentRequest.md) |  | 

### Return type

[**AssignDefaultIsolationSegmentRequest**](AssignDefaultIsolationSegmentRequest.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidUsersGet

> V3SpacesGuidUsersGet200Response V3SpacesGuidUsersGet(ctx, guid).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List users for a space



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
	guid := "guid_example" // string | 
	guids := "guids_example" // string |  (optional)
	usernames := "usernames_example" // string |  (optional)
	partialUsernames := "partialUsernames_example" // string |  (optional)
	origins := "origins_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidUsersGet(context.Background(), guid).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidUsersGet`: V3SpacesGuidUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **string** |  | 
 **usernames** | **string** |  | 
 **partialUsernames** | **string** |  | 
 **origins** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3SpacesGuidUsersGet200Response**](V3SpacesGuidUsersGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesPost

> Space V3SpacesPost(ctx).V3SpacesPostRequest(v3SpacesPostRequest).Execute()

Create a space



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
	v3SpacesPostRequest := *openapiclient.NewV3SpacesPostRequest("Name_example", *openapiclient.NewV3SpacesPostRequestRelationships()) // V3SpacesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesPost(context.Background()).V3SpacesPostRequest(v3SpacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesPost`: Space
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3SpacesPostRequest** | [**V3SpacesPostRequest**](V3SpacesPostRequest.md) |  | 

### Return type

[**Space**](Space.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGet

> V3StacksGet200Response V3StacksGet(ctx).Names(names).Default_(default_).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List all stacks



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
	names := "names_example" // string |  (optional)
	default_ := true // bool |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 50)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksGet(context.Background()).Names(names).Default_(default_).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3StacksGet`: V3StacksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3StacksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3StacksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **string** |  | 
 **default_** | **bool** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 50]
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3StacksGet200Response**](V3StacksGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGuidAppsGet

> V3StacksGuidAppsGet200Response V3StacksGuidAppsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List apps on a stack



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
	guid := "guid_example" // string | 
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 50)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksGuidAppsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksGuidAppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3StacksGuidAppsGet`: V3StacksGuidAppsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3StacksGuidAppsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3StacksGuidAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 50]
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3StacksGuidAppsGet200Response**](V3StacksGuidAppsGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGuidDelete

> V3StacksGuidDelete(ctx, guid).Execute()

Delete a stack



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3StacksGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3StacksGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGuidGet

> Stack V3StacksGuidGet(ctx, guid).Execute()

Get a stack by GUID



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3StacksGuidGet`: Stack
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3StacksGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3StacksGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Stack**](Stack.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGuidPatch

> Stack V3StacksGuidPatch(ctx, guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()

Update a stack



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
	guid := "guid_example" // string | 
	v3DeploymentsGuidPatchRequest := *openapiclient.NewV3DeploymentsGuidPatchRequest() // V3DeploymentsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksGuidPatch(context.Background(), guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3StacksGuidPatch`: Stack
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3StacksGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3StacksGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DeploymentsGuidPatchRequest** | [**V3DeploymentsGuidPatchRequest**](V3DeploymentsGuidPatchRequest.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksPost

> Stack V3StacksPost(ctx).V3StacksPostRequest(v3StacksPostRequest).Execute()

Create a stack



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
	v3StacksPostRequest := *openapiclient.NewV3StacksPostRequest("Name_example") // V3StacksPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksPost(context.Background()).V3StacksPostRequest(v3StacksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3StacksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3StacksPost`: Stack
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3StacksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3StacksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3StacksPostRequest** | [**V3StacksPostRequest**](V3StacksPostRequest.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TasksGet

> V3TasksGet200Response V3TasksGet(ctx).Guids(guids).Names(names).States(states).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List all tasks



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
	guids := "guids_example" // string |  (optional)
	names := "names_example" // string |  (optional)
	states := "states_example" // string |  (optional)
	appGuids := "appGuids_example" // string |  (optional)
	spaceGuids := "spaceGuids_example" // string |  (optional)
	organizationGuids := "organizationGuids_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional) (default to 50)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3TasksGet(context.Background()).Guids(guids).Names(names).States(states).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3TasksGet`: V3TasksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3TasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** |  | 
 **names** | **string** |  | 
 **states** | **string** |  | 
 **appGuids** | **string** |  | 
 **spaceGuids** | **string** |  | 
 **organizationGuids** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | [default to 50]
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3TasksGet200Response**](V3TasksGet200Response.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TasksGuidGet

> Task V3TasksGuidGet(ctx, guid).Execute()

Get a task



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3TasksGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3TasksGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3TasksGuidGet`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3TasksGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3TasksGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TasksGuidPatch

> Task V3TasksGuidPatch(ctx, guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()

Update a task



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
	guid := "guid_example" // string | 
	v3DeploymentsGuidPatchRequest := *openapiclient.NewV3DeploymentsGuidPatchRequest() // V3DeploymentsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3TasksGuidPatch(context.Background(), guid).V3DeploymentsGuidPatchRequest(v3DeploymentsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3TasksGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3TasksGuidPatch`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3TasksGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3TasksGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3DeploymentsGuidPatchRequest** | [**V3DeploymentsGuidPatchRequest**](V3DeploymentsGuidPatchRequest.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TasksGuidPost

> Task V3TasksGuidPost(ctx, guid).Execute()

Cancel a task



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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3TasksGuidPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3TasksGuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3TasksGuidPost`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3TasksGuidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3TasksGuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[GlobalAuditor](../README.md#GlobalAuditor), [OAuth2](../README.md#OAuth2), [AdminReadOnly](../README.md#AdminReadOnly), [OrgManager](../README.md#OrgManager), [Admin](../README.md#Admin), [SpaceDeveloper](../README.md#SpaceDeveloper), [SpaceManager](../README.md#SpaceManager), [SpaceAuditor](../README.md#SpaceAuditor), [SpaceSupporter](../README.md#SpaceSupporter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UsersGet

> V3UsersGet200Response V3UsersGet(ctx).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List users



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
	guids := []string{"Inner_example"} // []string | Comma-delimited list of user guids to filter by (optional)
	usernames := []string{"Inner_example"} // []string | Comma-delimited list of usernames to filter by. Mutually exclusive with partial_usernames (optional)
	partialUsernames := []string{"Inner_example"} // []string | Comma-delimited list of strings to search by. When using this query parameter, all the users that contain the string provided in their username will be returned. Mutually exclusive with usernames (optional)
	origins := []string{"Inner_example"} // []string | Comma-delimited list of user origins (user stores) to filter by, for example, users authenticated by UAA have the origin “uaa”; users authenticated by an LDAP provider have the origin “ldap”; when filtering by origins, usernames must be included (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3UsersGet(context.Background()).Guids(guids).Usernames(usernames).PartialUsernames(partialUsernames).Origins(origins).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UsersGet`: V3UsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3UsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of user guids to filter by | 
 **usernames** | **[]string** | Comma-delimited list of usernames to filter by. Mutually exclusive with partial_usernames | 
 **partialUsernames** | **[]string** | Comma-delimited list of strings to search by. When using this query parameter, all the users that contain the string provided in their username will be returned. Mutually exclusive with usernames | 
 **origins** | **[]string** | Comma-delimited list of user origins (user stores) to filter by, for example, users authenticated by UAA have the origin “uaa”; users authenticated by an LDAP provider have the origin “ldap”; when filtering by origins, usernames must be included | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**V3UsersGet200Response**](V3UsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UsersGuidDelete

> V3UsersGuidDelete(ctx, guid).Execute()

Delete a user



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
	guid := "guid_example" // string | Unique identifier for the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3UsersGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3UsersGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UsersGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UsersGuidGet

> User V3UsersGuidGet(ctx, guid).Execute()

Get a user



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
	guid := "guid_example" // string | Unique identifier for the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3UsersGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3UsersGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UsersGuidGet`: User
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3UsersGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UsersGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UsersGuidPatch

> User V3UsersGuidPatch(ctx, guid).V3UsersGuidPatchRequest(v3UsersGuidPatchRequest).Execute()

Update a user



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
	guid := "guid_example" // string | Unique identifier for the user
	v3UsersGuidPatchRequest := *openapiclient.NewV3UsersGuidPatchRequest() // V3UsersGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3UsersGuidPatch(context.Background(), guid).V3UsersGuidPatchRequest(v3UsersGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3UsersGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UsersGuidPatch`: User
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3UsersGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UsersGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3UsersGuidPatchRequest** | [**V3UsersGuidPatchRequest**](V3UsersGuidPatchRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UsersPost

> User V3UsersPost(ctx).V3UsersPostRequest(v3UsersPostRequest).Execute()

Create a user



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
	v3UsersPostRequest := *openapiclient.NewV3UsersPostRequest() // V3UsersPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3UsersPost(context.Background()).V3UsersPostRequest(v3UsersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UsersPost`: User
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3UsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3UsersPostRequest** | [**V3UsersPostRequest**](V3UsersPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

