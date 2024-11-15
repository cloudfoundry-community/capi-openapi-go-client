# V3AppsGuidDropletsCurrentGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to [**[]V3AppsGuidDropletsCurrentGet200ResponseBuildpacksInner**](V3AppsGuidDropletsCurrentGet200ResponseBuildpacksInner.md) |  | [optional] 
**Checksum** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseChecksum**](V3AppsGuidDropletsCurrentGet200ResponseChecksum.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**ExecutionMetadata** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **NullableString** |  | [optional] 
**Lifecycle** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseLifecycle**](V3AppsGuidDropletsCurrentGet200ResponseLifecycle.md) |  | [optional] 
**Links** | Pointer to [**map[string]V3AppsGuidDropletsCurrentGet200ResponseLinksValue**](V3AppsGuidDropletsCurrentGet200ResponseLinksValue.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**ProcessTypes** | Pointer to **map[string]string** |  | [optional] 
**Relationships** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseRelationships**](V3AppsGuidDropletsCurrentGet200ResponseRelationships.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV3AppsGuidDropletsCurrentGet200Response

`func NewV3AppsGuidDropletsCurrentGet200Response() *V3AppsGuidDropletsCurrentGet200Response`

NewV3AppsGuidDropletsCurrentGet200Response instantiates a new V3AppsGuidDropletsCurrentGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidDropletsCurrentGet200ResponseWithDefaults

`func NewV3AppsGuidDropletsCurrentGet200ResponseWithDefaults() *V3AppsGuidDropletsCurrentGet200Response`

NewV3AppsGuidDropletsCurrentGet200ResponseWithDefaults instantiates a new V3AppsGuidDropletsCurrentGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetBuildpacks() []V3AppsGuidDropletsCurrentGet200ResponseBuildpacksInner`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetBuildpacksOk() (*[]V3AppsGuidDropletsCurrentGet200ResponseBuildpacksInner, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetBuildpacks(v []V3AppsGuidDropletsCurrentGet200ResponseBuildpacksInner)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetChecksum

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetChecksum() V3AppsGuidDropletsCurrentGet200ResponseChecksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetChecksumOk() (*V3AppsGuidDropletsCurrentGet200ResponseChecksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetChecksum(v V3AppsGuidDropletsCurrentGet200ResponseChecksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *V3AppsGuidDropletsCurrentGet200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExecutionMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetExecutionMetadata() string`

GetExecutionMetadata returns the ExecutionMetadata field if non-nil, zero value otherwise.

### GetExecutionMetadataOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetExecutionMetadataOk() (*string, bool)`

GetExecutionMetadataOk returns a tuple with the ExecutionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetExecutionMetadata(v string)`

SetExecutionMetadata sets ExecutionMetadata field to given value.

### HasExecutionMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasExecutionMetadata() bool`

HasExecutionMetadata returns a boolean if a field has been set.

### GetGuid

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetImage

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *V3AppsGuidDropletsCurrentGet200Response) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetLifecycle

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetLifecycle() V3AppsGuidDropletsCurrentGet200ResponseLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetLifecycleOk() (*V3AppsGuidDropletsCurrentGet200ResponseLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetLifecycle(v V3AppsGuidDropletsCurrentGet200ResponseLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLinks

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetLinks() map[string]V3AppsGuidDropletsCurrentGet200ResponseLinksValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetLinksOk() (*map[string]V3AppsGuidDropletsCurrentGet200ResponseLinksValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetLinks(v map[string]V3AppsGuidDropletsCurrentGet200ResponseLinksValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessTypes

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetProcessTypes() map[string]string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetProcessTypesOk() (*map[string]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetProcessTypes(v map[string]string)`

SetProcessTypes sets ProcessTypes field to given value.

### HasProcessTypes

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasProcessTypes() bool`

HasProcessTypes returns a boolean if a field has been set.

### GetRelationships

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetStack

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetState

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V3AppsGuidDropletsCurrentGet200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V3AppsGuidDropletsCurrentGet200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


