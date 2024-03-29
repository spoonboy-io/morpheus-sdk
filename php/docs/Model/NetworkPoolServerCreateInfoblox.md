# # NetworkPoolServerCreateInfoblox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** | Type Code (Infoblox) |
**name** | **string** | Name |
**enabled** | **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**service_url** | **string** | URL |
**service_username** | **string** | Username | [optional]
**service_password** | **string** | Password | [optional]
**service_throttle_rate** | **int** | Throttle Rate | [optional] [default to 0]
**ignore_ssl** | **bool** | Disable SSL SNI Verification | [optional] [default to true]
**network_filter** | **string** | Network Filter | [optional]
**zone_filter** | **string** | Zone Filter | [optional]
**tenant_match** | **string** | Tenant Match | [optional]
**service_mode** | **string** | IP Mode | [optional] [default to 'static']
**config** | [**\OpenAPI\Client\Model\NetworkPoolServerCreateInfobloxConfig**](NetworkPoolServerCreateInfobloxConfig.md) |  | [optional]
**credential** | [**\OpenAPI\Client\Model\NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
