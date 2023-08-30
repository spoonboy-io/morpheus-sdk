# BlueprintKubernetesCreateSuccess
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | [optional] 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | [optional] 
**Kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | [optional] 
**Config** | [**BlueprintKubernetesCreateConfig**](BlueprintKubernetesCreateConfig.md) |  | [optional] 
**Visibility** | **String** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | [**SystemCollectionsHashtable**](.md) | Resource Permission Block | [optional] 
**Owner** | [**SystemCollectionsHashtable**](.md) | Owner | [optional] 
**Tenant** | [**SystemCollectionsHashtable**](.md) | Tenant | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintKubernetesCreateSuccess = Initialize-PSOpenAPIToolsBlueprintKubernetesCreateSuccess  -Name null `
 -Image null `
 -Type null `
 -Kubernetes null `
 -Config null `
 -Visibility null `
 -ResourcePermission null `
 -Owner null `
 -Tenant null
```

- Convert the resource to JSON
```powershell
$BlueprintKubernetesCreateSuccess | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

