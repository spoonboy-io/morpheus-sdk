# ApiRolesRolePersonas
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **String** | &#x60;code&#x60; of the persona | [optional] 
**Access** | **String** | The new access level. | 

## Examples

- Prepare the resource
```powershell
$ApiRolesRolePersonas = Initialize-PSOpenAPIToolsApiRolesRolePersonas  -Code null `
 -Access null
```

- Convert the resource to JSON
```powershell
$ApiRolesRolePersonas | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

