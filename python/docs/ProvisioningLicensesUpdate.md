# ProvisioningLicensesUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | [optional] 
**license_version** | **str** | License Version | [optional] 
**copies** | **int** | Copies - The number of times the key can be used. | [optional]  if omitted the server will use the default value of 1
**description** | **str** | Description | [optional] 
**virtual_images** | **[int]** | Virtual Images - Array of Virtual Image IDs to associate with license. | [optional] 
**tenants** | **[int]** | Tenants - Array of tenants that are allowed to use the key. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


