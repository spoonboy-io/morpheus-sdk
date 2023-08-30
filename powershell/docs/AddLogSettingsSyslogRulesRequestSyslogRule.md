# AddLogSettingsSyslogRulesRequestSyslogRule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name of rule | 
**Rule** | **String** | Rule configuration | 

## Examples

- Prepare the resource
```powershell
$AddLogSettingsSyslogRulesRequestSyslogRule = Initialize-PSOpenAPIToolsAddLogSettingsSyslogRulesRequestSyslogRule  -Name foo `
 -Rule *.* @@bar.com:8080
```

- Convert the resource to JSON
```powershell
$AddLogSettingsSyslogRulesRequestSyslogRule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

