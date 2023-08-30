# ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**load_balancer** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer.md) |  | [optional] 
**name** | **str** |  | [optional] 
**category** | **str, none_type** |  | [optional] 
**visibility** | **str** |  | [optional] 
**description** | **str, none_type** |  | [optional] 
**internal_id** | **str** |  | [optional] 
**external_id** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**vip_sticky** | **str, none_type** |  | [optional] 
**vip_balance** | **str** |  | [optional] 
**allow_nat** | **str, none_type** |  | [optional] 
**allow_snat** | **str, none_type** |  | [optional] 
**vip_client_ip_mode** | **str, none_type** |  | [optional] 
**vip_server_ip_mode** | **str, none_type** |  | [optional] 
**min_active** | **int** |  | [optional] 
**min_in_service** | **str, none_type** |  | [optional] 
**min_up_monitor** | **str, none_type** |  | [optional] 
**min_up_action** | **str, none_type** |  | [optional] 
**max_queue_depth** | **str, none_type** |  | [optional] 
**max_queue_time** | **str, none_type** |  | [optional] 
**number_active** | **int** |  | [optional] 
**number_in_service** | **int** |  | [optional] 
**health_score** | **int** |  | [optional] 
**performance_score** | **int** |  | [optional] 
**health_penalty** | **int** |  | [optional] 
**security_penalty** | **int** |  | [optional] 
**error_penalty** | **int** |  | [optional] 
**down_action** | **str, none_type** |  | [optional] 
**ramp_time** | **str, none_type** |  | [optional] 
**port** | **str, none_type** |  | [optional] 
**port_type** | **str, none_type** |  | [optional] 
**status** | **str** |  | [optional] 
**nodes** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**monitors** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**members** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**created_by** | **str, none_type** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


