# SpaceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**V3SpaceQuotasPostRequestApps**](V3SpaceQuotasPostRequestApps.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**SpaceQuotaLinks**](SpaceQuotaLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**SpaceQuotaRelationships**](SpaceQuotaRelationships.md) |  | [optional] 
**Routes** | Pointer to [**V3SpaceQuotasPostRequestRoutes**](V3SpaceQuotasPostRequestRoutes.md) |  | [optional] 
**Services** | Pointer to [**V3SpaceQuotasPostRequestServices**](V3SpaceQuotasPostRequestServices.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSpaceQuota

`func NewSpaceQuota() *SpaceQuota`

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

`func (o *SpaceQuota) GetApps() V3SpaceQuotasPostRequestApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SpaceQuota) GetAppsOk() (*V3SpaceQuotasPostRequestApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SpaceQuota) SetApps(v V3SpaceQuotasPostRequestApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SpaceQuota) HasApps() bool`

HasApps returns a boolean if a field has been set.

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

### HasCreatedAt

`func (o *SpaceQuota) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasGuid

`func (o *SpaceQuota) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

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

### HasLinks

`func (o *SpaceQuota) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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

### HasName

`func (o *SpaceQuota) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasRelationships

`func (o *SpaceQuota) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRoutes

`func (o *SpaceQuota) GetRoutes() V3SpaceQuotasPostRequestRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SpaceQuota) GetRoutesOk() (*V3SpaceQuotasPostRequestRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SpaceQuota) SetRoutes(v V3SpaceQuotasPostRequestRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *SpaceQuota) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *SpaceQuota) GetServices() V3SpaceQuotasPostRequestServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SpaceQuota) GetServicesOk() (*V3SpaceQuotasPostRequestServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SpaceQuota) SetServices(v V3SpaceQuotasPostRequestServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *SpaceQuota) HasServices() bool`

HasServices returns a boolean if a field has been set.

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

### HasUpdatedAt

`func (o *SpaceQuota) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


