# # InstanceScheduleCreateThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_threshold_id** | **int** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. | [optional]
**auto_up** | **bool** | Auto Upscale | [optional] [default to false]
**auto_down** | **bool** | Auto Downscale | [optional] [default to false]
**min_count** | **int** | The minimum number of nodes to scale down to | [optional]
**max_count** | **int** | The maximum number of nodes to scale up to | [optional]
**cpu_enabled** | **bool** | Enable CPU Threshold | [optional] [default to false]
**min_cpu** | **double** | Min CPU (%) | [optional] [default to 0]
**max_cpu** | **double** | Max CPU (%) | [optional] [default to 0]
**memory_enabled** | **bool** | Enable Memory Threshold | [optional] [default to false]
**min_memory** | **double** | Min Memory (%) | [optional] [default to 0]
**max_memory** | **double** | Max Memory (%) | [optional] [default to 0]
**disk_enabled** | **bool** | Enable Disk Threshold | [optional] [default to false]
**min_disk** | **double** | Min Disk (%) | [optional] [default to 0]
**max_disk** | **double** | Max Disk (%) | [optional] [default to 0]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
