# # UserSourceCreateUserSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the Identity Source |
**type** | **string** | IDM type code |
**description** | **string** | description | [optional]
**default_account_role** | [**\OpenAPI\Client\Model\UserSourceCreateUserSourceDefaultAccountRole**](UserSourceCreateUserSourceDefaultAccountRole.md) |  |
**role_mappings** | [**OneOfArrayMap**](OneOfArrayMap.md) |  | [optional]
**role_mapping_names** | **map[string,string]** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;. | [optional]
**allow_custom_mappings** | **bool** | Enable Role Mapping Permission | [optional]
**manual_role_assignment** | **bool** | Manual Role Assignment | [optional]
**config** | [**OneOfUserSourceCreateLDAPUserSourceCreateJumpCloudUserSourceCreateActiveDirectoryUserSourceCreateOktaUserSourceCreateOneLoginUserSourceCreateSamlUserSourceCreateCustomExternalUserSourceCreateCustomApi**](OneOfUserSourceCreateLDAPUserSourceCreateJumpCloudUserSourceCreateActiveDirectoryUserSourceCreateOktaUserSourceCreateOneLoginUserSourceCreateSamlUserSourceCreateCustomExternalUserSourceCreateCustomApi.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
