# BlueprintKubernetesCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | 
**Config** | [**BlueprintKubernetesCreateConfig**](BlueprintKubernetesCreateConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintKubernetesCreate = Initialize-PSOpenAPIToolsBlueprintKubernetesCreate  -Name null `
 -Image null `
 -Type null `
 -Labels null `
 -Kubernetes null `
 -Config null
```

- Convert the resource to JSON
```powershell
$BlueprintKubernetesCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

