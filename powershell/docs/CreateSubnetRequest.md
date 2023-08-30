# CreateSubnetRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | [**CreateSubnetRequestSubnet**](CreateSubnetRequestSubnet.md) |  | [optional] 
**ResourcePermission** | [**CreateSubnetRequestResourcePermission**](CreateSubnetRequestResourcePermission.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateSubnetRequest = Initialize-PSOpenAPIToolsCreateSubnetRequest  -Subnet null `
 -ResourcePermission null
```

- Convert the resource to JSON
```powershell
$CreateSubnetRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

