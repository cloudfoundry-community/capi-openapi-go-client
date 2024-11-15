# V3AppsGuidSidecarsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**MemoryInMb** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**ProcessTypes** | **[]string** |  | 

## Methods

### NewV3AppsGuidSidecarsPostRequest

`func NewV3AppsGuidSidecarsPostRequest(command string, name string, processTypes []string, ) *V3AppsGuidSidecarsPostRequest`

NewV3AppsGuidSidecarsPostRequest instantiates a new V3AppsGuidSidecarsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidSidecarsPostRequestWithDefaults

`func NewV3AppsGuidSidecarsPostRequestWithDefaults() *V3AppsGuidSidecarsPostRequest`

NewV3AppsGuidSidecarsPostRequestWithDefaults instantiates a new V3AppsGuidSidecarsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *V3AppsGuidSidecarsPostRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V3AppsGuidSidecarsPostRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V3AppsGuidSidecarsPostRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetMemoryInMb

`func (o *V3AppsGuidSidecarsPostRequest) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *V3AppsGuidSidecarsPostRequest) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *V3AppsGuidSidecarsPostRequest) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *V3AppsGuidSidecarsPostRequest) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetName

`func (o *V3AppsGuidSidecarsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3AppsGuidSidecarsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3AppsGuidSidecarsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProcessTypes

`func (o *V3AppsGuidSidecarsPostRequest) GetProcessTypes() []string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *V3AppsGuidSidecarsPostRequest) GetProcessTypesOk() (*[]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *V3AppsGuidSidecarsPostRequest) SetProcessTypes(v []string)`

SetProcessTypes sets ProcessTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


