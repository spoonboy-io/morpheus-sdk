# Alarm
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Alarm = Initialize-PSOpenAPIToolsAlarm  -Id null `
 -Name null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$Alarm | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

