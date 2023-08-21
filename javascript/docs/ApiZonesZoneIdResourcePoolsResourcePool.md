# MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name of Resource Pool | 
**defaultPool** | **Boolean** | Set as the Default Pool | [optional] [default to false]
**defaultImage** | **Boolean** | Set as the Default Image Target | [optional] [default to false]
**active** | **Boolean** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] [default to true]
**visibility** | **String** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to &#39;private&#39;]
**displayName** | **String** | Optional Display Name (VMware only) | [optional] 
**inventory** | **Boolean** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] [default to true]
**config** | [**AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig**](AnyOfAwsResourcePoolConfigCloudFoundryResourcePoolConfig.md) |  | 
**tenantPermissions** | [**[ApiZonesZoneIdFoldersIdFolderTenantPermissions]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**resourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 



## Enum: VisibilityEnum


* `public` (value: `"public"`)

* `private` (value: `"private"`)




