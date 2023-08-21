# # NetworkPoolCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name | [optional]
**type** | [**\OpenAPI\Client\Model\NetworkPoolCreateType**](NetworkPoolCreateType.md) |  | [optional]
**ip_ranges** | [**\OpenAPI\Client\Model\NetworkPoolCreateIpRanges[]**](NetworkPoolCreateIpRanges.md) | Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6. | [optional]
**config** | **object** | Configuration object with parameters that vary by pool type. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
