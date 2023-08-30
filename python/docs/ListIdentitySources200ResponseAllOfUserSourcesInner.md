# ListIdentitySources200ResponseAllOfUserSourcesInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**code** | **str** |  | [optional] 
**type** | **str** |  | [optional] 
**active** | **bool** |  | [optional] 
**deleted** | **bool** |  | [optional] 
**auto_sync_on_login** | **bool** |  | [optional] 
**external_login** | **bool** |  | [optional] 
**allow_custom_mappings** | **bool** |  | [optional] 
**manual_role_assignment** | **bool** |  | [optional] 
**account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**default_account_role** | [**IdentitySourcesLDAPConfigDefaultAccountRole**](IdentitySourcesLDAPConfigDefaultAccountRole.md) |  | [optional] 
**config** | [**IdentitySourcesCustomSSOConfigConfig**](IdentitySourcesCustomSSOConfigConfig.md) |  | [optional] 
**role_mappings** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**subdomain** | **str** |  | [optional] 
**login_url** | **str** |  | [optional] 
**provider_settings** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


