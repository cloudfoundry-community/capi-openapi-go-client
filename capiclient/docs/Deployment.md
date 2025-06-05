# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The time the deployment was created | 
**Droplet** | [**DeploymentDroplet**](DeploymentDroplet.md) |  | 
**Guid** | **string** | Unique identifier for the deployment | 
**Links** | [**DeploymentLinks**](DeploymentLinks.md) |  | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**NewProcesses** | [**[]DeploymentNewProcessesInner**](DeploymentNewProcessesInner.md) | List of new processes created by this deployment | 
**Options** | Pointer to [**DeploymentOptions**](DeploymentOptions.md) |  | [optional] 
**PreviousDroplet** | [**NullableDeploymentPreviousDroplet**](DeploymentPreviousDroplet.md) |  | 
**Relationships** | [**DeploymentRelationships**](DeploymentRelationships.md) |  | 
**Revision** | [**NullableDeploymentRevision**](DeploymentRevision.md) |  | 
**State** | **string** | Current state of the deployment | 
**Status** | [**DeploymentStatus**](DeploymentStatus.md) |  | 
**Strategy** | **string** | Deployment strategy being used | 
**UpdatedAt** | **time.Time** | The time the deployment was last updated | 

## Methods

### NewDeployment

`func NewDeployment(createdAt time.Time, droplet DeploymentDroplet, guid string, links DeploymentLinks, metadata V3AppsGuidTasksPostRequestMetadata, newProcesses []DeploymentNewProcessesInner, previousDroplet NullableDeploymentPreviousDroplet, relationships DeploymentRelationships, revision NullableDeploymentRevision, state string, status DeploymentStatus, strategy string, updatedAt time.Time, ) *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Deployment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Deployment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Deployment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDroplet

`func (o *Deployment) GetDroplet() DeploymentDroplet`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *Deployment) GetDropletOk() (*DeploymentDroplet, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *Deployment) SetDroplet(v DeploymentDroplet)`

SetDroplet sets Droplet field to given value.


### GetGuid

`func (o *Deployment) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Deployment) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Deployment) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *Deployment) GetLinks() DeploymentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Deployment) GetLinksOk() (*DeploymentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Deployment) SetLinks(v DeploymentLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *Deployment) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Deployment) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Deployment) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.


### GetNewProcesses

`func (o *Deployment) GetNewProcesses() []DeploymentNewProcessesInner`

GetNewProcesses returns the NewProcesses field if non-nil, zero value otherwise.

### GetNewProcessesOk

`func (o *Deployment) GetNewProcessesOk() (*[]DeploymentNewProcessesInner, bool)`

GetNewProcessesOk returns a tuple with the NewProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewProcesses

`func (o *Deployment) SetNewProcesses(v []DeploymentNewProcessesInner)`

SetNewProcesses sets NewProcesses field to given value.


### GetOptions

`func (o *Deployment) GetOptions() DeploymentOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Deployment) GetOptionsOk() (*DeploymentOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Deployment) SetOptions(v DeploymentOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Deployment) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreviousDroplet

`func (o *Deployment) GetPreviousDroplet() DeploymentPreviousDroplet`

GetPreviousDroplet returns the PreviousDroplet field if non-nil, zero value otherwise.

### GetPreviousDropletOk

`func (o *Deployment) GetPreviousDropletOk() (*DeploymentPreviousDroplet, bool)`

GetPreviousDropletOk returns a tuple with the PreviousDroplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousDroplet

`func (o *Deployment) SetPreviousDroplet(v DeploymentPreviousDroplet)`

SetPreviousDroplet sets PreviousDroplet field to given value.


### SetPreviousDropletNil

`func (o *Deployment) SetPreviousDropletNil(b bool)`

 SetPreviousDropletNil sets the value for PreviousDroplet to be an explicit nil

### UnsetPreviousDroplet
`func (o *Deployment) UnsetPreviousDroplet()`

UnsetPreviousDroplet ensures that no value is present for PreviousDroplet, not even an explicit nil
### GetRelationships

`func (o *Deployment) GetRelationships() DeploymentRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Deployment) GetRelationshipsOk() (*DeploymentRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Deployment) SetRelationships(v DeploymentRelationships)`

SetRelationships sets Relationships field to given value.


### GetRevision

`func (o *Deployment) GetRevision() DeploymentRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Deployment) GetRevisionOk() (*DeploymentRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Deployment) SetRevision(v DeploymentRevision)`

SetRevision sets Revision field to given value.


### SetRevisionNil

`func (o *Deployment) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *Deployment) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetState

`func (o *Deployment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Deployment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Deployment) SetState(v string)`

SetState sets State field to given value.


### GetStatus

`func (o *Deployment) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.


### GetStrategy

`func (o *Deployment) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Deployment) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Deployment) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.


### GetUpdatedAt

`func (o *Deployment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Deployment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Deployment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


