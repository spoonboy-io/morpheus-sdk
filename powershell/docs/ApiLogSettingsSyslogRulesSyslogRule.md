# ApiLogSettingsSyslogRulesSyslogRule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name of rule | 
**Rule** | **String** | Rule configuration | 

## Examples

- Prepare the resource
```powershell
$ApiLogSettingsSyslogRulesSyslogRule = Initialize-PSOpenAPIToolsApiLogSettingsSyslogRulesSyslogRule  -Name foo `
 -Rule *.* @@bar.com:8080
```

- Convert the resource to JSON
```powershell
$ApiLogSettingsSyslogRulesSyslogRule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

