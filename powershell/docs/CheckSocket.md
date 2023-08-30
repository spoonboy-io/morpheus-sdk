# CheckSocket
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the check | [optional] 
**Description** | **String** | Optional description field | [optional] 
**CheckType** | [**CheckSocketCheckType**](CheckSocketCheckType.md) |  | [optional] 
**CheckInterval** | **Int32** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**InUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations | [optional] [default to $true]
**Active** | **Boolean** | Used to determine if check should be scheduled to execute | [optional] [default to $true]
**Severity** | **String** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Config** | [**CheckSocketConfig**](CheckSocketConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckSocket = Initialize-PSOpenAPIToolsCheckSocket  -Name My Check `
 -Description This is my awesome check `
 -CheckType null `
 -CheckInterval null `
 -InUptime null `
 -Active null `
 -Severity null `
 -Config null
```

- Convert the resource to JSON
```powershell
$CheckSocket | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

