# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The time with zone when the object was created | 
**Data** | Pointer to [**PackageData**](PackageData.md) |  | [optional] 
**Guid** | **string** | Unique identifier for the package | 
**Links** | Pointer to [**PackageLinks**](PackageLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3PackagesPostRequestMetadata**](V3PackagesPostRequestMetadata.md) |  | [optional] 
**Relationships** | Pointer to [**V3PackagesPostRequestRelationships**](V3PackagesPostRequestRelationships.md) |  | [optional] 
**State** | **string** | State of the package; valid states are AWAITING_UPLOAD, PROCESSING_UPLOAD, READY, FAILED, COPYING, EXPIRED | 
**Type** | **string** | Package type; valid values are bits, docker | 
**UpdatedAt** | **time.Time** | The time with zone when the object was last updated | 

## Methods

### NewPackage

`func NewPackage(createdAt time.Time, guid string, state string, type_ string, updatedAt time.Time, ) *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Package) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Package) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Package) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetData

`func (o *Package) GetData() PackageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Package) GetDataOk() (*PackageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Package) SetData(v PackageData)`

SetData sets Data field to given value.

### HasData

`func (o *Package) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGuid

`func (o *Package) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Package) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Package) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *Package) GetLinks() PackageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Package) GetLinksOk() (*PackageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Package) SetLinks(v PackageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Package) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Package) GetMetadata() V3PackagesPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Package) GetMetadataOk() (*V3PackagesPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Package) SetMetadata(v V3PackagesPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Package) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationships

`func (o *Package) GetRelationships() V3PackagesPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Package) GetRelationshipsOk() (*V3PackagesPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Package) SetRelationships(v V3PackagesPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Package) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetState

`func (o *Package) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Package) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Package) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *Package) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Package) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Package) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Package) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Package) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Package) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


