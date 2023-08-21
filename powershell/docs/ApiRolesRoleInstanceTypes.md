# ApiRolesRoleInstanceTypes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** | &#x60;id&#x60; of the instance type | 
**Access** | **String** | The new access level. | 

## Examples

- Prepare the resource
```powershell
$ApiRolesRoleInstanceTypes = Initialize-PSOpenAPIToolsApiRolesRoleInstanceTypes  -Id null `
 -Access null
```

- Convert the resource to JSON
```powershell
$ApiRolesRoleInstanceTypes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

