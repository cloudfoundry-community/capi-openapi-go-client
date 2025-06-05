# SecurityGroupRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningSpaces** | Pointer to [**SecurityGroupRelationshipsRunningSpaces**](SecurityGroupRelationshipsRunningSpaces.md) |  | [optional] 
**StagingSpaces** | Pointer to [**SecurityGroupRelationshipsStagingSpaces**](SecurityGroupRelationshipsStagingSpaces.md) |  | [optional] 

## Methods

### NewSecurityGroupRelationships

`func NewSecurityGroupRelationships() *SecurityGroupRelationships`

NewSecurityGroupRelationships instantiates a new SecurityGroupRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRelationshipsWithDefaults

`func NewSecurityGroupRelationshipsWithDefaults() *SecurityGroupRelationships`

NewSecurityGroupRelationshipsWithDefaults instantiates a new SecurityGroupRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunningSpaces

`func (o *SecurityGroupRelationships) GetRunningSpaces() SecurityGroupRelationshipsRunningSpaces`

GetRunningSpaces returns the RunningSpaces field if non-nil, zero value otherwise.

### GetRunningSpacesOk

`func (o *SecurityGroupRelationships) GetRunningSpacesOk() (*SecurityGroupRelationshipsRunningSpaces, bool)`

GetRunningSpacesOk returns a tuple with the RunningSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningSpaces

`func (o *SecurityGroupRelationships) SetRunningSpaces(v SecurityGroupRelationshipsRunningSpaces)`

SetRunningSpaces sets RunningSpaces field to given value.

### HasRunningSpaces

`func (o *SecurityGroupRelationships) HasRunningSpaces() bool`

HasRunningSpaces returns a boolean if a field has been set.

### GetStagingSpaces

`func (o *SecurityGroupRelationships) GetStagingSpaces() SecurityGroupRelationshipsStagingSpaces`

GetStagingSpaces returns the StagingSpaces field if non-nil, zero value otherwise.

### GetStagingSpacesOk

`func (o *SecurityGroupRelationships) GetStagingSpacesOk() (*SecurityGroupRelationshipsStagingSpaces, bool)`

GetStagingSpacesOk returns a tuple with the StagingSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingSpaces

`func (o *SecurityGroupRelationships) SetStagingSpaces(v SecurityGroupRelationshipsStagingSpaces)`

SetStagingSpaces sets StagingSpaces field to given value.

### HasStagingSpaces

`func (o *SecurityGroupRelationships) HasStagingSpaces() bool`

HasStagingSpaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


