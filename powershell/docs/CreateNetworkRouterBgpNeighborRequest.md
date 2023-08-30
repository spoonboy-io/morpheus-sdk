# CreateNetworkRouterBgpNeighborRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkRouterBgpNeighbor** | [**SystemCollectionsHashtable**](.md) | For a full list of available options, see bgpOptionTypes in the specific Network Router Type | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkRouterBgpNeighborRequest = Initialize-PSOpenAPIToolsCreateNetworkRouterBgpNeighborRequest  -NetworkRouterBgpNeighbor null
```

- Convert the resource to JSON
```powershell
$CreateNetworkRouterBgpNeighborRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

