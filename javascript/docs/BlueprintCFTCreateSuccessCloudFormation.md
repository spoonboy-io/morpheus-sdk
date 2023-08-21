# MorpheusApi.BlueprintCFTCreateSuccessCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | **String** | Configuration Type | 
**json** | **String** | CloudFormation Template in JSON | [optional] 
**yaml** | **String** | CloudFormation Template in YAML | [optional] 
**git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_IAM | [optional] 
**CAPABILITY_NAMED_IAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] 
**CAPABILITY_AUTO_EXPAND** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] 
**installAgent** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Install Morpheus Agent | [optional] 
**cloudInitEnabled** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Cloud Init Enabled | [optional] 



## Enum: ConfigTypeEnum


* `json` (value: `"json"`)

* `yaml` (value: `"yaml"`)

* `git` (value: `"git"`)




