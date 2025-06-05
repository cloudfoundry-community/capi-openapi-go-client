# SpaceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | [**SpaceQuotaApps**](SpaceQuotaApps.md) |  | 
**CreatedAt** | **time.Time** | Timestamp when the space quota was created | 
**Guid** | **string** | Unique identifier for the space quota | 
**Links** | [**SpaceQuotaLinks**](SpaceQuotaLinks.md) |  | 
**Metadata** | Pointer to [**SpaceQuotaMetadata**](SpaceQuotaMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the space quota | 
**Relationships** | [**SpaceQuotaRelationships**](SpaceQuotaRelationships.md) |  | 
**Routes** | [**SpaceQuotaRoutes**](SpaceQuotaRoutes.md) |  | 
**Services** | [**SpaceQuotaServices**](SpaceQuotaServices.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp when the space quota was last updated | 

## Methods

### NewSpaceQuota

`func NewSpaceQuota(apps SpaceQuotaApps, createdAt time.Time, guid string, links SpaceQuotaLinks, name string, relationships SpaceQuotaRelationships, routes SpaceQuotaRoutes, services SpaceQuotaServices, updatedAt time.Time, ) *SpaceQuota`

NewSpaceQuota instantiates a new SpaceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaWithDefaults

`func NewSpaceQuotaWithDefaults() *SpaceQuota`

NewSpaceQuotaWithDefaults instantiates a new SpaceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *SpaceQuota) GetApps() SpaceQuotaApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SpaceQuota) GetAppsOk() (*SpaceQuotaApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SpaceQuota) SetApps(v SpaceQuotaApps)`

SetApps sets Apps field to given value.


### GetCreatedAt

`func (o *SpaceQuota) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceQuota) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceQuota) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGuid

`func (o *SpaceQuota) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *SpaceQuota) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *SpaceQuota) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *SpaceQuota) GetLinks() SpaceQuotaLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SpaceQuota) GetLinksOk() (*SpaceQuotaLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SpaceQuota) SetLinks(v SpaceQuotaLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *SpaceQuota) GetMetadata() SpaceQuotaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpaceQuota) GetMetadataOk() (*SpaceQuotaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpaceQuota) SetMetadata(v SpaceQuotaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpaceQuota) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SpaceQuota) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceQuota) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceQuota) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *SpaceQuota) GetRelationships() SpaceQuotaRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SpaceQuota) GetRelationshipsOk() (*SpaceQuotaRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SpaceQuota) SetRelationships(v SpaceQuotaRelationships)`

SetRelationships sets Relationships field to given value.


### GetRoutes

`func (o *SpaceQuota) GetRoutes() SpaceQuotaRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SpaceQuota) GetRoutesOk() (*SpaceQuotaRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SpaceQuota) SetRoutes(v SpaceQuotaRoutes)`

SetRoutes sets Routes field to given value.


### GetServices

`func (o *SpaceQuota) GetServices() SpaceQuotaServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SpaceQuota) GetServicesOk() (*SpaceQuotaServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SpaceQuota) SetServices(v SpaceQuotaServices)`

SetServices sets Services field to given value.


### GetUpdatedAt

`func (o *SpaceQuota) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SpaceQuota) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SpaceQuota) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


