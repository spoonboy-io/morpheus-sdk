# ApiMonitoringContactsContact
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the contact | 
**EmailAddress** | **String** | Email notification address | [optional] 
**SmsAddress** | **String** | SMS notification address | [optional] 
**SlackHook** | **String** | Slack Hook | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiMonitoringContactsContact = Initialize-PSOpenAPIToolsApiMonitoringContactsContact  -Name John Smith `
 -EmailAddress jsmith@email.com `
 -SmsAddress 555-555-5555 `
 -SlackHook https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
```

- Convert the resource to JSON
```powershell
$ApiMonitoringContactsContact | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

