# BlueprintTerraformCreateTerraform
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Json** | **String** | Terraform definition in JSON for &#x60;configType&#x60; &#x60;json&#x60; | [optional] 
**Tf** | **String** | Terraform definition for &#x60;configType&#x60; &#x60;tf&#x60; | [optional] 
**Git** | [**BlueprintTerraformCreateTerraformGit**](BlueprintTerraformCreateTerraformGit.md) |  | [optional] 
**TfvarSecret** | **String** | tfvar secret from Morpheus Cypher | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintTerraformCreateTerraform = Initialize-PSOpenAPIToolsBlueprintTerraformCreateTerraform  -ConfigType null `
 -Json null `
 -Tf null `
 -Git null `
 -TfvarSecret null
```

- Convert the resource to JSON
```powershell
$BlueprintTerraformCreateTerraform | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

