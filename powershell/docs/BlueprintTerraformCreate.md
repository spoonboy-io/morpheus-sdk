# BlueprintTerraformCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | 
**Config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintTerraformCreate = Initialize-PSOpenAPIToolsBlueprintTerraformCreate  -Name null `
 -Image null `
 -Type null `
 -Labels null `
 -Terraform null `
 -Config null
```

- Convert the resource to JSON
```powershell
$BlueprintTerraformCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

