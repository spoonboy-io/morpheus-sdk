# ClusterLayoutContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**VirtualImage** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewClusterLayoutContainerType

`func NewClusterLayoutContainerType() *ClusterLayoutContainerType`

NewClusterLayoutContainerType instantiates a new ClusterLayoutContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutContainerTypeWithDefaults

`func NewClusterLayoutContainerTypeWithDefaults() *ClusterLayoutContainerType`

NewClusterLayoutContainerTypeWithDefaults instantiates a new ClusterLayoutContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterLayoutContainerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterLayoutContainerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterLayoutContainerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterLayoutContainerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterLayoutContainerType) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterLayoutContainerType) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterLayoutContainerType) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterLayoutContainerType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ClusterLayoutContainerType) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ClusterLayoutContainerType) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *ClusterLayoutContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutContainerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutContainerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterLayoutContainerType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterLayoutContainerType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterLayoutContainerType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterLayoutContainerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClusterLayoutContainerType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClusterLayoutContainerType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetShortName

`func (o *ClusterLayoutContainerType) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ClusterLayoutContainerType) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ClusterLayoutContainerType) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ClusterLayoutContainerType) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *ClusterLayoutContainerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterLayoutContainerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterLayoutContainerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterLayoutContainerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ClusterLayoutContainerType) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ClusterLayoutContainerType) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ClusterLayoutContainerType) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ClusterLayoutContainerType) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *ClusterLayoutContainerType) GetProvisionType() InlineResponse20094Network`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ClusterLayoutContainerType) GetProvisionTypeOk() (*InlineResponse20094Network, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ClusterLayoutContainerType) SetProvisionType(v InlineResponse20094Network)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ClusterLayoutContainerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ClusterLayoutContainerType) GetVirtualImage() string`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ClusterLayoutContainerType) GetVirtualImageOk() (*string, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ClusterLayoutContainerType) SetVirtualImage(v string)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ClusterLayoutContainerType) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### SetVirtualImageNil

`func (o *ClusterLayoutContainerType) SetVirtualImageNil(b bool)`

 SetVirtualImageNil sets the value for VirtualImage to be an explicit nil

### UnsetVirtualImage
`func (o *ClusterLayoutContainerType) UnsetVirtualImage()`

UnsetVirtualImage ensures that no value is present for VirtualImage, not even an explicit nil
### GetCategory

`func (o *ClusterLayoutContainerType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterLayoutContainerType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterLayoutContainerType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterLayoutContainerType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterLayoutContainerType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterLayoutContainerType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterLayoutContainerType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterLayoutContainerType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainerPorts

`func (o *ClusterLayoutContainerType) GetContainerPorts() []map[string]interface{}`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ClusterLayoutContainerType) GetContainerPortsOk() (*[]map[string]interface{}, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ClusterLayoutContainerType) SetContainerPorts(v []map[string]interface{})`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ClusterLayoutContainerType) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetContainerScripts

`func (o *ClusterLayoutContainerType) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ClusterLayoutContainerType) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ClusterLayoutContainerType) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ClusterLayoutContainerType) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### GetContainerTemplates

`func (o *ClusterLayoutContainerType) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ClusterLayoutContainerType) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ClusterLayoutContainerType) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ClusterLayoutContainerType) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ClusterLayoutContainerType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ClusterLayoutContainerType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ClusterLayoutContainerType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ClusterLayoutContainerType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


