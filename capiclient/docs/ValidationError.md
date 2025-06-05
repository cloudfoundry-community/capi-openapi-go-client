# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backtrace** | Pointer to **[]string** | Stack trace (only in development environments) | [optional] 
**Code** | **float32** |  | 
**Description** | Pointer to **string** | Description of the error | [optional] 
**Detail** | **string** | Detailed description of the error | 
**Error** | Pointer to **string** | Name of the error | [optional] 
**TestModeInfo** | Pointer to [**ErrorTestModeInfo**](ErrorTestModeInfo.md) |  | [optional] 
**Title** | **string** |  | 
**Source** | Pointer to [**ValidationErrorAllOfSource**](ValidationErrorAllOfSource.md) |  | [optional] 

## Methods

### NewValidationError

`func NewValidationError(code float32, detail string, title string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacktrace

`func (o *ValidationError) GetBacktrace() []string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *ValidationError) GetBacktraceOk() (*[]string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *ValidationError) SetBacktrace(v []string)`

SetBacktrace sets Backtrace field to given value.

### HasBacktrace

`func (o *ValidationError) HasBacktrace() bool`

HasBacktrace returns a boolean if a field has been set.

### GetCode

`func (o *ValidationError) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationError) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationError) SetCode(v float32)`

SetCode sets Code field to given value.


### GetDescription

`func (o *ValidationError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidationError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidationError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ValidationError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetail

`func (o *ValidationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ValidationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ValidationError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetError

`func (o *ValidationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ValidationError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTestModeInfo

`func (o *ValidationError) GetTestModeInfo() ErrorTestModeInfo`

GetTestModeInfo returns the TestModeInfo field if non-nil, zero value otherwise.

### GetTestModeInfoOk

`func (o *ValidationError) GetTestModeInfoOk() (*ErrorTestModeInfo, bool)`

GetTestModeInfoOk returns a tuple with the TestModeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModeInfo

`func (o *ValidationError) SetTestModeInfo(v ErrorTestModeInfo)`

SetTestModeInfo sets TestModeInfo field to given value.

### HasTestModeInfo

`func (o *ValidationError) HasTestModeInfo() bool`

HasTestModeInfo returns a boolean if a field has been set.

### GetTitle

`func (o *ValidationError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ValidationError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ValidationError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSource

`func (o *ValidationError) GetSource() ValidationErrorAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ValidationError) GetSourceOk() (*ValidationErrorAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ValidationError) SetSource(v ValidationErrorAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ValidationError) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


