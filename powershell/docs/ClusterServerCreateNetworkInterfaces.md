# ClusterServerCreateNetworkInterfaces
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**ClusterServerCreateNetwork**](ClusterServerCreateNetwork.md) |  | 
**NetworkInterfaceTypeId** | **Int64** | The id of type of the network interface. | [optional] 
**IpAddress** | **String** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterServerCreateNetworkInterfaces = Initialize-PSOpenAPIToolsClusterServerCreateNetworkInterfaces  -Network null `
 -NetworkInterfaceTypeId null `
 -IpAddress null
```

- Convert the resource to JSON
```powershell
$ClusterServerCreateNetworkInterfaces | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

