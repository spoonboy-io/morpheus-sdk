# MorpheusApi.InstanceScheduleUpdateThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sourceThresholdId** | **Number** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. | [optional] 
**autoUp** | **Boolean** | Auto Upscale | [optional] 
**autoDown** | **Boolean** | Auto Downscale | [optional] 
**minCount** | **Number** | The minimum number of nodes to scale down to | [optional] 
**maxCount** | **Number** | The maximum number of nodes to scale up to | [optional] 
**cpuEnabled** | **Boolean** | Enable CPU Threshold | [optional] 
**minCpu** | **Number** | Min CPU (%) | [optional] 
**maxCpu** | **Number** | Max CPU (%) | [optional] 
**memoryEnabled** | **Boolean** | Enable Memory Threshold | [optional] 
**minMemory** | **Number** | Min Memory (%) | [optional] 
**maxMemory** | **Number** | Max Memory (%) | [optional] 
**diskEnabled** | **Boolean** | Enable Disk Threshold | [optional] 
**minDisk** | **Number** | Min Disk (%) | [optional] 
**maxDisk** | **Number** | Max Disk (%) | [optional] 


