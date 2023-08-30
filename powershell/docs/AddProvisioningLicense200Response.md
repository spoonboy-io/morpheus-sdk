# AddProvisioningLicense200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**License** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddProvisioningLicense200Response = Initialize-PSOpenAPIToolsAddProvisioningLicense200Response  -Success null `
 -License null
```

- Convert the resource to JSON
```powershell
$AddProvisioningLicense200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

