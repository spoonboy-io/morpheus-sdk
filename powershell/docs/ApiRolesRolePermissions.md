# ApiRolesRolePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **String** | &#x60;code&#x60; of the feature permission | 
**Access** | **String** | The new access level. | 

## Examples

- Prepare the resource
```powershell
$ApiRolesRolePermissions = Initialize-PSOpenAPIToolsApiRolesRolePermissions  -Code null `
 -Access null
```

- Convert the resource to JSON
```powershell
$ApiRolesRolePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

