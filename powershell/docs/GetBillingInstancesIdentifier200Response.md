# GetBillingInstancesIdentifier200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | [**BillingInstance**](BillingInstance.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBillingInstancesIdentifier200Response = Initialize-PSOpenAPIToolsGetBillingInstancesIdentifier200Response  -BillingInfo null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetBillingInstancesIdentifier200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

