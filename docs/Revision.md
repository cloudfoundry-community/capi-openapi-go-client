# Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Deployable** | **bool** | Whether this revision can be deployed | 
**Description** | **NullableString** |  | 
**Droplet** | [**V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData**](V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData.md) |  | 
**Guid** | **string** |  | 
**Links** | [**RevisionLinks**](RevisionLinks.md) |  | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**Processes** | [**map[string]RevisionProcessesValue**](RevisionProcessesValue.md) |  | 
**Relationships** | Pointer to [**RevisionRelationships**](RevisionRelationships.md) |  | [optional] 
**Sidecars** | [**[]Sidecar**](Sidecar.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** | Revision version number | 

## Methods

### NewRevision

`func NewRevision(createdAt time.Time, deployable bool, description NullableString, droplet V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData, guid string, links RevisionLinks, metadata V3AppsGuidTasksPostRequestMetadata, processes map[string]RevisionProcessesValue, sidecars []Sidecar, updatedAt time.Time, version int32, ) *Revision`

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


### SetDescriptionNil

`func (o *Revision) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Revision) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDroplet

`func (o *Revision) GetDroplet() V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *Revision) GetDropletOk() (*V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *Revision) SetDroplet(v V3AppsGuidRelationshipsCurrentDropletPatch200ResponseData)`

SetDroplet sets Droplet field to given value.


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


### GetRelationships

`func (o *Revision) GetRelationships() RevisionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Revision) GetRelationshipsOk() (*RevisionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Revision) SetRelationships(v RevisionRelationships)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


