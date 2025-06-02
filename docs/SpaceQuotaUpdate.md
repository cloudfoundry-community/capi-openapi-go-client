# SpaceQuotaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**SpaceQuotaCreateApps**](SpaceQuotaCreateApps.md) |  | [optional] 
**Metadata** | Pointer to [**SpaceQuotaCreateMetadata**](SpaceQuotaCreateMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable name for the space quota | [optional] 
**Routes** | Pointer to [**SpaceQuotaCreateRoutes**](SpaceQuotaCreateRoutes.md) |  | [optional] 
**Services** | Pointer to [**SpaceQuotaUpdateServices**](SpaceQuotaUpdateServices.md) |  | [optional] 

## Methods

### NewSpaceQuotaUpdate

`func NewSpaceQuotaUpdate() *SpaceQuotaUpdate`

NewSpaceQuotaUpdate instantiates a new SpaceQuotaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaUpdateWithDefaults

`func NewSpaceQuotaUpdateWithDefaults() *SpaceQuotaUpdate`

NewSpaceQuotaUpdateWithDefaults instantiates a new SpaceQuotaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *SpaceQuotaUpdate) GetApps() SpaceQuotaCreateApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SpaceQuotaUpdate) GetAppsOk() (*SpaceQuotaCreateApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SpaceQuotaUpdate) SetApps(v SpaceQuotaCreateApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SpaceQuotaUpdate) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetMetadata

`func (o *SpaceQuotaUpdate) GetMetadata() SpaceQuotaCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpaceQuotaUpdate) GetMetadataOk() (*SpaceQuotaCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpaceQuotaUpdate) SetMetadata(v SpaceQuotaCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpaceQuotaUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SpaceQuotaUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceQuotaUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceQuotaUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceQuotaUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *SpaceQuotaUpdate) GetRoutes() SpaceQuotaCreateRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SpaceQuotaUpdate) GetRoutesOk() (*SpaceQuotaCreateRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SpaceQuotaUpdate) SetRoutes(v SpaceQuotaCreateRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *SpaceQuotaUpdate) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *SpaceQuotaUpdate) GetServices() SpaceQuotaUpdateServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SpaceQuotaUpdate) GetServicesOk() (*SpaceQuotaUpdateServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SpaceQuotaUpdate) SetServices(v SpaceQuotaUpdateServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *SpaceQuotaUpdate) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


