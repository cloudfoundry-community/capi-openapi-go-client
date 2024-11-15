# V3SpaceQuotasPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**V3SpaceQuotasPostRequestApps**](V3SpaceQuotasPostRequestApps.md) |  | [optional] 
**Name** | **string** |  | 
**Relationships** | [**V3SpaceQuotasPostRequestRelationships**](V3SpaceQuotasPostRequestRelationships.md) |  | 
**Routes** | Pointer to [**V3SpaceQuotasPostRequestRoutes**](V3SpaceQuotasPostRequestRoutes.md) |  | [optional] 
**Services** | Pointer to [**V3SpaceQuotasPostRequestServices**](V3SpaceQuotasPostRequestServices.md) |  | [optional] 

## Methods

### NewV3SpaceQuotasPostRequest

`func NewV3SpaceQuotasPostRequest(name string, relationships V3SpaceQuotasPostRequestRelationships, ) *V3SpaceQuotasPostRequest`

NewV3SpaceQuotasPostRequest instantiates a new V3SpaceQuotasPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SpaceQuotasPostRequestWithDefaults

`func NewV3SpaceQuotasPostRequestWithDefaults() *V3SpaceQuotasPostRequest`

NewV3SpaceQuotasPostRequestWithDefaults instantiates a new V3SpaceQuotasPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *V3SpaceQuotasPostRequest) GetApps() V3SpaceQuotasPostRequestApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *V3SpaceQuotasPostRequest) GetAppsOk() (*V3SpaceQuotasPostRequestApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *V3SpaceQuotasPostRequest) SetApps(v V3SpaceQuotasPostRequestApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *V3SpaceQuotasPostRequest) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetName

`func (o *V3SpaceQuotasPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3SpaceQuotasPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3SpaceQuotasPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *V3SpaceQuotasPostRequest) GetRelationships() V3SpaceQuotasPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3SpaceQuotasPostRequest) GetRelationshipsOk() (*V3SpaceQuotasPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3SpaceQuotasPostRequest) SetRelationships(v V3SpaceQuotasPostRequestRelationships)`

SetRelationships sets Relationships field to given value.


### GetRoutes

`func (o *V3SpaceQuotasPostRequest) GetRoutes() V3SpaceQuotasPostRequestRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *V3SpaceQuotasPostRequest) GetRoutesOk() (*V3SpaceQuotasPostRequestRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *V3SpaceQuotasPostRequest) SetRoutes(v V3SpaceQuotasPostRequestRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *V3SpaceQuotasPostRequest) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *V3SpaceQuotasPostRequest) GetServices() V3SpaceQuotasPostRequestServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *V3SpaceQuotasPostRequest) GetServicesOk() (*V3SpaceQuotasPostRequestServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *V3SpaceQuotasPostRequest) SetServices(v V3SpaceQuotasPostRequestServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *V3SpaceQuotasPostRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


