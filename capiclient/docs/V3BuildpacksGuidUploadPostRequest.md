# V3BuildpacksGuidUploadPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Docker image reference for Cloud Native Buildpack (experimental) | [optional] 
**Password** | Pointer to **string** | Password for private registry authentication | [optional] 
**Username** | Pointer to **string** | Username for private registry authentication | [optional] 

## Methods

### NewV3BuildpacksGuidUploadPostRequest

`func NewV3BuildpacksGuidUploadPostRequest() *V3BuildpacksGuidUploadPostRequest`

NewV3BuildpacksGuidUploadPostRequest instantiates a new V3BuildpacksGuidUploadPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3BuildpacksGuidUploadPostRequestWithDefaults

`func NewV3BuildpacksGuidUploadPostRequestWithDefaults() *V3BuildpacksGuidUploadPostRequest`

NewV3BuildpacksGuidUploadPostRequestWithDefaults instantiates a new V3BuildpacksGuidUploadPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V3BuildpacksGuidUploadPostRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V3BuildpacksGuidUploadPostRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V3BuildpacksGuidUploadPostRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V3BuildpacksGuidUploadPostRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPassword

`func (o *V3BuildpacksGuidUploadPostRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V3BuildpacksGuidUploadPostRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V3BuildpacksGuidUploadPostRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V3BuildpacksGuidUploadPostRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *V3BuildpacksGuidUploadPostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V3BuildpacksGuidUploadPostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V3BuildpacksGuidUploadPostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V3BuildpacksGuidUploadPostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


