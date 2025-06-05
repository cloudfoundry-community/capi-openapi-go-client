# V3Get200ResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUsageEvents** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Apps** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**AuditEvents** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Buildpacks** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Builds** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Deployments** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Domains** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Droplets** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**EnvironmentVariableGroups** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**FeatureFlags** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Info** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**IsolationSegments** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**OrganizationQuotas** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Organizations** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Packages** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Processes** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**ResourceMatches** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Roles** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Routes** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**SecurityGroups** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Self** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**ServiceBrokers** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**ServiceInstances** | Pointer to [**NullableV3Get200ResponseLinksServiceInstances**](V3Get200ResponseLinksServiceInstances.md) |  | [optional] 
**ServiceOfferings** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**ServicePlans** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**ServiceUsageEvents** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**SpaceQuotas** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Spaces** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Stacks** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Tasks** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Users** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 

## Methods

### NewV3Get200ResponseLinks

`func NewV3Get200ResponseLinks() *V3Get200ResponseLinks`

NewV3Get200ResponseLinks instantiates a new V3Get200ResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3Get200ResponseLinksWithDefaults

`func NewV3Get200ResponseLinksWithDefaults() *V3Get200ResponseLinks`

NewV3Get200ResponseLinksWithDefaults instantiates a new V3Get200ResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUsageEvents

`func (o *V3Get200ResponseLinks) GetAppUsageEvents() Get200ResponseLinksLogCache`

GetAppUsageEvents returns the AppUsageEvents field if non-nil, zero value otherwise.

### GetAppUsageEventsOk

`func (o *V3Get200ResponseLinks) GetAppUsageEventsOk() (*Get200ResponseLinksLogCache, bool)`

GetAppUsageEventsOk returns a tuple with the AppUsageEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUsageEvents

`func (o *V3Get200ResponseLinks) SetAppUsageEvents(v Get200ResponseLinksLogCache)`

SetAppUsageEvents sets AppUsageEvents field to given value.

### HasAppUsageEvents

`func (o *V3Get200ResponseLinks) HasAppUsageEvents() bool`

HasAppUsageEvents returns a boolean if a field has been set.

### GetApps

`func (o *V3Get200ResponseLinks) GetApps() Get200ResponseLinksLogCache`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *V3Get200ResponseLinks) GetAppsOk() (*Get200ResponseLinksLogCache, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *V3Get200ResponseLinks) SetApps(v Get200ResponseLinksLogCache)`

SetApps sets Apps field to given value.

### HasApps

`func (o *V3Get200ResponseLinks) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAuditEvents

`func (o *V3Get200ResponseLinks) GetAuditEvents() Get200ResponseLinksLogCache`

GetAuditEvents returns the AuditEvents field if non-nil, zero value otherwise.

### GetAuditEventsOk

`func (o *V3Get200ResponseLinks) GetAuditEventsOk() (*Get200ResponseLinksLogCache, bool)`

GetAuditEventsOk returns a tuple with the AuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEvents

`func (o *V3Get200ResponseLinks) SetAuditEvents(v Get200ResponseLinksLogCache)`

SetAuditEvents sets AuditEvents field to given value.

### HasAuditEvents

`func (o *V3Get200ResponseLinks) HasAuditEvents() bool`

HasAuditEvents returns a boolean if a field has been set.

### GetBuildpacks

`func (o *V3Get200ResponseLinks) GetBuildpacks() Get200ResponseLinksLogCache`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *V3Get200ResponseLinks) GetBuildpacksOk() (*Get200ResponseLinksLogCache, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *V3Get200ResponseLinks) SetBuildpacks(v Get200ResponseLinksLogCache)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *V3Get200ResponseLinks) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetBuilds

`func (o *V3Get200ResponseLinks) GetBuilds() Get200ResponseLinksLogCache`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *V3Get200ResponseLinks) GetBuildsOk() (*Get200ResponseLinksLogCache, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *V3Get200ResponseLinks) SetBuilds(v Get200ResponseLinksLogCache)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *V3Get200ResponseLinks) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetDeployments

`func (o *V3Get200ResponseLinks) GetDeployments() Get200ResponseLinksLogCache`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *V3Get200ResponseLinks) GetDeploymentsOk() (*Get200ResponseLinksLogCache, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *V3Get200ResponseLinks) SetDeployments(v Get200ResponseLinksLogCache)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *V3Get200ResponseLinks) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetDomains

`func (o *V3Get200ResponseLinks) GetDomains() Get200ResponseLinksLogCache`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *V3Get200ResponseLinks) GetDomainsOk() (*Get200ResponseLinksLogCache, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *V3Get200ResponseLinks) SetDomains(v Get200ResponseLinksLogCache)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *V3Get200ResponseLinks) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetDroplets

`func (o *V3Get200ResponseLinks) GetDroplets() Get200ResponseLinksLogCache`

GetDroplets returns the Droplets field if non-nil, zero value otherwise.

### GetDropletsOk

`func (o *V3Get200ResponseLinks) GetDropletsOk() (*Get200ResponseLinksLogCache, bool)`

GetDropletsOk returns a tuple with the Droplets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplets

`func (o *V3Get200ResponseLinks) SetDroplets(v Get200ResponseLinksLogCache)`

SetDroplets sets Droplets field to given value.

### HasDroplets

`func (o *V3Get200ResponseLinks) HasDroplets() bool`

HasDroplets returns a boolean if a field has been set.

### GetEnvironmentVariableGroups

`func (o *V3Get200ResponseLinks) GetEnvironmentVariableGroups() Get200ResponseLinksLogCache`

GetEnvironmentVariableGroups returns the EnvironmentVariableGroups field if non-nil, zero value otherwise.

### GetEnvironmentVariableGroupsOk

`func (o *V3Get200ResponseLinks) GetEnvironmentVariableGroupsOk() (*Get200ResponseLinksLogCache, bool)`

GetEnvironmentVariableGroupsOk returns a tuple with the EnvironmentVariableGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariableGroups

`func (o *V3Get200ResponseLinks) SetEnvironmentVariableGroups(v Get200ResponseLinksLogCache)`

SetEnvironmentVariableGroups sets EnvironmentVariableGroups field to given value.

### HasEnvironmentVariableGroups

`func (o *V3Get200ResponseLinks) HasEnvironmentVariableGroups() bool`

HasEnvironmentVariableGroups returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *V3Get200ResponseLinks) GetFeatureFlags() Get200ResponseLinksLogCache`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *V3Get200ResponseLinks) GetFeatureFlagsOk() (*Get200ResponseLinksLogCache, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *V3Get200ResponseLinks) SetFeatureFlags(v Get200ResponseLinksLogCache)`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *V3Get200ResponseLinks) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetInfo

`func (o *V3Get200ResponseLinks) GetInfo() Get200ResponseLinksLogCache`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V3Get200ResponseLinks) GetInfoOk() (*Get200ResponseLinksLogCache, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V3Get200ResponseLinks) SetInfo(v Get200ResponseLinksLogCache)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V3Get200ResponseLinks) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIsolationSegments

`func (o *V3Get200ResponseLinks) GetIsolationSegments() Get200ResponseLinksLogCache`

GetIsolationSegments returns the IsolationSegments field if non-nil, zero value otherwise.

### GetIsolationSegmentsOk

`func (o *V3Get200ResponseLinks) GetIsolationSegmentsOk() (*Get200ResponseLinksLogCache, bool)`

GetIsolationSegmentsOk returns a tuple with the IsolationSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationSegments

`func (o *V3Get200ResponseLinks) SetIsolationSegments(v Get200ResponseLinksLogCache)`

SetIsolationSegments sets IsolationSegments field to given value.

### HasIsolationSegments

`func (o *V3Get200ResponseLinks) HasIsolationSegments() bool`

HasIsolationSegments returns a boolean if a field has been set.

### GetOrganizationQuotas

`func (o *V3Get200ResponseLinks) GetOrganizationQuotas() Get200ResponseLinksLogCache`

GetOrganizationQuotas returns the OrganizationQuotas field if non-nil, zero value otherwise.

### GetOrganizationQuotasOk

`func (o *V3Get200ResponseLinks) GetOrganizationQuotasOk() (*Get200ResponseLinksLogCache, bool)`

GetOrganizationQuotasOk returns a tuple with the OrganizationQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationQuotas

`func (o *V3Get200ResponseLinks) SetOrganizationQuotas(v Get200ResponseLinksLogCache)`

SetOrganizationQuotas sets OrganizationQuotas field to given value.

### HasOrganizationQuotas

`func (o *V3Get200ResponseLinks) HasOrganizationQuotas() bool`

HasOrganizationQuotas returns a boolean if a field has been set.

### GetOrganizations

`func (o *V3Get200ResponseLinks) GetOrganizations() Get200ResponseLinksLogCache`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *V3Get200ResponseLinks) GetOrganizationsOk() (*Get200ResponseLinksLogCache, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *V3Get200ResponseLinks) SetOrganizations(v Get200ResponseLinksLogCache)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *V3Get200ResponseLinks) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPackages

`func (o *V3Get200ResponseLinks) GetPackages() Get200ResponseLinksLogCache`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *V3Get200ResponseLinks) GetPackagesOk() (*Get200ResponseLinksLogCache, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *V3Get200ResponseLinks) SetPackages(v Get200ResponseLinksLogCache)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *V3Get200ResponseLinks) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetProcesses

`func (o *V3Get200ResponseLinks) GetProcesses() Get200ResponseLinksLogCache`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *V3Get200ResponseLinks) GetProcessesOk() (*Get200ResponseLinksLogCache, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *V3Get200ResponseLinks) SetProcesses(v Get200ResponseLinksLogCache)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *V3Get200ResponseLinks) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetResourceMatches

`func (o *V3Get200ResponseLinks) GetResourceMatches() Get200ResponseLinksLogCache`

GetResourceMatches returns the ResourceMatches field if non-nil, zero value otherwise.

### GetResourceMatchesOk

`func (o *V3Get200ResponseLinks) GetResourceMatchesOk() (*Get200ResponseLinksLogCache, bool)`

GetResourceMatchesOk returns a tuple with the ResourceMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceMatches

`func (o *V3Get200ResponseLinks) SetResourceMatches(v Get200ResponseLinksLogCache)`

SetResourceMatches sets ResourceMatches field to given value.

### HasResourceMatches

`func (o *V3Get200ResponseLinks) HasResourceMatches() bool`

HasResourceMatches returns a boolean if a field has been set.

### GetRoles

`func (o *V3Get200ResponseLinks) GetRoles() Get200ResponseLinksLogCache`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *V3Get200ResponseLinks) GetRolesOk() (*Get200ResponseLinksLogCache, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *V3Get200ResponseLinks) SetRoles(v Get200ResponseLinksLogCache)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *V3Get200ResponseLinks) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRoutes

`func (o *V3Get200ResponseLinks) GetRoutes() Get200ResponseLinksLogCache`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *V3Get200ResponseLinks) GetRoutesOk() (*Get200ResponseLinksLogCache, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *V3Get200ResponseLinks) SetRoutes(v Get200ResponseLinksLogCache)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *V3Get200ResponseLinks) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *V3Get200ResponseLinks) GetSecurityGroups() Get200ResponseLinksLogCache`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *V3Get200ResponseLinks) GetSecurityGroupsOk() (*Get200ResponseLinksLogCache, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *V3Get200ResponseLinks) SetSecurityGroups(v Get200ResponseLinksLogCache)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *V3Get200ResponseLinks) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSelf

`func (o *V3Get200ResponseLinks) GetSelf() Get200ResponseLinksLogCache`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *V3Get200ResponseLinks) GetSelfOk() (*Get200ResponseLinksLogCache, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *V3Get200ResponseLinks) SetSelf(v Get200ResponseLinksLogCache)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *V3Get200ResponseLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetServiceBrokers

`func (o *V3Get200ResponseLinks) GetServiceBrokers() Get200ResponseLinksLogCache`

GetServiceBrokers returns the ServiceBrokers field if non-nil, zero value otherwise.

### GetServiceBrokersOk

`func (o *V3Get200ResponseLinks) GetServiceBrokersOk() (*Get200ResponseLinksLogCache, bool)`

GetServiceBrokersOk returns a tuple with the ServiceBrokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBrokers

`func (o *V3Get200ResponseLinks) SetServiceBrokers(v Get200ResponseLinksLogCache)`

SetServiceBrokers sets ServiceBrokers field to given value.

### HasServiceBrokers

`func (o *V3Get200ResponseLinks) HasServiceBrokers() bool`

HasServiceBrokers returns a boolean if a field has been set.

### GetServiceInstances

`func (o *V3Get200ResponseLinks) GetServiceInstances() V3Get200ResponseLinksServiceInstances`

GetServiceInstances returns the ServiceInstances field if non-nil, zero value otherwise.

### GetServiceInstancesOk

`func (o *V3Get200ResponseLinks) GetServiceInstancesOk() (*V3Get200ResponseLinksServiceInstances, bool)`

GetServiceInstancesOk returns a tuple with the ServiceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstances

`func (o *V3Get200ResponseLinks) SetServiceInstances(v V3Get200ResponseLinksServiceInstances)`

SetServiceInstances sets ServiceInstances field to given value.

### HasServiceInstances

`func (o *V3Get200ResponseLinks) HasServiceInstances() bool`

HasServiceInstances returns a boolean if a field has been set.

### SetServiceInstancesNil

`func (o *V3Get200ResponseLinks) SetServiceInstancesNil(b bool)`

 SetServiceInstancesNil sets the value for ServiceInstances to be an explicit nil

### UnsetServiceInstances
`func (o *V3Get200ResponseLinks) UnsetServiceInstances()`

UnsetServiceInstances ensures that no value is present for ServiceInstances, not even an explicit nil
### GetServiceOfferings

`func (o *V3Get200ResponseLinks) GetServiceOfferings() Get200ResponseLinksLogCache`

GetServiceOfferings returns the ServiceOfferings field if non-nil, zero value otherwise.

### GetServiceOfferingsOk

`func (o *V3Get200ResponseLinks) GetServiceOfferingsOk() (*Get200ResponseLinksLogCache, bool)`

GetServiceOfferingsOk returns a tuple with the ServiceOfferings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOfferings

`func (o *V3Get200ResponseLinks) SetServiceOfferings(v Get200ResponseLinksLogCache)`

SetServiceOfferings sets ServiceOfferings field to given value.

### HasServiceOfferings

`func (o *V3Get200ResponseLinks) HasServiceOfferings() bool`

HasServiceOfferings returns a boolean if a field has been set.

### GetServicePlans

`func (o *V3Get200ResponseLinks) GetServicePlans() Get200ResponseLinksLogCache`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *V3Get200ResponseLinks) GetServicePlansOk() (*Get200ResponseLinksLogCache, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlans

`func (o *V3Get200ResponseLinks) SetServicePlans(v Get200ResponseLinksLogCache)`

SetServicePlans sets ServicePlans field to given value.

### HasServicePlans

`func (o *V3Get200ResponseLinks) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### GetServiceUsageEvents

`func (o *V3Get200ResponseLinks) GetServiceUsageEvents() Get200ResponseLinksLogCache`

GetServiceUsageEvents returns the ServiceUsageEvents field if non-nil, zero value otherwise.

### GetServiceUsageEventsOk

`func (o *V3Get200ResponseLinks) GetServiceUsageEventsOk() (*Get200ResponseLinksLogCache, bool)`

GetServiceUsageEventsOk returns a tuple with the ServiceUsageEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsageEvents

`func (o *V3Get200ResponseLinks) SetServiceUsageEvents(v Get200ResponseLinksLogCache)`

SetServiceUsageEvents sets ServiceUsageEvents field to given value.

### HasServiceUsageEvents

`func (o *V3Get200ResponseLinks) HasServiceUsageEvents() bool`

HasServiceUsageEvents returns a boolean if a field has been set.

### GetSpaceQuotas

`func (o *V3Get200ResponseLinks) GetSpaceQuotas() Get200ResponseLinksLogCache`

GetSpaceQuotas returns the SpaceQuotas field if non-nil, zero value otherwise.

### GetSpaceQuotasOk

`func (o *V3Get200ResponseLinks) GetSpaceQuotasOk() (*Get200ResponseLinksLogCache, bool)`

GetSpaceQuotasOk returns a tuple with the SpaceQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceQuotas

`func (o *V3Get200ResponseLinks) SetSpaceQuotas(v Get200ResponseLinksLogCache)`

SetSpaceQuotas sets SpaceQuotas field to given value.

### HasSpaceQuotas

`func (o *V3Get200ResponseLinks) HasSpaceQuotas() bool`

HasSpaceQuotas returns a boolean if a field has been set.

### GetSpaces

`func (o *V3Get200ResponseLinks) GetSpaces() Get200ResponseLinksLogCache`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *V3Get200ResponseLinks) GetSpacesOk() (*Get200ResponseLinksLogCache, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *V3Get200ResponseLinks) SetSpaces(v Get200ResponseLinksLogCache)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *V3Get200ResponseLinks) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.

### GetStacks

`func (o *V3Get200ResponseLinks) GetStacks() Get200ResponseLinksLogCache`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *V3Get200ResponseLinks) GetStacksOk() (*Get200ResponseLinksLogCache, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *V3Get200ResponseLinks) SetStacks(v Get200ResponseLinksLogCache)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *V3Get200ResponseLinks) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetTasks

`func (o *V3Get200ResponseLinks) GetTasks() Get200ResponseLinksLogCache`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *V3Get200ResponseLinks) GetTasksOk() (*Get200ResponseLinksLogCache, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *V3Get200ResponseLinks) SetTasks(v Get200ResponseLinksLogCache)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *V3Get200ResponseLinks) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUsers

`func (o *V3Get200ResponseLinks) GetUsers() Get200ResponseLinksLogCache`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V3Get200ResponseLinks) GetUsersOk() (*Get200ResponseLinksLogCache, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V3Get200ResponseLinks) SetUsers(v Get200ResponseLinksLogCache)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V3Get200ResponseLinks) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


