# # PolicyGroupUpdatePolicyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **string** | The policy type | [optional]
**config** | [**OneOfObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObject**](OneOfObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObjectObject.md) | A map of config values. The expected values vary by policyType. | [optional]
**enabled** | **bool** | Set to false to disable | [optional] [default to true]
**ref_type** | **string** | Scope object type | [optional]
**ref_id** | **int** | Scope object ID (&#x60;group&#x60;) | [optional]
**accounts** | **int[]** | Array of tenants to scope the policy to | [optional]
**each_user** | **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
