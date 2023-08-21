# BlueprintKubernetesCreateConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specs** | [**BlueprintKubernetesCreateConfigSpecs[]**](BlueprintKubernetesCreateConfigSpecs.md) | Array of Kubernetes specs in Morpheus | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintKubernetesCreateConfig = Initialize-PSOpenAPIToolsBlueprintKubernetesCreateConfig  -Specs null
```

- Convert the resource to JSON
```powershell
$BlueprintKubernetesCreateConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

