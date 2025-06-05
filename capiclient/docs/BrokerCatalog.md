# BrokerCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**BrokerCatalogFeatures**](BrokerCatalogFeatures.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximumPollingDuration** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to [**BrokerCatalogMetadata**](BrokerCatalogMetadata.md) |  | [optional] 

## Methods

### NewBrokerCatalog

`func NewBrokerCatalog() *BrokerCatalog`

NewBrokerCatalog instantiates a new BrokerCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerCatalogWithDefaults

`func NewBrokerCatalogWithDefaults() *BrokerCatalog`

NewBrokerCatalogWithDefaults instantiates a new BrokerCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *BrokerCatalog) GetFeatures() BrokerCatalogFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BrokerCatalog) GetFeaturesOk() (*BrokerCatalogFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BrokerCatalog) SetFeatures(v BrokerCatalogFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BrokerCatalog) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *BrokerCatalog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrokerCatalog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrokerCatalog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BrokerCatalog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximumPollingDuration

`func (o *BrokerCatalog) GetMaximumPollingDuration() int32`

GetMaximumPollingDuration returns the MaximumPollingDuration field if non-nil, zero value otherwise.

### GetMaximumPollingDurationOk

`func (o *BrokerCatalog) GetMaximumPollingDurationOk() (*int32, bool)`

GetMaximumPollingDurationOk returns a tuple with the MaximumPollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPollingDuration

`func (o *BrokerCatalog) SetMaximumPollingDuration(v int32)`

SetMaximumPollingDuration sets MaximumPollingDuration field to given value.

### HasMaximumPollingDuration

`func (o *BrokerCatalog) HasMaximumPollingDuration() bool`

HasMaximumPollingDuration returns a boolean if a field has been set.

### SetMaximumPollingDurationNil

`func (o *BrokerCatalog) SetMaximumPollingDurationNil(b bool)`

 SetMaximumPollingDurationNil sets the value for MaximumPollingDuration to be an explicit nil

### UnsetMaximumPollingDuration
`func (o *BrokerCatalog) UnsetMaximumPollingDuration()`

UnsetMaximumPollingDuration ensures that no value is present for MaximumPollingDuration, not even an explicit nil
### GetMetadata

`func (o *BrokerCatalog) GetMetadata() BrokerCatalogMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerCatalog) GetMetadataOk() (*BrokerCatalogMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerCatalog) SetMetadata(v BrokerCatalogMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BrokerCatalog) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


