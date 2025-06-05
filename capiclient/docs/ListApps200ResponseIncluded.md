# ListApps200ResponseIncluded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Spaces** | Pointer to [**[]Space**](Space.md) |  | [optional] 

## Methods

### NewListApps200ResponseIncluded

`func NewListApps200ResponseIncluded() *ListApps200ResponseIncluded`

NewListApps200ResponseIncluded instantiates a new ListApps200ResponseIncluded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApps200ResponseIncludedWithDefaults

`func NewListApps200ResponseIncludedWithDefaults() *ListApps200ResponseIncluded`

NewListApps200ResponseIncludedWithDefaults instantiates a new ListApps200ResponseIncluded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *ListApps200ResponseIncluded) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *ListApps200ResponseIncluded) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *ListApps200ResponseIncluded) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *ListApps200ResponseIncluded) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetSpaces

`func (o *ListApps200ResponseIncluded) GetSpaces() []Space`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *ListApps200ResponseIncluded) GetSpacesOk() (*[]Space, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *ListApps200ResponseIncluded) SetSpaces(v []Space)`

SetSpaces sets Spaces field to given value.

### HasSpaces

`func (o *ListApps200ResponseIncluded) HasSpaces() bool`

HasSpaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


