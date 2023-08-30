# CatalogItemTypeInstanceScribeConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resource_pool_id** | **str** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**availability_options** | **str** | Availability Options | [optional] 
**availability_set** | **str** | Availability Set | [optional] 
**availability_zone** | **int** | Availability Zone | [optional] 
**azurefloating_ip** | **str** | Assign Public IP | [optional] 
**boot_diagnostics** | **str** | Boot Diagnostics | [optional] 
**os_guest_diagnostics** | **str** | OS Guest Diagnostics | [optional] 
**diagnostics_storage_account** | **str** | Diagnostics Storage Account | [optional] 
**create_user** | **bool** | Create User | [optional]  if omitted the server will use the default value of True
**no_agent** | **bool, none_type** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional]  if omitted the server will use the default value of False
**host_id** | **str** | Specific host to deploy to if so desired. | [optional] 
**smbios_asset_tag** | **str** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**nested_virtualization** | **str** | Enable Nested Virtualization | [optional]  if omitted the server will use the default value of "off"
**vmware_folder_id** | **str** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 
**google_zone_id** | **int** | External ID of the Google zone to use for instance. | [optional] 
**external_ip_type** | **int** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**network_tags** | **str** | Network Tags | [optional] 
**service_account** | **str** | Service Account | [optional] 
**access_scope** | **str** | Access Scope | [optional] 
**is_ec2** | **str** | Amazon Cloud Type | [optional]  if omitted the server will use the default value of "false"
**availability_id** | **str** | Amazon Zone | [optional] 
**security_id** | **str** | Security Group | [optional] 
**public_ip_type** | **str** | Public IP | [optional] 
**instance_profile** | **str** | IAM Profile | [optional] 
**kms_key_id** | **str** | KMS Key ID | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


