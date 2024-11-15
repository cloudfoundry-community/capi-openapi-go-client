# OrganizationQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**AppsQuota**](AppsQuota.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time with zone when the organization quota was created | [optional] 
**Domains** | Pointer to [**DomainsQuota**](DomainsQuota.md) |  | [optional] 
**Guid** | Pointer to **string** | Unique identifier for the organization quota | [optional] 
**Links** | Pointer to [**OrganizationQuotaLinks**](OrganizationQuotaLinks.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the quota | [optional] 
**Relationships** | Pointer to [**OrganizationQuotaRelationships**](OrganizationQuotaRelationships.md) |  | [optional] 
**Routes** | Pointer to [**RoutesQuota**](RoutesQuota.md) |  | [optional] 
**Services** | Pointer to [**ServicesQuota**](ServicesQuota.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time with zone when the organization quota was last updated | [optional] 

## Methods

### NewOrganizationQuota

`func NewOrganizationQuota() *OrganizationQuota`

NewOrganizationQuota instantiates a new OrganizationQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaWithDefaults

`func NewOrganizationQuotaWithDefaults() *OrganizationQuota`

NewOrganizationQuotaWithDefaults instantiates a new OrganizationQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OrganizationQuota) GetApps() AppsQuota`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OrganizationQuota) GetAppsOk() (*AppsQuota, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OrganizationQuota) SetApps(v AppsQuota)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OrganizationQuota) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationQuota) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationQuota) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationQuota) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationQuota) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDomains

`func (o *OrganizationQuota) GetDomains() DomainsQuota`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationQuota) GetDomainsOk() (*DomainsQuota, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationQuota) SetDomains(v DomainsQuota)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationQuota) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGuid

`func (o *OrganizationQuota) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *OrganizationQuota) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *OrganizationQuota) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *OrganizationQuota) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *OrganizationQuota) GetLinks() OrganizationQuotaLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationQuota) GetLinksOk() (*OrganizationQuotaLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationQuota) SetLinks(v OrganizationQuotaLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganizationQuota) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *OrganizationQuota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationQuota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationQuota) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationQuota) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *OrganizationQuota) GetRelationships() OrganizationQuotaRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrganizationQuota) GetRelationshipsOk() (*OrganizationQuotaRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrganizationQuota) SetRelationships(v OrganizationQuotaRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrganizationQuota) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRoutes

`func (o *OrganizationQuota) GetRoutes() RoutesQuota`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OrganizationQuota) GetRoutesOk() (*RoutesQuota, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OrganizationQuota) SetRoutes(v RoutesQuota)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *OrganizationQuota) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *OrganizationQuota) GetServices() ServicesQuota`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationQuota) GetServicesOk() (*ServicesQuota, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationQuota) SetServices(v ServicesQuota)`

SetServices sets Services field to given value.

### HasServices

`func (o *OrganizationQuota) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationQuota) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationQuota) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationQuota) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationQuota) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


