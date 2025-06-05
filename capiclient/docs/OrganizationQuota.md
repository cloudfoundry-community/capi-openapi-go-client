# OrganizationQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | [**OrganizationQuotaApps**](OrganizationQuotaApps.md) |  | 
**CreatedAt** | **time.Time** | Timestamp when the organization quota was created | 
**Domains** | [**OrganizationQuotaDomains**](OrganizationQuotaDomains.md) |  | 
**Guid** | **string** | Unique identifier for the organization quota | 
**Links** | [**OrganizationQuotaLinks**](OrganizationQuotaLinks.md) |  | 
**Metadata** | Pointer to [**OrganizationQuotaMetadata**](OrganizationQuotaMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the organization quota | 
**Relationships** | [**OrganizationQuotaRelationships**](OrganizationQuotaRelationships.md) |  | 
**Routes** | [**OrganizationQuotaRoutes**](OrganizationQuotaRoutes.md) |  | 
**Services** | [**OrganizationQuotaServices**](OrganizationQuotaServices.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp when the organization quota was last updated | 

## Methods

### NewOrganizationQuota

`func NewOrganizationQuota(apps OrganizationQuotaApps, createdAt time.Time, domains OrganizationQuotaDomains, guid string, links OrganizationQuotaLinks, name string, relationships OrganizationQuotaRelationships, routes OrganizationQuotaRoutes, services OrganizationQuotaServices, updatedAt time.Time, ) *OrganizationQuota`

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

`func (o *OrganizationQuota) GetApps() OrganizationQuotaApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OrganizationQuota) GetAppsOk() (*OrganizationQuotaApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OrganizationQuota) SetApps(v OrganizationQuotaApps)`

SetApps sets Apps field to given value.


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


### GetDomains

`func (o *OrganizationQuota) GetDomains() OrganizationQuotaDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationQuota) GetDomainsOk() (*OrganizationQuotaDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationQuota) SetDomains(v OrganizationQuotaDomains)`

SetDomains sets Domains field to given value.


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


### GetMetadata

`func (o *OrganizationQuota) GetMetadata() OrganizationQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrganizationQuota) GetMetadataOk() (*OrganizationQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrganizationQuota) SetMetadata(v OrganizationQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrganizationQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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


### GetRoutes

`func (o *OrganizationQuota) GetRoutes() OrganizationQuotaRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OrganizationQuota) GetRoutesOk() (*OrganizationQuotaRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OrganizationQuota) SetRoutes(v OrganizationQuotaRoutes)`

SetRoutes sets Routes field to given value.


### GetServices

`func (o *OrganizationQuota) GetServices() OrganizationQuotaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationQuota) GetServicesOk() (*OrganizationQuotaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationQuota) SetServices(v OrganizationQuotaServices)`

SetServices sets Services field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


