# MorpheusApi.InstanceScheduleCreateThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sourceThresholdId** | **Number** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. | [optional] 
**autoUp** | **Boolean** | Auto Upscale | [optional] [default to false]
**autoDown** | **Boolean** | Auto Downscale | [optional] [default to false]
**minCount** | **Number** | The minimum number of nodes to scale down to | [optional] 
**maxCount** | **Number** | The maximum number of nodes to scale up to | [optional] 
**cpuEnabled** | **Boolean** | Enable CPU Threshold | [optional] [default to false]
**minCpu** | **Number** | Min CPU (%) | [optional] [default to 0]
**maxCpu** | **Number** | Max CPU (%) | [optional] [default to 0]
**memoryEnabled** | **Boolean** | Enable Memory Threshold | [optional] [default to false]
**minMemory** | **Number** | Min Memory (%) | [optional] [default to 0]
**maxMemory** | **Number** | Max Memory (%) | [optional] [default to 0]
**diskEnabled** | **Boolean** | Enable Disk Threshold | [optional] [default to false]
**minDisk** | **Number** | Min Disk (%) | [optional] [default to 0]
**maxDisk** | **Number** | Max Disk (%) | [optional] [default to 0]


