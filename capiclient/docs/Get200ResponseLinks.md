# Get200ResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSsh** | Pointer to [**Get200ResponseLinksAppSsh**](Get200ResponseLinksAppSsh.md) |  | [optional] 
**CloudControllerV2** | Pointer to [**Get200ResponseLinksCloudControllerV2**](Get200ResponseLinksCloudControllerV2.md) |  | [optional] 
**CloudControllerV3** | Pointer to [**Get200ResponseLinksCloudControllerV2**](Get200ResponseLinksCloudControllerV2.md) |  | [optional] 
**Credhub** | Pointer to **map[string]interface{}** |  | [optional] 
**LogCache** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**LogStream** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Logging** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Login** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**NetworkPolicyV0** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**NetworkPolicyV1** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Routing** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Self** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 
**Uaa** | Pointer to [**Get200ResponseLinksLogCache**](Get200ResponseLinksLogCache.md) |  | [optional] 

## Methods

### NewGet200ResponseLinks

`func NewGet200ResponseLinks() *Get200ResponseLinks`

NewGet200ResponseLinks instantiates a new Get200ResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGet200ResponseLinksWithDefaults

`func NewGet200ResponseLinksWithDefaults() *Get200ResponseLinks`

NewGet200ResponseLinksWithDefaults instantiates a new Get200ResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSsh

`func (o *Get200ResponseLinks) GetAppSsh() Get200ResponseLinksAppSsh`

GetAppSsh returns the AppSsh field if non-nil, zero value otherwise.

### GetAppSshOk

`func (o *Get200ResponseLinks) GetAppSshOk() (*Get200ResponseLinksAppSsh, bool)`

GetAppSshOk returns a tuple with the AppSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSsh

`func (o *Get200ResponseLinks) SetAppSsh(v Get200ResponseLinksAppSsh)`

SetAppSsh sets AppSsh field to given value.

### HasAppSsh

`func (o *Get200ResponseLinks) HasAppSsh() bool`

HasAppSsh returns a boolean if a field has been set.

### GetCloudControllerV2

`func (o *Get200ResponseLinks) GetCloudControllerV2() Get200ResponseLinksCloudControllerV2`

GetCloudControllerV2 returns the CloudControllerV2 field if non-nil, zero value otherwise.

### GetCloudControllerV2Ok

`func (o *Get200ResponseLinks) GetCloudControllerV2Ok() (*Get200ResponseLinksCloudControllerV2, bool)`

GetCloudControllerV2Ok returns a tuple with the CloudControllerV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudControllerV2

`func (o *Get200ResponseLinks) SetCloudControllerV2(v Get200ResponseLinksCloudControllerV2)`

SetCloudControllerV2 sets CloudControllerV2 field to given value.

### HasCloudControllerV2

`func (o *Get200ResponseLinks) HasCloudControllerV2() bool`

HasCloudControllerV2 returns a boolean if a field has been set.

### GetCloudControllerV3

`func (o *Get200ResponseLinks) GetCloudControllerV3() Get200ResponseLinksCloudControllerV2`

GetCloudControllerV3 returns the CloudControllerV3 field if non-nil, zero value otherwise.

### GetCloudControllerV3Ok

`func (o *Get200ResponseLinks) GetCloudControllerV3Ok() (*Get200ResponseLinksCloudControllerV2, bool)`

GetCloudControllerV3Ok returns a tuple with the CloudControllerV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudControllerV3

`func (o *Get200ResponseLinks) SetCloudControllerV3(v Get200ResponseLinksCloudControllerV2)`

SetCloudControllerV3 sets CloudControllerV3 field to given value.

### HasCloudControllerV3

`func (o *Get200ResponseLinks) HasCloudControllerV3() bool`

HasCloudControllerV3 returns a boolean if a field has been set.

### GetCredhub

`func (o *Get200ResponseLinks) GetCredhub() map[string]interface{}`

GetCredhub returns the Credhub field if non-nil, zero value otherwise.

### GetCredhubOk

`func (o *Get200ResponseLinks) GetCredhubOk() (*map[string]interface{}, bool)`

GetCredhubOk returns a tuple with the Credhub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredhub

`func (o *Get200ResponseLinks) SetCredhub(v map[string]interface{})`

SetCredhub sets Credhub field to given value.

### HasCredhub

`func (o *Get200ResponseLinks) HasCredhub() bool`

HasCredhub returns a boolean if a field has been set.

### SetCredhubNil

`func (o *Get200ResponseLinks) SetCredhubNil(b bool)`

 SetCredhubNil sets the value for Credhub to be an explicit nil

### UnsetCredhub
`func (o *Get200ResponseLinks) UnsetCredhub()`

UnsetCredhub ensures that no value is present for Credhub, not even an explicit nil
### GetLogCache

`func (o *Get200ResponseLinks) GetLogCache() Get200ResponseLinksLogCache`

GetLogCache returns the LogCache field if non-nil, zero value otherwise.

### GetLogCacheOk

`func (o *Get200ResponseLinks) GetLogCacheOk() (*Get200ResponseLinksLogCache, bool)`

GetLogCacheOk returns a tuple with the LogCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCache

`func (o *Get200ResponseLinks) SetLogCache(v Get200ResponseLinksLogCache)`

SetLogCache sets LogCache field to given value.

### HasLogCache

`func (o *Get200ResponseLinks) HasLogCache() bool`

HasLogCache returns a boolean if a field has been set.

### GetLogStream

`func (o *Get200ResponseLinks) GetLogStream() Get200ResponseLinksLogCache`

GetLogStream returns the LogStream field if non-nil, zero value otherwise.

### GetLogStreamOk

`func (o *Get200ResponseLinks) GetLogStreamOk() (*Get200ResponseLinksLogCache, bool)`

GetLogStreamOk returns a tuple with the LogStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStream

`func (o *Get200ResponseLinks) SetLogStream(v Get200ResponseLinksLogCache)`

SetLogStream sets LogStream field to given value.

### HasLogStream

`func (o *Get200ResponseLinks) HasLogStream() bool`

HasLogStream returns a boolean if a field has been set.

### GetLogging

`func (o *Get200ResponseLinks) GetLogging() Get200ResponseLinksLogCache`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *Get200ResponseLinks) GetLoggingOk() (*Get200ResponseLinksLogCache, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *Get200ResponseLinks) SetLogging(v Get200ResponseLinksLogCache)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *Get200ResponseLinks) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetLogin

`func (o *Get200ResponseLinks) GetLogin() Get200ResponseLinksLogCache`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Get200ResponseLinks) GetLoginOk() (*Get200ResponseLinksLogCache, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Get200ResponseLinks) SetLogin(v Get200ResponseLinksLogCache)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *Get200ResponseLinks) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetNetworkPolicyV0

`func (o *Get200ResponseLinks) GetNetworkPolicyV0() Get200ResponseLinksLogCache`

GetNetworkPolicyV0 returns the NetworkPolicyV0 field if non-nil, zero value otherwise.

### GetNetworkPolicyV0Ok

`func (o *Get200ResponseLinks) GetNetworkPolicyV0Ok() (*Get200ResponseLinksLogCache, bool)`

GetNetworkPolicyV0Ok returns a tuple with the NetworkPolicyV0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyV0

`func (o *Get200ResponseLinks) SetNetworkPolicyV0(v Get200ResponseLinksLogCache)`

SetNetworkPolicyV0 sets NetworkPolicyV0 field to given value.

### HasNetworkPolicyV0

`func (o *Get200ResponseLinks) HasNetworkPolicyV0() bool`

HasNetworkPolicyV0 returns a boolean if a field has been set.

### GetNetworkPolicyV1

`func (o *Get200ResponseLinks) GetNetworkPolicyV1() Get200ResponseLinksLogCache`

GetNetworkPolicyV1 returns the NetworkPolicyV1 field if non-nil, zero value otherwise.

### GetNetworkPolicyV1Ok

`func (o *Get200ResponseLinks) GetNetworkPolicyV1Ok() (*Get200ResponseLinksLogCache, bool)`

GetNetworkPolicyV1Ok returns a tuple with the NetworkPolicyV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyV1

`func (o *Get200ResponseLinks) SetNetworkPolicyV1(v Get200ResponseLinksLogCache)`

SetNetworkPolicyV1 sets NetworkPolicyV1 field to given value.

### HasNetworkPolicyV1

`func (o *Get200ResponseLinks) HasNetworkPolicyV1() bool`

HasNetworkPolicyV1 returns a boolean if a field has been set.

### GetRouting

`func (o *Get200ResponseLinks) GetRouting() Get200ResponseLinksLogCache`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *Get200ResponseLinks) GetRoutingOk() (*Get200ResponseLinksLogCache, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *Get200ResponseLinks) SetRouting(v Get200ResponseLinksLogCache)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *Get200ResponseLinks) HasRouting() bool`

HasRouting returns a boolean if a field has been set.

### GetSelf

`func (o *Get200ResponseLinks) GetSelf() Get200ResponseLinksLogCache`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Get200ResponseLinks) GetSelfOk() (*Get200ResponseLinksLogCache, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Get200ResponseLinks) SetSelf(v Get200ResponseLinksLogCache)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Get200ResponseLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetUaa

`func (o *Get200ResponseLinks) GetUaa() Get200ResponseLinksLogCache`

GetUaa returns the Uaa field if non-nil, zero value otherwise.

### GetUaaOk

`func (o *Get200ResponseLinks) GetUaaOk() (*Get200ResponseLinksLogCache, bool)`

GetUaaOk returns a tuple with the Uaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaa

`func (o *Get200ResponseLinks) SetUaa(v Get200ResponseLinksLogCache)`

SetUaa sets Uaa field to given value.

### HasUaa

`func (o *Get200ResponseLinks) HasUaa() bool`

HasUaa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


