# CheckType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**DefaultInterval** | **Int64** |  | [optional] 
**MetricName** | **String** |  | [optional] 
**InUptime** | **Boolean** |  | [optional] 
**CreateIncident** | **Boolean** |  | [optional] 
**PushOnly** | **Boolean** |  | [optional] 
**TunnelSupported** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckType = Initialize-PSOpenAPIToolsCheckType  -Id null `
 -Code null `
 -Name null `
 -DefaultInterval null `
 -MetricName null `
 -InUptime null `
 -CreateIncident null `
 -PushOnly null `
 -TunnelSupported null
```

- Convert the resource to JSON
```powershell
$CheckType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

