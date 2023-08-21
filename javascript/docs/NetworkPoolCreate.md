# MorpheusApi.NetworkPoolCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**type** | [**NetworkPoolCreateType**](NetworkPoolCreateType.md) |  | [optional] 
**ipRanges** | [**[NetworkPoolCreateIpRanges]**](NetworkPoolCreateIpRanges.md) | Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6. | [optional] 
**config** | **Object** | Configuration object with parameters that vary by pool type. | [optional] 


