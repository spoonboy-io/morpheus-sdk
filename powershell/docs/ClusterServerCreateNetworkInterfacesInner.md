# ClusterServerCreateNetworkInterfacesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**ClusterServerCreateNetworkInterfacesInnerNetwork**](ClusterServerCreateNetworkInterfacesInnerNetwork.md) |  | 
**NetworkInterfaceTypeId** | **Int64** | The id of type of the network interface. | [optional] 
**IpAddress** | **String** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterServerCreateNetworkInterfacesInner = Initialize-PSOpenAPIToolsClusterServerCreateNetworkInterfacesInner  -Network null `
 -NetworkInterfaceTypeId null `
 -IpAddress null
```

- Convert the resource to JSON
```powershell
$ClusterServerCreateNetworkInterfacesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

