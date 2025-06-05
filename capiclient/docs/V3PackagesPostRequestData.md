# V3PackagesPostRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Docker image URL or OCI image reference (docker/cnb types) | [optional] 
**Password** | Pointer to **string** | Password for private registry authentication | [optional] 
**Username** | Pointer to **string** | Username for private registry authentication | [optional] 

## Methods

### NewV3PackagesPostRequestData

`func NewV3PackagesPostRequestData() *V3PackagesPostRequestData`

NewV3PackagesPostRequestData instantiates a new V3PackagesPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3PackagesPostRequestDataWithDefaults

`func NewV3PackagesPostRequestDataWithDefaults() *V3PackagesPostRequestData`

NewV3PackagesPostRequestDataWithDefaults instantiates a new V3PackagesPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V3PackagesPostRequestData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V3PackagesPostRequestData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V3PackagesPostRequestData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V3PackagesPostRequestData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPassword

`func (o *V3PackagesPostRequestData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V3PackagesPostRequestData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V3PackagesPostRequestData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V3PackagesPostRequestData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *V3PackagesPostRequestData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V3PackagesPostRequestData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V3PackagesPostRequestData) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V3PackagesPostRequestData) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


