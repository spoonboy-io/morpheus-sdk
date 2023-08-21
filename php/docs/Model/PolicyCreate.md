# # PolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the policy |
**description** | **string** | A description for the policy | [optional]
**policy_type** | [**\OpenAPI\Client\Model\PolicyCreatePolicyType**](PolicyCreatePolicyType.md) |  |
**config** | [**AnyOfObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObject**](AnyOfObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObject.md) | A map of config values. The expected values vary by policy type. See &#x60;Retrieves all Policy Types&#x60; endpoint for &#x60;fieldName&#x60;(s) of required options. |
**enabled** | **bool** | Set to false to disable | [optional] [default to true]
**ref_type** | **string** | Scope object type.  If none specified, will default to Global (null) | [optional] [default to 'null']
**ref_id** | **int** | Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc) | [optional]
**accounts** | **int[]** | Array of tenants to scope the policy to | [optional]
**each_user** | **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
