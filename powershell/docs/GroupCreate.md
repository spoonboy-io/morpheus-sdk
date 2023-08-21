# GroupCreate
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
$GroupCreate = Initialize-PSOpenAPIToolsGroupCreate  -Name null `
 -Code null `
 -Location null `
 -Config null
```

- Convert the resource to JSON
```powershell
$GroupCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

