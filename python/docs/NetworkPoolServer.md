# NetworkPoolServer


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Network Pool Server ID | [optional] 
**type** | [**NetworkPoolServerType**](NetworkPoolServerType.md) |  | [optional] 
**name** | **str** | Name | [optional] 
**enabled** | **bool** |  | [optional] 
**service_url** | **str, none_type** | Service URL | [optional] 
**service_host** | **str, none_type** | Service Host | [optional] 
**service_port** | **int, none_type** | Service Port | [optional] 
**service_mode** | **str, none_type** | Service Mode | [optional] 
**service_username** | **str, none_type** | Service Username | [optional] 
**service_password** | **str, none_type** | Service Password | [optional] 
**service_password_hash** | **str** |  | [optional] 
**service_throttle_rate** | **int, none_type** | Throttle Rate | [optional]  if omitted the server will use the default value of 0
**ignore_ssl** | **bool, none_type** | Disable SSL SNI Verification | [optional]  if omitted the server will use the default value of True
**status** | **str** | Status | [optional] 
**status_message** | **str, none_type** | Status Message | [optional] 
**status_date** | **datetime, none_type** |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Config object varies with pool server type. | [optional] 
**network_filter** | **str, none_type** | Network Filter | [optional] 
**zone_filter** | **str, none_type** | Zone Filter | [optional] 
**tenant_match** | **str, none_type** | Tenant Match | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**account** | [**NetworkPoolServerAccount**](NetworkPoolServerAccount.md) |  | [optional] 
**integration** | [**NetworkPoolServerIntegration**](NetworkPoolServerIntegration.md) |  | [optional] 
**pools** | [**[ListDeploys200ResponseAllOfAppDeploysInnerInstance]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**credential** | [**Creds2**](Creds2.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


