# MorpheusApi.BlueprintTerraformCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the blueprint | [optional] 
**image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**type** | **String** | Blueprint Type | [optional] 
**terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | [optional] 
**config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 
**visibility** | **String** | Private or Public Access | [optional] [default to &#39;private&#39;]
**resourcePermission** | **Object** | Resource Permission Block | [optional] 
**owner** | **Object** | Owner | [optional] 
**tenant** | **Object** | Tenant | [optional] 



## Enum: TypeEnum


* `terraform` (value: `"terraform"`)





## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




