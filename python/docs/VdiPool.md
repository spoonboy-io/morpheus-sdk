# VdiPool


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**description** | **str, none_type** |  | [optional] 
**min_idle** | **int** |  | [optional] 
**max_idle** | **int** |  | [optional] 
**initial_pool_size** | **int** |  | [optional] 
**max_pool_size** | **int** |  | [optional] 
**allocation_timeout_minutes** | **int** |  | [optional] 
**persistent_user** | **bool, none_type** |  | [optional] 
**recyclable** | **bool, none_type** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**auto_create_local_user_on_reservation** | **bool** |  | [optional] 
**allow_hypervisor_console** | **bool, none_type** |  | [optional] 
**allow_copy** | **bool, none_type** |  | [optional] 
**allow_printer** | **bool, none_type** |  | [optional] 
**allow_fileshare** | **bool, none_type** |  | [optional] 
**guest_console_jump_host** | **str, none_type** |  | [optional] 
**guest_console_jump_port** | **str, none_type** |  | [optional] 
**guest_console_jump_username** | **str, none_type** |  | [optional] 
**guest_console_jump_password** | **str, none_type** |  | [optional] 
**guest_console_jump_keypair** | **str, none_type** |  | [optional] 
**gateway** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**icon_path** | **str** |  | [optional] 
**logo** | **str** |  | [optional] 
**apps** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**owner** | [**VdiPoolOwner**](VdiPoolOwner.md) |  | [optional] 
**config** | [**VdiPoolConfig**](VdiPoolConfig.md) |  | [optional] 
**group** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**cloud** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**used_count** | **int** |  | [optional] 
**reserved_count** | **int** |  | [optional] 
**preparing_count** | **int** |  | [optional] 
**idle_count** | **int** |  | [optional] 
**status** | **str** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


