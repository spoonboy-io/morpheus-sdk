# MorpheusApi.BlueprintTerraformCreateTerraform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | **String** | Configuration Type | 
**json** | **String** | Terraform definition in JSON for &#x60;configType&#x60; &#x60;json&#x60; | [optional] 
**tf** | **String** | Terraform definition for &#x60;configType&#x60; &#x60;tf&#x60; | [optional] 
**git** | [**BlueprintTerraformCreateTerraformGit**](BlueprintTerraformCreateTerraformGit.md) |  | [optional] 
**tfvarSecret** | **String** | tfvar secret from Morpheus Cypher | [optional] 



## Enum: ConfigTypeEnum


* `tf` (value: `"tf"`)

* `spec` (value: `"spec"`)

* `git` (value: `"git"`)

* `json` (value: `"json"`)




