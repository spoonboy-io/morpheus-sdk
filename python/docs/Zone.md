# Zone


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**uuid** | **str** |  | [optional] 
**external_id** | **str, none_type** |  | [optional] 
**name** | **str** |  | [optional] 
**code** | **str** |  | [optional] 
**location** | **str, none_type** |  | [optional] 
**owner** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**account_id** | **int** |  | [optional] 
**account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**visibility** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**status** | **str** |  | [optional] 
**status_message** | **str, none_type** |  | [optional] 
**status_date** | **datetime, none_type** |  | [optional] 
**cost_status** | **str, none_type** |  | [optional] 
**cost_status_message** | **str, none_type** |  | [optional] 
**cost_status_date** | **datetime, none_type** |  | [optional] 
**cost_last_sync_duration** | **int, none_type** |  | [optional] 
**cost_last_sync** | **datetime, none_type** |  | [optional] 
**zone_type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**zone_type_id** | **int** |  | [optional] 
**guidance_mode** | **str, none_type** |  | [optional] 
**storage_mode** | **str** |  | [optional] 
**agent_mode** | **str** |  | [optional] 
**user_data_linux** | **str, none_type** |  | [optional] 
**user_data_windows** | **str, none_type** |  | [optional] 
**console_keymap** | **str, none_type** |  | [optional] 
**container_mode** | **str** |  | [optional] 
**costing_mode** | **str, none_type** |  | [optional] 
**service_version** | **str, none_type** |  | [optional] 
**security_mode** | **str** |  | [optional] 
**inventory_level** | **str** |  | [optional] 
**timezone** | **str, none_type** |  | [optional] 
**api_proxy** | **str, none_type** |  | [optional] 
**provisioning_proxy** | **str, none_type** |  | [optional] 
**network_domain** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**domain_name** | **str** |  | [optional] 
**region_code** | **str, none_type** |  | [optional] 
**auto_recover_power_state** | **bool** |  | [optional] 
**scale_priority** | **int** |  | [optional] 
**config** | [**ListClouds200ResponseAllOfZonesInner**](ListClouds200ResponseAllOfZonesInner.md) |  | [optional] 
**credential** | [**ZoneCredential**](ZoneCredential.md) |  | [optional] 
**image_path** | **str, none_type** | Logo image URL | [optional] 
**dark_image_path** | **str, none_type** | Dark logo image URL | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**last_sync** | **datetime, none_type** |  | [optional] 
**last_sync_duration** | **int, none_type** |  | [optional] 
**next_run_date** | **datetime, none_type** |  | [optional] 
**groups** | [**[ZoneGroupsInner]**](ZoneGroupsInner.md) |  | [optional] 
**security_server** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**network_server** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**stats** | [**ZoneStats**](ZoneStats.md) |  | [optional] 
**server_count** | **int** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


