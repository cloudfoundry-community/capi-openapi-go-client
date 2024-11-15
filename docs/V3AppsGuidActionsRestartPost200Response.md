# V3AppsGuidActionsRestartPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Lifecycle** | Pointer to [**V3AppsGuidActionsRestartPost200ResponseLifecycle**](V3AppsGuidActionsRestartPost200ResponseLifecycle.md) |  | [optional] 
**Links** | Pointer to [**map[string]Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**V3AppsGuidActionsRestartPost200ResponseRelationships**](V3AppsGuidActionsRestartPost200ResponseRelationships.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV3AppsGuidActionsRestartPost200Response

`func NewV3AppsGuidActionsRestartPost200Response() *V3AppsGuidActionsRestartPost200Response`

NewV3AppsGuidActionsRestartPost200Response instantiates a new V3AppsGuidActionsRestartPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidActionsRestartPost200ResponseWithDefaults

`func NewV3AppsGuidActionsRestartPost200ResponseWithDefaults() *V3AppsGuidActionsRestartPost200Response`

NewV3AppsGuidActionsRestartPost200ResponseWithDefaults instantiates a new V3AppsGuidActionsRestartPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *V3AppsGuidActionsRestartPost200Response) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *V3AppsGuidActionsRestartPost200Response) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *V3AppsGuidActionsRestartPost200Response) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLifecycle

`func (o *V3AppsGuidActionsRestartPost200Response) GetLifecycle() V3AppsGuidActionsRestartPost200ResponseLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetLifecycleOk() (*V3AppsGuidActionsRestartPost200ResponseLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3AppsGuidActionsRestartPost200Response) SetLifecycle(v V3AppsGuidActionsRestartPost200ResponseLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3AppsGuidActionsRestartPost200Response) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLinks

`func (o *V3AppsGuidActionsRestartPost200Response) GetLinks() map[string]Get200ResponseLinksLogCache`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetLinksOk() (*map[string]Get200ResponseLinksLogCache, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *V3AppsGuidActionsRestartPost200Response) SetLinks(v map[string]Get200ResponseLinksLogCache)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *V3AppsGuidActionsRestartPost200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *V3AppsGuidActionsRestartPost200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3AppsGuidActionsRestartPost200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3AppsGuidActionsRestartPost200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *V3AppsGuidActionsRestartPost200Response) GetRelationships() V3AppsGuidActionsRestartPost200ResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetRelationshipsOk() (*V3AppsGuidActionsRestartPost200ResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3AppsGuidActionsRestartPost200Response) SetRelationships(v V3AppsGuidActionsRestartPost200ResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3AppsGuidActionsRestartPost200Response) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetState

`func (o *V3AppsGuidActionsRestartPost200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V3AppsGuidActionsRestartPost200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V3AppsGuidActionsRestartPost200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V3AppsGuidActionsRestartPost200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V3AppsGuidActionsRestartPost200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


