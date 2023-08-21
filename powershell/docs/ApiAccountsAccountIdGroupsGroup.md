# ApiAccountsAccountIdGroupsGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to the subtenant for the group | 
**Description** | **String** | Optional description field if you want to put more info there | [optional] 
**Code** | **String** | Optional code for use with policies | [optional] 
**Location** | **String** | location | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiAccountsAccountIdGroupsGroup = Initialize-PSOpenAPIToolsApiAccountsAccountIdGroupsGroup  -Name null `
 -Description null `
 -Code null `
 -Location null
```

- Convert the resource to JSON
```powershell
$ApiAccountsAccountIdGroupsGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

