# Sidecar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Command to execute for the sidecar process | 
**CreatedAt** | **time.Time** | Timestamp when the sidecar was created | 
**Guid** | **string** | Unique identifier for the sidecar | 
**Links** | [**SidecarLinks**](SidecarLinks.md) |  | 
**MemoryInMb** | **int32** | Memory allocation for the sidecar in MB | 
**Metadata** | Pointer to [**SidecarMetadata**](SidecarMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the sidecar | 
**Origin** | **string** | How the sidecar was created | 
**ProcessTypes** | **[]string** | Process types this sidecar runs alongside | 
**Relationships** | [**SidecarRelationships**](SidecarRelationships.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp when the sidecar was last updated | 

## Methods

### NewSidecar

`func NewSidecar(command string, createdAt time.Time, guid string, links SidecarLinks, memoryInMb int32, name string, origin string, processTypes []string, relationships SidecarRelationships, updatedAt time.Time, ) *Sidecar`

NewSidecar instantiates a new Sidecar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSidecarWithDefaults

`func NewSidecarWithDefaults() *Sidecar`

NewSidecarWithDefaults instantiates a new Sidecar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Sidecar) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Sidecar) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Sidecar) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetCreatedAt

`func (o *Sidecar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Sidecar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Sidecar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGuid

`func (o *Sidecar) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Sidecar) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Sidecar) SetGuid(v string)`

SetGuid sets Guid field to given value.


### GetLinks

`func (o *Sidecar) GetLinks() SidecarLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Sidecar) GetLinksOk() (*SidecarLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Sidecar) SetLinks(v SidecarLinks)`

SetLinks sets Links field to given value.


### GetMemoryInMb

`func (o *Sidecar) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *Sidecar) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *Sidecar) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.


### GetMetadata

`func (o *Sidecar) GetMetadata() SidecarMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Sidecar) GetMetadataOk() (*SidecarMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Sidecar) SetMetadata(v SidecarMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Sidecar) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Sidecar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sidecar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sidecar) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *Sidecar) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Sidecar) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Sidecar) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetProcessTypes

`func (o *Sidecar) GetProcessTypes() []string`

GetProcessTypes returns the ProcessTypes field if non-nil, zero value otherwise.

### GetProcessTypesOk

`func (o *Sidecar) GetProcessTypesOk() (*[]string, bool)`

GetProcessTypesOk returns a tuple with the ProcessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessTypes

`func (o *Sidecar) SetProcessTypes(v []string)`

SetProcessTypes sets ProcessTypes field to given value.


### GetRelationships

`func (o *Sidecar) GetRelationships() SidecarRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Sidecar) GetRelationshipsOk() (*SidecarRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Sidecar) SetRelationships(v SidecarRelationships)`

SetRelationships sets Relationships field to given value.


### GetUpdatedAt

`func (o *Sidecar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Sidecar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Sidecar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


