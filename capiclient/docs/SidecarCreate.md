# SidecarCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Command to execute for the sidecar process | 
**MemoryInMb** | Pointer to **int32** | Memory allocation for the sidecar in MB | [optional] [default to 256]
**Metadata** | Pointer to [**SidecarCreateMetadata**](SidecarCreateMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the sidecar | 
**ProcessTypes** | **[]string** | Process types this sidecar runs alongside | 

## Methods

### NewSidecarCreate

`func NewSidecarCreate(command string, name string, processTypes []string, ) *SidecarCreate`

NewSidecarCreate instantiates a new SidecarCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidecarCreateWithDefaults

`func NewSidecarCreateWithDefaults() *SidecarCreate`

NewSidecarCreateWithDefaults instantiates a new SidecarCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *SidecarCreate) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SidecarCreate) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SidecarCreate) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetMemoryInMb

`func (o *SidecarCreate) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *SidecarCreate) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *SidecarCreate) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *SidecarCreate) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetMetadata

`func (o *SidecarCreate) GetMetadata() SidecarCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SidecarCreate) GetMetadataOk() (*SidecarCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SidecarCreate) SetMetadata(v SidecarCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SidecarCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SidecarCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SidecarCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SidecarCreate) SetName(v string)`

SetName sets Name field to given value.


### GetProcessTypes

`func (o *SidecarCreate) GetProcessTypes() []string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *SidecarCreate) GetProcessTypesOk() (*[]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *SidecarCreate) SetProcessTypes(v []string)`

SetProcessTypes sets ProcessTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


