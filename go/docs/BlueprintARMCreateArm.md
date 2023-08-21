# BlueprintARMCreateArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintARMCreateArmGit**](blueprintARMCreate_arm_git.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | Install Morpheus Agent | [optional] 
**CloudInitEnabled** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | Cloud Init Enabled | [optional] 

## Methods

### NewBlueprintARMCreateArm

`func NewBlueprintARMCreateArm(configType string, ) *BlueprintARMCreateArm`

NewBlueprintARMCreateArm instantiates a new BlueprintARMCreateArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintARMCreateArmWithDefaults

`func NewBlueprintARMCreateArmWithDefaults() *BlueprintARMCreateArm`

NewBlueprintARMCreateArmWithDefaults instantiates a new BlueprintARMCreateArm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintARMCreateArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintARMCreateArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintARMCreateArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintARMCreateArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintARMCreateArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintARMCreateArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintARMCreateArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintARMCreateArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintARMCreateArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintARMCreateArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintARMCreateArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintARMCreateArm) GetGit() BlueprintARMCreateArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintARMCreateArm) GetGitOk() (*BlueprintARMCreateArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintARMCreateArm) SetGit(v BlueprintARMCreateArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintARMCreateArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *BlueprintARMCreateArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *BlueprintARMCreateArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *BlueprintARMCreateArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *BlueprintARMCreateArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintARMCreateArm) GetInstallAgent() OneOfbooleanstring`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintARMCreateArm) GetInstallAgentOk() (*OneOfbooleanstring, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintARMCreateArm) SetInstallAgent(v OneOfbooleanstring)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintARMCreateArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintARMCreateArm) GetCloudInitEnabled() OneOfbooleanstring`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintARMCreateArm) GetCloudInitEnabledOk() (*OneOfbooleanstring, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintARMCreateArm) SetCloudInitEnabled(v OneOfbooleanstring)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintARMCreateArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


