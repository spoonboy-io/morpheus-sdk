# BlueprintTerraformCreateConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specs** | [**BlueprintKubernetesCreateConfigSpecsInner[]**](BlueprintKubernetesCreateConfigSpecsInner.md) | Array of Terraform specs in Morpheus | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintTerraformCreateConfig = Initialize-PSOpenAPIToolsBlueprintTerraformCreateConfig  -Specs null
```

- Convert the resource to JSON
```powershell
$BlueprintTerraformCreateConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

