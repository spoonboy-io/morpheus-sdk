# IntegrationDockerRepo


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**type** | **str** |  | [optional]  if omitted the server will use the default value of "docker.registry"
**integration_type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**url** | **str** |  | [optional] 
**username** | **str** |  | [optional] 
**password** | **str** |  | [optional] 
**password_hash** | **str** |  | [optional] 
**is_plugin** | **bool** |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}, none_type** |  | [optional] 
**status** | **str, none_type** |  | [optional] 
**status_date** | **datetime, none_type** |  | [optional] 
**status_message** | **str, none_type** |  | [optional] 
**last_sync** | **str, none_type** |  | [optional] 
**last_sync_duration** | **str, none_type** |  | [optional] 
**credential** | [**Creds**](Creds.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


