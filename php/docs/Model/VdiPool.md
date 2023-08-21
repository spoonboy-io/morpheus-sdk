# # VdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional]
**name** | **string** |  | [optional]
**description** | **string** |  | [optional]
**min_idle** | **int** |  | [optional]
**max_idle** | **int** |  | [optional]
**initial_pool_size** | **int** |  | [optional]
**max_pool_size** | **int** |  | [optional]
**allocation_timeout_minutes** | **int** |  | [optional]
**persistent_user** | **bool** |  | [optional]
**recyclable** | **bool** |  | [optional]
**enabled** | **bool** |  | [optional]
**auto_create_local_user_on_reservation** | **bool** |  | [optional]
**allow_hypervisor_console** | **bool** |  | [optional]
**allow_copy** | **bool** |  | [optional]
**allow_printer** | **bool** |  | [optional]
**allow_fileshare** | **bool** |  | [optional]
**guest_console_jump_host** | **string** |  | [optional]
**guest_console_jump_port** | **string** |  | [optional]
**guest_console_jump_username** | **string** |  | [optional]
**guest_console_jump_password** | **string** |  | [optional]
**guest_console_jump_keypair** | **string** |  | [optional]
**gateway** | [**\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional]
**icon_path** | **string** |  | [optional]
**logo** | **string** |  | [optional]
**apps** | [**\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]**](InlineResponse20040AppDeployInstance.md) |  | [optional]
**owner** | [**\OpenAPI\Client\Model\VdiPoolOwner**](VdiPoolOwner.md) |  | [optional]
**config** | [**\OpenAPI\Client\Model\VdiPoolConfig**](VdiPoolConfig.md) |  | [optional]
**group** | [**\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional]
**cloud** | [**\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional]
**used_count** | **int** |  | [optional]
**reserved_count** | **int** |  | [optional]
**preparing_count** | **int** |  | [optional]
**idle_count** | **int** |  | [optional]
**status** | **string** |  | [optional]
**date_created** | [**\DateTime**](\DateTime.md) |  | [optional]
**last_updated** | [**\DateTime**](\DateTime.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
