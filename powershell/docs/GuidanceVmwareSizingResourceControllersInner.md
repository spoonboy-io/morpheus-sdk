# GuidanceVmwareSizingResourceControllersInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**MaxDevices** | **Int64** |  | [optional] 
**ReservedUnitNumber** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingResourceControllersInner = Initialize-PSOpenAPIToolsGuidanceVmwareSizingResourceControllersInner  -Id null `
 -Name null `
 -Type null `
 -MaxDevices null `
 -ReservedUnitNumber null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingResourceControllersInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

