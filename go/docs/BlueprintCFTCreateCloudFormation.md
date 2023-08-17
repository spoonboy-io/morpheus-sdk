# BlueprintCFTCreateCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | CloudFormation Template in JSON | [optional] 
**Yaml** | Pointer to **string** | CloudFormation Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_IAM | [optional] [default to false]
**CAPABILITY_NAMED_IAM** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] [default to false]
**CAPABILITY_AUTO_EXPAND** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] [default to false]
**InstallAgent** | Pointer to **bool** | Install Morpheus Agent | [optional] [default to false]
**CloudInitEnabled** | Pointer to **bool** | Cloud Init Enabled | [optional] [default to false]

## Methods

### NewBlueprintCFTCreateCloudFormation

`func NewBlueprintCFTCreateCloudFormation(configType string, ) *BlueprintCFTCreateCloudFormation`

NewBlueprintCFTCreateCloudFormation instantiates a new BlueprintCFTCreateCloudFormation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCFTCreateCloudFormationWithDefaults

`func NewBlueprintCFTCreateCloudFormationWithDefaults() *BlueprintCFTCreateCloudFormation`

NewBlueprintCFTCreateCloudFormationWithDefaults instantiates a new BlueprintCFTCreateCloudFormation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintCFTCreateCloudFormation) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintCFTCreateCloudFormation) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintCFTCreateCloudFormation) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintCFTCreateCloudFormation) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintCFTCreateCloudFormation) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintCFTCreateCloudFormation) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintCFTCreateCloudFormation) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintCFTCreateCloudFormation) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintCFTCreateCloudFormation) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintCFTCreateCloudFormation) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintCFTCreateCloudFormation) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintCFTCreateCloudFormation) GetGit() BlueprintCFTCreateCloudFormationGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintCFTCreateCloudFormation) GetGitOk() (*BlueprintCFTCreateCloudFormationGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintCFTCreateCloudFormation) SetGit(v BlueprintCFTCreateCloudFormationGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintCFTCreateCloudFormation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetIAM

`func (o *BlueprintCFTCreateCloudFormation) GetIAM() bool`

GetIAM returns the IAM field if non-nil, zero value otherwise.

### GetIAMOk

`func (o *BlueprintCFTCreateCloudFormation) GetIAMOk() (*bool, bool)`

GetIAMOk returns a tuple with the IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAM

`func (o *BlueprintCFTCreateCloudFormation) SetIAM(v bool)`

SetIAM sets IAM field to given value.

### HasIAM

`func (o *BlueprintCFTCreateCloudFormation) HasIAM() bool`

HasIAM returns a boolean if a field has been set.

### GetCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateCloudFormation) GetCAPABILITY_NAMED_IAM() bool`

GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field if non-nil, zero value otherwise.

### GetCAPABILITY_NAMED_IAMOk

`func (o *BlueprintCFTCreateCloudFormation) GetCAPABILITY_NAMED_IAMOk() (*bool, bool)`

GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateCloudFormation) SetCAPABILITY_NAMED_IAM(v bool)`

SetCAPABILITY_NAMED_IAM sets CAPABILITY_NAMED_IAM field to given value.

### HasCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateCloudFormation) HasCAPABILITY_NAMED_IAM() bool`

HasCAPABILITY_NAMED_IAM returns a boolean if a field has been set.

### GetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateCloudFormation) GetCAPABILITY_AUTO_EXPAND() bool`

GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field if non-nil, zero value otherwise.

### GetCAPABILITY_AUTO_EXPANDOk

`func (o *BlueprintCFTCreateCloudFormation) GetCAPABILITY_AUTO_EXPANDOk() (*bool, bool)`

GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateCloudFormation) SetCAPABILITY_AUTO_EXPAND(v bool)`

SetCAPABILITY_AUTO_EXPAND sets CAPABILITY_AUTO_EXPAND field to given value.

### HasCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateCloudFormation) HasCAPABILITY_AUTO_EXPAND() bool`

HasCAPABILITY_AUTO_EXPAND returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintCFTCreateCloudFormation) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintCFTCreateCloudFormation) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintCFTCreateCloudFormation) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintCFTCreateCloudFormation) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintCFTCreateCloudFormation) GetCloudInitEnabled() bool`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintCFTCreateCloudFormation) GetCloudInitEnabledOk() (*bool, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintCFTCreateCloudFormation) SetCloudInitEnabled(v bool)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintCFTCreateCloudFormation) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


