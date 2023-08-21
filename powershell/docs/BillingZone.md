# BillingZone
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | **String** |  | [optional] 
**ZoneId** | **Int64** |  | [optional] 
**ZoneUUID** | **String** |  | [optional] 
**ZoneCode** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**PriceUnit** | **String** |  | [optional] 
**ComputeServers** | [**BillingComputeServers**](BillingComputeServers.md) |  | [optional] 
**Instances** | [**BillingInstances**](BillingInstances.md) |  | [optional] 
**DiscoveredServers** | [**BillingComputeServers**](BillingComputeServers.md) |  | [optional] 
**LoadBalancers** | [**BillingLoadBalancers**](BillingLoadBalancers.md) |  | [optional] 
**VirtualImages** | [**BillingVirtualImages**](BillingVirtualImages.md) |  | [optional] 
**Snapshots** | [**BillingSnapshots**](BillingSnapshots.md) |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingZone = Initialize-PSOpenAPIToolsBillingZone  -ZoneName null `
 -ZoneId null `
 -ZoneUUID null `
 -ZoneCode null `
 -StartDate null `
 -EndDate null `
 -PriceUnit null `
 -ComputeServers null `
 -Instances null `
 -DiscoveredServers null `
 -LoadBalancers null `
 -VirtualImages null `
 -Snapshots null `
 -Price null `
 -Cost null
```

- Convert the resource to JSON
```powershell
$BillingZone | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

