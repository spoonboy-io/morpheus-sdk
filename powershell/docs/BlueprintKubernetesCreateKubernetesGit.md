# BlueprintKubernetesCreateKubernetesGit
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoId** | **Int64** | Morpheus SCM Repository ID | [optional] 
**Path** | **String** | Path to kubernetes Files in the Repository | [optional] 
**IntegrationId** | **Int64** | Morpheus SCM Integration ID | [optional] 
**Branch** | **String** | Branch Name | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintKubernetesCreateKubernetesGit = Initialize-PSOpenAPIToolsBlueprintKubernetesCreateKubernetesGit  -RepoId null `
 -Path null `
 -IntegrationId null `
 -Branch null
```

- Convert the resource to JSON
```powershell
$BlueprintKubernetesCreateKubernetesGit | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

