# InstanceScheduleUpdateThreshold


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_threshold_id** | **int** | Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. | [optional] 
**auto_up** | **bool** | Auto Upscale | [optional] 
**auto_down** | **bool** | Auto Downscale | [optional] 
**min_count** | **int** | The minimum number of nodes to scale down to | [optional] 
**max_count** | **int** | The maximum number of nodes to scale up to | [optional] 
**cpu_enabled** | **bool** | Enable CPU Threshold | [optional] 
**min_cpu** | **float** | Min CPU (%) | [optional] 
**max_cpu** | **float** | Max CPU (%) | [optional] 
**memory_enabled** | **bool** | Enable Memory Threshold | [optional] 
**min_memory** | **float** | Min Memory (%) | [optional] 
**max_memory** | **float** | Max Memory (%) | [optional] 
**disk_enabled** | **bool** | Enable Disk Threshold | [optional] 
**min_disk** | **float** | Min Disk (%) | [optional] 
**max_disk** | **float** | Max Disk (%) | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


