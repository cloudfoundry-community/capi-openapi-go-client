# IsolationSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**IsolationSegmentLinks**](IsolationSegmentLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIsolationSegment

`func NewIsolationSegment() *IsolationSegment`

NewIsolationSegment instantiates a new IsolationSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolationSegmentWithDefaults

`func NewIsolationSegmentWithDefaults() *IsolationSegment`

NewIsolationSegmentWithDefaults instantiates a new IsolationSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IsolationSegment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IsolationSegment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IsolationSegment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IsolationSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGuid

`func (o *IsolationSegment) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IsolationSegment) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IsolationSegment) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IsolationSegment) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *IsolationSegment) GetLinks() IsolationSegmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IsolationSegment) GetLinksOk() (*IsolationSegmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IsolationSegment) SetLinks(v IsolationSegmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IsolationSegment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *IsolationSegment) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IsolationSegment) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IsolationSegment) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IsolationSegment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IsolationSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IsolationSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IsolationSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IsolationSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IsolationSegment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IsolationSegment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IsolationSegment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IsolationSegment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


