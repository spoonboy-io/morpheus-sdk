# GuidanceSettings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuAvgCutoffPower** | **Int32** | Power Shutdown Average CPU (%). Lower limit for average CPU usage | [optional] 
**CpuMaxCutoffPower** | **Int32** | Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage | [optional] 
**NetworkCutoffPower** | **Int32** | Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth | [optional] 
**CpuUpAvgStandardCutoffRightSize** | **Int32** | CPU Up-size Average CPU (%). Upper limit for CPU usage | [optional] 
**CpuUpMaxStandardCutoffRightSize** | **Int32** | CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage | [optional] 
**MemoryUpAvgStandardCutoffRightSize** | **Int32** | Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage | [optional] 
**MemoryDownAvgStandardCutoffRightSize** | **Int32** | Memory Down-size Maximum Free Memory (%). Upper limit for average free memory | [optional] 
**MemoryDownMaxStandardCutoffRightSize** | **Int32** | Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceSettings = Initialize-PSOpenAPIToolsGuidanceSettings  -CpuAvgCutoffPower 75 `
 -CpuMaxCutoffPower 500 `
 -NetworkCutoffPower 2000 `
 -CpuUpAvgStandardCutoffRightSize 50 `
 -CpuUpMaxStandardCutoffRightSize 99 `
 -MemoryUpAvgStandardCutoffRightSize 10 `
 -MemoryDownAvgStandardCutoffRightSize 60 `
 -MemoryDownMaxStandardCutoffRightSize 30
```

- Convert the resource to JSON
```powershell
$GuidanceSettings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

