# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | When the app was created | 
**Guid** | **string** | Unique identifier for the app | 
**Lifecycle** | [**AppLifecycle**](AppLifecycle.md) |  | 
**Links** | [**AppLinks**](AppLinks.md) |  | 
**Metadata** | [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | 
**Name** | **string** | Name of the app | 
**Relationships** | [**AppRelationships**](AppRelationships.md) |  | 
**State** | **string** | Current desired state of the app | 
**UpdatedAt** | **time.Time** | When the app was last updated | 

## Methods

### NewApp

`func NewApp(createdAt time.Time, guid string, lifecycle AppLifecycle, links AppLinks, metadata V3AppsGuidTasksPostRequestMetadata, name string, relationships AppRelationships, state string, updatedAt time.Time, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *App) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *App) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *App) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGuid

`func (o *App) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *App) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *App) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLifecycle

`func (o *App) GetLifecycle() AppLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *App) GetLifecycleOk() (*AppLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *App) SetLifecycle(v AppLifecycle)`

SetLifecycle sets Lifecycle field to given value.


### GetLinks

`func (o *App) GetLinks() AppLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *App) GetLinksOk() (*AppLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *App) SetLinks(v AppLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *App) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *App) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *App) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *App) GetRelationships() AppRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *App) GetRelationshipsOk() (*AppRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *App) SetRelationships(v AppRelationships)`

SetRelationships sets Relationships field to given value.


### GetState

`func (o *App) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *App) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *App) SetState(v string)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *App) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *App) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *App) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


