# # ApiBudgetsBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** |  |
**description** | **string** |  | [optional]
**scope** | **string** |  | [optional] [default to 'account']
**period** | **string** |  | [optional] [default to 'year']
**year** | **int** |  | [optional]
**start_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**end_date** | [**\DateTime**](\DateTime.md) |  | [optional]
**interval** | **string** |  | [optional] [default to 'year']
**scope_tenant_id** | **int** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;tenant | [optional]
**scope_group_id** | **int** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;group | [optional]
**scope_cloud_id** | **int** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;cloud | [optional]
**scope_user_id** | **int** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;user | [optional]
**costs** | **int[]** |  | [optional]
**enabled** | **bool** |  | [optional] [default to true]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
