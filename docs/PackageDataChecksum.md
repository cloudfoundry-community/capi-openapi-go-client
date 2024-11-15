# PackageDataChecksum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The checksum type, for example: sha256 | [optional] 
**Value** | Pointer to **NullableString** | The checksum value; this will be populated after bits are uploaded | [optional] 

## Methods

### NewPackageDataChecksum

`func NewPackageDataChecksum() *PackageDataChecksum`

NewPackageDataChecksum instantiates a new PackageDataChecksum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDataChecksumWithDefaults

`func NewPackageDataChecksumWithDefaults() *PackageDataChecksum`

NewPackageDataChecksumWithDefaults instantiates a new PackageDataChecksum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PackageDataChecksum) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PackageDataChecksum) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PackageDataChecksum) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PackageDataChecksum) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *PackageDataChecksum) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PackageDataChecksum) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PackageDataChecksum) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PackageDataChecksum) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PackageDataChecksum) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PackageDataChecksum) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


