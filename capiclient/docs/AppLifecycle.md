# AppLifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppLifecycleData**](AppLifecycleData.md) |  | 
**Type** | **string** | Lifecycle type | 

## Methods

### NewAppLifecycle

`func NewAppLifecycle(data AppLifecycleData, type_ string, ) *AppLifecycle`

NewAppLifecycle instantiates a new AppLifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLifecycleWithDefaults

`func NewAppLifecycleWithDefaults() *AppLifecycle`

NewAppLifecycleWithDefaults instantiates a new AppLifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppLifecycle) GetData() AppLifecycleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppLifecycle) GetDataOk() (*AppLifecycleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppLifecycle) SetData(v AppLifecycleData)`

SetData sets Data field to given value.


### GetType

`func (o *AppLifecycle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLifecycle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLifecycle) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


