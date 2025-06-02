# OrganizationQuotaUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**OrganizationQuotaCreateApps**](OrganizationQuotaCreateApps.md) |  | [optional] 
**Domains** | Pointer to [**OrganizationQuotaCreateDomains**](OrganizationQuotaCreateDomains.md) |  | [optional] 
**Metadata** | Pointer to [**OrganizationQuotaCreateMetadata**](OrganizationQuotaCreateMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable name for the organization quota | [optional] 
**Routes** | Pointer to [**OrganizationQuotaCreateRoutes**](OrganizationQuotaCreateRoutes.md) |  | [optional] 
**Services** | Pointer to [**OrganizationQuotaUpdateServices**](OrganizationQuotaUpdateServices.md) |  | [optional] 

## Methods

### NewOrganizationQuotaUpdate

`func NewOrganizationQuotaUpdate() *OrganizationQuotaUpdate`

NewOrganizationQuotaUpdate instantiates a new OrganizationQuotaUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationQuotaUpdateWithDefaults

`func NewOrganizationQuotaUpdateWithDefaults() *OrganizationQuotaUpdate`

NewOrganizationQuotaUpdateWithDefaults instantiates a new OrganizationQuotaUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OrganizationQuotaUpdate) GetApps() OrganizationQuotaCreateApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OrganizationQuotaUpdate) GetAppsOk() (*OrganizationQuotaCreateApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OrganizationQuotaUpdate) SetApps(v OrganizationQuotaCreateApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OrganizationQuotaUpdate) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDomains

`func (o *OrganizationQuotaUpdate) GetDomains() OrganizationQuotaCreateDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationQuotaUpdate) GetDomainsOk() (*OrganizationQuotaCreateDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationQuotaUpdate) SetDomains(v OrganizationQuotaCreateDomains)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationQuotaUpdate) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetMetadata

`func (o *OrganizationQuotaUpdate) GetMetadata() OrganizationQuotaCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrganizationQuotaUpdate) GetMetadataOk() (*OrganizationQuotaCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrganizationQuotaUpdate) SetMetadata(v OrganizationQuotaCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrganizationQuotaUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrganizationQuotaUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationQuotaUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationQuotaUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationQuotaUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *OrganizationQuotaUpdate) GetRoutes() OrganizationQuotaCreateRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *OrganizationQuotaUpdate) GetRoutesOk() (*OrganizationQuotaCreateRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *OrganizationQuotaUpdate) SetRoutes(v OrganizationQuotaCreateRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *OrganizationQuotaUpdate) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *OrganizationQuotaUpdate) GetServices() OrganizationQuotaUpdateServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OrganizationQuotaUpdate) GetServicesOk() (*OrganizationQuotaUpdateServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OrganizationQuotaUpdate) SetServices(v OrganizationQuotaUpdateServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *OrganizationQuotaUpdate) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


