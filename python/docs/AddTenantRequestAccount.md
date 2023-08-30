# AddTenantRequestAccount

Payload for creating a new tenant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | 
**description** | **str, none_type** | Description | [optional] 
**role** | [**AddTenantRequestAccountRole**](AddTenantRequestAccountRole.md) |  | [optional] 
**subdomain** | **str, none_type** | The subdomain. This will be part of the login URL and username for sub tenant users. | [optional] 
**currency** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


