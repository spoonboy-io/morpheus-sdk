# BlueprintHelmCreateHelmGit
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoId** | **Int64** | Morpheus SCM Repository ID | 
**Path** | **String** | Path to Helm Files in the Repository | 
**IntegrationId** | **Int64** | Morpheus SCM Integration ID | 
**Branch** | **String** | Branch Name | 

## Examples

- Prepare the resource
```powershell
$BlueprintHelmCreateHelmGit = Initialize-PSOpenAPIToolsBlueprintHelmCreateHelmGit  -RepoId null `
 -Path null `
 -IntegrationId null `
 -Branch null
```

- Convert the resource to JSON
```powershell
$BlueprintHelmCreateHelmGit | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

