# AddVDIGatewaysRequestVdiGateway
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | VDI Gateway name | 
**Description** | **String** | Description | [optional] 
**GatewayUrl** | **String** | Gateway URL | 

## Examples

- Prepare the resource
```powershell
$AddVDIGatewaysRequestVdiGateway = Initialize-PSOpenAPIToolsAddVDIGatewaysRequestVdiGateway  -Name null `
 -Description null `
 -GatewayUrl https://fqdn.com
```

- Convert the resource to JSON
```powershell
$AddVDIGatewaysRequestVdiGateway | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

