# RelationshipError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backtrace** | Pointer to **[]string** | Stack trace (only in development environments) | [optional] 
**Code** | **float32** |  | 
**Description** | Pointer to **string** | Description of the error | [optional] 
**Detail** | **interface{}** |  | 
**Error** | Pointer to **string** | Name of the error | [optional] 
**TestModeInfo** | Pointer to [**ErrorTestModeInfo**](ErrorTestModeInfo.md) |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewRelationshipError

`func NewRelationshipError(code float32, detail interface{}, title string, ) *RelationshipError`

NewRelationshipError instantiates a new RelationshipError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipErrorWithDefaults

`func NewRelationshipErrorWithDefaults() *RelationshipError`

NewRelationshipErrorWithDefaults instantiates a new RelationshipError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacktrace

`func (o *RelationshipError) GetBacktrace() []string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *RelationshipError) GetBacktraceOk() (*[]string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *RelationshipError) SetBacktrace(v []string)`

SetBacktrace sets Backtrace field to given value.

### HasBacktrace

`func (o *RelationshipError) HasBacktrace() bool`

HasBacktrace returns a boolean if a field has been set.

### GetCode

`func (o *RelationshipError) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RelationshipError) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RelationshipError) SetCode(v float32)`

SetCode sets Code field to given value.


### GetDescription

`func (o *RelationshipError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelationshipError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelationshipError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelationshipError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *RelationshipError) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RelationshipError) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RelationshipError) SetDetail(v interface{})`

SetDetail sets Detail field to given value.


### SetDetailNil

`func (o *RelationshipError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *RelationshipError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetError

`func (o *RelationshipError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RelationshipError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RelationshipError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RelationshipError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTestModeInfo

`func (o *RelationshipError) GetTestModeInfo() ErrorTestModeInfo`

GetTestModeInfo returns the TestModeInfo field if non-nil, zero value otherwise.

### GetTestModeInfoOk

`func (o *RelationshipError) GetTestModeInfoOk() (*ErrorTestModeInfo, bool)`

GetTestModeInfoOk returns a tuple with the TestModeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModeInfo

`func (o *RelationshipError) SetTestModeInfo(v ErrorTestModeInfo)`

SetTestModeInfo sets TestModeInfo field to given value.

### HasTestModeInfo

`func (o *RelationshipError) HasTestModeInfo() bool`

HasTestModeInfo returns a boolean if a field has been set.

### GetTitle

`func (o *RelationshipError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RelationshipError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RelationshipError) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


