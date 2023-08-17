# AppSecurityGroups
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppSecurityGroups = Initialize-PSOpenAPIToolsAppSecurityGroups  -Id null `
 -AccountId null `
 -Name null `
 -Description null
```

- Convert the resource to JSON
```powershell
$AppSecurityGroups | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

