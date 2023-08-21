

# InstanceScheduleUpdateThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sourceThresholdId** | **Long** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. |  [optional]
**autoUp** | **Boolean** | Auto Upscale |  [optional]
**autoDown** | **Boolean** | Auto Downscale |  [optional]
**minCount** | **Integer** | The minimum number of nodes to scale down to |  [optional]
**maxCount** | **Integer** | The maximum number of nodes to scale up to |  [optional]
**cpuEnabled** | **Boolean** | Enable CPU Threshold |  [optional]
**minCpu** | **Double** | Min CPU (%) |  [optional]
**maxCpu** | **Double** | Max CPU (%) |  [optional]
**memoryEnabled** | **Boolean** | Enable Memory Threshold |  [optional]
**minMemory** | **Double** | Min Memory (%) |  [optional]
**maxMemory** | **Double** | Max Memory (%) |  [optional]
**diskEnabled** | **Boolean** | Enable Disk Threshold |  [optional]
**minDisk** | **Double** | Min Disk (%) |  [optional]
**maxDisk** | **Double** | Max Disk (%) |  [optional]



