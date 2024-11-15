# Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deployable** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Droplet** | Pointer to [**V3AppsPostRequestRelationshipsSpaceData**](V3AppsPostRequestRelationshipsSpaceData.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**RevisionLinks**](RevisionLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Processes** | Pointer to [**map[string]RevisionProcessesValue**](RevisionProcessesValue.md) |  | [optional] 
**Relationships** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseRelationships**](V3AppsGuidDropletsCurrentGet200ResponseRelationships.md) |  | [optional] 
**Sidecars** | Pointer to [**[]Sidecar**](Sidecar.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewRevision

`func NewRevision() *Revision`

NewRevision instantiates a new Revision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionWithDefaults

`func NewRevisionWithDefaults() *Revision`

NewRevisionWithDefaults instantiates a new Revision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Revision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Revision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Revision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Revision) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployable

`func (o *Revision) GetDeployable() bool`

GetDeployable returns the Deployable field if non-nil, zero value otherwise.

### GetDeployableOk

`func (o *Revision) GetDeployableOk() (*bool, bool)`

GetDeployableOk returns a tuple with the Deployable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployable

`func (o *Revision) SetDeployable(v bool)`

SetDeployable sets Deployable field to given value.

### HasDeployable

`func (o *Revision) HasDeployable() bool`

HasDeployable returns a boolean if a field has been set.

### GetDescription

`func (o *Revision) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Revision) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Revision) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Revision) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDroplet

`func (o *Revision) GetDroplet() V3AppsPostRequestRelationshipsSpaceData`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *Revision) GetDropletOk() (*V3AppsPostRequestRelationshipsSpaceData, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *Revision) SetDroplet(v V3AppsPostRequestRelationshipsSpaceData)`

SetDroplet sets Droplet field to given value.

### HasDroplet

`func (o *Revision) HasDroplet() bool`

HasDroplet returns a boolean if a field has been set.

### GetGuid

`func (o *Revision) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Revision) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Revision) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Revision) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetLinks

`func (o *Revision) GetLinks() RevisionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Revision) GetLinksOk() (*RevisionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Revision) SetLinks(v RevisionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Revision) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Revision) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Revision) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Revision) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Revision) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcesses

`func (o *Revision) GetProcesses() map[string]RevisionProcessesValue`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *Revision) GetProcessesOk() (*map[string]RevisionProcessesValue, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *Revision) SetProcesses(v map[string]RevisionProcessesValue)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *Revision) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetRelationships

`func (o *Revision) GetRelationships() V3AppsGuidDropletsCurrentGet200ResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Revision) GetRelationshipsOk() (*V3AppsGuidDropletsCurrentGet200ResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Revision) SetRelationships(v V3AppsGuidDropletsCurrentGet200ResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Revision) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetSidecars

`func (o *Revision) GetSidecars() []Sidecar`

GetSidecars returns the Sidecars field if non-nil, zero value otherwise.

### GetSidecarsOk

`func (o *Revision) GetSidecarsOk() (*[]Sidecar, bool)`

GetSidecarsOk returns a tuple with the Sidecars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidecars

`func (o *Revision) SetSidecars(v []Sidecar)`

SetSidecars sets Sidecars field to given value.

### HasSidecars

`func (o *Revision) HasSidecars() bool`

HasSidecars returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Revision) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Revision) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Revision) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Revision) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *Revision) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Revision) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Revision) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Revision) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


