# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Destinations** | Pointer to [**[]Destination**](Destination.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**RouteLinks**](RouteLinks.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidDropletsCurrentGet200ResponseMetadata**](V3AppsGuidDropletsCurrentGet200ResponseMetadata.md) |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**V3RoutesPostRequestRelationships**](V3RoutesPostRequestRelationships.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Route) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Route) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Route) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Route) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestinations

`func (o *Route) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *Route) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *Route) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *Route) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetGuid

`func (o *Route) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Route) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Route) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Route) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetHost

`func (o *Route) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Route) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Route) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Route) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *Route) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Route) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetLinks

`func (o *Route) GetLinks() RouteLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Route) GetLinksOk() (*RouteLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Route) SetLinks(v RouteLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Route) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *Route) GetMetadata() V3AppsGuidDropletsCurrentGet200ResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Route) GetMetadataOk() (*V3AppsGuidDropletsCurrentGet200ResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Route) SetMetadata(v V3AppsGuidDropletsCurrentGet200ResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Route) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPath

`func (o *Route) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Route) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Route) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Route) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Route) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Route) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPort

`func (o *Route) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Route) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Route) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Route) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *Route) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Route) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *Route) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Route) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Route) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Route) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRelationships

`func (o *Route) GetRelationships() V3RoutesPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Route) GetRelationshipsOk() (*V3RoutesPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Route) SetRelationships(v V3RoutesPostRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Route) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Route) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Route) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Route) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Route) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Route) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Route) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Route) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Route) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


