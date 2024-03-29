# AddIncidentRequestIncident
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | **String** | Description of the resolution to this incident | [optional] 
**Comment** | **String** | Comment on this incident, updates summary field | [optional] 
**Status** | **String** | Set status | [optional] 
**Severity** | **String** | Set severity | [optional] 
**Name** | **String** | Set display name | 
**StartDate** | **System.DateTime** | Set start time | [optional] 
**EndDate** | **System.DateTime** | Set start time | [optional] 
**InUptime** | **Boolean** | Set &#39;In Availability&#39; | [optional] 

## Examples

- Prepare the resource
```powershell
$AddIncidentRequestIncident = Initialize-PSOpenAPIToolsAddIncidentRequestIncident  -Resolution I plugged it back in `
 -Comment This is a summary of the incident `
 -Status null `
 -Severity null `
 -Name Incident Name `
 -StartDate 2019-10-20T19:42Z `
 -EndDate 2019-10-20T19:42Z `
 -InUptime true
```

- Convert the resource to JSON
```powershell
$AddIncidentRequestIncident | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

