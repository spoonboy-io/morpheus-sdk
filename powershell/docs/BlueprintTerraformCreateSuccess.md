# BlueprintTerraformCreateSuccess
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | [optional] 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | [optional] 
**Terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | [optional] 
**Config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 
**Visibility** | **String** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | [**SystemCollectionsHashtable**](.md) | Resource Permission Block | [optional] 
**Owner** | [**SystemCollectionsHashtable**](.md) | Owner | [optional] 
**Tenant** | [**SystemCollectionsHashtable**](.md) | Tenant | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintTerraformCreateSuccess = Initialize-PSOpenAPIToolsBlueprintTerraformCreateSuccess  -Name null `
 -Image null `
 -Type null `
 -Terraform null `
 -Config null `
 -Visibility null `
 -ResourcePermission null `
 -Owner null `
 -Tenant null
```

- Convert the resource to JSON
```powershell
$BlueprintTerraformCreateSuccess | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

