# BlueprintKubernetesCreateKubernetes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Yaml** | **String** | Kubernetes Spec in YAML | [optional] 
**Git** | [**BlueprintKubernetesCreateKubernetesGit**](BlueprintKubernetesCreateKubernetesGit.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintKubernetesCreateKubernetes = Initialize-PSOpenAPIToolsBlueprintKubernetesCreateKubernetes  -ConfigType null `
 -Yaml null `
 -Git null
```

- Convert the resource to JSON
```powershell
$BlueprintKubernetesCreateKubernetes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

