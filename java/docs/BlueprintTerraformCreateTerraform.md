

# BlueprintTerraformCreateTerraform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | [**ConfigTypeEnum**](#ConfigTypeEnum) | Configuration Type | 
**json** | **String** | Terraform definition in JSON for &#x60;configType&#x60; &#x60;json&#x60; |  [optional]
**tf** | **String** | Terraform definition for &#x60;configType&#x60; &#x60;tf&#x60; |  [optional]
**git** | [**BlueprintTerraformCreateTerraformGit**](BlueprintTerraformCreateTerraformGit.md) |  |  [optional]
**tfvarSecret** | **String** | tfvar secret from Morpheus Cypher |  [optional]



## Enum: ConfigTypeEnum

Name | Value
---- | -----
TF | &quot;tf&quot;
SPEC | &quot;spec&quot;
GIT | &quot;git&quot;
JSON | &quot;json&quot;



