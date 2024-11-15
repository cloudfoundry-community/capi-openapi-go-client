# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyOrganizationQuotaToOrganizations**](DefaultAPI.md#ApplyOrganizationQuotaToOrganizations) | **Post** /v3/organization_quotas/{quota_guid}/relationships/organizations | Apply an organization quota to organizations
[**AssignDefaultIsolationSegment**](DefaultAPI.md#AssignDefaultIsolationSegment) | **Patch** /v3/organizations/{guid}/relationships/default_isolation_segment | Assign Default Isolation Segment
[**CreateOrganization**](DefaultAPI.md#CreateOrganization) | **Post** /v3/organizations | Create an Organization
[**CreateOrganizationQuota**](DefaultAPI.md#CreateOrganizationQuota) | **Post** /v3/organization_quotas | Create an organization quota
[**DeleteOrganization**](DefaultAPI.md#DeleteOrganization) | **Delete** /v3/organizations/{guid} | Delete an Organization
[**DeleteOrganizationQuota**](DefaultAPI.md#DeleteOrganizationQuota) | **Delete** /v3/organization_quotas/{guid} | Delete an organization quota
[**GetDefaultDomain**](DefaultAPI.md#GetDefaultDomain) | **Get** /v3/organizations/{guid}/domains/default | Get Default Domain
[**GetDefaultIsolationSegment**](DefaultAPI.md#GetDefaultIsolationSegment) | **Get** /v3/organizations/{guid}/relationships/default_isolation_segment | Get Default Isolation Segment
[**GetOrganization**](DefaultAPI.md#GetOrganization) | **Get** /v3/organizations/{guid} | Get an Organization
[**GetOrganizationQuota**](DefaultAPI.md#GetOrganizationQuota) | **Get** /v3/organization_quotas/{guid} | Get an organization quota
[**GetUsageSummary**](DefaultAPI.md#GetUsageSummary) | **Get** /v3/organizations/{guid}/usage_summary | Get Usage Summary
[**ListOrganizationQuotas**](DefaultAPI.md#ListOrganizationQuotas) | **Get** /v3/organization_quotas | List organization quotas
[**ListOrganizations**](DefaultAPI.md#ListOrganizations) | **Get** /v3/organizations | List Organizations
[**ListUsersForOrganization**](DefaultAPI.md#ListUsersForOrganization) | **Get** /v3/organizations/{guid}/users | List Users for an Organization
[**RootGet**](DefaultAPI.md#RootGet) | **Get** / | Global API Root
[**UpdateOrganization**](DefaultAPI.md#UpdateOrganization) | **Patch** /v3/organizations/{guid} | Update an Organization
[**UpdateOrganizationQuota**](DefaultAPI.md#UpdateOrganizationQuota) | **Patch** /v3/organization_quotas/{guid} | Update an organization quota
[**V3AdminActionsClearBuildpackCachePost**](DefaultAPI.md#V3AdminActionsClearBuildpackCachePost) | **Post** /v3/admin/actions/clear_buildpack_cache | Clear buildpack cache
[**V3AppsGet**](DefaultAPI.md#V3AppsGet) | **Get** /v3/apps | List apps
[**V3AppsGuidActionsClearBuildpackCachePost**](DefaultAPI.md#V3AppsGuidActionsClearBuildpackCachePost) | **Post** /v3/apps/{guid}/actions/clear_buildpack_cache | Clear buildpack cache for application
[**V3AppsGuidActionsRestartPost**](DefaultAPI.md#V3AppsGuidActionsRestartPost) | **Post** /v3/apps/{guid}/actions/restart | Restart an app
[**V3AppsGuidActionsStartPost**](DefaultAPI.md#V3AppsGuidActionsStartPost) | **Post** /v3/apps/{guid}/actions/start | Start an app
[**V3AppsGuidActionsStopPost**](DefaultAPI.md#V3AppsGuidActionsStopPost) | **Post** /v3/apps/{guid}/actions/stop | Stop an app
[**V3AppsGuidBuildsGet**](DefaultAPI.md#V3AppsGuidBuildsGet) | **Get** /v3/apps/{guid}/builds | List builds for an app
[**V3AppsGuidDelete**](DefaultAPI.md#V3AppsGuidDelete) | **Delete** /v3/apps/{guid} | Delete an app
[**V3AppsGuidDropletsCurrentGet**](DefaultAPI.md#V3AppsGuidDropletsCurrentGet) | **Get** /v3/apps/{guid}/droplets/current | Get current droplet
[**V3AppsGuidDropletsGet**](DefaultAPI.md#V3AppsGuidDropletsGet) | **Get** /v3/apps/{guid}/droplets | List droplets for an app
[**V3AppsGuidEnvGet**](DefaultAPI.md#V3AppsGuidEnvGet) | **Get** /v3/apps/{guid}/env | Get environment for an app
[**V3AppsGuidEnvironmentVariablesGet**](DefaultAPI.md#V3AppsGuidEnvironmentVariablesGet) | **Get** /v3/apps/{guid}/environment_variables | Get environment variables for an app
[**V3AppsGuidEnvironmentVariablesPatch**](DefaultAPI.md#V3AppsGuidEnvironmentVariablesPatch) | **Patch** /v3/apps/{guid}/environment_variables | Update environment variables for an app
[**V3AppsGuidFeaturesGet**](DefaultAPI.md#V3AppsGuidFeaturesGet) | **Get** /v3/apps/{guid}/features | List app features
[**V3AppsGuidFeaturesNameGet**](DefaultAPI.md#V3AppsGuidFeaturesNameGet) | **Get** /v3/apps/{guid}/features/{name} | Get an app feature
[**V3AppsGuidFeaturesNamePatch**](DefaultAPI.md#V3AppsGuidFeaturesNamePatch) | **Patch** /v3/apps/{guid}/features/{name} | Update an app feature
[**V3AppsGuidGet**](DefaultAPI.md#V3AppsGuidGet) | **Get** /v3/apps/{guid} | Retrieve a specific app
[**V3AppsGuidPatch**](DefaultAPI.md#V3AppsGuidPatch) | **Patch** /v3/apps/{guid} | Update an app
[**V3AppsGuidPermissionsGet**](DefaultAPI.md#V3AppsGuidPermissionsGet) | **Get** /v3/apps/{guid}/permissions | Get permissions for an app
[**V3AppsGuidProcessesGet**](DefaultAPI.md#V3AppsGuidProcessesGet) | **Get** /v3/apps/{guid}/processes | List processes for app
[**V3AppsGuidRelationshipsCurrentDropletGet**](DefaultAPI.md#V3AppsGuidRelationshipsCurrentDropletGet) | **Get** /v3/apps/{guid}/relationships/current_droplet | Get current droplet association for an app
[**V3AppsGuidRevisionsDeployedGet**](DefaultAPI.md#V3AppsGuidRevisionsDeployedGet) | **Get** /v3/apps/{guid}/revisions/deployed | List deployed revisions for an app
[**V3AppsGuidRevisionsGet**](DefaultAPI.md#V3AppsGuidRevisionsGet) | **Get** /v3/apps/{guid}/revisions | List revisions for an app
[**V3AppsGuidRoutesGet**](DefaultAPI.md#V3AppsGuidRoutesGet) | **Get** /v3/apps/{guid}/routes | Retrieve all routes for an app
[**V3AppsGuidSidecarsGet**](DefaultAPI.md#V3AppsGuidSidecarsGet) | **Get** /v3/apps/{guid}/sidecars | List sidecars for an app
[**V3AppsGuidSidecarsPost**](DefaultAPI.md#V3AppsGuidSidecarsPost) | **Post** /v3/apps/{guid}/sidecars | Create a sidecar associated with an app
[**V3AppsGuidSshEnabledGet**](DefaultAPI.md#V3AppsGuidSshEnabledGet) | **Get** /v3/apps/{guid}/ssh_enabled | Get SSH enabled for an app
[**V3AppsGuidTasksPost**](DefaultAPI.md#V3AppsGuidTasksPost) | **Post** /v3/apps/{guid}/tasks | Create a task
[**V3AppsPost**](DefaultAPI.md#V3AppsPost) | **Post** /v3/apps | Create an app
[**V3BuildpacksGet**](DefaultAPI.md#V3BuildpacksGet) | **Get** /v3/buildpacks | List buildpacks
[**V3BuildpacksGuidPatch**](DefaultAPI.md#V3BuildpacksGuidPatch) | **Patch** /v3/buildpacks/{guid} | Update a buildpack
[**V3BuildpacksGuidUploadPost**](DefaultAPI.md#V3BuildpacksGuidUploadPost) | **Post** /v3/buildpacks/{guid}/upload | Upload buildpack bits
[**V3BuildpacksPost**](DefaultAPI.md#V3BuildpacksPost) | **Post** /v3/buildpacks | Create a buildpack
[**V3BuildsGet**](DefaultAPI.md#V3BuildsGet) | **Get** /v3/builds | List builds
[**V3BuildsGuidPatch**](DefaultAPI.md#V3BuildsGuidPatch) | **Patch** /v3/builds/{guid} | Update a build
[**V3BuildsPost**](DefaultAPI.md#V3BuildsPost) | **Post** /v3/builds | Create a build
[**V3DeploymentsGet**](DefaultAPI.md#V3DeploymentsGet) | **Get** /v3/deployments | List deployments
[**V3DeploymentsGuidActionsCancelPost**](DefaultAPI.md#V3DeploymentsGuidActionsCancelPost) | **Post** /v3/deployments/{guid}/actions/cancel | Cancel a deployment
[**V3DeploymentsGuidActionsContinuePost**](DefaultAPI.md#V3DeploymentsGuidActionsContinuePost) | **Post** /v3/deployments/{guid}/actions/continue | Continue a deployment
[**V3DeploymentsGuidGet**](DefaultAPI.md#V3DeploymentsGuidGet) | **Get** /v3/deployments/{guid} | Get a deployment
[**V3DeploymentsPost**](DefaultAPI.md#V3DeploymentsPost) | **Post** /v3/deployments | Create a deployment
[**V3DropletsGet**](DefaultAPI.md#V3DropletsGet) | **Get** /v3/droplets | List droplets
[**V3DropletsGuidGet**](DefaultAPI.md#V3DropletsGuidGet) | **Get** /v3/droplets/{guid} | Get a droplet
[**V3DropletsPost**](DefaultAPI.md#V3DropletsPost) | **Post** /v3/droplets | Create a droplet
[**V3EnvironmentVariableGroupsNameGet**](DefaultAPI.md#V3EnvironmentVariableGroupsNameGet) | **Get** /v3/environment_variable_groups/{name} | Get an environment variable group
[**V3EnvironmentVariableGroupsNamePatch**](DefaultAPI.md#V3EnvironmentVariableGroupsNamePatch) | **Patch** /v3/environment_variable_groups/{name} | Update environment variable group
[**V3FeatureFlagsGet**](DefaultAPI.md#V3FeatureFlagsGet) | **Get** /v3/feature_flags | List feature flags
[**V3FeatureFlagsNameGet**](DefaultAPI.md#V3FeatureFlagsNameGet) | **Get** /v3/feature_flags/{name} | Get a feature flag
[**V3FeatureFlagsNamePatch**](DefaultAPI.md#V3FeatureFlagsNamePatch) | **Patch** /v3/feature_flags/{name} | Update a feature flag
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
[**V3PackagesGet**](DefaultAPI.md#V3PackagesGet) | **Get** /v3/packages | List packages
[**V3PackagesGuidDelete**](DefaultAPI.md#V3PackagesGuidDelete) | **Delete** /v3/packages/{guid} | Delete a package
[**V3PackagesGuidDropletsGet**](DefaultAPI.md#V3PackagesGuidDropletsGet) | **Get** /v3/packages/{guid}/droplets | List droplets for a package
[**V3PackagesGuidGet**](DefaultAPI.md#V3PackagesGuidGet) | **Get** /v3/packages/{guid} | Get a package
[**V3PackagesGuidPatch**](DefaultAPI.md#V3PackagesGuidPatch) | **Patch** /v3/packages/{guid} | Update a package
[**V3PackagesGuidUploadPost**](DefaultAPI.md#V3PackagesGuidUploadPost) | **Post** /v3/packages/{guid}/upload | Upload package bits
[**V3PackagesPost**](DefaultAPI.md#V3PackagesPost) | **Post** /v3/packages | Create a package
[**V3ProcessesGet**](DefaultAPI.md#V3ProcessesGet) | **Get** /v3/processes | List processes
[**V3ProcessesGuidActionsScalePost**](DefaultAPI.md#V3ProcessesGuidActionsScalePost) | **Post** /v3/processes/{guid}/actions/scale | Scale a process
[**V3ProcessesGuidInstancesIndexDelete**](DefaultAPI.md#V3ProcessesGuidInstancesIndexDelete) | **Delete** /v3/processes/{guid}/instances/{index} | Terminate a process instance
[**V3ProcessesGuidPatch**](DefaultAPI.md#V3ProcessesGuidPatch) | **Patch** /v3/processes/{guid} | Update a process
[**V3ProcessesGuidSidecarsGet**](DefaultAPI.md#V3ProcessesGuidSidecarsGet) | **Get** /v3/processes/{guid}/sidecars | List sidecars for a process
[**V3ProcessesGuidStatsGet**](DefaultAPI.md#V3ProcessesGuidStatsGet) | **Get** /v3/processes/{guid}/stats | Get stats for a process
[**V3ResourceMatchesPost**](DefaultAPI.md#V3ResourceMatchesPost) | **Post** /v3/resource_matches | Create a resource match
[**V3RevisionsGuidEnvironmentVariablesGet**](DefaultAPI.md#V3RevisionsGuidEnvironmentVariablesGet) | **Get** /v3/revisions/{guid}/environment_variables | Get environment variables for a revision
[**V3RevisionsGuidPatch**](DefaultAPI.md#V3RevisionsGuidPatch) | **Patch** /v3/revisions/{guid} | Update a revision
[**V3RolesGet**](DefaultAPI.md#V3RolesGet) | **Get** /v3/roles | List roles
[**V3RolesGuidDelete**](DefaultAPI.md#V3RolesGuidDelete) | **Delete** /v3/roles/{guid} | Delete a role
[**V3RolesGuidGet**](DefaultAPI.md#V3RolesGuidGet) | **Get** /v3/roles/{guid} | Get a role
[**V3RolesPost**](DefaultAPI.md#V3RolesPost) | **Post** /v3/roles | Create a role
[**V3RoutesGet**](DefaultAPI.md#V3RoutesGet) | **Get** /v3/routes | List routes
[**V3RoutesGuidGet**](DefaultAPI.md#V3RoutesGuidGet) | **Get** /v3/routes/{guid} | Get a route
[**V3RoutesPost**](DefaultAPI.md#V3RoutesPost) | **Post** /v3/routes | Create a route
[**V3SecurityGroupsGet**](DefaultAPI.md#V3SecurityGroupsGet) | **Get** /v3/security_groups | List security groups
[**V3SecurityGroupsGuidDelete**](DefaultAPI.md#V3SecurityGroupsGuidDelete) | **Delete** /v3/security_groups/{guid} | Delete a security group
[**V3SecurityGroupsGuidGet**](DefaultAPI.md#V3SecurityGroupsGuidGet) | **Get** /v3/security_groups/{guid} | Get a security group
[**V3SecurityGroupsGuidPatch**](DefaultAPI.md#V3SecurityGroupsGuidPatch) | **Patch** /v3/security_groups/{guid} | Update a security group
[**V3SecurityGroupsPost**](DefaultAPI.md#V3SecurityGroupsPost) | **Post** /v3/security_groups | Create a security group
[**V3ServiceBrokersGet**](DefaultAPI.md#V3ServiceBrokersGet) | **Get** /v3/service_brokers | List service brokers
[**V3ServiceBrokersGuidDelete**](DefaultAPI.md#V3ServiceBrokersGuidDelete) | **Delete** /v3/service_brokers/{guid} | Delete a service broker
[**V3ServiceBrokersGuidGet**](DefaultAPI.md#V3ServiceBrokersGuidGet) | **Get** /v3/service_brokers/{guid} | Get a service broker
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
[**V3ServiceInstancesGuidPatch**](DefaultAPI.md#V3ServiceInstancesGuidPatch) | **Patch** /v3/service_instances/{guid} | Update a service instance
[**V3ServiceInstancesGuidRelationshipsSharedSpacesPost**](DefaultAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesPost) | **Post** /v3/service_instances/{guid}/relationships/shared_spaces | Share a service instance to other spaces
[**V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete**](DefaultAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesSpaceGuidDelete) | **Delete** /v3/service_instances/{guid}/relationships/shared_spaces/{space_guid} | Unshare a service instance from another space
[**V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet**](DefaultAPI.md#V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet) | **Get** /v3/service_instances/{guid}/relationships/shared_spaces/usage_summary | Get usage summary in shared spaces
[**V3ServiceOfferingsGet**](DefaultAPI.md#V3ServiceOfferingsGet) | **Get** /v3/service_offerings | List service offerings
[**V3ServiceOfferingsGuidDelete**](DefaultAPI.md#V3ServiceOfferingsGuidDelete) | **Delete** /v3/service_offerings/{guid} | Delete a service offering
[**V3ServiceOfferingsGuidGet**](DefaultAPI.md#V3ServiceOfferingsGuidGet) | **Get** /v3/service_offerings/{guid} | Get a service offering
[**V3ServiceOfferingsGuidPatch**](DefaultAPI.md#V3ServiceOfferingsGuidPatch) | **Patch** /v3/service_offerings/{guid} | Update a service offering
[**V3ServiceOfferingsPost**](DefaultAPI.md#V3ServiceOfferingsPost) | **Post** /v3/service_offerings | Create a service offering
[**V3ServicePlansGet**](DefaultAPI.md#V3ServicePlansGet) | **Get** /v3/service_plans | List service plans
[**V3ServicePlansGuidDelete**](DefaultAPI.md#V3ServicePlansGuidDelete) | **Delete** /v3/service_plans/{guid} | Delete a service plan
[**V3ServicePlansGuidGet**](DefaultAPI.md#V3ServicePlansGuidGet) | **Get** /v3/service_plans/{guid} | Get a service plan
[**V3ServicePlansGuidPatch**](DefaultAPI.md#V3ServicePlansGuidPatch) | **Patch** /v3/service_plans/{guid} | Update a service plan
[**V3ServicePlansGuidVisibilityGet**](DefaultAPI.md#V3ServicePlansGuidVisibilityGet) | **Get** /v3/service_plans/{guid}/visibility | Get a service plan visibility
[**V3ServicePlansGuidVisibilityOrganizationGuidDelete**](DefaultAPI.md#V3ServicePlansGuidVisibilityOrganizationGuidDelete) | **Delete** /v3/service_plans/{guid}/visibility/{organization_guid} | Remove organization from a service plan visibility
[**V3ServicePlansGuidVisibilityPatch**](DefaultAPI.md#V3ServicePlansGuidVisibilityPatch) | **Patch** /v3/service_plans/{guid}/visibility | Update a service plan visibility
[**V3ServicePlansGuidVisibilityPost**](DefaultAPI.md#V3ServicePlansGuidVisibilityPost) | **Post** /v3/service_plans/{guid}/visibility | Apply a service plan visibility
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
[**V3SidecarsGuidDelete**](DefaultAPI.md#V3SidecarsGuidDelete) | **Delete** /v3/sidecars/{guid} | Delete a sidecar
[**V3SidecarsGuidGet**](DefaultAPI.md#V3SidecarsGuidGet) | **Get** /v3/sidecars/{guid} | Get a sidecar
[**V3SidecarsGuidPatch**](DefaultAPI.md#V3SidecarsGuidPatch) | **Patch** /v3/sidecars/{guid} | Update a sidecar
[**V3SpaceQuotasGet**](DefaultAPI.md#V3SpaceQuotasGet) | **Get** /v3/space_quotas | List space quotas
[**V3SpaceQuotasGuidDelete**](DefaultAPI.md#V3SpaceQuotasGuidDelete) | **Delete** /v3/space_quotas/{guid} | Delete a space quota
[**V3SpaceQuotasGuidGet**](DefaultAPI.md#V3SpaceQuotasGuidGet) | **Get** /v3/space_quotas/{guid} | Get a space quota
[**V3SpaceQuotasGuidPatch**](DefaultAPI.md#V3SpaceQuotasGuidPatch) | **Patch** /v3/space_quotas/{guid} | Update a space quota
[**V3SpaceQuotasPost**](DefaultAPI.md#V3SpaceQuotasPost) | **Post** /v3/space_quotas | Create a space quota
[**V3SpaceQuotasQuotaGuidRelationshipsSpacesPost**](DefaultAPI.md#V3SpaceQuotasQuotaGuidRelationshipsSpacesPost) | **Post** /v3/space_quotas/{quota_guid}/relationships/spaces | Apply a space quota to spaces
[**V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete**](DefaultAPI.md#V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete) | **Delete** /v3/space_quotas/{quota_guid}/relationships/spaces/{space_guid} | Remove a space quota from a space
[**V3SpacesGet**](DefaultAPI.md#V3SpacesGet) | **Get** /v3/spaces | List spaces
[**V3SpacesGuidDelete**](DefaultAPI.md#V3SpacesGuidDelete) | **Delete** /v3/spaces/{guid} | Delete a space
[**V3SpacesGuidFeaturesGet**](DefaultAPI.md#V3SpacesGuidFeaturesGet) | **Get** /v3/spaces/{guid}/features | List space features
[**V3SpacesGuidFeaturesNameGet**](DefaultAPI.md#V3SpacesGuidFeaturesNameGet) | **Get** /v3/spaces/{guid}/features/{name} | Get a space feature
[**V3SpacesGuidFeaturesPatch**](DefaultAPI.md#V3SpacesGuidFeaturesPatch) | **Patch** /v3/spaces/{guid}/features | Update space features
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



## ApplyOrganizationQuotaToOrganizations

> ApplyOrganizationQuotaToOrganizations201Response ApplyOrganizationQuotaToOrganizations(ctx, quotaGuid).ApplyOrganizationQuotaToOrganizationsRequest(applyOrganizationQuotaToOrganizationsRequest).Execute()

Apply an organization quota to organizations

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
	quotaGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	applyOrganizationQuotaToOrganizationsRequest := *openapiclient.NewApplyOrganizationQuotaToOrganizationsRequest() // ApplyOrganizationQuotaToOrganizationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApplyOrganizationQuotaToOrganizations(context.Background(), quotaGuid).ApplyOrganizationQuotaToOrganizationsRequest(applyOrganizationQuotaToOrganizationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApplyOrganizationQuotaToOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyOrganizationQuotaToOrganizations`: ApplyOrganizationQuotaToOrganizations201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApplyOrganizationQuotaToOrganizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyOrganizationQuotaToOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyOrganizationQuotaToOrganizationsRequest** | [**ApplyOrganizationQuotaToOrganizationsRequest**](ApplyOrganizationQuotaToOrganizationsRequest.md) |  | 

### Return type

[**ApplyOrganizationQuotaToOrganizations201Response**](ApplyOrganizationQuotaToOrganizations201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignDefaultIsolationSegment

> map[string]interface{} AssignDefaultIsolationSegment(ctx, guid).V3AppsPostRequestRelationshipsSpace(v3AppsPostRequestRelationshipsSpace).Execute()

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
	v3AppsPostRequestRelationshipsSpace := *openapiclient.NewV3AppsPostRequestRelationshipsSpace() // V3AppsPostRequestRelationshipsSpace |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AssignDefaultIsolationSegment(context.Background(), guid).V3AppsPostRequestRelationshipsSpace(v3AppsPostRequestRelationshipsSpace).Execute()
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

 **v3AppsPostRequestRelationshipsSpace** | [**V3AppsPostRequestRelationshipsSpace**](V3AppsPostRequestRelationshipsSpace.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationQuota

> OrganizationQuota CreateOrganizationQuota(ctx).CreateOrganizationQuotaRequest(createOrganizationQuotaRequest).Execute()

Create an organization quota

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
	createOrganizationQuotaRequest := *openapiclient.NewCreateOrganizationQuotaRequest() // CreateOrganizationQuotaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateOrganizationQuota(context.Background()).CreateOrganizationQuotaRequest(createOrganizationQuotaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationQuotaRequest** | [**CreateOrganizationQuotaRequest**](CreateOrganizationQuotaRequest.md) |  | 

### Return type

[**OrganizationQuota**](OrganizationQuota.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DeleteOrganizationQuota(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationQuota`: DeleteOrganizationQuota202Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DeleteOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOrganizationQuota202Response**](DeleteOrganizationQuota202Response.md)

### Authorization

No authorization required

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


## GetOrganizationQuota

> OrganizationQuota GetOrganizationQuota(ctx, guid).Execute()

Get an organization quota

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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetOrganizationQuota(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationQuota**](OrganizationQuota.md)

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


## ListOrganizationQuotas

> ListOrganizationQuotas200Response ListOrganizationQuotas(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List organization quotas

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
	guids := []string{"Inner_example"} // []string |  (optional)
	names := []string{"Inner_example"} // []string |  (optional)
	organizationGuids := []string{"Inner_example"} // []string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	createdAts := time.Now() // time.Time |  (optional)
	updatedAts := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListOrganizationQuotas(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListOrganizationQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationQuotas`: ListOrganizationQuotas200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListOrganizationQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** |  | 
 **names** | **[]string** |  | 
 **organizationGuids** | **[]string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **createdAts** | **time.Time** |  | 
 **updatedAts** | **time.Time** |  | 

### Return type

[**ListOrganizationQuotas200Response**](ListOrganizationQuotas200Response.md)

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationQuota

> OrganizationQuota UpdateOrganizationQuota(ctx, guid).UpdateOrganizationQuotaRequest(updateOrganizationQuotaRequest).Execute()

Update an organization quota

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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateOrganizationQuotaRequest := *openapiclient.NewUpdateOrganizationQuotaRequest() // UpdateOrganizationQuotaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateOrganizationQuota(context.Background(), guid).UpdateOrganizationQuotaRequest(updateOrganizationQuotaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateOrganizationQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationQuota`: OrganizationQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateOrganizationQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationQuotaRequest** | [**UpdateOrganizationQuotaRequest**](UpdateOrganizationQuotaRequest.md) |  | 

### Return type

[**OrganizationQuota**](OrganizationQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AdminActionsClearBuildpackCachePost

> V3AdminActionsClearBuildpackCachePost(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AdminActionsClearBuildpackCachePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AdminActionsClearBuildpackCachePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3AdminActionsClearBuildpackCachePostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGet

> App V3AppsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

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
	page := int32(56) // int32 | Page number for pagination. (optional)
	perPage := int32(56) // int32 | Number of results per page. (optional)
	orderBy := "orderBy_example" // string | Field by which to order results. (optional)
	names := "names_example" // string | Filter results by app names. (optional)
	guids := "guids_example" // string | Filter results by app GUIDs. (optional)
	organizationGuids := "organizationGuids_example" // string | Filter results by organization GUIDs. (optional)
	spaceGuids := "spaceGuids_example" // string | Filter results by space GUIDs. (optional)
	stacks := "stacks_example" // string | Filter results by stack names. (optional)
	include := "include_example" // string | Include related resources in the response. (optional)
	lifecycleType := "lifecycleType_example" // string | Filter results by lifecycle type. (optional)
	labelSelector := "labelSelector_example" // string | Filter results by label selector. (optional)
	createdAts := "createdAts_example" // string | Filter results by creation timestamps. (optional)
	updatedAts := "updatedAts_example" // string | Filter results by update timestamps. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).Names(names).Guids(guids).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Stacks(stacks).Include(include).LifecycleType(lifecycleType).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGet`: App
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number for pagination. | 
 **perPage** | **int32** | Number of results per page. | 
 **orderBy** | **string** | Field by which to order results. | 
 **names** | **string** | Filter results by app names. | 
 **guids** | **string** | Filter results by app GUIDs. | 
 **organizationGuids** | **string** | Filter results by organization GUIDs. | 
 **spaceGuids** | **string** | Filter results by space GUIDs. | 
 **stacks** | **string** | Filter results by stack names. | 
 **include** | **string** | Include related resources in the response. | 
 **lifecycleType** | **string** | Filter results by lifecycle type. | 
 **labelSelector** | **string** | Filter results by label selector. | 
 **createdAts** | **string** | Filter results by creation timestamps. | 
 **updatedAts** | **string** | Filter results by update timestamps. | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsClearBuildpackCachePost

> V3AppsGuidActionsClearBuildpackCachePost(ctx, guid).Execute()

Clear buildpack cache for application

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
	r, err := apiClient.DefaultAPI.V3AppsGuidActionsClearBuildpackCachePost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidActionsClearBuildpackCachePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3AppsGuidActionsClearBuildpackCachePostRequest struct via the builder pattern


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


## V3AppsGuidActionsRestartPost

> V3AppsGuidActionsRestartPost200Response V3AppsGuidActionsRestartPost(ctx, guid).Execute()

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidActionsRestartPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidActionsRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsRestartPost`: V3AppsGuidActionsRestartPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidActionsRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidActionsRestartPost200Response**](V3AppsGuidActionsRestartPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsStartPost

> V3AppsGuidActionsRestartPost200Response V3AppsGuidActionsStartPost(ctx, guid).Execute()

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidActionsStartPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidActionsStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsStartPost`: V3AppsGuidActionsRestartPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidActionsStartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidActionsRestartPost200Response**](V3AppsGuidActionsRestartPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidActionsStopPost

> V3AppsGuidActionsRestartPost200Response V3AppsGuidActionsStopPost(ctx, guid).Execute()

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
	guid := "guid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidActionsStopPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidActionsStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidActionsStopPost`: V3AppsGuidActionsRestartPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidActionsStopPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidActionsStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidActionsRestartPost200Response**](V3AppsGuidActionsRestartPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidBuildsGet

> V3AppsGuidBuildsGet(ctx, guid).Execute()

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
	guid := "guid_example" // string | The GUID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AppsGuidBuildsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
	guid := "guid_example" // string | The GUID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AppsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidDropletsCurrentGet

> V3AppsGuidDropletsCurrentGet200Response V3AppsGuidDropletsCurrentGet(ctx, guid).Execute()

Get current droplet

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidDropletsCurrentGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidDropletsCurrentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidDropletsCurrentGet`: V3AppsGuidDropletsCurrentGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidDropletsCurrentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidDropletsCurrentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidDropletsCurrentGet200Response**](V3AppsGuidDropletsCurrentGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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


## V3AppsGuidEnvGet

> V3AppsGuidEnvGet200Response V3AppsGuidEnvGet(ctx, guid).Execute()

Get environment for an app

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidEnvGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidEnvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvGet`: V3AppsGuidEnvGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidEnvGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3AppsGuidEnvGet200Response**](V3AppsGuidEnvGet200Response.md)

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

Get environment variables for an app

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidEnvironmentVariablesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidEnvironmentVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesGet`: V3AppsGuidEnvironmentVariablesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidEnvironmentVariablesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

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

> V3AppsGuidEnvironmentVariablesGet200Response V3AppsGuidEnvironmentVariablesPatch(ctx, guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()

Update environment variables for an app

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
	v3AppsGuidEnvironmentVariablesPatchRequest := *openapiclient.NewV3AppsGuidEnvironmentVariablesPatchRequest() // V3AppsGuidEnvironmentVariablesPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidEnvironmentVariablesPatch(context.Background(), guid).V3AppsGuidEnvironmentVariablesPatchRequest(v3AppsGuidEnvironmentVariablesPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidEnvironmentVariablesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidEnvironmentVariablesPatch`: V3AppsGuidEnvironmentVariablesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidEnvironmentVariablesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidEnvironmentVariablesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidEnvironmentVariablesPatchRequest** | [**V3AppsGuidEnvironmentVariablesPatchRequest**](V3AppsGuidEnvironmentVariablesPatchRequest.md) |  | 

### Return type

[**V3AppsGuidEnvironmentVariablesGet200Response**](V3AppsGuidEnvironmentVariablesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesGet

> V3AppsGuidFeaturesGet(ctx, guid).Execute()

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
	guid := "guid_example" // string | The GUID of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AppsGuidFeaturesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNameGet

> V3AppsGuidFeaturesNameGet(ctx, guid, name).Execute()

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
	guid := "guid_example" // string | The GUID of the app.
	name := "name_example" // string | The name of the feature.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AppsGuidFeaturesNameGet(context.Background(), guid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidFeaturesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 
**name** | **string** | The name of the feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidFeaturesNamePatch

> V3AppsGuidFeaturesNamePatch(ctx, guid, name).V3AppsGuidFeaturesNamePatchRequest(v3AppsGuidFeaturesNamePatchRequest).Execute()

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
	guid := "guid_example" // string | The GUID of the app.
	name := "name_example" // string | The name of the feature.
	v3AppsGuidFeaturesNamePatchRequest := *openapiclient.NewV3AppsGuidFeaturesNamePatchRequest() // V3AppsGuidFeaturesNamePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3AppsGuidFeaturesNamePatch(context.Background(), guid, name).V3AppsGuidFeaturesNamePatchRequest(v3AppsGuidFeaturesNamePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidFeaturesNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 
**name** | **string** | The name of the feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidFeaturesNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v3AppsGuidFeaturesNamePatchRequest** | [**V3AppsGuidFeaturesNamePatchRequest**](V3AppsGuidFeaturesNamePatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidGet

> App V3AppsGuidGet(ctx, guid).Include(include).Execute()

Retrieve a specific app



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
	guid := "guid_example" // string | The GUID of the app.
	include := []string{"Include_example"} // []string | Include related resources in the response; valid values are space and space.organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidGet`: App
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Include related resources in the response; valid values are space and space.organization. | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

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
	guid := "guid_example" // string | The GUID of the app.
	v3AppsGuidPatchRequest := *openapiclient.NewV3AppsGuidPatchRequest() // V3AppsGuidPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidPatch(context.Background(), guid).V3AppsGuidPatchRequest(v3AppsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPatch`: App
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidPatchRequest** | [**V3AppsGuidPatchRequest**](V3AppsGuidPatchRequest.md) |  | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidPermissionsGet

> V3AppsGuidPermissionsGet200Response V3AppsGuidPermissionsGet(ctx, guid).Execute()

Get permissions for an app

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidPermissionsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidPermissionsGet`: V3AppsGuidPermissionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

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

> ProcessList V3AppsGuidProcessesGet(ctx, guid).Guids(guids).Types(types).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List processes for app

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
	types := "types_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidProcessesGet(context.Background(), guid).Guids(guids).Types(types).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidProcessesGet`: ProcessList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidProcessesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **guids** | **string** |  | 
 **types** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**ProcessList**](ProcessList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRelationshipsCurrentDropletGet

> V3AppsGuidRelationshipsCurrentDropletGet200Response V3AppsGuidRelationshipsCurrentDropletGet(ctx, guid).Execute()

Get current droplet association for an app

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidRelationshipsCurrentDropletGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidRelationshipsCurrentDropletGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRelationshipsCurrentDropletGet`: V3AppsGuidRelationshipsCurrentDropletGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidRelationshipsCurrentDropletGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidRoutesGet

> RouteList V3AppsGuidRoutesGet(ctx, guid).DomainGuids(domainGuids).Hosts(hosts).OrganizationGuids(organizationGuids).Paths(paths).Ports(ports).SpaceGuids(spaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()

Retrieve all routes for an app



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
	guid := "guid_example" // string | The app GUID
	domainGuids := []string{"Inner_example"} // []string | Comma-delimited list of domain guids to filter by (optional)
	hosts := []string{"Inner_example"} // []string | Comma-delimited list of hostnames to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	paths := []string{"Inner_example"} // []string | Comma-delimited list of paths to filter by (e.g. /path1,/path2) (optional)
	ports := []int32{int32(123)} // []int32 | Comma-delimited list of ports to filter by (e.g. 3306,5432) (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidRoutesGet(context.Background(), guid).DomainGuids(domainGuids).Hosts(hosts).OrganizationGuids(organizationGuids).Paths(paths).Ports(ports).SpaceGuids(spaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidRoutesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidRoutesGet`: RouteList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidRoutesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The app GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidRoutesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainGuids** | **[]string** | Comma-delimited list of domain guids to filter by | 
 **hosts** | **[]string** | Comma-delimited list of hostnames to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **paths** | **[]string** | Comma-delimited list of paths to filter by (e.g. /path1,/path2) | 
 **ports** | **[]int32** | Comma-delimited list of ports to filter by (e.g. 3306,5432) | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 

### Return type

[**RouteList**](RouteList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidSidecarsGet

> V3AppsGuidSidecarsGet200Response V3AppsGuidSidecarsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List sidecars for an app



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
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidSidecarsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidSidecarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSidecarsGet`: V3AppsGuidSidecarsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidSidecarsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSidecarsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3AppsGuidSidecarsGet200Response**](V3AppsGuidSidecarsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidSidecarsPost

> Sidecar V3AppsGuidSidecarsPost(ctx, guid).V3AppsGuidSidecarsPostRequest(v3AppsGuidSidecarsPostRequest).Execute()

Create a sidecar associated with an app



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
	v3AppsGuidSidecarsPostRequest := *openapiclient.NewV3AppsGuidSidecarsPostRequest("Command_example", "Name_example", []string{"ProcessTypes_example"}) // V3AppsGuidSidecarsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidSidecarsPost(context.Background(), guid).V3AppsGuidSidecarsPostRequest(v3AppsGuidSidecarsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidSidecarsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSidecarsPost`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidSidecarsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3AppsGuidSidecarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3AppsGuidSidecarsPostRequest** | [**V3AppsGuidSidecarsPostRequest**](V3AppsGuidSidecarsPostRequest.md) |  | 

### Return type

[**Sidecar**](Sidecar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3AppsGuidSshEnabledGet

> V3AppsGuidSshEnabledGet200Response V3AppsGuidSshEnabledGet(ctx, guid).Execute()

Get SSH enabled for an app

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
	resp, r, err := apiClient.DefaultAPI.V3AppsGuidSshEnabledGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsGuidSshEnabledGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsGuidSshEnabledGet`: V3AppsGuidSshEnabledGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsGuidSshEnabledGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
	v3AppsPostRequest := *openapiclient.NewV3AppsPostRequest("Name_example", "SpaceGuid_example") // V3AppsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3AppsPost(context.Background()).V3AppsPostRequest(v3AppsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3AppsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3AppsPost`: App
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3AppsPost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGet

> V3BuildpacksGet(ctx).Execute()

List buildpacks



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
	r, err := apiClient.DefaultAPI.V3BuildpacksGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildpacksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGuidPatch

> V3BuildpacksGuidPatch(ctx, guid).Body(body).Execute()

Update a buildpack



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
	guid := "guid_example" // string | The GUID of the buildpack.
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3BuildpacksGuidPatch(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildpacksGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksGuidUploadPost

> V3BuildpacksGuidUploadPost(ctx, guid).Bits(bits).Execute()

Upload buildpack bits



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
	guid := "guid_example" // string | The GUID of the buildpack.
	bits := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3BuildpacksGuidUploadPost(context.Background(), guid).Bits(bits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildpacksGuidUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the buildpack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksGuidUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bits** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildpacksPost

> V3BuildpacksPost(ctx).Body(body).Execute()

Create a buildpack



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3BuildpacksPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildpacksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildpacksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsGet

> V3BuildsGet(ctx).Execute()

List builds



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
	r, err := apiClient.DefaultAPI.V3BuildsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsGuidPatch

> V3BuildsGuidPatch(ctx, guid).Body(body).Execute()

Update a build



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
	guid := "guid_example" // string | The GUID of the build.
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3BuildsGuidPatch(context.Background(), guid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3BuildsPost

> V3BuildsPost(ctx).Body(body).Execute()

Create a build



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3BuildsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3BuildsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3BuildsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGet

> V3DeploymentsGet(ctx).Execute()

List deployments



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
	r, err := apiClient.DefaultAPI.V3DeploymentsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidActionsCancelPost

> V3DeploymentsGuidActionsCancelPost(ctx, guid).Execute()

Cancel a deployment



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
	guid := "guid_example" // string | The GUID of the deployment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3DeploymentsGuidActionsCancelPost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DeploymentsGuidActionsCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidActionsCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidActionsContinuePost

> V3DeploymentsGuidActionsContinuePost(ctx, guid).Execute()

Continue a deployment



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
	guid := "guid_example" // string | The GUID of the deployment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3DeploymentsGuidActionsContinuePost(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DeploymentsGuidActionsContinuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidActionsContinuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsGuidGet

> V3DeploymentsGuidGet(ctx, guid).Execute()

Get a deployment



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
	guid := "guid_example" // string | The GUID of the deployment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3DeploymentsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DeploymentsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The GUID of the deployment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeploymentsPost

> V3DeploymentsPost(ctx).Body(body).Execute()

Create a deployment



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3DeploymentsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3DeploymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3DeploymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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


## V3DropletsPost

> Droplet V3DropletsPost(ctx).V3DropletsPostRequest(v3DropletsPostRequest).Execute()

Create a droplet



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3DropletsPost(context.Background()).V3DropletsPostRequest(v3DropletsPostRequest).Execute()
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

### Return type

[**Droplet**](Droplet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnvironmentVariableGroupsNameGet

> EnvironmentVariableGroup V3EnvironmentVariableGroupsNameGet(ctx, name).Execute()

Get an environment variable group



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
	name := "name_example" // string | The name of the environment variable group (running or staging)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3EnvironmentVariableGroupsNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3EnvironmentVariableGroupsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnvironmentVariableGroupsNameGet`: EnvironmentVariableGroup
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3EnvironmentVariableGroupsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the environment variable group (running or staging) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnvironmentVariableGroupsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentVariableGroup**](EnvironmentVariableGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3EnvironmentVariableGroupsNamePatch

> EnvironmentVariableGroup V3EnvironmentVariableGroupsNamePatch(ctx, name).V3EnvironmentVariableGroupsNamePatchRequest(v3EnvironmentVariableGroupsNamePatchRequest).Execute()

Update environment variable group



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
	name := "name_example" // string | The name of the environment variable group (running or staging)
	v3EnvironmentVariableGroupsNamePatchRequest := *openapiclient.NewV3EnvironmentVariableGroupsNamePatchRequest() // V3EnvironmentVariableGroupsNamePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3EnvironmentVariableGroupsNamePatch(context.Background(), name).V3EnvironmentVariableGroupsNamePatchRequest(v3EnvironmentVariableGroupsNamePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3EnvironmentVariableGroupsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3EnvironmentVariableGroupsNamePatch`: EnvironmentVariableGroup
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3EnvironmentVariableGroupsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the environment variable group (running or staging) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3EnvironmentVariableGroupsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3EnvironmentVariableGroupsNamePatchRequest** | [**V3EnvironmentVariableGroupsNamePatchRequest**](V3EnvironmentVariableGroupsNamePatchRequest.md) |  | 

### Return type

[**EnvironmentVariableGroup**](EnvironmentVariableGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3FeatureFlagsGet

> V3FeatureFlagsGet200Response V3FeatureFlagsGet(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).UpdatedAts(updatedAts).Execute()

List feature flags



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
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to name ascending; prepend with - to sort descending. (optional)
	updatedAts := time.Now() // time.Time | Timestamp to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3FeatureFlagsGet(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3FeatureFlagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsGet`: V3FeatureFlagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3FeatureFlagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by. Defaults to name ascending; prepend with - to sort descending. | 
 **updatedAts** | **time.Time** | Timestamp to filter by | 

### Return type

[**V3FeatureFlagsGet200Response**](V3FeatureFlagsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3FeatureFlagsNameGet

> FeatureFlag V3FeatureFlagsNameGet(ctx, name).Execute()

Get a feature flag



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
	name := "name_example" // string | The name of the feature flag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3FeatureFlagsNameGet(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3FeatureFlagsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsNameGet`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3FeatureFlagsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3FeatureFlagsNamePatch

> FeatureFlag V3FeatureFlagsNamePatch(ctx, name).V3FeatureFlagsNamePatchRequest(v3FeatureFlagsNamePatchRequest).Execute()

Update a feature flag



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
	name := "name_example" // string | The name of the feature flag
	v3FeatureFlagsNamePatchRequest := *openapiclient.NewV3FeatureFlagsNamePatchRequest() // V3FeatureFlagsNamePatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3FeatureFlagsNamePatch(context.Background(), name).V3FeatureFlagsNamePatchRequest(v3FeatureFlagsNamePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3FeatureFlagsNamePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3FeatureFlagsNamePatch`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3FeatureFlagsNamePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the feature flag | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3FeatureFlagsNamePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3FeatureFlagsNamePatchRequest** | [**V3FeatureFlagsNamePatchRequest**](V3FeatureFlagsNamePatchRequest.md) |  | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGet

> V3PackagesGet200Response V3PackagesGet(ctx).Guids(guids).States(states).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List packages



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
	guids := []string{"Inner_example"} // []string | Comma-delimited list of package guids to filter by (optional)
	states := []string{"Inner_example"} // []string | Comma-delimited list of package states to filter by (optional)
	types := []string{"Inner_example"} // []string | Comma-delimited list of package types to filter by (optional)
	appGuids := []string{"Inner_example"} // []string | Comma-delimited list of app guids to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	page := int32(56) // int32 | Page to display (optional)
	perPage := int32(56) // int32 | Number of results per page (optional)
	orderBy := "orderBy_example" // string | Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := []time.Time{time.Now()} // []time.Time | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGet(context.Background()).Guids(guids).States(states).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGet`: V3PackagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of package guids to filter by | 
 **states** | **[]string** | Comma-delimited list of package states to filter by | 
 **types** | **[]string** | Comma-delimited list of package types to filter by | 
 **appGuids** | **[]string** | Comma-delimited list of app guids to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **page** | **int32** | Page to display | 
 **perPage** | **int32** | Number of results per page | 
 **orderBy** | **string** | Value to sort by; defaults to ascending. Prepend with - to sort descending. Valid values are created_at, updated_at | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | [**[]time.Time**](time.Time.md) | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**V3PackagesGet200Response**](V3PackagesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidDelete

> V3PackagesGuidDelete202Response V3PackagesGuidDelete(ctx, guid).Execute()

Delete a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidDelete`: V3PackagesGuidDelete202Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGuidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3PackagesGuidDelete202Response**](V3PackagesGuidDelete202Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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


## V3PackagesGuidGet

> Package V3PackagesGuidGet(ctx, guid).Execute()

Get a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidGet`: Package
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Package**](Package.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidPatch

> Package V3PackagesGuidPatch(ctx, guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()

Update a package



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	v3PackagesGuidPatchRequest := *openapiclient.NewV3PackagesGuidPatchRequest() // V3PackagesGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGuidPatch(context.Background(), guid).V3PackagesGuidPatchRequest(v3PackagesGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidPatch`: Package
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3PackagesGuidPatchRequest** | [**V3PackagesGuidPatchRequest**](V3PackagesGuidPatchRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesGuidUploadPost

> Package V3PackagesGuidUploadPost(ctx, guid).Bits(bits).Resources(resources).Execute()

Upload package bits



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
	guid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	bits := os.NewFile(1234, "some_file") // *os.File | A binary zip file containing the package bits (optional)
	resources := []openapiclient.V3PackagesGuidUploadPostRequestResourcesInner{*openapiclient.NewV3PackagesGuidUploadPostRequestResourcesInner()} // []V3PackagesGuidUploadPostRequestResourcesInner |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesGuidUploadPost(context.Background(), guid).Bits(bits).Resources(resources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesGuidUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesGuidUploadPost`: Package
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesGuidUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesGuidUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bits** | ***os.File** | A binary zip file containing the package bits | 
 **resources** | [**[]V3PackagesGuidUploadPostRequestResourcesInner**](V3PackagesGuidUploadPostRequestResourcesInner.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3PackagesPost

> Package V3PackagesPost(ctx).V3PackagesPostRequest(v3PackagesPostRequest).Execute()

Create a package



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
	v3PackagesPostRequest := *openapiclient.NewV3PackagesPostRequest() // V3PackagesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3PackagesPost(context.Background()).V3PackagesPostRequest(v3PackagesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3PackagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3PackagesPost`: Package
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3PackagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3PackagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3PackagesPostRequest** | [**V3PackagesPostRequest**](V3PackagesPostRequest.md) |  | 

### Return type

[**Package**](Package.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGet

> ProcessList V3ProcessesGet(ctx).Guids(guids).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List processes

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
	types := "types_example" // string |  (optional)
	appGuids := "appGuids_example" // string |  (optional)
	spaceGuids := "spaceGuids_example" // string |  (optional)
	organizationGuids := "organizationGuids_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	labelSelector := "labelSelector_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ProcessesGet(context.Background()).Guids(guids).Types(types).AppGuids(appGuids).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGet`: ProcessList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ProcessesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** |  | 
 **types** | **string** |  | 
 **appGuids** | **string** |  | 
 **spaceGuids** | **string** |  | 
 **organizationGuids** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **labelSelector** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**ProcessList**](ProcessList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidActionsScalePost

> Process V3ProcessesGuidActionsScalePost(ctx, guid).ProcessScale(processScale).Execute()

Scale a process

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
	processScale := *openapiclient.NewProcessScale() // ProcessScale |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ProcessesGuidActionsScalePost(context.Background(), guid).ProcessScale(processScale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGuidActionsScalePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidActionsScalePost`: Process
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ProcessesGuidActionsScalePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidActionsScalePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processScale** | [**ProcessScale**](ProcessScale.md) |  | 

### Return type

[**Process**](Process.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidInstancesIndexDelete

> V3ProcessesGuidInstancesIndexDelete(ctx, guid, index).Execute()

Terminate a process instance

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
	index := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ProcessesGuidInstancesIndexDelete(context.Background(), guid, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGuidInstancesIndexDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 
**index** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidInstancesIndexDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidPatch

> Process V3ProcessesGuidPatch(ctx, guid).ProcessUpdate(processUpdate).Execute()

Update a process

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
	processUpdate := *openapiclient.NewProcessUpdate() // ProcessUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ProcessesGuidPatch(context.Background(), guid).ProcessUpdate(processUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidPatch`: Process
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ProcessesGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processUpdate** | [**ProcessUpdate**](ProcessUpdate.md) |  | 

### Return type

[**Process**](Process.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidSidecarsGet

> V3AppsGuidSidecarsGet200Response V3ProcessesGuidSidecarsGet(ctx, guid).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List sidecars for a process



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
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ProcessesGuidSidecarsGet(context.Background(), guid).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGuidSidecarsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidSidecarsGet`: V3AppsGuidSidecarsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ProcessesGuidSidecarsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidSidecarsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3AppsGuidSidecarsGet200Response**](V3AppsGuidSidecarsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ProcessesGuidStatsGet

> []ProcessStats V3ProcessesGuidStatsGet(ctx, guid).Execute()

Get stats for a process

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
	resp, r, err := apiClient.DefaultAPI.V3ProcessesGuidStatsGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ProcessesGuidStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ProcessesGuidStatsGet`: []ProcessStats
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ProcessesGuidStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ProcessesGuidStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProcessStats**](ProcessStats.md)

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RevisionsGuidPatch

> Revision V3RevisionsGuidPatch(ctx, guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()

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
	v3RevisionsGuidPatchRequest := *openapiclient.NewV3RevisionsGuidPatchRequest() // V3RevisionsGuidPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RevisionsGuidPatch(context.Background(), guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()
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

 **v3RevisionsGuidPatchRequest** | [**V3RevisionsGuidPatchRequest**](V3RevisionsGuidPatchRequest.md) |  | 

### Return type

[**Revision**](Revision.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesGet

> RolesList V3RolesGet(ctx).Guids(guids).Types(types).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).UserGuids(userGuids).Page(page).PerPage(perPage).OrderBy(orderBy).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

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
	guids := []string{"Inner_example"} // []string | Comma-delimited list of role guids to filter by (optional)
	types := []string{"Inner_example"} // []string | Comma-delimited list of role types to filter by (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	organizationGuids := []string{"Inner_example"} // []string | Comma-delimited list of organization guids to filter by (optional)
	userGuids := []string{"Inner_example"} // []string | Comma-delimited list of user guids to filter by (optional)
	page := int32(56) // int32 | Page to display, valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page, valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by, defaults to ascending, prepend with - to sort descending (optional)
	include := []string{"Inner_example"} // []string | Optionally include a list of unique related resources in the response (optional)
	createdAts := []string{"Inner_example"} // []string | Timestamp to filter by, supports filtering with relational operators (optional)
	updatedAts := []string{"Inner_example"} // []string | Timestamp to filter by, supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RolesGet(context.Background()).Guids(guids).Types(types).SpaceGuids(spaceGuids).OrganizationGuids(organizationGuids).UserGuids(userGuids).Page(page).PerPage(perPage).OrderBy(orderBy).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesGet`: RolesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of role guids to filter by | 
 **types** | **[]string** | Comma-delimited list of role types to filter by | 
 **spaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **organizationGuids** | **[]string** | Comma-delimited list of organization guids to filter by | 
 **userGuids** | **[]string** | Comma-delimited list of user guids to filter by | 
 **page** | **int32** | Page to display, valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page, valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by, defaults to ascending, prepend with - to sort descending | 
 **include** | **[]string** | Optionally include a list of unique related resources in the response | 
 **createdAts** | **[]string** | Timestamp to filter by, supports filtering with relational operators | 
 **updatedAts** | **[]string** | Timestamp to filter by, supports filtering with relational operators | 

### Return type

[**RolesList**](RolesList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesGuidDelete

> V3RolesGuidDelete(ctx, guid).Execute()

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
	guid := "guid_example" // string | Unique identifier for the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3RolesGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RolesGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
	guid := "guid_example" // string | Unique identifier for the role
	include := []string{"Inner_example"} // []string | Optionally include a list of unique related resources in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RolesGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RolesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesGuidGet`: Role
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RolesGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | Unique identifier for the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | Optionally include a list of unique related resources in the response | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RolesPost

> Role V3RolesPost(ctx).V3RolesPostRequest(v3RolesPostRequest).Execute()

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
	v3RolesPostRequest := *openapiclient.NewV3RolesPostRequest() // V3RolesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RolesPost(context.Background()).V3RolesPostRequest(v3RolesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RolesPost`: Role
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3RolesPostRequest** | [**V3RolesPostRequest**](V3RolesPostRequest.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RoutesGet

> V3RoutesGet200Response V3RoutesGet(ctx).AppGuids(appGuids).DomainGuids(domainGuids).Hosts(hosts).OrganizationGuids(organizationGuids).Paths(paths).Ports(ports).SpaceGuids(spaceGuids).ServiceInstanceGuids(serviceInstanceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List routes



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
	appGuids := "appGuids_example" // string | Comma-delimited list of app guids to filter by (optional)
	domainGuids := "domainGuids_example" // string | Comma-delimited list of domain guids to filter by (optional)
	hosts := "hosts_example" // string | Comma-delimited list of hostnames to filter by (optional)
	organizationGuids := "organizationGuids_example" // string | Comma-delimited list of organization guids to filter by (optional)
	paths := "paths_example" // string | Comma-delimited list of paths to filter by (e.g. /path1,/path2) (optional)
	ports := "ports_example" // string | Comma-delimited list of ports to filter by (e.g. 3306,5432) (optional)
	spaceGuids := "spaceGuids_example" // string | Comma-delimited list of space guids to filter by (optional)
	serviceInstanceGuids := "serviceInstanceGuids_example" // string | Comma-delimited list of service instance guids to filter by (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	include := "include_example" // string | Optionally include a list of unique related resources in the response. Valid values are domain, space.organization, space (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RoutesGet(context.Background()).AppGuids(appGuids).DomainGuids(domainGuids).Hosts(hosts).OrganizationGuids(organizationGuids).Paths(paths).Ports(ports).SpaceGuids(spaceGuids).ServiceInstanceGuids(serviceInstanceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).LabelSelector(labelSelector).Include(include).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RoutesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGet`: V3RoutesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RoutesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3RoutesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appGuids** | **string** | Comma-delimited list of app guids to filter by | 
 **domainGuids** | **string** | Comma-delimited list of domain guids to filter by | 
 **hosts** | **string** | Comma-delimited list of hostnames to filter by | 
 **organizationGuids** | **string** | Comma-delimited list of organization guids to filter by | 
 **paths** | **string** | Comma-delimited list of paths to filter by (e.g. /path1,/path2) | 
 **ports** | **string** | Comma-delimited list of ports to filter by (e.g. 3306,5432) | 
 **spaceGuids** | **string** | Comma-delimited list of space guids to filter by | 
 **serviceInstanceGuids** | **string** | Comma-delimited list of service instance guids to filter by | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **include** | **string** | Optionally include a list of unique related resources in the response. Valid values are domain, space.organization, space | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**V3RoutesGet200Response**](V3RoutesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	guid := "guid_example" // string | The route GUID
	include := "include_example" // string | Optionally include additional related resources in the response. Valid values are domain, space.organization, space (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RoutesGuidGet(context.Background(), guid).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RoutesGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesGuidGet`: Route
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RoutesGuidGet`: %v\n", resp)
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

 **include** | **string** | Optionally include additional related resources in the response. Valid values are domain, space.organization, space | 

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
	openapiclient "github.com/cloudfoundry-community/capi-openapi-go-client/capiclient"
)

func main() {
	v3RoutesPostRequest := *openapiclient.NewV3RoutesPostRequest() // V3RoutesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3RoutesPost(context.Background()).V3RoutesPostRequest(v3RoutesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3RoutesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RoutesPost`: Route
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3RoutesPost`: %v\n", resp)
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


## V3SecurityGroupsGet

> SecurityGroupList V3SecurityGroupsGet(ctx).Guids(guids).Names(names).GloballyEnabledRunning(globallyEnabledRunning).GloballyEnabledStaging(globallyEnabledStaging).RunningSpaceGuids(runningSpaceGuids).StagingSpaceGuids(stagingSpaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

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
	guids := []string{"Inner_example"} // []string | Comma-delimited list of security group guids to filter by (optional)
	names := []string{"Inner_example"} // []string | Comma-delimited list of security group names to filter by (optional)
	globallyEnabledRunning := true // bool | If true, only include the security groups that are enabled for running (optional)
	globallyEnabledStaging := true // bool | If true, only include the security groups that are enabled for staging (optional)
	runningSpaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	stagingSpaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space guids to filter by (optional)
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SecurityGroupsGet(context.Background()).Guids(guids).Names(names).GloballyEnabledRunning(globallyEnabledRunning).GloballyEnabledStaging(globallyEnabledStaging).RunningSpaceGuids(runningSpaceGuids).StagingSpaceGuids(stagingSpaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SecurityGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGet`: SecurityGroupList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SecurityGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SecurityGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **[]string** | Comma-delimited list of security group guids to filter by | 
 **names** | **[]string** | Comma-delimited list of security group names to filter by | 
 **globallyEnabledRunning** | **bool** | If true, only include the security groups that are enabled for running | 
 **globallyEnabledStaging** | **bool** | If true, only include the security groups that are enabled for staging | 
 **runningSpaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **stagingSpaceGuids** | **[]string** | Comma-delimited list of space guids to filter by | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**SecurityGroupList**](SecurityGroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SecurityGroupsGuidDelete

> V3SecurityGroupsGuidDelete(ctx, guid).Execute()

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
	guid := "guid_example" // string | The security group GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3SecurityGroupsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SecurityGroupsGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
	guid := "guid_example" // string | The security group GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SecurityGroupsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SecurityGroupsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidGet`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SecurityGroupsGuidGet`: %v\n", resp)
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

No authorization required

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
	guid := "guid_example" // string | The security group GUID
	securityGroupUpdate := *openapiclient.NewSecurityGroupUpdate() // SecurityGroupUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SecurityGroupsGuidPatch(context.Background(), guid).SecurityGroupUpdate(securityGroupUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SecurityGroupsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsGuidPatch`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SecurityGroupsGuidPatch`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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
	securityGroupCreate := *openapiclient.NewSecurityGroupCreate() // SecurityGroupCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SecurityGroupsPost(context.Background()).SecurityGroupCreate(securityGroupCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SecurityGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SecurityGroupsPost`: SecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SecurityGroupsPost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGet

> ServiceBrokerList V3ServiceBrokersGet(ctx).Names(names).Page(page).PerPage(perPage).SpaceGuids(spaceGuids).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

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
	page := int32(56) // int32 | Page to display; valid values are integers >= 1 (optional)
	perPage := int32(56) // int32 | Number of results per page; valid values are 1 through 5000 (optional)
	spaceGuids := []string{"Inner_example"} // []string | Comma-delimited list of space GUIDs to filter by (optional)
	orderBy := "orderBy_example" // string | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name (optional)
	labelSelector := "labelSelector_example" // string | A query string containing a list of label selector requirements (optional)
	createdAts := "createdAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)
	updatedAts := "updatedAts_example" // string | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGet(context.Background()).Names(names).Page(page).PerPage(perPage).SpaceGuids(spaceGuids).OrderBy(orderBy).LabelSelector(labelSelector).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceBrokersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceBrokersGet`: ServiceBrokerList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceBrokersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Comma-delimited list of service broker names to filter by | 
 **page** | **int32** | Page to display; valid values are integers &gt;&#x3D; 1 | 
 **perPage** | **int32** | Number of results per page; valid values are 1 through 5000 | 
 **spaceGuids** | **[]string** | Comma-delimited list of space GUIDs to filter by | 
 **orderBy** | **string** | Value to sort by. Defaults to ascending; prepend with - to sort descending. Valid values are created_at, updated_at, name | 
 **labelSelector** | **string** | A query string containing a list of label selector requirements | 
 **createdAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 
 **updatedAts** | **string** | Timestamp to filter by. When filtering on equality, several comma-delimited timestamps may be passed. Also supports filtering with relational operators | 

### Return type

[**ServiceBrokerList**](ServiceBrokerList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	guid := "guid_example" // string | The service broker GUID

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
**guid** | **string** | The service broker GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

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
	guid := "guid_example" // string | The service broker GUID

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
**guid** | **string** | The service broker GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersGuidPatch

> ServiceBroker V3ServiceBrokersGuidPatch(ctx, guid).ServiceBrokerUpdate(serviceBrokerUpdate).Execute()

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
	guid := "guid_example" // string | The service broker GUID
	serviceBrokerUpdate := *openapiclient.NewServiceBrokerUpdate() // ServiceBrokerUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersGuidPatch(context.Background(), guid).ServiceBrokerUpdate(serviceBrokerUpdate).Execute()
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
**guid** | **string** | The service broker GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceBrokersGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceBrokerUpdate** | [**ServiceBrokerUpdate**](ServiceBrokerUpdate.md) |  | 

### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceBrokersPost

> ServiceBroker V3ServiceBrokersPost(ctx).ServiceBrokerCreate(serviceBrokerCreate).Execute()

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
	serviceBrokerCreate := *openapiclient.NewServiceBrokerCreate() // ServiceBrokerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceBrokersPost(context.Background()).ServiceBrokerCreate(serviceBrokerCreate).Execute()
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
 **serviceBrokerCreate** | [**ServiceBrokerCreate**](ServiceBrokerCreate.md) |  | 

### Return type

[**ServiceBroker**](ServiceBroker.md)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceCredentialBindingsGuidPatch

> ServiceCredentialBinding V3ServiceCredentialBindingsGuidPatch(ctx, guid).V3ServiceCredentialBindingsGuidPatchRequest(v3ServiceCredentialBindingsGuidPatchRequest).Execute()

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
	v3ServiceCredentialBindingsGuidPatchRequest := *openapiclient.NewV3ServiceCredentialBindingsGuidPatchRequest() // V3ServiceCredentialBindingsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceCredentialBindingsGuidPatch(context.Background(), guid).V3ServiceCredentialBindingsGuidPatchRequest(v3ServiceCredentialBindingsGuidPatchRequest).Execute()
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

 **v3ServiceCredentialBindingsGuidPatchRequest** | [**V3ServiceCredentialBindingsGuidPatchRequest**](V3ServiceCredentialBindingsGuidPatchRequest.md) |  | 

### Return type

[**ServiceCredentialBinding**](ServiceCredentialBinding.md)

### Authorization

No authorization required

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
	v3ServiceCredentialBindingsPostRequest := *openapiclient.NewV3ServiceCredentialBindingsPostRequest(*openapiclient.NewV3ServiceCredentialBindingsPostRequestRelationships(*openapiclient.NewV3AppsPostRequestRelationshipsSpace()), "Type_example") // V3ServiceCredentialBindingsPostRequest | 

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidRelationshipsSharedSpacesPost

> V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest V3ServiceInstancesGuidRelationshipsSharedSpacesPost(ctx, guid).V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest(v3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest).Execute()

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
	v3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest := *openapiclient.NewV3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest() // V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesPost(context.Background(), guid).V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest(v3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidRelationshipsSharedSpacesPost`: V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest
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

 **v3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest** | [**V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest**](V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest.md) |  | 

### Return type

[**V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest**](V3ServiceInstancesGuidRelationshipsSharedSpacesPostRequest.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet

> UsageSummary V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet(ctx, guid).Execute()

Get usage summary in shared spaces



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
	resp, r, err := apiClient.DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | GUID of the service instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServiceInstancesGuidRelationshipsSharedSpacesUsageSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageSummary**](UsageSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityGet

> ServicePlanVisibility V3ServicePlansGuidVisibilityGet(ctx, guid).Execute()

Get a service plan visibility



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
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGuidVisibilityGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidVisibilityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityGet`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGuidVisibilityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityOrganizationGuidDelete

> V3ServicePlansGuidVisibilityOrganizationGuidDelete(ctx, guid, organizationGuid).Execute()

Remove organization from a service plan visibility



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
	organizationGuid := "organizationGuid_example" // string | The organization GUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3ServicePlansGuidVisibilityOrganizationGuidDelete(context.Background(), guid, organizationGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidVisibilityOrganizationGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 
**organizationGuid** | **string** | The organization GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityOrganizationGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityPatch

> ServicePlanVisibility V3ServicePlansGuidVisibilityPatch(ctx, guid).ServicePlanVisibilityUpdate(servicePlanVisibilityUpdate).Execute()

Update a service plan visibility



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
	servicePlanVisibilityUpdate := *openapiclient.NewServicePlanVisibilityUpdate() // ServicePlanVisibilityUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGuidVisibilityPatch(context.Background(), guid).ServicePlanVisibilityUpdate(servicePlanVisibilityUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidVisibilityPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityPatch`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGuidVisibilityPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePlanVisibilityUpdate** | [**ServicePlanVisibilityUpdate**](ServicePlanVisibilityUpdate.md) |  | 

### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ServicePlansGuidVisibilityPost

> ServicePlanVisibility V3ServicePlansGuidVisibilityPost(ctx, guid).ServicePlanVisibilityApply(servicePlanVisibilityApply).Execute()

Apply a service plan visibility



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
	servicePlanVisibilityApply := *openapiclient.NewServicePlanVisibilityApply() // ServicePlanVisibilityApply | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3ServicePlansGuidVisibilityPost(context.Background(), guid).ServicePlanVisibilityApply(servicePlanVisibilityApply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3ServicePlansGuidVisibilityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ServicePlansGuidVisibilityPost`: ServicePlanVisibility
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3ServicePlansGuidVisibilityPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** | The service plan GUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3ServicePlansGuidVisibilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePlanVisibilityApply** | [**ServicePlanVisibilityApply**](ServicePlanVisibilityApply.md) |  | 

### Return type

[**ServicePlanVisibility**](ServicePlanVisibility.md)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidDelete

> V3SidecarsGuidDelete(ctx, guid).Execute()

Delete a sidecar



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
	r, err := apiClient.DefaultAPI.V3SidecarsGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SidecarsGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3SidecarsGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidGet

> Sidecar V3SidecarsGuidGet(ctx, guid).Execute()

Get a sidecar



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
	resp, r, err := apiClient.DefaultAPI.V3SidecarsGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SidecarsGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidGet`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SidecarsGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sidecar**](Sidecar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SidecarsGuidPatch

> Sidecar V3SidecarsGuidPatch(ctx, guid).V3SidecarsGuidPatchRequest(v3SidecarsGuidPatchRequest).Execute()

Update a sidecar



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
	v3SidecarsGuidPatchRequest := *openapiclient.NewV3SidecarsGuidPatchRequest() // V3SidecarsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SidecarsGuidPatch(context.Background(), guid).V3SidecarsGuidPatchRequest(v3SidecarsGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SidecarsGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SidecarsGuidPatch`: Sidecar
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SidecarsGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SidecarsGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SidecarsGuidPatchRequest** | [**V3SidecarsGuidPatchRequest**](V3SidecarsGuidPatchRequest.md) |  | 

### Return type

[**Sidecar**](Sidecar.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGet

> V3SpaceQuotasGet200Response V3SpaceQuotasGet(ctx).Guids(guids).Names(names).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()

List space quotas



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
	organizationGuids := "organizationGuids_example" // string |  (optional)
	spaceGuids := "spaceGuids_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	perPage := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	createdAts := "createdAts_example" // string |  (optional)
	updatedAts := "updatedAts_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpaceQuotasGet(context.Background()).Guids(guids).Names(names).OrganizationGuids(organizationGuids).SpaceGuids(spaceGuids).Page(page).PerPage(perPage).OrderBy(orderBy).CreatedAts(createdAts).UpdatedAts(updatedAts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGet`: V3SpaceQuotasGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpaceQuotasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guids** | **string** |  | 
 **names** | **string** |  | 
 **organizationGuids** | **string** |  | 
 **spaceGuids** | **string** |  | 
 **page** | **int32** |  | 
 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **createdAts** | **string** |  | 
 **updatedAts** | **string** |  | 

### Return type

[**V3SpaceQuotasGet200Response**](V3SpaceQuotasGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidDelete

> V3SpaceQuotasGuidDelete(ctx, guid).Execute()

Delete a space quota



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
	r, err := apiClient.DefaultAPI.V3SpaceQuotasGuidDelete(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasGuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidGet

> SpaceQuota V3SpaceQuotasGuidGet(ctx, guid).Execute()

Get a space quota



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
	resp, r, err := apiClient.DefaultAPI.V3SpaceQuotasGuidGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasGuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidGet`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpaceQuotasGuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasGuidPatch

> SpaceQuota V3SpaceQuotasGuidPatch(ctx, guid).V3SpaceQuotasGuidPatchRequest(v3SpaceQuotasGuidPatchRequest).Execute()

Update a space quota



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
	v3SpaceQuotasGuidPatchRequest := *openapiclient.NewV3SpaceQuotasGuidPatchRequest() // V3SpaceQuotasGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpaceQuotasGuidPatch(context.Background(), guid).V3SpaceQuotasGuidPatchRequest(v3SpaceQuotasGuidPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasGuidPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasGuidPatch`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpaceQuotasGuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasGuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3SpaceQuotasGuidPatchRequest** | [**V3SpaceQuotasGuidPatchRequest**](V3SpaceQuotasGuidPatchRequest.md) |  | 

### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasPost

> SpaceQuota V3SpaceQuotasPost(ctx).V3SpaceQuotasPostRequest(v3SpaceQuotasPostRequest).Execute()

Create a space quota



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
	v3SpaceQuotasPostRequest := *openapiclient.NewV3SpaceQuotasPostRequest("Name_example", *openapiclient.NewV3SpaceQuotasPostRequestRelationships()) // V3SpaceQuotasPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpaceQuotasPost(context.Background()).V3SpaceQuotasPostRequest(v3SpaceQuotasPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasPost`: SpaceQuota
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpaceQuotasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v3SpaceQuotasPostRequest** | [**V3SpaceQuotasPostRequest**](V3SpaceQuotasPostRequest.md) |  | 

### Return type

[**SpaceQuota**](SpaceQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasQuotaGuidRelationshipsSpacesPost

> V3SpaceQuotasQuotaGuidRelationshipsSpacesPost200Response V3SpaceQuotasQuotaGuidRelationshipsSpacesPost(ctx, quotaGuid).V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest(v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest).Execute()

Apply a space quota to spaces



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
	quotaGuid := "quotaGuid_example" // string | 
	v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest := *openapiclient.NewV3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest() // V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpaceQuotasQuotaGuidRelationshipsSpacesPost(context.Background(), quotaGuid).V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest(v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasQuotaGuidRelationshipsSpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpaceQuotasQuotaGuidRelationshipsSpacesPost`: V3SpaceQuotasQuotaGuidRelationshipsSpacesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpaceQuotasQuotaGuidRelationshipsSpacesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasQuotaGuidRelationshipsSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest** | [**V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest**](V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest.md) |  | 

### Return type

[**V3SpaceQuotasQuotaGuidRelationshipsSpacesPost200Response**](V3SpaceQuotasQuotaGuidRelationshipsSpacesPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete

> V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete(ctx, quotaGuid, spaceGuid).Execute()

Remove a space quota from a space



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
	quotaGuid := "quotaGuid_example" // string | 
	spaceGuid := "spaceGuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete(context.Background(), quotaGuid, spaceGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quotaGuid** | **string** |  | 
**spaceGuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpaceQuotasQuotaGuidRelationshipsSpacesSpaceGuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidFeaturesGet

> V3SpacesGuidFeaturesGet200Response V3SpacesGuidFeaturesGet(ctx, guid).Execute()

List space features



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
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidFeaturesGet(context.Background(), guid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesGet`: V3SpacesGuidFeaturesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidFeaturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V3SpacesGuidFeaturesGet200Response**](V3SpacesGuidFeaturesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidFeaturesNameGet

> SpaceFeature V3SpacesGuidFeaturesNameGet(ctx, guid, name).Execute()

Get a space feature



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidFeaturesNameGet(context.Background(), guid, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidFeaturesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesNameGet`: SpaceFeature
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidFeaturesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SpaceFeature**](SpaceFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidFeaturesPatch

> SpaceFeature V3SpacesGuidFeaturesPatch(ctx, guid, name).V3SpacesGuidFeaturesPatchRequest(v3SpacesGuidFeaturesPatchRequest).Execute()

Update space features



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
	name := "name_example" // string | 
	v3SpacesGuidFeaturesPatchRequest := *openapiclient.NewV3SpacesGuidFeaturesPatchRequest() // V3SpacesGuidFeaturesPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidFeaturesPatch(context.Background(), guid, name).V3SpacesGuidFeaturesPatchRequest(v3SpacesGuidFeaturesPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidFeaturesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidFeaturesPatch`: SpaceFeature
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V3SpacesGuidFeaturesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3SpacesGuidFeaturesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v3SpacesGuidFeaturesPatchRequest** | [**V3SpacesGuidFeaturesPatchRequest**](V3SpacesGuidFeaturesPatchRequest.md) |  | 

### Return type

[**SpaceFeature**](SpaceFeature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidRelationshipsIsolationSegmentGet

> V3AppsPostRequestRelationshipsSpace V3SpacesGuidRelationshipsIsolationSegmentGet(ctx, guid).Execute()

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
	// response from `V3SpacesGuidRelationshipsIsolationSegmentGet`: V3AppsPostRequestRelationshipsSpace
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

[**V3AppsPostRequestRelationshipsSpace**](V3AppsPostRequestRelationshipsSpace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SpacesGuidRelationshipsIsolationSegmentPatch

> V3AppsPostRequestRelationshipsSpace V3SpacesGuidRelationshipsIsolationSegmentPatch(ctx, guid).V3AppsPostRequestRelationshipsSpace(v3AppsPostRequestRelationshipsSpace).Execute()

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
	v3AppsPostRequestRelationshipsSpace := *openapiclient.NewV3AppsPostRequestRelationshipsSpace() // V3AppsPostRequestRelationshipsSpace | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentPatch(context.Background(), guid).V3AppsPostRequestRelationshipsSpace(v3AppsPostRequestRelationshipsSpace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V3SpacesGuidRelationshipsIsolationSegmentPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SpacesGuidRelationshipsIsolationSegmentPatch`: V3AppsPostRequestRelationshipsSpace
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

 **v3AppsPostRequestRelationshipsSpace** | [**V3AppsPostRequestRelationshipsSpace**](V3AppsPostRequestRelationshipsSpace.md) |  | 

### Return type

[**V3AppsPostRequestRelationshipsSpace**](V3AppsPostRequestRelationshipsSpace.md)

### Authorization

No authorization required

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

No authorization required

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
	v3SpacesPostRequest := *openapiclient.NewV3SpacesPostRequest("Name_example", *openapiclient.NewV3SpaceQuotasPostRequestRelationships()) // V3SpacesPostRequest | 

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3StacksGuidPatch

> Stack V3StacksGuidPatch(ctx, guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()

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
	v3RevisionsGuidPatchRequest := *openapiclient.NewV3RevisionsGuidPatchRequest() // V3RevisionsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3StacksGuidPatch(context.Background(), guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()
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

 **v3RevisionsGuidPatchRequest** | [**V3RevisionsGuidPatchRequest**](V3RevisionsGuidPatchRequest.md) |  | 

### Return type

[**Stack**](Stack.md)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3TasksGuidPatch

> Task V3TasksGuidPatch(ctx, guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()

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
	v3RevisionsGuidPatchRequest := *openapiclient.NewV3RevisionsGuidPatchRequest() // V3RevisionsGuidPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V3TasksGuidPatch(context.Background(), guid).V3RevisionsGuidPatchRequest(v3RevisionsGuidPatchRequest).Execute()
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

 **v3RevisionsGuidPatchRequest** | [**V3RevisionsGuidPatchRequest**](V3RevisionsGuidPatchRequest.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

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

No authorization required

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

