# GetClusterUpgradeVersions200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **String[]** |  | [optional] 
**CurrentVersion** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetClusterUpgradeVersions200Response = Initialize-PSOpenAPIToolsGetClusterUpgradeVersions200Response  -Versions null `
 -CurrentVersion null
```

- Convert the resource to JSON
```powershell
$GetClusterUpgradeVersions200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

