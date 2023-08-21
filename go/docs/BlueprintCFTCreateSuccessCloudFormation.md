# BlueprintCFTCreateSuccessCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | CloudFormation Template in JSON | [optional] 
**Yaml** | Pointer to **string** | CloudFormation Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintCFTCreateCloudFormationGit**](blueprintCFTCreate_cloudFormation_git.md) |  | [optional] 
**IAM** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | CloudFormation Attribute CAPABILITY_IAM | [optional] 
**CAPABILITY_NAMED_IAM** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] 
**CAPABILITY_AUTO_EXPAND** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] 
**InstallAgent** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | Install Morpheus Agent | [optional] 
**CloudInitEnabled** | Pointer to [**OneOfbooleanstring**](oneOf&lt;boolean,string&gt;.md) | Cloud Init Enabled | [optional] 

## Methods

### NewBlueprintCFTCreateSuccessCloudFormation

`func NewBlueprintCFTCreateSuccessCloudFormation(configType string, ) *BlueprintCFTCreateSuccessCloudFormation`

NewBlueprintCFTCreateSuccessCloudFormation instantiates a new BlueprintCFTCreateSuccessCloudFormation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCFTCreateSuccessCloudFormationWithDefaults

`func NewBlueprintCFTCreateSuccessCloudFormationWithDefaults() *BlueprintCFTCreateSuccessCloudFormation`

NewBlueprintCFTCreateSuccessCloudFormationWithDefaults instantiates a new BlueprintCFTCreateSuccessCloudFormation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetGit() BlueprintCFTCreateCloudFormationGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetGitOk() (*BlueprintCFTCreateCloudFormationGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetGit(v BlueprintCFTCreateCloudFormationGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetIAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetIAM() OneOfbooleanstring`

GetIAM returns the IAM field if non-nil, zero value otherwise.

### GetIAMOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetIAMOk() (*OneOfbooleanstring, bool)`

GetIAMOk returns a tuple with the IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetIAM(v OneOfbooleanstring)`

SetIAM sets IAM field to given value.

### HasIAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasIAM() bool`

HasIAM returns a boolean if a field has been set.

### GetCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCAPABILITY_NAMED_IAM() OneOfbooleanstring`

GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field if non-nil, zero value otherwise.

### GetCAPABILITY_NAMED_IAMOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCAPABILITY_NAMED_IAMOk() (*OneOfbooleanstring, bool)`

GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetCAPABILITY_NAMED_IAM(v OneOfbooleanstring)`

SetCAPABILITY_NAMED_IAM sets CAPABILITY_NAMED_IAM field to given value.

### HasCAPABILITY_NAMED_IAM

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasCAPABILITY_NAMED_IAM() bool`

HasCAPABILITY_NAMED_IAM returns a boolean if a field has been set.

### GetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCAPABILITY_AUTO_EXPAND() OneOfbooleanstring`

GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field if non-nil, zero value otherwise.

### GetCAPABILITY_AUTO_EXPANDOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCAPABILITY_AUTO_EXPANDOk() (*OneOfbooleanstring, bool)`

GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetCAPABILITY_AUTO_EXPAND(v OneOfbooleanstring)`

SetCAPABILITY_AUTO_EXPAND sets CAPABILITY_AUTO_EXPAND field to given value.

### HasCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasCAPABILITY_AUTO_EXPAND() bool`

HasCAPABILITY_AUTO_EXPAND returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetInstallAgent() OneOfbooleanstring`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetInstallAgentOk() (*OneOfbooleanstring, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetInstallAgent(v OneOfbooleanstring)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCloudInitEnabled() OneOfbooleanstring`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintCFTCreateSuccessCloudFormation) GetCloudInitEnabledOk() (*OneOfbooleanstring, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintCFTCreateSuccessCloudFormation) SetCloudInitEnabled(v OneOfbooleanstring)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintCFTCreateSuccessCloudFormation) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


