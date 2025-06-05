# ErrorTestModeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backtrace** | Pointer to **[]string** | Stack trace of the error | [optional] 
**Source** | Pointer to [**ErrorTestModeInfoSource**](ErrorTestModeInfoSource.md) |  | [optional] 

## Methods

### NewErrorTestModeInfo

`func NewErrorTestModeInfo() *ErrorTestModeInfo`

NewErrorTestModeInfo instantiates a new ErrorTestModeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTestModeInfoWithDefaults

`func NewErrorTestModeInfoWithDefaults() *ErrorTestModeInfo`

NewErrorTestModeInfoWithDefaults instantiates a new ErrorTestModeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacktrace

`func (o *ErrorTestModeInfo) GetBacktrace() []string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *ErrorTestModeInfo) GetBacktraceOk() (*[]string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *ErrorTestModeInfo) SetBacktrace(v []string)`

SetBacktrace sets Backtrace field to given value.

### HasBacktrace

`func (o *ErrorTestModeInfo) HasBacktrace() bool`

HasBacktrace returns a boolean if a field has been set.

### GetSource

`func (o *ErrorTestModeInfo) GetSource() ErrorTestModeInfoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorTestModeInfo) GetSourceOk() (*ErrorTestModeInfoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorTestModeInfo) SetSource(v ErrorTestModeInfoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorTestModeInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


