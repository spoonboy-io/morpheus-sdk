# # ProvisioningLicensesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name |
**license_type** | **string** | License Type - The license type code. |
**license_key** | **string** | License Key - The license key, to be kept a secret. |
**org_name** | **string** | Org Name - The Organization Name (if applicable) related to the license key | [optional]
**full_name** | **string** | Full Name - The Full Name (if applicable) related to the license key | [optional]
**license_version** | **string** | License Version | [optional]
**copies** | **int** | Copies - The number of times the key can be used. | [optional] [default to 1]
**description** | **string** | Description | [optional]
**virtual_images** | **int[]** | Virtual Images - Array of Virtual Image IDs to associate with license. | [optional]
**tenants** | **int[]** | Tenants - Array of tenants that are allowed to use the key. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
