# CreateStaticRouteRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkRoute** | [**NetworkStaticRouteCreate**](NetworkStaticRouteCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateStaticRouteRequest = Initialize-PSOpenAPIToolsCreateStaticRouteRequest  -NetworkRoute null
```

- Convert the resource to JSON
```powershell
$CreateStaticRouteRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

