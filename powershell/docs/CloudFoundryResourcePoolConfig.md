# CloudFoundryResourcePoolConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managers** | **String[]** | Array of manager usernames | [optional] 
**Developers** | **String[]** | Array of developer usernames | [optional] 
**Auditors** | **String[]** | Array of auditor usernames | [optional] 

## Examples

- Prepare the resource
```powershell
$CloudFoundryResourcePoolConfig = Initialize-PSOpenAPIToolsCloudFoundryResourcePoolConfig  -Managers null `
 -Developers null `
 -Auditors null
```

- Convert the resource to JSON
```powershell
$CloudFoundryResourcePoolConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

