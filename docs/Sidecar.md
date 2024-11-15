# Sidecar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**MemoryInMb** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**ProcessTypes** | Pointer to **[]string** |  | [optional] 
**Relationships** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseRelationships**](V3AppsGuidDropletsCurrentGet200ResponseRelationships.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSidecar

`func NewSidecar() *Sidecar`

NewSidecar instantiates a new Sidecar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidecarWithDefaults

`func NewSidecarWithDefaults() *Sidecar`

NewSidecarWithDefaults instantiates a new Sidecar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Sidecar) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Sidecar) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Sidecar) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Sidecar) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Sidecar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Sidecar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Sidecar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Sidecar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *Sidecar) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Sidecar) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Sidecar) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Sidecar) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetMemoryInMb

`func (o *Sidecar) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *Sidecar) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *Sidecar) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *Sidecar) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetName

`func (o *Sidecar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sidecar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sidecar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sidecar) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *Sidecar) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Sidecar) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Sidecar) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Sidecar) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetProcessTypes

`func (o *Sidecar) GetProcessTypes() []string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *Sidecar) GetProcessTypesOk() (*[]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *Sidecar) SetProcessTypes(v []string)`

SetProcessTypes sets ProcessTypes field to given value.

### HasProcessTypes

`func (o *Sidecar) HasProcessTypes() bool`

HasProcessTypes returns a boolean if a field has been set.

### GetRelationships

`func (o *Sidecar) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Sidecar) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Sidecar) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Sidecar) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Sidecar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Sidecar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Sidecar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Sidecar) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


