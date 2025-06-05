# SpaceQuotaCreateRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | [**SpaceQuotaRelationshipsOrganization**](SpaceQuotaRelationshipsOrganization.md) |  | 
**Spaces** | Pointer to [**SpaceQuotaCreateRelationshipsSpaces**](SpaceQuotaCreateRelationshipsSpaces.md) |  | [optional] 

## Methods

### NewSpaceQuotaCreateRelationships

`func NewSpaceQuotaCreateRelationships(organization SpaceQuotaRelationshipsOrganization, ) *SpaceQuotaCreateRelationships`

NewSpaceQuotaCreateRelationships instantiates a new SpaceQuotaCreateRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaCreateRelationshipsWithDefaults

`func NewSpaceQuotaCreateRelationshipsWithDefaults() *SpaceQuotaCreateRelationships`

NewSpaceQuotaCreateRelationshipsWithDefaults instantiates a new SpaceQuotaCreateRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SpaceQuotaCreateRelationships) GetOrganization() SpaceQuotaRelationshipsOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SpaceQuotaCreateRelationships) GetOrganizationOk() (*SpaceQuotaRelationshipsOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SpaceQuotaCreateRelationships) SetOrganization(v SpaceQuotaRelationshipsOrganization)`

SetOrganization sets Organization field to given value.


### GetSpaces

`func (o *SpaceQuotaCreateRelationships) GetSpaces() SpaceQuotaCreateRelationshipsSpaces`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *SpaceQuotaCreateRelationships) GetSpacesOk() (*SpaceQuotaCreateRelationshipsSpaces, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *SpaceQuotaCreateRelationships) SetSpaces(v SpaceQuotaCreateRelationshipsSpaces)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *SpaceQuotaCreateRelationships) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


