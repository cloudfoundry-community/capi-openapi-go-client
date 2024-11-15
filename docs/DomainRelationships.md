# DomainRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to [**V3DropletsPostRequestRelationshipsApp**](V3DropletsPostRequestRelationshipsApp.md) |  | [optional] 
**SharedOrganizations** | Pointer to [**DomainRelationshipsSharedOrganizations**](DomainRelationshipsSharedOrganizations.md) |  | [optional] 

## Methods

### NewDomainRelationships

`func NewDomainRelationships() *DomainRelationships`

NewDomainRelationships instantiates a new DomainRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRelationshipsWithDefaults

`func NewDomainRelationshipsWithDefaults() *DomainRelationships`

NewDomainRelationshipsWithDefaults instantiates a new DomainRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *DomainRelationships) GetOrganization() V3DropletsPostRequestRelationshipsApp`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DomainRelationships) GetOrganizationOk() (*V3DropletsPostRequestRelationshipsApp, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DomainRelationships) SetOrganization(v V3DropletsPostRequestRelationshipsApp)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DomainRelationships) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSharedOrganizations

`func (o *DomainRelationships) GetSharedOrganizations() DomainRelationshipsSharedOrganizations`

GetSharedOrganizations returns the SharedOrganizations field if non-nil, zero value otherwise.

### GetSharedOrganizationsOk

`func (o *DomainRelationships) GetSharedOrganizationsOk() (*DomainRelationshipsSharedOrganizations, bool)`

GetSharedOrganizationsOk returns a tuple with the SharedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOrganizations

`func (o *DomainRelationships) SetSharedOrganizations(v DomainRelationshipsSharedOrganizations)`

SetSharedOrganizations sets SharedOrganizations field to given value.

### HasSharedOrganizations

`func (o *DomainRelationships) HasSharedOrganizations() bool`

HasSharedOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


