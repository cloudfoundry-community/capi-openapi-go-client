# OrganizationQuotaCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**OrganizationQuotaCreateApps**](OrganizationQuotaCreateApps.md) |  | [optional] 
**Domains** | Pointer to [**OrganizationQuotaCreateDomains**](OrganizationQuotaCreateDomains.md) |  | [optional] 
**Metadata** | Pointer to [**OrganizationQuotaCreateMetadata**](OrganizationQuotaCreateMetadata.md) |  | [optional] 
**Name** | **string** | Human-readable name for the organization quota | 
**Relationships** | Pointer to [**OrganizationQuotaCreateRelationships**](OrganizationQuotaCreateRelationships.md) |  | [optional] 
**Routes** | Pointer to [**OrganizationQuotaCreateRoutes**](OrganizationQuotaCreateRoutes.md) |  | [optional] 
**Services** | Pointer to [**OrganizationQuotaCreateServices**](OrganizationQuotaCreateServices.md) |  | [optional] 

## Methods

### NewOrganizationQuotaCreate

`func NewOrganizationQuotaCreate(name string, ) *OrganizationQuotaCreate`

NewOrganizationQuotaCreate instantiates a new OrganizationQuotaCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaCreateWithDefaults

`func NewOrganizationQuotaCreateWithDefaults() *OrganizationQuotaCreate`

NewOrganizationQuotaCreateWithDefaults instantiates a new OrganizationQuotaCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OrganizationQuotaCreate) GetApps() OrganizationQuotaCreateApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OrganizationQuotaCreate) GetAppsOk() (*OrganizationQuotaCreateApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OrganizationQuotaCreate) SetApps(v OrganizationQuotaCreateApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OrganizationQuotaCreate) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDomains

`func (o *OrganizationQuotaCreate) GetDomains() OrganizationQuotaCreateDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationQuotaCreate) GetDomainsOk() (*OrganizationQuotaCreateDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationQuotaCreate) SetDomains(v OrganizationQuotaCreateDomains)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationQuotaCreate) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetMetadata

`func (o *OrganizationQuotaCreate) GetMetadata() OrganizationQuotaCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrganizationQuotaCreate) GetMetadataOk() (*OrganizationQuotaCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrganizationQuotaCreate) SetMetadata(v OrganizationQuotaCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrganizationQuotaCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrganizationQuotaCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationQuotaCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationQuotaCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRelationships

`func (o *OrganizationQuotaCreate) GetRelationships() OrganizationQuotaCreateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrganizationQuotaCreate) GetRelationshipsOk() (*OrganizationQuotaCreateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrganizationQuotaCreate) SetRelationships(v OrganizationQuotaCreateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrganizationQuotaCreate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRoutes

`func (o *OrganizationQuotaCreate) GetRoutes() OrganizationQuotaCreateRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OrganizationQuotaCreate) GetRoutesOk() (*OrganizationQuotaCreateRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OrganizationQuotaCreate) SetRoutes(v OrganizationQuotaCreateRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *OrganizationQuotaCreate) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *OrganizationQuotaCreate) GetServices() OrganizationQuotaCreateServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationQuotaCreate) GetServicesOk() (*OrganizationQuotaCreateServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationQuotaCreate) SetServices(v OrganizationQuotaCreateServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *OrganizationQuotaCreate) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


