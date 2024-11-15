/*
Cloud Controller API

API specification for managing environment variable groups in Cloud Controller.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package capiclient

import (
	"encoding/json"
)

// checks if the V3Get200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V3Get200ResponseLinks{}

// V3Get200ResponseLinks struct for V3Get200ResponseLinks
type V3Get200ResponseLinks struct {
	AppUsageEvents *Get200ResponseLinksLogCache `json:"app_usage_events,omitempty"`
	Apps *Get200ResponseLinksLogCache `json:"apps,omitempty"`
	AuditEvents *Get200ResponseLinksLogCache `json:"audit_events,omitempty"`
	Buildpacks *Get200ResponseLinksLogCache `json:"buildpacks,omitempty"`
	Builds *Get200ResponseLinksLogCache `json:"builds,omitempty"`
	Deployments *Get200ResponseLinksLogCache `json:"deployments,omitempty"`
	Domains *Get200ResponseLinksLogCache `json:"domains,omitempty"`
	Droplets *Get200ResponseLinksLogCache `json:"droplets,omitempty"`
	EnvironmentVariableGroups *Get200ResponseLinksLogCache `json:"environment_variable_groups,omitempty"`
	FeatureFlags *Get200ResponseLinksLogCache `json:"feature_flags,omitempty"`
	Info *Get200ResponseLinksLogCache `json:"info,omitempty"`
	IsolationSegments *Get200ResponseLinksLogCache `json:"isolation_segments,omitempty"`
	OrganizationQuotas *Get200ResponseLinksLogCache `json:"organization_quotas,omitempty"`
	Organizations *Get200ResponseLinksLogCache `json:"organizations,omitempty"`
	Packages *Get200ResponseLinksLogCache `json:"packages,omitempty"`
	Processes *Get200ResponseLinksLogCache `json:"processes,omitempty"`
	ResourceMatches *Get200ResponseLinksLogCache `json:"resource_matches,omitempty"`
	Roles *Get200ResponseLinksLogCache `json:"roles,omitempty"`
	Routes *Get200ResponseLinksLogCache `json:"routes,omitempty"`
	SecurityGroups *Get200ResponseLinksLogCache `json:"security_groups,omitempty"`
	Self *Get200ResponseLinksLogCache `json:"self,omitempty"`
	ServiceBrokers *Get200ResponseLinksLogCache `json:"service_brokers,omitempty"`
	ServiceInstances NullableV3Get200ResponseLinksServiceInstances `json:"service_instances,omitempty"`
	ServiceOfferings *Get200ResponseLinksLogCache `json:"service_offerings,omitempty"`
	ServicePlans *Get200ResponseLinksLogCache `json:"service_plans,omitempty"`
	ServiceUsageEvents *Get200ResponseLinksLogCache `json:"service_usage_events,omitempty"`
	SpaceQuotas *Get200ResponseLinksLogCache `json:"space_quotas,omitempty"`
	Spaces *Get200ResponseLinksLogCache `json:"spaces,omitempty"`
	Stacks *Get200ResponseLinksLogCache `json:"stacks,omitempty"`
	Tasks *Get200ResponseLinksLogCache `json:"tasks,omitempty"`
	Users *Get200ResponseLinksLogCache `json:"users,omitempty"`
}

// NewV3Get200ResponseLinks instantiates a new V3Get200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV3Get200ResponseLinks() *V3Get200ResponseLinks {
	this := V3Get200ResponseLinks{}
	return &this
}

// NewV3Get200ResponseLinksWithDefaults instantiates a new V3Get200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV3Get200ResponseLinksWithDefaults() *V3Get200ResponseLinks {
	this := V3Get200ResponseLinks{}
	return &this
}

// GetAppUsageEvents returns the AppUsageEvents field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetAppUsageEvents() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.AppUsageEvents) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.AppUsageEvents
}

// GetAppUsageEventsOk returns a tuple with the AppUsageEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetAppUsageEventsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.AppUsageEvents) {
		return nil, false
	}
	return o.AppUsageEvents, true
}

// HasAppUsageEvents returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasAppUsageEvents() bool {
	if o != nil && !IsNil(o.AppUsageEvents) {
		return true
	}

	return false
}

// SetAppUsageEvents gets a reference to the given Get200ResponseLinksLogCache and assigns it to the AppUsageEvents field.
func (o *V3Get200ResponseLinks) SetAppUsageEvents(v Get200ResponseLinksLogCache) {
	o.AppUsageEvents = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetApps() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Apps) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetAppsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Apps field.
func (o *V3Get200ResponseLinks) SetApps(v Get200ResponseLinksLogCache) {
	o.Apps = &v
}

// GetAuditEvents returns the AuditEvents field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetAuditEvents() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.AuditEvents) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.AuditEvents
}

// GetAuditEventsOk returns a tuple with the AuditEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetAuditEventsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.AuditEvents) {
		return nil, false
	}
	return o.AuditEvents, true
}

// HasAuditEvents returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasAuditEvents() bool {
	if o != nil && !IsNil(o.AuditEvents) {
		return true
	}

	return false
}

// SetAuditEvents gets a reference to the given Get200ResponseLinksLogCache and assigns it to the AuditEvents field.
func (o *V3Get200ResponseLinks) SetAuditEvents(v Get200ResponseLinksLogCache) {
	o.AuditEvents = &v
}

// GetBuildpacks returns the Buildpacks field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetBuildpacks() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Buildpacks) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Buildpacks
}

// GetBuildpacksOk returns a tuple with the Buildpacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetBuildpacksOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Buildpacks) {
		return nil, false
	}
	return o.Buildpacks, true
}

// HasBuildpacks returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasBuildpacks() bool {
	if o != nil && !IsNil(o.Buildpacks) {
		return true
	}

	return false
}

// SetBuildpacks gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Buildpacks field.
func (o *V3Get200ResponseLinks) SetBuildpacks(v Get200ResponseLinksLogCache) {
	o.Buildpacks = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetBuilds() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Builds) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetBuildsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Builds field.
func (o *V3Get200ResponseLinks) SetBuilds(v Get200ResponseLinksLogCache) {
	o.Builds = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetDeployments() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Deployments) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetDeploymentsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Deployments) {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasDeployments() bool {
	if o != nil && !IsNil(o.Deployments) {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Deployments field.
func (o *V3Get200ResponseLinks) SetDeployments(v Get200ResponseLinksLogCache) {
	o.Deployments = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetDomains() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Domains) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetDomainsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Domains field.
func (o *V3Get200ResponseLinks) SetDomains(v Get200ResponseLinksLogCache) {
	o.Domains = &v
}

// GetDroplets returns the Droplets field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetDroplets() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Droplets) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Droplets
}

// GetDropletsOk returns a tuple with the Droplets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetDropletsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Droplets) {
		return nil, false
	}
	return o.Droplets, true
}

// HasDroplets returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasDroplets() bool {
	if o != nil && !IsNil(o.Droplets) {
		return true
	}

	return false
}

// SetDroplets gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Droplets field.
func (o *V3Get200ResponseLinks) SetDroplets(v Get200ResponseLinksLogCache) {
	o.Droplets = &v
}

// GetEnvironmentVariableGroups returns the EnvironmentVariableGroups field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetEnvironmentVariableGroups() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.EnvironmentVariableGroups) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.EnvironmentVariableGroups
}

// GetEnvironmentVariableGroupsOk returns a tuple with the EnvironmentVariableGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetEnvironmentVariableGroupsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.EnvironmentVariableGroups) {
		return nil, false
	}
	return o.EnvironmentVariableGroups, true
}

// HasEnvironmentVariableGroups returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasEnvironmentVariableGroups() bool {
	if o != nil && !IsNil(o.EnvironmentVariableGroups) {
		return true
	}

	return false
}

// SetEnvironmentVariableGroups gets a reference to the given Get200ResponseLinksLogCache and assigns it to the EnvironmentVariableGroups field.
func (o *V3Get200ResponseLinks) SetEnvironmentVariableGroups(v Get200ResponseLinksLogCache) {
	o.EnvironmentVariableGroups = &v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetFeatureFlags() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetFeatureFlagsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given Get200ResponseLinksLogCache and assigns it to the FeatureFlags field.
func (o *V3Get200ResponseLinks) SetFeatureFlags(v Get200ResponseLinksLogCache) {
	o.FeatureFlags = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetInfo() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Info) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetInfoOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Info field.
func (o *V3Get200ResponseLinks) SetInfo(v Get200ResponseLinksLogCache) {
	o.Info = &v
}

// GetIsolationSegments returns the IsolationSegments field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetIsolationSegments() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.IsolationSegments) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.IsolationSegments
}

// GetIsolationSegmentsOk returns a tuple with the IsolationSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetIsolationSegmentsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.IsolationSegments) {
		return nil, false
	}
	return o.IsolationSegments, true
}

// HasIsolationSegments returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasIsolationSegments() bool {
	if o != nil && !IsNil(o.IsolationSegments) {
		return true
	}

	return false
}

// SetIsolationSegments gets a reference to the given Get200ResponseLinksLogCache and assigns it to the IsolationSegments field.
func (o *V3Get200ResponseLinks) SetIsolationSegments(v Get200ResponseLinksLogCache) {
	o.IsolationSegments = &v
}

// GetOrganizationQuotas returns the OrganizationQuotas field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetOrganizationQuotas() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.OrganizationQuotas) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.OrganizationQuotas
}

// GetOrganizationQuotasOk returns a tuple with the OrganizationQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetOrganizationQuotasOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.OrganizationQuotas) {
		return nil, false
	}
	return o.OrganizationQuotas, true
}

// HasOrganizationQuotas returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasOrganizationQuotas() bool {
	if o != nil && !IsNil(o.OrganizationQuotas) {
		return true
	}

	return false
}

// SetOrganizationQuotas gets a reference to the given Get200ResponseLinksLogCache and assigns it to the OrganizationQuotas field.
func (o *V3Get200ResponseLinks) SetOrganizationQuotas(v Get200ResponseLinksLogCache) {
	o.OrganizationQuotas = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetOrganizations() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Organizations) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetOrganizationsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Organizations field.
func (o *V3Get200ResponseLinks) SetOrganizations(v Get200ResponseLinksLogCache) {
	o.Organizations = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetPackages() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Packages) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetPackagesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Packages field.
func (o *V3Get200ResponseLinks) SetPackages(v Get200ResponseLinksLogCache) {
	o.Packages = &v
}

// GetProcesses returns the Processes field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetProcesses() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Processes) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetProcessesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasProcesses() bool {
	if o != nil && !IsNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Processes field.
func (o *V3Get200ResponseLinks) SetProcesses(v Get200ResponseLinksLogCache) {
	o.Processes = &v
}

// GetResourceMatches returns the ResourceMatches field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetResourceMatches() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.ResourceMatches) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.ResourceMatches
}

// GetResourceMatchesOk returns a tuple with the ResourceMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetResourceMatchesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.ResourceMatches) {
		return nil, false
	}
	return o.ResourceMatches, true
}

// HasResourceMatches returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasResourceMatches() bool {
	if o != nil && !IsNil(o.ResourceMatches) {
		return true
	}

	return false
}

// SetResourceMatches gets a reference to the given Get200ResponseLinksLogCache and assigns it to the ResourceMatches field.
func (o *V3Get200ResponseLinks) SetResourceMatches(v Get200ResponseLinksLogCache) {
	o.ResourceMatches = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetRoles() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Roles) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetRolesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Roles field.
func (o *V3Get200ResponseLinks) SetRoles(v Get200ResponseLinksLogCache) {
	o.Roles = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetRoutes() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Routes) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetRoutesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Routes field.
func (o *V3Get200ResponseLinks) SetRoutes(v Get200ResponseLinksLogCache) {
	o.Routes = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetSecurityGroups() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetSecurityGroupsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given Get200ResponseLinksLogCache and assigns it to the SecurityGroups field.
func (o *V3Get200ResponseLinks) SetSecurityGroups(v Get200ResponseLinksLogCache) {
	o.SecurityGroups = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetSelf() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Self) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetSelfOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Self field.
func (o *V3Get200ResponseLinks) SetSelf(v Get200ResponseLinksLogCache) {
	o.Self = &v
}

// GetServiceBrokers returns the ServiceBrokers field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetServiceBrokers() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.ServiceBrokers) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.ServiceBrokers
}

// GetServiceBrokersOk returns a tuple with the ServiceBrokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetServiceBrokersOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.ServiceBrokers) {
		return nil, false
	}
	return o.ServiceBrokers, true
}

// HasServiceBrokers returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasServiceBrokers() bool {
	if o != nil && !IsNil(o.ServiceBrokers) {
		return true
	}

	return false
}

// SetServiceBrokers gets a reference to the given Get200ResponseLinksLogCache and assigns it to the ServiceBrokers field.
func (o *V3Get200ResponseLinks) SetServiceBrokers(v Get200ResponseLinksLogCache) {
	o.ServiceBrokers = &v
}

// GetServiceInstances returns the ServiceInstances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V3Get200ResponseLinks) GetServiceInstances() V3Get200ResponseLinksServiceInstances {
	if o == nil || IsNil(o.ServiceInstances.Get()) {
		var ret V3Get200ResponseLinksServiceInstances
		return ret
	}
	return *o.ServiceInstances.Get()
}

// GetServiceInstancesOk returns a tuple with the ServiceInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V3Get200ResponseLinks) GetServiceInstancesOk() (*V3Get200ResponseLinksServiceInstances, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceInstances.Get(), o.ServiceInstances.IsSet()
}

// HasServiceInstances returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasServiceInstances() bool {
	if o != nil && o.ServiceInstances.IsSet() {
		return true
	}

	return false
}

// SetServiceInstances gets a reference to the given NullableV3Get200ResponseLinksServiceInstances and assigns it to the ServiceInstances field.
func (o *V3Get200ResponseLinks) SetServiceInstances(v V3Get200ResponseLinksServiceInstances) {
	o.ServiceInstances.Set(&v)
}
// SetServiceInstancesNil sets the value for ServiceInstances to be an explicit nil
func (o *V3Get200ResponseLinks) SetServiceInstancesNil() {
	o.ServiceInstances.Set(nil)
}

// UnsetServiceInstances ensures that no value is present for ServiceInstances, not even an explicit nil
func (o *V3Get200ResponseLinks) UnsetServiceInstances() {
	o.ServiceInstances.Unset()
}

// GetServiceOfferings returns the ServiceOfferings field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetServiceOfferings() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.ServiceOfferings) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.ServiceOfferings
}

// GetServiceOfferingsOk returns a tuple with the ServiceOfferings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetServiceOfferingsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.ServiceOfferings) {
		return nil, false
	}
	return o.ServiceOfferings, true
}

// HasServiceOfferings returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasServiceOfferings() bool {
	if o != nil && !IsNil(o.ServiceOfferings) {
		return true
	}

	return false
}

// SetServiceOfferings gets a reference to the given Get200ResponseLinksLogCache and assigns it to the ServiceOfferings field.
func (o *V3Get200ResponseLinks) SetServiceOfferings(v Get200ResponseLinksLogCache) {
	o.ServiceOfferings = &v
}

// GetServicePlans returns the ServicePlans field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetServicePlans() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.ServicePlans) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.ServicePlans
}

// GetServicePlansOk returns a tuple with the ServicePlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetServicePlansOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.ServicePlans) {
		return nil, false
	}
	return o.ServicePlans, true
}

// HasServicePlans returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasServicePlans() bool {
	if o != nil && !IsNil(o.ServicePlans) {
		return true
	}

	return false
}

// SetServicePlans gets a reference to the given Get200ResponseLinksLogCache and assigns it to the ServicePlans field.
func (o *V3Get200ResponseLinks) SetServicePlans(v Get200ResponseLinksLogCache) {
	o.ServicePlans = &v
}

// GetServiceUsageEvents returns the ServiceUsageEvents field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetServiceUsageEvents() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.ServiceUsageEvents) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.ServiceUsageEvents
}

// GetServiceUsageEventsOk returns a tuple with the ServiceUsageEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetServiceUsageEventsOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.ServiceUsageEvents) {
		return nil, false
	}
	return o.ServiceUsageEvents, true
}

// HasServiceUsageEvents returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasServiceUsageEvents() bool {
	if o != nil && !IsNil(o.ServiceUsageEvents) {
		return true
	}

	return false
}

// SetServiceUsageEvents gets a reference to the given Get200ResponseLinksLogCache and assigns it to the ServiceUsageEvents field.
func (o *V3Get200ResponseLinks) SetServiceUsageEvents(v Get200ResponseLinksLogCache) {
	o.ServiceUsageEvents = &v
}

// GetSpaceQuotas returns the SpaceQuotas field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetSpaceQuotas() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.SpaceQuotas) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.SpaceQuotas
}

// GetSpaceQuotasOk returns a tuple with the SpaceQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetSpaceQuotasOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.SpaceQuotas) {
		return nil, false
	}
	return o.SpaceQuotas, true
}

// HasSpaceQuotas returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasSpaceQuotas() bool {
	if o != nil && !IsNil(o.SpaceQuotas) {
		return true
	}

	return false
}

// SetSpaceQuotas gets a reference to the given Get200ResponseLinksLogCache and assigns it to the SpaceQuotas field.
func (o *V3Get200ResponseLinks) SetSpaceQuotas(v Get200ResponseLinksLogCache) {
	o.SpaceQuotas = &v
}

// GetSpaces returns the Spaces field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetSpaces() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Spaces) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Spaces
}

// GetSpacesOk returns a tuple with the Spaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetSpacesOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Spaces) {
		return nil, false
	}
	return o.Spaces, true
}

// HasSpaces returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasSpaces() bool {
	if o != nil && !IsNil(o.Spaces) {
		return true
	}

	return false
}

// SetSpaces gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Spaces field.
func (o *V3Get200ResponseLinks) SetSpaces(v Get200ResponseLinksLogCache) {
	o.Spaces = &v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetStacks() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Stacks) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetStacksOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Stacks) {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasStacks() bool {
	if o != nil && !IsNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Stacks field.
func (o *V3Get200ResponseLinks) SetStacks(v Get200ResponseLinksLogCache) {
	o.Stacks = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetTasks() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Tasks) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetTasksOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Tasks field.
func (o *V3Get200ResponseLinks) SetTasks(v Get200ResponseLinksLogCache) {
	o.Tasks = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V3Get200ResponseLinks) GetUsers() Get200ResponseLinksLogCache {
	if o == nil || IsNil(o.Users) {
		var ret Get200ResponseLinksLogCache
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V3Get200ResponseLinks) GetUsersOk() (*Get200ResponseLinksLogCache, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V3Get200ResponseLinks) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given Get200ResponseLinksLogCache and assigns it to the Users field.
func (o *V3Get200ResponseLinks) SetUsers(v Get200ResponseLinksLogCache) {
	o.Users = &v
}

func (o V3Get200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V3Get200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppUsageEvents) {
		toSerialize["app_usage_events"] = o.AppUsageEvents
	}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.AuditEvents) {
		toSerialize["audit_events"] = o.AuditEvents
	}
	if !IsNil(o.Buildpacks) {
		toSerialize["buildpacks"] = o.Buildpacks
	}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.Deployments) {
		toSerialize["deployments"] = o.Deployments
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.Droplets) {
		toSerialize["droplets"] = o.Droplets
	}
	if !IsNil(o.EnvironmentVariableGroups) {
		toSerialize["environment_variable_groups"] = o.EnvironmentVariableGroups
	}
	if !IsNil(o.FeatureFlags) {
		toSerialize["feature_flags"] = o.FeatureFlags
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.IsolationSegments) {
		toSerialize["isolation_segments"] = o.IsolationSegments
	}
	if !IsNil(o.OrganizationQuotas) {
		toSerialize["organization_quotas"] = o.OrganizationQuotas
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.Processes) {
		toSerialize["processes"] = o.Processes
	}
	if !IsNil(o.ResourceMatches) {
		toSerialize["resource_matches"] = o.ResourceMatches
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["security_groups"] = o.SecurityGroups
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.ServiceBrokers) {
		toSerialize["service_brokers"] = o.ServiceBrokers
	}
	if o.ServiceInstances.IsSet() {
		toSerialize["service_instances"] = o.ServiceInstances.Get()
	}
	if !IsNil(o.ServiceOfferings) {
		toSerialize["service_offerings"] = o.ServiceOfferings
	}
	if !IsNil(o.ServicePlans) {
		toSerialize["service_plans"] = o.ServicePlans
	}
	if !IsNil(o.ServiceUsageEvents) {
		toSerialize["service_usage_events"] = o.ServiceUsageEvents
	}
	if !IsNil(o.SpaceQuotas) {
		toSerialize["space_quotas"] = o.SpaceQuotas
	}
	if !IsNil(o.Spaces) {
		toSerialize["spaces"] = o.Spaces
	}
	if !IsNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableV3Get200ResponseLinks struct {
	value *V3Get200ResponseLinks
	isSet bool
}

func (v NullableV3Get200ResponseLinks) Get() *V3Get200ResponseLinks {
	return v.value
}

func (v *NullableV3Get200ResponseLinks) Set(val *V3Get200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableV3Get200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableV3Get200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV3Get200ResponseLinks(val *V3Get200ResponseLinks) *NullableV3Get200ResponseLinks {
	return &NullableV3Get200ResponseLinks{value: val, isSet: true}
}

func (v NullableV3Get200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV3Get200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

