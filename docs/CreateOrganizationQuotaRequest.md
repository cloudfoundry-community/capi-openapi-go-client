# CreateOrganizationQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**CreateOrganizationQuotaRequestApps**](CreateOrganizationQuotaRequestApps.md) |  | [optional] 
**Domains** | Pointer to [**CreateOrganizationQuotaRequestDomains**](CreateOrganizationQuotaRequestDomains.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the quota | [optional] 
**Relationships** | Pointer to [**CreateOrganizationQuotaRequestRelationships**](CreateOrganizationQuotaRequestRelationships.md) |  | [optional] 
**Routes** | Pointer to [**CreateOrganizationQuotaRequestRoutes**](CreateOrganizationQuotaRequestRoutes.md) |  | [optional] 
**Services** | Pointer to [**CreateOrganizationQuotaRequestServices**](CreateOrganizationQuotaRequestServices.md) |  | [optional] 

## Methods

### NewCreateOrganizationQuotaRequest

`func NewCreateOrganizationQuotaRequest() *CreateOrganizationQuotaRequest`

NewCreateOrganizationQuotaRequest instantiates a new CreateOrganizationQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationQuotaRequestWithDefaults

`func NewCreateOrganizationQuotaRequestWithDefaults() *CreateOrganizationQuotaRequest`

NewCreateOrganizationQuotaRequestWithDefaults instantiates a new CreateOrganizationQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *CreateOrganizationQuotaRequest) GetApps() CreateOrganizationQuotaRequestApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *CreateOrganizationQuotaRequest) GetAppsOk() (*CreateOrganizationQuotaRequestApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *CreateOrganizationQuotaRequest) SetApps(v CreateOrganizationQuotaRequestApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *CreateOrganizationQuotaRequest) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDomains

`func (o *CreateOrganizationQuotaRequest) GetDomains() CreateOrganizationQuotaRequestDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateOrganizationQuotaRequest) GetDomainsOk() (*CreateOrganizationQuotaRequestDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateOrganizationQuotaRequest) SetDomains(v CreateOrganizationQuotaRequestDomains)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateOrganizationQuotaRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationQuotaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationQuotaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationQuotaRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrganizationQuotaRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelationships

`func (o *CreateOrganizationQuotaRequest) GetRelationships() CreateOrganizationQuotaRequestRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CreateOrganizationQuotaRequest) GetRelationshipsOk() (*CreateOrganizationQuotaRequestRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CreateOrganizationQuotaRequest) SetRelationships(v CreateOrganizationQuotaRequestRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CreateOrganizationQuotaRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRoutes

`func (o *CreateOrganizationQuotaRequest) GetRoutes() CreateOrganizationQuotaRequestRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *CreateOrganizationQuotaRequest) GetRoutesOk() (*CreateOrganizationQuotaRequestRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *CreateOrganizationQuotaRequest) SetRoutes(v CreateOrganizationQuotaRequestRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *CreateOrganizationQuotaRequest) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *CreateOrganizationQuotaRequest) GetServices() CreateOrganizationQuotaRequestServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateOrganizationQuotaRequest) GetServicesOk() (*CreateOrganizationQuotaRequestServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateOrganizationQuotaRequest) SetServices(v CreateOrganizationQuotaRequestServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *CreateOrganizationQuotaRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


