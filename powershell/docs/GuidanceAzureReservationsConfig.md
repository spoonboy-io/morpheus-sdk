# GuidanceAzureReservationsConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**DetailList** | [**GuidanceAzureReservationsConfigDetailListInner[]**](GuidanceAzureReservationsConfigDetailListInner.md) |  | [optional] 
**Services** | [**GuidanceAzureReservationsConfigServices**](GuidanceAzureReservationsConfigServices.md) |  | [optional] 
**Summary** | [**GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptionsValueTermOptionsValueSummary**](GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptionsValueTermOptionsValueSummary.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceAzureReservationsConfig = Initialize-PSOpenAPIToolsGuidanceAzureReservationsConfig  -Success null `
 -DetailList null `
 -Services null `
 -Summary null
```

- Convert the resource to JSON
```powershell
$GuidanceAzureReservationsConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

