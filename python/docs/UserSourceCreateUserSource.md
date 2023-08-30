# UserSourceCreateUserSource


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the Identity Source | 
**type** | **str** | IDM type code | 
**default_account_role** | [**UserSourceCreateUserSourceDefaultAccountRole**](UserSourceCreateUserSourceDefaultAccountRole.md) |  | 
**description** | **str** | description | [optional] 
**role_mappings** | [**UserSourceCreateUserSourceRoleMappings**](UserSourceCreateUserSourceRoleMappings.md) |  | [optional] 
**role_mapping_names** | **{str: (str,)}** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;.  | [optional] 
**allow_custom_mappings** | **bool** | Enable Role Mapping Permission | [optional] 
**manual_role_assignment** | **bool** | Manual Role Assignment | [optional] 
**config** | [**UserSourceCreateUserSourceConfig**](UserSourceCreateUserSourceConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


