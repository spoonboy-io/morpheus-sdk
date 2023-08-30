# ProvisioningSettingsUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**allow_zone_selection** | **bool** | Use this to enable / disable allowing cloud selection | [optional] 
**allow_server_selection** | **bool** | Use this to enable / disable allowing host selection | [optional] 
**require_environments** | **bool** | Use this to enable / disable requiring environment selection | [optional] 
**show_pricing** | **bool** | Use this to enable / disable showing pricing | [optional] 
**hide_datastore_stats** | **bool** | Use this to enable / disable hiding datastore stats | [optional] 
**cross_tenant_naming_policies** | **bool** | Use this to enable / disable cross-tenant naming policies | [optional] 
**reuse_sequence** | **bool** | Use this to enable / disable reusing naming sequence numbers | [optional] 
**cloud_init_username** | **str** | Cloud-init username | [optional] 
**cloud_init_password** | **str** | Cloud-init password | [optional] 
**cloud_init_key_pair** | [**ProvisioningSettingsUpdateCloudInitKeyPair**](ProvisioningSettingsUpdateCloudInitKeyPair.md) |  | [optional] 
**deploy_storage_provider** | [**ProvisioningSettingsUpdateDeployStorageProvider**](ProvisioningSettingsUpdateDeployStorageProvider.md) |  | [optional] 
**windows_password** | **str** | Windows administrator password | [optional] 
**pxe_root_password** | **str** | PXE Boot default root password | [optional] 
**default_template_type** | [**ProvisioningSettingsUpdateDefaultTemplateType**](ProvisioningSettingsUpdateDefaultTemplateType.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


