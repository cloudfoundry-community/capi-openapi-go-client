# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**AuthenticationCredentials**](AuthenticationCredentials.md) |  | 
**Type** | **string** |  | 

## Methods

### NewAuthentication

`func NewAuthentication(credentials AuthenticationCredentials, type_ string, ) *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *Authentication) GetCredentials() AuthenticationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Authentication) GetCredentialsOk() (*AuthenticationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Authentication) SetCredentials(v AuthenticationCredentials)`

SetCredentials sets Credentials field to given value.


### GetType

`func (o *Authentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Authentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Authentication) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


