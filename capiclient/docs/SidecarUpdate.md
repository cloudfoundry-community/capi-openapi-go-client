# SidecarUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Command to execute for the sidecar process | [optional] 
**MemoryInMb** | Pointer to **int32** | Memory allocation for the sidecar in MB | [optional] 
**Metadata** | Pointer to [**SidecarCreateMetadata**](SidecarCreateMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable name for the sidecar | [optional] 
**ProcessTypes** | Pointer to **[]string** | Process types this sidecar runs alongside | [optional] 

## Methods

### NewSidecarUpdate

`func NewSidecarUpdate() *SidecarUpdate`

NewSidecarUpdate instantiates a new SidecarUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidecarUpdateWithDefaults

`func NewSidecarUpdateWithDefaults() *SidecarUpdate`

NewSidecarUpdateWithDefaults instantiates a new SidecarUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *SidecarUpdate) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SidecarUpdate) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SidecarUpdate) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SidecarUpdate) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetMemoryInMb

`func (o *SidecarUpdate) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *SidecarUpdate) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *SidecarUpdate) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *SidecarUpdate) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetMetadata

`func (o *SidecarUpdate) GetMetadata() SidecarCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SidecarUpdate) GetMetadataOk() (*SidecarCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SidecarUpdate) SetMetadata(v SidecarCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SidecarUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SidecarUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SidecarUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SidecarUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SidecarUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessTypes

`func (o *SidecarUpdate) GetProcessTypes() []string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *SidecarUpdate) GetProcessTypesOk() (*[]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *SidecarUpdate) SetProcessTypes(v []string)`

SetProcessTypes sets ProcessTypes field to given value.

### HasProcessTypes

`func (o *SidecarUpdate) HasProcessTypes() bool`

HasProcessTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


