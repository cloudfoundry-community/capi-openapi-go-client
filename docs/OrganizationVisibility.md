# OrganizationVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guid** | Pointer to **string** | Unique identifier for the organization where the plan is available | [optional] 
**Name** | Pointer to **string** | Name of the organization where the plan is available | [optional] 

## Methods

### NewOrganizationVisibility

`func NewOrganizationVisibility() *OrganizationVisibility`

NewOrganizationVisibility instantiates a new OrganizationVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationVisibilityWithDefaults

`func NewOrganizationVisibilityWithDefaults() *OrganizationVisibility`

NewOrganizationVisibilityWithDefaults instantiates a new OrganizationVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuid

`func (o *OrganizationVisibility) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *OrganizationVisibility) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *OrganizationVisibility) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *OrganizationVisibility) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetName

`func (o *OrganizationVisibility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationVisibility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationVisibility) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationVisibility) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


