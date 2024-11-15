# V3AppsGuidTasksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | Command that will be executed; optional if a template.process.guid is provided | [optional] 
**DiskInMb** | Pointer to **int32** | Amount of disk to allocate for the task in MB | [optional] 
**DropletGuid** | Pointer to **string** | The guid of the droplet that will be used to run the command | [optional] 
**LogRateLimitInBytesPerSecond** | Pointer to **int32** | Amount of log rate to allocate for the task in bytes | [optional] 
**MemoryInMb** | Pointer to **int32** | Amount of memory to allocate for the task in MB | [optional] 
**Metadata** | Pointer to [**V3AppsGuidTasksPostRequestMetadata**](V3AppsGuidTasksPostRequestMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the task; auto-generated if not provided | [optional] 
**Template** | Pointer to [**V3AppsGuidTasksPostRequestTemplate**](V3AppsGuidTasksPostRequestTemplate.md) |  | [optional] 

## Methods

### NewV3AppsGuidTasksPostRequest

`func NewV3AppsGuidTasksPostRequest() *V3AppsGuidTasksPostRequest`

NewV3AppsGuidTasksPostRequest instantiates a new V3AppsGuidTasksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AppsGuidTasksPostRequestWithDefaults

`func NewV3AppsGuidTasksPostRequestWithDefaults() *V3AppsGuidTasksPostRequest`

NewV3AppsGuidTasksPostRequestWithDefaults instantiates a new V3AppsGuidTasksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *V3AppsGuidTasksPostRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V3AppsGuidTasksPostRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V3AppsGuidTasksPostRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V3AppsGuidTasksPostRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDiskInMb

`func (o *V3AppsGuidTasksPostRequest) GetDiskInMb() int32`

GetDiskInMb returns the DiskInMb field if non-nil, zero value otherwise.

### GetDiskInMbOk

`func (o *V3AppsGuidTasksPostRequest) GetDiskInMbOk() (*int32, bool)`

GetDiskInMbOk returns a tuple with the DiskInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskInMb

`func (o *V3AppsGuidTasksPostRequest) SetDiskInMb(v int32)`

SetDiskInMb sets DiskInMb field to given value.

### HasDiskInMb

`func (o *V3AppsGuidTasksPostRequest) HasDiskInMb() bool`

HasDiskInMb returns a boolean if a field has been set.

### GetDropletGuid

`func (o *V3AppsGuidTasksPostRequest) GetDropletGuid() string`

GetDropletGuid returns the DropletGuid field if non-nil, zero value otherwise.

### GetDropletGuidOk

`func (o *V3AppsGuidTasksPostRequest) GetDropletGuidOk() (*string, bool)`

GetDropletGuidOk returns a tuple with the DropletGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropletGuid

`func (o *V3AppsGuidTasksPostRequest) SetDropletGuid(v string)`

SetDropletGuid sets DropletGuid field to given value.

### HasDropletGuid

`func (o *V3AppsGuidTasksPostRequest) HasDropletGuid() bool`

HasDropletGuid returns a boolean if a field has been set.

### GetLogRateLimitInBytesPerSecond

`func (o *V3AppsGuidTasksPostRequest) GetLogRateLimitInBytesPerSecond() int32`

GetLogRateLimitInBytesPerSecond returns the LogRateLimitInBytesPerSecond field if non-nil, zero value otherwise.

### GetLogRateLimitInBytesPerSecondOk

`func (o *V3AppsGuidTasksPostRequest) GetLogRateLimitInBytesPerSecondOk() (*int32, bool)`

GetLogRateLimitInBytesPerSecondOk returns a tuple with the LogRateLimitInBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRateLimitInBytesPerSecond

`func (o *V3AppsGuidTasksPostRequest) SetLogRateLimitInBytesPerSecond(v int32)`

SetLogRateLimitInBytesPerSecond sets LogRateLimitInBytesPerSecond field to given value.

### HasLogRateLimitInBytesPerSecond

`func (o *V3AppsGuidTasksPostRequest) HasLogRateLimitInBytesPerSecond() bool`

HasLogRateLimitInBytesPerSecond returns a boolean if a field has been set.

### GetMemoryInMb

`func (o *V3AppsGuidTasksPostRequest) GetMemoryInMb() int32`

GetMemoryInMb returns the MemoryInMb field if non-nil, zero value otherwise.

### GetMemoryInMbOk

`func (o *V3AppsGuidTasksPostRequest) GetMemoryInMbOk() (*int32, bool)`

GetMemoryInMbOk returns a tuple with the MemoryInMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInMb

`func (o *V3AppsGuidTasksPostRequest) SetMemoryInMb(v int32)`

SetMemoryInMb sets MemoryInMb field to given value.

### HasMemoryInMb

`func (o *V3AppsGuidTasksPostRequest) HasMemoryInMb() bool`

HasMemoryInMb returns a boolean if a field has been set.

### GetMetadata

`func (o *V3AppsGuidTasksPostRequest) GetMetadata() V3AppsGuidTasksPostRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V3AppsGuidTasksPostRequest) GetMetadataOk() (*V3AppsGuidTasksPostRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V3AppsGuidTasksPostRequest) SetMetadata(v V3AppsGuidTasksPostRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V3AppsGuidTasksPostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *V3AppsGuidTasksPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3AppsGuidTasksPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3AppsGuidTasksPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3AppsGuidTasksPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *V3AppsGuidTasksPostRequest) GetTemplate() V3AppsGuidTasksPostRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V3AppsGuidTasksPostRequest) GetTemplateOk() (*V3AppsGuidTasksPostRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V3AppsGuidTasksPostRequest) SetTemplate(v V3AppsGuidTasksPostRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V3AppsGuidTasksPostRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


