# Droplet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buildpacks** | Pointer to [**[]Buildpack**](Buildpack.md) |  | [optional] 
**Checksum** | Pointer to [**DropletChecksum**](DropletChecksum.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ExecutionMetadata** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Lifecycle** | Pointer to [**DropletLifecycle**](DropletLifecycle.md) |  | [optional] 
**Links** | Pointer to [**DropletLinks**](DropletLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3DropletsGuidPatchRequestMetadata**](V3DropletsGuidPatchRequestMetadata.md) |  | [optional] 
**ProcessTypes** | Pointer to **map[string]string** |  | [optional] 
**Relationships** | Pointer to [**V3DropletsPostRequestRelationships**](V3DropletsPostRequestRelationships.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDroplet

`func NewDroplet() *Droplet`

NewDroplet instantiates a new Droplet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropletWithDefaults

`func NewDropletWithDefaults() *Droplet`

NewDropletWithDefaults instantiates a new Droplet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildpacks

`func (o *Droplet) GetBuildpacks() []Buildpack`

GetBuildpacks returns the Buildpacks field if non-nil, zero value otherwise.

### GetBuildpacksOk

`func (o *Droplet) GetBuildpacksOk() (*[]Buildpack, bool)`

GetBuildpacksOk returns a tuple with the Buildpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpacks

`func (o *Droplet) SetBuildpacks(v []Buildpack)`

SetBuildpacks sets Buildpacks field to given value.

### HasBuildpacks

`func (o *Droplet) HasBuildpacks() bool`

HasBuildpacks returns a boolean if a field has been set.

### GetChecksum

`func (o *Droplet) GetChecksum() DropletChecksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *Droplet) GetChecksumOk() (*DropletChecksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *Droplet) SetChecksum(v DropletChecksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *Droplet) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Droplet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Droplet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Droplet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Droplet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *Droplet) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Droplet) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Droplet) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Droplet) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecutionMetadata

`func (o *Droplet) GetExecutionMetadata() string`

GetExecutionMetadata returns the ExecutionMetadata field if non-nil, zero value otherwise.

### GetExecutionMetadataOk

`func (o *Droplet) GetExecutionMetadataOk() (*string, bool)`

GetExecutionMetadataOk returns a tuple with the ExecutionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMetadata

`func (o *Droplet) SetExecutionMetadata(v string)`

SetExecutionMetadata sets ExecutionMetadata field to given value.

### HasExecutionMetadata

`func (o *Droplet) HasExecutionMetadata() bool`

HasExecutionMetadata returns a boolean if a field has been set.

### GetGuid

`func (o *Droplet) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Droplet) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Droplet) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Droplet) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetImage

`func (o *Droplet) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Droplet) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Droplet) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Droplet) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLifecycle

`func (o *Droplet) GetLifecycle() DropletLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *Droplet) GetLifecycleOk() (*DropletLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *Droplet) SetLifecycle(v DropletLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *Droplet) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLinks

`func (o *Droplet) GetLinks() DropletLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Droplet) GetLinksOk() (*DropletLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Droplet) SetLinks(v DropletLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Droplet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Droplet) GetMetadata() V3DropletsGuidPatchRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Droplet) GetMetadataOk() (*V3DropletsGuidPatchRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Droplet) SetMetadata(v V3DropletsGuidPatchRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Droplet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessTypes

`func (o *Droplet) GetProcessTypes() map[string]string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *Droplet) GetProcessTypesOk() (*map[string]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *Droplet) SetProcessTypes(v map[string]string)`

SetProcessTypes sets ProcessTypes field to given value.

### HasProcessTypes

`func (o *Droplet) HasProcessTypes() bool`

HasProcessTypes returns a boolean if a field has been set.

### GetRelationships

`func (o *Droplet) GetRelationships() V3DropletsPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Droplet) GetRelationshipsOk() (*V3DropletsPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Droplet) SetRelationships(v V3DropletsPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Droplet) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetStack

`func (o *Droplet) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *Droplet) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *Droplet) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *Droplet) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetState

`func (o *Droplet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Droplet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Droplet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Droplet) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Droplet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Droplet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Droplet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Droplet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


