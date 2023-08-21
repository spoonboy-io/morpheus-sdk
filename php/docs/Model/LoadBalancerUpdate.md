# # LoadBalancerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name | [optional]
**description** | **string** | Description | [optional]
**enabled** | **bool** | Activate (true) or disable (false) | [optional]
**config** | **object** | Configuration object with parameters that vary by load balancer type. | [optional]
**visibility** | **string** | private or public | [optional] [default to 'public']
**tenants** | [**\OpenAPI\Client\Model\LoadBalancerCreateTenants[]**](LoadBalancerCreateTenants.md) | Array of tenant account ids that are allowed access | [optional]
**resource_permission** | [**\OpenAPI\Client\Model\LoadBalancerCreateResourcePermission**](LoadBalancerCreateResourcePermission.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
