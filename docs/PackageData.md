# PackageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to [**PackageDataChecksum**](PackageDataChecksum.md) |  | [optional] 
**Error** | Pointer to **NullableString** | If an error occurs this field will contain the error message | [optional] 
**Image** | Pointer to **string** | The registry address of the image (for Docker packages) | [optional] 
**Password** | Pointer to **string** | The password for the image&#39;s registry (for Docker packages) | [optional] 
**Username** | Pointer to **string** | The username for the image&#39;s registry (for Docker packages) | [optional] 

## Methods

### NewPackageData

`func NewPackageData() *PackageData`

NewPackageData instantiates a new PackageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageDataWithDefaults

`func NewPackageDataWithDefaults() *PackageData`

NewPackageDataWithDefaults instantiates a new PackageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *PackageData) GetChecksum() PackageDataChecksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PackageData) GetChecksumOk() (*PackageDataChecksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PackageData) SetChecksum(v PackageDataChecksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *PackageData) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetError

`func (o *PackageData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PackageData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PackageData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PackageData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *PackageData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PackageData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetImage

`func (o *PackageData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PackageData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PackageData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PackageData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPassword

`func (o *PackageData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PackageData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PackageData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PackageData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *PackageData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PackageData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PackageData) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PackageData) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


