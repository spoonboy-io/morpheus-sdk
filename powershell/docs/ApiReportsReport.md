# ApiReportsReport
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** | Code value for the report type you want to run | 
**StartDate** | **String** |  | [optional] 
**EndDate** | **String** |  | [optional] 
**StartMonth** | **String** |  | [optional] 
**EndMonth** | **String** |  | [optional] 
**GroupId** | **Decimal** | The Group ID filter for the report | [optional] 
**CloudId** | **Decimal** | The Cloud ID filter for the report | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiReportsReport = Initialize-PSOpenAPIToolsApiReportsReport  -Type appCost `
 -StartDate 2019-01-01 `
 -EndDate 2020-01-01 `
 -StartMonth 01/2019 `
 -EndMonth 01/2020 `
 -GroupId 1 `
 -CloudId 7
```

- Convert the resource to JSON
```powershell
$ApiReportsReport | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

