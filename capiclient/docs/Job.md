# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp when the job was created | 
**Errors** | Pointer to [**[]Error**](Error.md) | Errors that occurred during job processing | [optional] 
**Guid** | **string** | Unique identifier for the job | 
**Links** | [**JobLinks**](JobLinks.md) |  | 
**Metadata** | Pointer to [**JobMetadata**](JobMetadata.md) |  | [optional] 
**Operation** | **string** | Description of the operation being performed | 
**Resource** | Pointer to [**JobResource**](JobResource.md) |  | [optional] 
**State** | **string** | Current state of the job | 
**UpdatedAt** | **time.Time** | Timestamp when the job was last updated | 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | Warnings that occurred during job processing | [optional] 

## Methods

### NewJob

`func NewJob(createdAt time.Time, guid string, links JobLinks, operation string, state string, updatedAt time.Time, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrors

`func (o *Job) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Job) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Job) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Job) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetGuid

`func (o *Job) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Job) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Job) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *Job) GetLinks() JobLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Job) GetLinksOk() (*JobLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Job) SetLinks(v JobLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *Job) GetMetadata() JobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Job) GetMetadataOk() (*JobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Job) SetMetadata(v JobMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Job) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOperation

`func (o *Job) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Job) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Job) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetResource

`func (o *Job) GetResource() JobResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Job) GetResourceOk() (*JobResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Job) SetResource(v JobResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Job) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetState

`func (o *Job) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Job) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Job) SetState(v string)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWarnings

`func (o *Job) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Job) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Job) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Job) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


