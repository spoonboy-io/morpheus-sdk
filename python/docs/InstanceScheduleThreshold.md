# InstanceScheduleThreshold


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_up** | **bool** | Auto Upscale | [optional]  if omitted the server will use the default value of False
**auto_down** | **bool** | Auto Downscale | [optional]  if omitted the server will use the default value of False
**min_count** | **int** | The minimum number of nodes to scale down to | [optional] 
**max_count** | **int** | The maximum number of nodes to scale up to | [optional] 
**scale_increment** | **int** | The number of nodes that are added or removed at one time when scaling up or down | [optional]  if omitted the server will use the default value of 1
**cpu_enabled** | **bool** | Enable CPU Threshold | [optional]  if omitted the server will use the default value of False
**min_cpu** | **float** | Min CPU (%) | [optional]  if omitted the server will use the default value of 0
**max_cpu** | **float** | Max CPU (%) | [optional]  if omitted the server will use the default value of 0
**memory_enabled** | **bool** | Enable Memory Threshold | [optional]  if omitted the server will use the default value of False
**min_memory** | **float** | Min Memory (%) | [optional]  if omitted the server will use the default value of 0
**max_memory** | **float** | Max Memory (%) | [optional]  if omitted the server will use the default value of 0
**disk_enabled** | **bool** | Enable Disk Threshold | [optional]  if omitted the server will use the default value of False
**min_disk** | **float** | Min Disk (%) | [optional]  if omitted the server will use the default value of 0
**max_disk** | **float** | Max Disk (%) | [optional]  if omitted the server will use the default value of 0
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


