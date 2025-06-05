# Stack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildRootfsImage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ServiceUsageEventLinks**](ServiceUsageEventLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RunRootfsImage** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStack

`func NewStack() *Stack`

NewStack instantiates a new Stack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackWithDefaults

`func NewStackWithDefaults() *Stack`

NewStackWithDefaults instantiates a new Stack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildRootfsImage

`func (o *Stack) GetBuildRootfsImage() string`

GetBuildRootfsImage returns the BuildRootfsImage field if non-nil, zero value otherwise.

### GetBuildRootfsImageOk

`func (o *Stack) GetBuildRootfsImageOk() (*string, bool)`

GetBuildRootfsImageOk returns a tuple with the BuildRootfsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRootfsImage

`func (o *Stack) SetBuildRootfsImage(v string)`

SetBuildRootfsImage sets BuildRootfsImage field to given value.

### HasBuildRootfsImage

`func (o *Stack) HasBuildRootfsImage() bool`

HasBuildRootfsImage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Stack) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Stack) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Stack) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Stack) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *Stack) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Stack) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Stack) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Stack) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *Stack) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Stack) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Stack) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Stack) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuid

`func (o *Stack) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Stack) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Stack) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Stack) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *Stack) GetLinks() ServiceUsageEventLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Stack) GetLinksOk() (*ServiceUsageEventLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Stack) SetLinks(v ServiceUsageEventLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Stack) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Stack) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Stack) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Stack) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Stack) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Stack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Stack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunRootfsImage

`func (o *Stack) GetRunRootfsImage() string`

GetRunRootfsImage returns the RunRootfsImage field if non-nil, zero value otherwise.

### GetRunRootfsImageOk

`func (o *Stack) GetRunRootfsImageOk() (*string, bool)`

GetRunRootfsImageOk returns a tuple with the RunRootfsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunRootfsImage

`func (o *Stack) SetRunRootfsImage(v string)`

SetRunRootfsImage sets RunRootfsImage field to given value.

### HasRunRootfsImage

`func (o *Stack) HasRunRootfsImage() bool`

HasRunRootfsImage returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Stack) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Stack) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Stack) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Stack) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


