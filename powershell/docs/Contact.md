# Contact
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**EmailAddress** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**SmsAddress** | **String** |  | [optional] 
**SlackHook** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Contact = Initialize-PSOpenAPIToolsContact  -Id null `
 -EmailAddress null `
 -Name null `
 -SmsAddress null `
 -SlackHook null
```

- Convert the resource to JSON
```powershell
$Contact | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

