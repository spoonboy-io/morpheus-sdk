# CreateNetworkPoolServer200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**NetworkPoolServer** | [**NetworkPoolServer**](NetworkPoolServer.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkPoolServer200Response = Initialize-PSOpenAPIToolsCreateNetworkPoolServer200Response  -Success null `
 -NetworkPoolServer null
```

- Convert the resource to JSON
```powershell
$CreateNetworkPoolServer200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

