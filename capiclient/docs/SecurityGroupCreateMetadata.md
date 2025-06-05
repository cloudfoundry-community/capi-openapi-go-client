# SecurityGroupCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** | Key-value pairs for storing additional information | [optional] 
**Labels** | Pointer to **map[string]string** | Key-value pairs for organizing and filtering | [optional] 

## Methods

### NewSecurityGroupCreateMetadata

`func NewSecurityGroupCreateMetadata() *SecurityGroupCreateMetadata`

NewSecurityGroupCreateMetadata instantiates a new SecurityGroupCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupCreateMetadataWithDefaults

`func NewSecurityGroupCreateMetadataWithDefaults() *SecurityGroupCreateMetadata`

NewSecurityGroupCreateMetadataWithDefaults instantiates a new SecurityGroupCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SecurityGroupCreateMetadata) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SecurityGroupCreateMetadata) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SecurityGroupCreateMetadata) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SecurityGroupCreateMetadata) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLabels

`func (o *SecurityGroupCreateMetadata) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SecurityGroupCreateMetadata) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SecurityGroupCreateMetadata) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SecurityGroupCreateMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


