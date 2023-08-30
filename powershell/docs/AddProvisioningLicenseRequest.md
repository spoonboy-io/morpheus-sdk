# AddProvisioningLicenseRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | [**ProvisioningLicensesCreate**](ProvisioningLicensesCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddProvisioningLicenseRequest = Initialize-PSOpenAPIToolsAddProvisioningLicenseRequest  -License null
```

- Convert the resource to JSON
```powershell
$AddProvisioningLicenseRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

