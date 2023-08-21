# # NetworkPoolServerUpdateBluecat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name | [optional]
**enabled** | **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**service_url** | **string** | URL | [optional]
**service_username** | **string** | Username | [optional]
**service_password** | **string** | Password | [optional]
**service_throttle_rate** | **int** | Throttle Rate | [optional] [default to 0]
**ignore_ssl** | **bool** | Disable SSL SNI Verification | [optional]
**network_filter** | **string** | Network Filter | [optional]
**config** | [**\OpenAPI\Client\Model\NetworkPoolServerCreateBluecatConfig**](NetworkPoolServerCreateBluecatConfig.md) |  | [optional]
**credential** | [**\OpenAPI\Client\Model\NetworkPoolServerCreateBluecatCredential**](NetworkPoolServerCreateBluecatCredential.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
