# ApiScaleThresholdsScaleThreshold
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the scale threshold | 
**AutoUp** | **Boolean** | Auto Upscale | [optional] [default to $false]
**AutoDown** | **Boolean** | Auto Downscale | [optional] [default to $false]
**MinCount** | **Int32** | The minimum number of nodes to scale down to | [optional] 
**MaxCount** | **Int32** | The maximum number of nodes to scale up to | [optional] 
**CpuEnabled** | **Boolean** | Enable CPU Threshold | [optional] [default to $false]
**MinCpu** | **Double** | Min CPU (%) | [optional] [default to 0]
**MaxCpu** | **Double** | Max CPU (%) | [optional] [default to 0]
**MemoryEnabled** | **Boolean** | Enable Memory Threshold | [optional] [default to $false]
**MinMemory** | **Double** | Min Memory (%) | [optional] [default to 0]
**MaxMemory** | **Double** | Max Memory (%) | [optional] [default to 0]
**DiskEnabled** | **Boolean** | Enable Disk Threshold | [optional] [default to $false]
**MinDisk** | **Double** | Min Disk (%) | [optional] [default to 0]
**MaxDisk** | **Double** | Max Disk (%) | [optional] [default to 0]

## Examples

- Prepare the resource
```powershell
$ApiScaleThresholdsScaleThreshold = Initialize-PSOpenAPIToolsApiScaleThresholdsScaleThreshold  -Name Sample Threshold `
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
$ApiScaleThresholdsScaleThreshold | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

