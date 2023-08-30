# NetworkPoolCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**type** | [**NetworkPoolCreateType**](NetworkPoolCreateType.md) |  | [optional] 
**ip_ranges** | [**[NetworkPoolCreateIpRangesInner]**](NetworkPoolCreateIpRangesInner.md) | Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6. | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration object with parameters that vary by pool type. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


