# ServiceBrokerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**Authentication**](Authentication.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**ServiceBrokerRelationships**](ServiceBrokerRelationships.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceBrokerCreate

`func NewServiceBrokerCreate() *ServiceBrokerCreate`

NewServiceBrokerCreate instantiates a new ServiceBrokerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBrokerCreateWithDefaults

`func NewServiceBrokerCreateWithDefaults() *ServiceBrokerCreate`

NewServiceBrokerCreateWithDefaults instantiates a new ServiceBrokerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *ServiceBrokerCreate) GetAuthentication() Authentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ServiceBrokerCreate) GetAuthenticationOk() (*Authentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ServiceBrokerCreate) SetAuthentication(v Authentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ServiceBrokerCreate) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetName

`func (o *ServiceBrokerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceBrokerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceBrokerCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceBrokerCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *ServiceBrokerCreate) GetRelationships() ServiceBrokerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ServiceBrokerCreate) GetRelationshipsOk() (*ServiceBrokerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ServiceBrokerCreate) SetRelationships(v ServiceBrokerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ServiceBrokerCreate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceBrokerCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceBrokerCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceBrokerCreate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceBrokerCreate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


