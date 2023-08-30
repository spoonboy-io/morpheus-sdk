# ImageBuilds


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**site** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**zone** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**name** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**boot_script** | [**ImageBuildsBootScript**](ImageBuildsBootScript.md) |  | [optional] 
**boot_command** | **str, none_type** |  | [optional] 
**preseed_script** | [**ImageBuildsBootScript**](ImageBuildsBootScript.md) |  | [optional] 
**scripts** | [**[ImageBuildsScriptsInner]**](ImageBuildsScriptsInner.md) |  | [optional] 
**ssh_username** | **str** |  | [optional] 
**ssh_password** | **str** |  | [optional] 
**storage_provider** | **str, none_type** |  | [optional] 
**build_output_name** | **str** |  | [optional] 
**conversion_formats** | **str, none_type** |  | [optional] 
**is_cloud_init** | **bool** |  | [optional] 
**vm_tools_installed** | **bool** |  | [optional] 
**keep_results** | **int** |  | [optional] 
**config** | [**ImageBuildsConfig**](ImageBuildsConfig.md) |  | [optional] 
**last_result** | [**ImageBuildsLastResult**](ImageBuildsLastResult.md) |  | [optional] 
**execution_count** | **int** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


