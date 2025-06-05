# AsyncServiceError

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

### NewAsyncServiceError

`func NewAsyncServiceError(code float32, detail interface{}, title string, ) *AsyncServiceError`

NewAsyncServiceError instantiates a new AsyncServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncServiceErrorWithDefaults

`func NewAsyncServiceErrorWithDefaults() *AsyncServiceError`

NewAsyncServiceErrorWithDefaults instantiates a new AsyncServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacktrace

`func (o *AsyncServiceError) GetBacktrace() []string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *AsyncServiceError) GetBacktraceOk() (*[]string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *AsyncServiceError) SetBacktrace(v []string)`

SetBacktrace sets Backtrace field to given value.

### HasBacktrace

`func (o *AsyncServiceError) HasBacktrace() bool`

HasBacktrace returns a boolean if a field has been set.

### GetCode

`func (o *AsyncServiceError) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AsyncServiceError) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AsyncServiceError) SetCode(v float32)`

SetCode sets Code field to given value.


### GetDescription

`func (o *AsyncServiceError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AsyncServiceError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AsyncServiceError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AsyncServiceError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *AsyncServiceError) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AsyncServiceError) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AsyncServiceError) SetDetail(v interface{})`

SetDetail sets Detail field to given value.


### SetDetailNil

`func (o *AsyncServiceError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *AsyncServiceError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetError

`func (o *AsyncServiceError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AsyncServiceError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AsyncServiceError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AsyncServiceError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTestModeInfo

`func (o *AsyncServiceError) GetTestModeInfo() ErrorTestModeInfo`

GetTestModeInfo returns the TestModeInfo field if non-nil, zero value otherwise.

### GetTestModeInfoOk

`func (o *AsyncServiceError) GetTestModeInfoOk() (*ErrorTestModeInfo, bool)`

GetTestModeInfoOk returns a tuple with the TestModeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModeInfo

`func (o *AsyncServiceError) SetTestModeInfo(v ErrorTestModeInfo)`

SetTestModeInfo sets TestModeInfo field to given value.

### HasTestModeInfo

`func (o *AsyncServiceError) HasTestModeInfo() bool`

HasTestModeInfo returns a boolean if a field has been set.

### GetTitle

`func (o *AsyncServiceError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AsyncServiceError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AsyncServiceError) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


