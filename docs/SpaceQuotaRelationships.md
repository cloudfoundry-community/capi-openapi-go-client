# SpaceQuotaRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**V3AppsPostRequestRelationshipsSpace**](V3AppsPostRequestRelationshipsSpace.md) |  | [optional] 
**Spaces** | Pointer to [**V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest**](V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest.md) |  | [optional] 

## Methods

### NewSpaceQuotaRelationships

`func NewSpaceQuotaRelationships() *SpaceQuotaRelationships`

NewSpaceQuotaRelationships instantiates a new SpaceQuotaRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceQuotaRelationshipsWithDefaults

`func NewSpaceQuotaRelationshipsWithDefaults() *SpaceQuotaRelationships`

NewSpaceQuotaRelationshipsWithDefaults instantiates a new SpaceQuotaRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *SpaceQuotaRelationships) GetOrganization() V3AppsPostRequestRelationshipsSpace`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SpaceQuotaRelationships) GetOrganizationOk() (*V3AppsPostRequestRelationshipsSpace, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SpaceQuotaRelationships) SetOrganization(v V3AppsPostRequestRelationshipsSpace)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SpaceQuotaRelationships) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSpaces

`func (o *SpaceQuotaRelationships) GetSpaces() V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *SpaceQuotaRelationships) GetSpacesOk() (*V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *SpaceQuotaRelationships) SetSpaces(v V3IsolationSegmentsGuidRelationshipsOrganizationsPostRequest)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *SpaceQuotaRelationships) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


