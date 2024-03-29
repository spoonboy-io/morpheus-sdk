# NetworkPoolServerUpdatePhpIpam


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**enabled** | **bool** | Can be used to enable / disable the network pool server. | [optional]  if omitted the server will use the default value of True
**service_url** | **str, none_type** | URL | [optional] 
**service_username** | **str, none_type** | Username | [optional] 
**service_password** | **str, none_type** | Password | [optional] 
**service_throttle_rate** | **int, none_type** | Throttle Rate | [optional]  if omitted the server will use the default value of 0
**ignore_ssl** | **bool** | Disable SSL SNI Verification | [optional] 
**network_filter** | **str, none_type** | Network Filter | [optional] 
**config** | [**NetworkPoolServerUpdatePhpIpamConfig**](NetworkPoolServerUpdatePhpIpamConfig.md) |  | [optional] 
**credential** | [**NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


