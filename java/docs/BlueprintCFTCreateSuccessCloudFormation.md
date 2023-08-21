

# BlueprintCFTCreateSuccessCloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | [**ConfigTypeEnum**](#ConfigTypeEnum) | Configuration Type | 
**json** | **String** | CloudFormation Template in JSON |  [optional]
**yaml** | **String** | CloudFormation Template in YAML |  [optional]
**git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  |  [optional]
**IAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_IAM |  [optional]
**CAPABILITY_NAMED_IAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_NAMED_IAM |  [optional]
**CAPABILITY_AUTO_EXPAND** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_AUTO_EXPAND |  [optional]
**installAgent** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Install Morpheus Agent |  [optional]
**cloudInitEnabled** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Cloud Init Enabled |  [optional]



## Enum: ConfigTypeEnum

Name | Value
---- | -----
JSON | &quot;json&quot;
YAML | &quot;yaml&quot;
GIT | &quot;git&quot;



