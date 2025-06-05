# ResourceMatchObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to [**ResourceMatchObjectChecksum**](ResourceMatchObjectChecksum.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SizeInBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourceMatchObject

`func NewResourceMatchObject() *ResourceMatchObject`

NewResourceMatchObject instantiates a new ResourceMatchObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMatchObjectWithDefaults

`func NewResourceMatchObjectWithDefaults() *ResourceMatchObject`

NewResourceMatchObjectWithDefaults instantiates a new ResourceMatchObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *ResourceMatchObject) GetChecksum() ResourceMatchObjectChecksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ResourceMatchObject) GetChecksumOk() (*ResourceMatchObjectChecksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ResourceMatchObject) SetChecksum(v ResourceMatchObjectChecksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ResourceMatchObject) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetMode

`func (o *ResourceMatchObject) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ResourceMatchObject) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ResourceMatchObject) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ResourceMatchObject) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *ResourceMatchObject) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceMatchObject) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceMatchObject) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ResourceMatchObject) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSizeInBytes

`func (o *ResourceMatchObject) GetSizeInBytes() int32`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *ResourceMatchObject) GetSizeInBytesOk() (*int32, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *ResourceMatchObject) SetSizeInBytes(v int32)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *ResourceMatchObject) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


