

# BlueprintCFTCreateCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | [**ConfigTypeEnum**](#ConfigTypeEnum) | Configuration Type | 
**json** | **String** | CloudFormation Template in JSON |  [optional]
**yaml** | **String** | CloudFormation Template in YAML |  [optional]
**git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  |  [optional]
**IAM** | **Boolean** | CloudFormation Attribute CAPABILITY_IAM |  [optional]
**CAPABILITY_NAMED_IAM** | **Boolean** | CloudFormation Attribute CAPABILITY_NAMED_IAM |  [optional]
**CAPABILITY_AUTO_EXPAND** | **Boolean** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND |  [optional]
**installAgent** | **Boolean** | Install Morpheus Agent |  [optional]
**cloudInitEnabled** | **Boolean** | Cloud Init Enabled |  [optional]



## Enum: ConfigTypeEnum

Name | Value
---- | -----
JSON | &quot;json&quot;
YAML | &quot;yaml&quot;
GIT | &quot;git&quot;



