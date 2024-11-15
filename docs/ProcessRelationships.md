# ProcessRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**ToOneRelationship**](ToOneRelationship.md) |  | [optional] 
**Revision** | Pointer to [**ToOneRelationship**](ToOneRelationship.md) |  | [optional] 

## Methods

### NewProcessRelationships

`func NewProcessRelationships() *ProcessRelationships`

NewProcessRelationships instantiates a new ProcessRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessRelationshipsWithDefaults

`func NewProcessRelationshipsWithDefaults() *ProcessRelationships`

NewProcessRelationshipsWithDefaults instantiates a new ProcessRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *ProcessRelationships) GetApp() ToOneRelationship`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ProcessRelationships) GetAppOk() (*ToOneRelationship, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ProcessRelationships) SetApp(v ToOneRelationship)`

SetApp sets App field to given value.

### HasApp

`func (o *ProcessRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetRevision

`func (o *ProcessRelationships) GetRevision() ToOneRelationship`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ProcessRelationships) GetRevisionOk() (*ToOneRelationship, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ProcessRelationships) SetRevision(v ToOneRelationship)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ProcessRelationships) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


