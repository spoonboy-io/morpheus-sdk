# AddRolesRequestRolePersonasInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **String** | &#x60;code&#x60; of the persona | [optional] 
**Access** | **String** | The new access level. | 

## Examples

- Prepare the resource
```powershell
$AddRolesRequestRolePersonasInner = Initialize-PSOpenAPIToolsAddRolesRequestRolePersonasInner  -Code null `
 -Access null
```

- Convert the resource to JSON
```powershell
$AddRolesRequestRolePersonasInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

