# ApiScaleThresholdsIdScaleThreshold
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the scale threshold | [optional] 
**AutoUp** | **Boolean** | Auto Upscale | [optional] [default to $false]
**AutoDown** | **Boolean** | Auto Downscale | [optional] [default to $false]
**MinCount** | **Int32** | The minimum number of nodes to scale down to | [optional] 
**MaxCount** | **Int32** | The maximum number of nodes to scale up to | [optional] 
**CpuEnabled** | **Boolean** | Enable CPU Threshold | [optional] [default to $false]
**MinCpu** | **Decimal** | Min CPU (%) | [optional] [default to 0]
**MaxCpu** | **Decimal** | Max CPU (%) | [optional] [default to 0]
**MemoryEnabled** | **Boolean** | Enable Memory Threshold | [optional] [default to $false]
**MinMemory** | **Decimal** | Min Memory (%) | [optional] [default to 0]
**MaxMemory** | **Decimal** | Max Memory (%) | [optional] [default to 0]
**DiskEnabled** | **Boolean** | Enable Disk Threshold | [optional] [default to $false]
**MinDisk** | **Decimal** | Min Disk (%) | [optional] [default to 0]
**MaxDisk** | **Decimal** | Max Disk (%) | [optional] [default to 0]

## Examples

- Prepare the resource
```powershell
$ApiScaleThresholdsIdScaleThreshold = Initialize-PSOpenAPIToolsApiScaleThresholdsIdScaleThreshold  -Name Sample Threshold `
 -AutoUp null `
 -AutoDown null `
 -MinCount 1 `
 -MaxCount 3 `
 -CpuEnabled null `
 -MinCpu null `
 -MaxCpu null `
 -MemoryEnabled null `
 -MinMemory null `
 -MaxMemory null `
 -DiskEnabled null `
 -MinDisk null `
 -MaxDisk null
```

- Convert the resource to JSON
```powershell
$ApiScaleThresholdsIdScaleThreshold | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

