# ClusterMastersInterfaces
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**UniqueId** | **String** |  | [optional] 
**PublicIpAddress** | **String** |  | [optional] 
**PublicIpv6Address** | **String** |  | [optional] 
**IpAddress** | **String** |  | [optional] 
**Ipv6Address** | **String** |  | [optional] 
**IpSubnet** | **String** |  | [optional] 
**Ipv6Subnet** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Dhcp** | **Boolean** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**PoolAssigned** | **Boolean** |  | [optional] 
**PrimaryInterface** | **Boolean** |  | [optional] 
**Network** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Subnet** | **String** |  | [optional] 
**NetworkGroup** | **String** |  | [optional] 
**NetworkPosition** | **String** |  | [optional] 
**NetworkPool** | **String** |  | [optional] 
**NetworkDomain** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**IpMode** | **String** |  | [optional] 
**MacAddress** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterMastersInterfaces = Initialize-PSOpenAPIToolsClusterMastersInterfaces  -Id null `
 -RefType null `
 -RefId null `
 -Name null `
 -InternalId null `
 -ExternalId null `
 -UniqueId null `
 -PublicIpAddress null `
 -PublicIpv6Address null `
 -IpAddress null `
 -Ipv6Address null `
 -IpSubnet null `
 -Ipv6Subnet null `
 -Description null `
 -Dhcp null `
 -Active null `
 -PoolAssigned null `
 -PrimaryInterface null `
 -Network null `
 -Subnet null `
 -NetworkGroup null `
 -NetworkPosition null `
 -NetworkPool null `
 -NetworkDomain null `
 -Type null `
 -IpMode null `
 -MacAddress null
```

- Convert the resource to JSON
```powershell
$ClusterMastersInterfaces | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

