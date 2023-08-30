# IntegrationSaltMaster


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**type** | **str** |  | [optional]  if omitted the server will use the default value of "saltMaster"
**integration_type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**url** | **str** |  | [optional] 
**port** | **str** |  | [optional] 
**username** | **str** |  | [optional] 
**password** | **str** |  | [optional] 
**password_hash** | **str** |  | [optional] 
**path** | **str** |  | [optional] 
**version** | **str** |  | [optional] 
**windows_version** | **str** |  | [optional] 
**repo_url** | **str** |  | [optional] 
**service_mode** | **str** |  | [optional] 
**is_plugin** | **bool** |  | [optional] 
**config** | [**IntegrationSaltMasterConfig**](IntegrationSaltMasterConfig.md) |  | [optional] 
**status** | **str** |  | [optional] 
**status_date** | **datetime** |  | [optional] 
**status_message** | **str** |  | [optional] 
**last_sync** | **str, none_type** |  | [optional] 
**last_sync_duration** | **str, none_type** |  | [optional] 
**credential** | [**Creds**](Creds.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


