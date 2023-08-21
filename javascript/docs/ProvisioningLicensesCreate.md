# MorpheusApi.ProvisioningLicensesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | 
**licenseType** | **String** | License Type - The license type code. | 
**licenseKey** | **String** | License Key - The license key, to be kept a secret. | 
**orgName** | **String** | Org Name - The Organization Name (if applicable) related to the license key | [optional] 
**fullName** | **String** | Full Name - The Full Name (if applicable) related to the license key | [optional] 
**licenseVersion** | **String** | License Version | [optional] 
**copies** | **Number** | Copies - The number of times the key can be used. | [optional] [default to 1]
**description** | **String** | Description | [optional] 
**virtualImages** | **[Number]** | Virtual Images - Array of Virtual Image IDs to associate with license. | [optional] 
**tenants** | **[Number]** | Tenants - Array of tenants that are allowed to use the key. | [optional] 


