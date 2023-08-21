# # NetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Network Pool Server ID | [optional]
**type** | [**\OpenAPI\Client\Model\NetworkPoolServerType**](NetworkPoolServerType.md) |  | [optional]
**name** | **string** | Name | [optional]
**enabled** | **bool** |  | [optional]
**service_url** | **string** | Service URL | [optional]
**service_host** | **string** | Service Host | [optional]
**service_port** | **int** | Service Port | [optional]
**service_mode** | **string** | Service Mode | [optional]
**service_username** | **string** | Service Username | [optional]
**service_password** | **string** | Service Password | [optional]
**service_password_hash** | **string** |  | [optional]
**service_throttle_rate** | **int** | Throttle Rate | [optional] [default to 0]
**ignore_ssl** | **bool** | Disable SSL SNI Verification | [optional] [default to true]
**status** | **string** | Status | [optional]
**status_message** | **string** | Status Message | [optional]
**status_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**config** | **object** | Config object varies with pool server type. | [optional]
**network_filter** | **string** | Network Filter | [optional]
**zone_filter** | **string** | Zone Filter | [optional]
**tenant_match** | **string** | Tenant Match | [optional]
**date_created** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) |  | [optional]
**account** | [**\OpenAPI\Client\Model\NetworkPoolServerAccount**](NetworkPoolServerAccount.md) |  | [optional]
**integration** | [**\OpenAPI\Client\Model\NetworkPoolServerIntegration**](NetworkPoolServerIntegration.md) |  | [optional]
**pools** | [**\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]**](InlineResponse20040AppDeployInstance.md) |  | [optional]
**credential** | [**\OpenAPI\Client\Model\Creds2**](Creds2.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
