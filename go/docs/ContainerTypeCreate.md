# ContainerTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Node type name | 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | **string** | The short name is a name with no spaces used for display in your container list. | 
**Description** | Pointer to **string** | Node type description | [optional] 
**ContainerVersion** | **string** | Version of the node type | 
**ProvisionTypeCode** | **string** | Provision type code, eg. &#x60;amazon&#x60;, etc. | 
**Scripts** | Pointer to **[]int64** | Array of script IDs. | [optional] 
**Templates** | Pointer to **[]int64** | Array of file template IDs. | [optional] 
**VirtualImageId** | Pointer to **int64** | Virtual image ID | [optional] 
**StatTypeCode** | Pointer to **string** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**LogTypeCode** | Pointer to **string** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**ServerType** | Pointer to **string** | Server type.  Always pass \&quot;vm\&quot;. | [optional] 
**ContainerPorts** | Pointer to [**[]ContainerTypeCreateContainerPorts**](ContainerTypeCreateContainerPorts.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL | [optional] 
**EnvironmentVariables** | Pointer to [**[]ClusterLayoutCreateEnvironmentVariables**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. | [optional] 

## Methods

### NewContainerTypeCreate

`func NewContainerTypeCreate(name string, shortName string, containerVersion string, provisionTypeCode string, ) *ContainerTypeCreate`

NewContainerTypeCreate instantiates a new ContainerTypeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTypeCreateWithDefaults

`func NewContainerTypeCreateWithDefaults() *ContainerTypeCreate`

NewContainerTypeCreateWithDefaults instantiates a new ContainerTypeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerTypeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTypeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTypeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *ContainerTypeCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ContainerTypeCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ContainerTypeCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ContainerTypeCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *ContainerTypeCreate) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ContainerTypeCreate) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ContainerTypeCreate) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetDescription

`func (o *ContainerTypeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerTypeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerTypeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerTypeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ContainerTypeCreate) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ContainerTypeCreate) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ContainerTypeCreate) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.


### GetProvisionTypeCode

`func (o *ContainerTypeCreate) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *ContainerTypeCreate) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *ContainerTypeCreate) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.


### GetScripts

`func (o *ContainerTypeCreate) GetScripts() []int64`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ContainerTypeCreate) GetScriptsOk() (*[]int64, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ContainerTypeCreate) SetScripts(v []int64)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *ContainerTypeCreate) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetTemplates

`func (o *ContainerTypeCreate) GetTemplates() []int64`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ContainerTypeCreate) GetTemplatesOk() (*[]int64, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ContainerTypeCreate) SetTemplates(v []int64)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *ContainerTypeCreate) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetVirtualImageId

`func (o *ContainerTypeCreate) GetVirtualImageId() int64`

GetVirtualImageId returns the VirtualImageId field if non-nil, zero value otherwise.

### GetVirtualImageIdOk

`func (o *ContainerTypeCreate) GetVirtualImageIdOk() (*int64, bool)`

GetVirtualImageIdOk returns a tuple with the VirtualImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImageId

`func (o *ContainerTypeCreate) SetVirtualImageId(v int64)`

SetVirtualImageId sets VirtualImageId field to given value.

### HasVirtualImageId

`func (o *ContainerTypeCreate) HasVirtualImageId() bool`

HasVirtualImageId returns a boolean if a field has been set.

### GetStatTypeCode

`func (o *ContainerTypeCreate) GetStatTypeCode() string`

GetStatTypeCode returns the StatTypeCode field if non-nil, zero value otherwise.

### GetStatTypeCodeOk

`func (o *ContainerTypeCreate) GetStatTypeCodeOk() (*string, bool)`

GetStatTypeCodeOk returns a tuple with the StatTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTypeCode

`func (o *ContainerTypeCreate) SetStatTypeCode(v string)`

SetStatTypeCode sets StatTypeCode field to given value.

### HasStatTypeCode

`func (o *ContainerTypeCreate) HasStatTypeCode() bool`

HasStatTypeCode returns a boolean if a field has been set.

### GetLogTypeCode

`func (o *ContainerTypeCreate) GetLogTypeCode() string`

GetLogTypeCode returns the LogTypeCode field if non-nil, zero value otherwise.

### GetLogTypeCodeOk

`func (o *ContainerTypeCreate) GetLogTypeCodeOk() (*string, bool)`

GetLogTypeCodeOk returns a tuple with the LogTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeCode

`func (o *ContainerTypeCreate) SetLogTypeCode(v string)`

SetLogTypeCode sets LogTypeCode field to given value.

### HasLogTypeCode

`func (o *ContainerTypeCreate) HasLogTypeCode() bool`

HasLogTypeCode returns a boolean if a field has been set.

### GetServerType

`func (o *ContainerTypeCreate) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ContainerTypeCreate) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ContainerTypeCreate) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ContainerTypeCreate) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetContainerPorts

`func (o *ContainerTypeCreate) GetContainerPorts() []ContainerTypeCreateContainerPorts`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ContainerTypeCreate) GetContainerPortsOk() (*[]ContainerTypeCreateContainerPorts, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ContainerTypeCreate) SetContainerPorts(v []ContainerTypeCreateContainerPorts)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ContainerTypeCreate) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ContainerTypeCreate) GetEnvironmentVariables() []ClusterLayoutCreateEnvironmentVariables`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ContainerTypeCreate) GetEnvironmentVariablesOk() (*[]ClusterLayoutCreateEnvironmentVariables, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ContainerTypeCreate) SetEnvironmentVariables(v []ClusterLayoutCreateEnvironmentVariables)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ContainerTypeCreate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetConfig

`func (o *ContainerTypeCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerTypeCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerTypeCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ContainerTypeCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


