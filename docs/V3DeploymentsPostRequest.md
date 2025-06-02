# V3DeploymentsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Droplet** | Pointer to [**V3DeploymentsPostRequestDroplet**](V3DeploymentsPostRequestDroplet.md) |  | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Options** | Pointer to [**V3DeploymentsPostRequestOptions**](V3DeploymentsPostRequestOptions.md) |  | [optional] 
**Relationships** | [**V3DeploymentsPostRequestRelationships**](V3DeploymentsPostRequestRelationships.md) |  | 
**Revision** | Pointer to [**V3DeploymentsPostRequestRevision**](V3DeploymentsPostRequestRevision.md) |  | [optional] 
**Strategy** | Pointer to **string** | Deployment strategy to use | [optional] [default to "rolling"]

## Methods

### NewV3DeploymentsPostRequest

`func NewV3DeploymentsPostRequest(relationships V3DeploymentsPostRequestRelationships, ) *V3DeploymentsPostRequest`

NewV3DeploymentsPostRequest instantiates a new V3DeploymentsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3DeploymentsPostRequestWithDefaults

`func NewV3DeploymentsPostRequestWithDefaults() *V3DeploymentsPostRequest`

NewV3DeploymentsPostRequestWithDefaults instantiates a new V3DeploymentsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDroplet

`func (o *V3DeploymentsPostRequest) GetDroplet() V3DeploymentsPostRequestDroplet`

GetDroplet returns the Droplet field if non-nil, zero value otherwise.

### GetDropletOk

`func (o *V3DeploymentsPostRequest) GetDropletOk() (*V3DeploymentsPostRequestDroplet, bool)`

GetDropletOk returns a tuple with the Droplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplet

`func (o *V3DeploymentsPostRequest) SetDroplet(v V3DeploymentsPostRequestDroplet)`

SetDroplet sets Droplet field to given value.

### HasDroplet

`func (o *V3DeploymentsPostRequest) HasDroplet() bool`

HasDroplet returns a boolean if a field has been set.

### GetMetadata

`func (o *V3DeploymentsPostRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3DeploymentsPostRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3DeploymentsPostRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3DeploymentsPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOptions

`func (o *V3DeploymentsPostRequest) GetOptions() V3DeploymentsPostRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *V3DeploymentsPostRequest) GetOptionsOk() (*V3DeploymentsPostRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *V3DeploymentsPostRequest) SetOptions(v V3DeploymentsPostRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *V3DeploymentsPostRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRelationships

`func (o *V3DeploymentsPostRequest) GetRelationships() V3DeploymentsPostRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *V3DeploymentsPostRequest) GetRelationshipsOk() (*V3DeploymentsPostRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *V3DeploymentsPostRequest) SetRelationships(v V3DeploymentsPostRequestRelationships)`

SetRelationships sets Relationships field to given value.


### GetRevision

`func (o *V3DeploymentsPostRequest) GetRevision() V3DeploymentsPostRequestRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V3DeploymentsPostRequest) GetRevisionOk() (*V3DeploymentsPostRequestRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V3DeploymentsPostRequest) SetRevision(v V3DeploymentsPostRequestRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V3DeploymentsPostRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStrategy

`func (o *V3DeploymentsPostRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V3DeploymentsPostRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V3DeploymentsPostRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *V3DeploymentsPostRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


