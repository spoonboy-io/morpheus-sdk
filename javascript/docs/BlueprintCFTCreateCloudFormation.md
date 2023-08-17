# MorpheusApi.BlueprintCFTCreateCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | **String** | Configuration Type | 
**json** | **String** | CloudFormation Template in JSON | [optional] 
**yaml** | **String** | CloudFormation Template in YAML | [optional] 
**git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | **Boolean** | CloudFormation Attribute CAPABILITY_IAM | [optional] [default to false]
**CAPABILITY_NAMED_IAM** | **Boolean** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] [default to false]
**CAPABILITY_AUTO_EXPAND** | **Boolean** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] [default to false]
**installAgent** | **Boolean** | Install Morpheus Agent | [optional] [default to false]
**cloudInitEnabled** | **Boolean** | Cloud Init Enabled | [optional] [default to false]



## Enum: ConfigTypeEnum


* `json` (value: `"json"`)

* `yaml` (value: `"yaml"`)

* `git` (value: `"git"`)




