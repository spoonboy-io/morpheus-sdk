# Deployment
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**VersionCount** | **Int64** |  | [optional] 
**Versions** | [**DeploymentVersions[]**](DeploymentVersions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Deployment = Initialize-PSOpenAPIToolsDeployment  -Id null `
 -Name null `
 -Description null `
 -AccountId null `
 -ExternalId null `
 -DateCreated null `
 -LastUpdated null `
 -VersionCount null `
 -Versions null
```

- Convert the resource to JSON
```powershell
$Deployment | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

