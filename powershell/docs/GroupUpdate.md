# GroupUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the group | 
**Code** | **String** | Optional code for use with policies | [optional] 
**Location** | **String** | Optional location argument for your group | [optional] 
**Config** | [**GroupCreateConfig**](GroupCreateConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GroupUpdate = Initialize-PSOpenAPIToolsGroupUpdate  -Name null `
 -Code null `
 -Location null `
 -Config null
```

- Convert the resource to JSON
```powershell
$GroupUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

