# NetworkRouterRouteCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source** | **str** | Source address or range | 
**destination** | **str** | Destination address or range | 
**network_mtu** | **str** | MTU | 
**name** | **str** | Route name | [optional] 
**description** | **str** | Route description | [optional] 
**enabled** | **bool** | Can be used to enable / disable the route (true, false). Default is off | [optional]  if omitted the server will use the default value of False
**default_route** | **bool** | Can be used to set as default route (true, false). Default is off | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


