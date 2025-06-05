# AppLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentDroplet** | Pointer to [**Link**](Link.md) |  | [optional] 
**DeployedRevisions** | Pointer to [**Link**](Link.md) |  | [optional] 
**Droplets** | Pointer to [**Link**](Link.md) |  | [optional] 
**EnvironmentVariables** | Pointer to [**Link**](Link.md) |  | [optional] 
**Features** | Pointer to [**Link**](Link.md) |  | [optional] 
**Packages** | Pointer to [**Link**](Link.md) |  | [optional] 
**Processes** | Pointer to [**Link**](Link.md) |  | [optional] 
**Revisions** | Pointer to [**Link**](Link.md) |  | [optional] 
**Self** | Pointer to [**Link**](Link.md) |  | [optional] 
**Space** | Pointer to [**Link**](Link.md) |  | [optional] 
**Start** | Pointer to [**LinkWithMethod**](LinkWithMethod.md) |  | [optional] 
**Stop** | Pointer to [**LinkWithMethod**](LinkWithMethod.md) |  | [optional] 
**Tasks** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewAppLinks

`func NewAppLinks() *AppLinks`

NewAppLinks instantiates a new AppLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinksWithDefaults

`func NewAppLinksWithDefaults() *AppLinks`

NewAppLinksWithDefaults instantiates a new AppLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentDroplet

`func (o *AppLinks) GetCurrentDroplet() Link`

GetCurrentDroplet returns the CurrentDroplet field if non-nil, zero value otherwise.

### GetCurrentDropletOk

`func (o *AppLinks) GetCurrentDropletOk() (*Link, bool)`

GetCurrentDropletOk returns a tuple with the CurrentDroplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDroplet

`func (o *AppLinks) SetCurrentDroplet(v Link)`

SetCurrentDroplet sets CurrentDroplet field to given value.

### HasCurrentDroplet

`func (o *AppLinks) HasCurrentDroplet() bool`

HasCurrentDroplet returns a boolean if a field has been set.

### GetDeployedRevisions

`func (o *AppLinks) GetDeployedRevisions() Link`

GetDeployedRevisions returns the DeployedRevisions field if non-nil, zero value otherwise.

### GetDeployedRevisionsOk

`func (o *AppLinks) GetDeployedRevisionsOk() (*Link, bool)`

GetDeployedRevisionsOk returns a tuple with the DeployedRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedRevisions

`func (o *AppLinks) SetDeployedRevisions(v Link)`

SetDeployedRevisions sets DeployedRevisions field to given value.

### HasDeployedRevisions

`func (o *AppLinks) HasDeployedRevisions() bool`

HasDeployedRevisions returns a boolean if a field has been set.

### GetDroplets

`func (o *AppLinks) GetDroplets() Link`

GetDroplets returns the Droplets field if non-nil, zero value otherwise.

### GetDropletsOk

`func (o *AppLinks) GetDropletsOk() (*Link, bool)`

GetDropletsOk returns a tuple with the Droplets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroplets

`func (o *AppLinks) SetDroplets(v Link)`

SetDroplets sets Droplets field to given value.

### HasDroplets

`func (o *AppLinks) HasDroplets() bool`

HasDroplets returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *AppLinks) GetEnvironmentVariables() Link`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *AppLinks) GetEnvironmentVariablesOk() (*Link, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *AppLinks) SetEnvironmentVariables(v Link)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *AppLinks) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetFeatures

`func (o *AppLinks) GetFeatures() Link`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AppLinks) GetFeaturesOk() (*Link, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AppLinks) SetFeatures(v Link)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AppLinks) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPackages

`func (o *AppLinks) GetPackages() Link`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AppLinks) GetPackagesOk() (*Link, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AppLinks) SetPackages(v Link)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AppLinks) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetProcesses

`func (o *AppLinks) GetProcesses() Link`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *AppLinks) GetProcessesOk() (*Link, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *AppLinks) SetProcesses(v Link)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *AppLinks) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetRevisions

`func (o *AppLinks) GetRevisions() Link`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *AppLinks) GetRevisionsOk() (*Link, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *AppLinks) SetRevisions(v Link)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *AppLinks) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSelf

`func (o *AppLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AppLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AppLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AppLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSpace

`func (o *AppLinks) GetSpace() Link`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *AppLinks) GetSpaceOk() (*Link, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *AppLinks) SetSpace(v Link)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *AppLinks) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetStart

`func (o *AppLinks) GetStart() LinkWithMethod`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AppLinks) GetStartOk() (*LinkWithMethod, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AppLinks) SetStart(v LinkWithMethod)`

SetStart sets Start field to given value.

### HasStart

`func (o *AppLinks) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *AppLinks) GetStop() LinkWithMethod`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *AppLinks) GetStopOk() (*LinkWithMethod, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *AppLinks) SetStop(v LinkWithMethod)`

SetStop sets Stop field to given value.

### HasStop

`func (o *AppLinks) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetTasks

`func (o *AppLinks) GetTasks() Link`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AppLinks) GetTasksOk() (*Link, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AppLinks) SetTasks(v Link)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AppLinks) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


