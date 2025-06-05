# NotFoundError

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

### NewNotFoundError

`func NewNotFoundError(code float32, detail interface{}, title string, ) *NotFoundError`

NewNotFoundError instantiates a new NotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorWithDefaults

`func NewNotFoundErrorWithDefaults() *NotFoundError`

NewNotFoundErrorWithDefaults instantiates a new NotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacktrace

`func (o *NotFoundError) GetBacktrace() []string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *NotFoundError) GetBacktraceOk() (*[]string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *NotFoundError) SetBacktrace(v []string)`

SetBacktrace sets Backtrace field to given value.

### HasBacktrace

`func (o *NotFoundError) HasBacktrace() bool`

HasBacktrace returns a boolean if a field has been set.

### GetCode

`func (o *NotFoundError) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundError) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundError) SetCode(v float32)`

SetCode sets Code field to given value.


### GetDescription

`func (o *NotFoundError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotFoundError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotFoundError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotFoundError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *NotFoundError) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *NotFoundError) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *NotFoundError) SetDetail(v interface{})`

SetDetail sets Detail field to given value.


### SetDetailNil

`func (o *NotFoundError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *NotFoundError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetError

`func (o *NotFoundError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NotFoundError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NotFoundError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NotFoundError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTestModeInfo

`func (o *NotFoundError) GetTestModeInfo() ErrorTestModeInfo`

GetTestModeInfo returns the TestModeInfo field if non-nil, zero value otherwise.

### GetTestModeInfoOk

`func (o *NotFoundError) GetTestModeInfoOk() (*ErrorTestModeInfo, bool)`

GetTestModeInfoOk returns a tuple with the TestModeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModeInfo

`func (o *NotFoundError) SetTestModeInfo(v ErrorTestModeInfo)`

SetTestModeInfo sets TestModeInfo field to given value.

### HasTestModeInfo

`func (o *NotFoundError) HasTestModeInfo() bool`

HasTestModeInfo returns a boolean if a field has been set.

### GetTitle

`func (o *NotFoundError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotFoundError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotFoundError) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


