# UpdateOrganizationQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**AppsQuota**](AppsQuota.md) |  | [optional] 
**Domains** | Pointer to [**DomainsQuota**](DomainsQuota.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the quota | [optional] 
**Routes** | Pointer to [**RoutesQuota**](RoutesQuota.md) |  | [optional] 
**Services** | Pointer to [**ServicesQuota**](ServicesQuota.md) |  | [optional] 

## Methods

### NewUpdateOrganizationQuotaRequest

`func NewUpdateOrganizationQuotaRequest() *UpdateOrganizationQuotaRequest`

NewUpdateOrganizationQuotaRequest instantiates a new UpdateOrganizationQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationQuotaRequestWithDefaults

`func NewUpdateOrganizationQuotaRequestWithDefaults() *UpdateOrganizationQuotaRequest`

NewUpdateOrganizationQuotaRequestWithDefaults instantiates a new UpdateOrganizationQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *UpdateOrganizationQuotaRequest) GetApps() AppsQuota`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *UpdateOrganizationQuotaRequest) GetAppsOk() (*AppsQuota, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *UpdateOrganizationQuotaRequest) SetApps(v AppsQuota)`

SetApps sets Apps field to given value.

### HasApps

`func (o *UpdateOrganizationQuotaRequest) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateOrganizationQuotaRequest) GetDomains() DomainsQuota`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateOrganizationQuotaRequest) GetDomainsOk() (*DomainsQuota, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateOrganizationQuotaRequest) SetDomains(v DomainsQuota)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateOrganizationQuotaRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationQuotaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationQuotaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationQuotaRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationQuotaRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutes

`func (o *UpdateOrganizationQuotaRequest) GetRoutes() RoutesQuota`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *UpdateOrganizationQuotaRequest) GetRoutesOk() (*RoutesQuota, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *UpdateOrganizationQuotaRequest) SetRoutes(v RoutesQuota)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *UpdateOrganizationQuotaRequest) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetServices

`func (o *UpdateOrganizationQuotaRequest) GetServices() ServicesQuota`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *UpdateOrganizationQuotaRequest) GetServicesOk() (*ServicesQuota, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *UpdateOrganizationQuotaRequest) SetServices(v ServicesQuota)`

SetServices sets Services field to given value.

### HasServices

`func (o *UpdateOrganizationQuotaRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


