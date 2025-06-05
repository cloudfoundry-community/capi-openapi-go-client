# GetAppEnvironmentVariables200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**GetAppEnvironmentVariables200ResponseLinks**](GetAppEnvironmentVariables200ResponseLinks.md) |  | [optional] 
**Var** | Pointer to **map[string]string** | User-provided environment variables | [optional] 

## Methods

### NewGetAppEnvironmentVariables200Response

`func NewGetAppEnvironmentVariables200Response() *GetAppEnvironmentVariables200Response`

NewGetAppEnvironmentVariables200Response instantiates a new GetAppEnvironmentVariables200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppEnvironmentVariables200ResponseWithDefaults

`func NewGetAppEnvironmentVariables200ResponseWithDefaults() *GetAppEnvironmentVariables200Response`

NewGetAppEnvironmentVariables200ResponseWithDefaults instantiates a new GetAppEnvironmentVariables200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetAppEnvironmentVariables200Response) GetLinks() GetAppEnvironmentVariables200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAppEnvironmentVariables200Response) GetLinksOk() (*GetAppEnvironmentVariables200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAppEnvironmentVariables200Response) SetLinks(v GetAppEnvironmentVariables200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAppEnvironmentVariables200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetVar

`func (o *GetAppEnvironmentVariables200Response) GetVar() map[string]string`

GetVar returns the Var field if non-nil, zero value otherwise.

### GetVarOk

`func (o *GetAppEnvironmentVariables200Response) GetVarOk() (*map[string]string, bool)`

GetVarOk returns a tuple with the Var field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar

`func (o *GetAppEnvironmentVariables200Response) SetVar(v map[string]string)`

SetVar sets Var field to given value.

### HasVar

`func (o *GetAppEnvironmentVariables200Response) HasVar() bool`

HasVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


