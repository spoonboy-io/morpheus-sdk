# CreateNetworkGroupRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkGroup** | [**NetworkGroupsCreate**](NetworkGroupsCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkGroupRequest = Initialize-PSOpenAPIToolsCreateNetworkGroupRequest  -NetworkGroup null
```

- Convert the resource to JSON
```powershell
$CreateNetworkGroupRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

