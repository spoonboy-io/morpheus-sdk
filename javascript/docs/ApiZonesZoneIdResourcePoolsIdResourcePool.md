# MorpheusApi.ApiZonesZoneIdResourcePoolsIdResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | **Boolean** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**visibility** | **String** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to &#39;private&#39;]
**displayName** | **String** | Optional Display Name (VMware only) | [optional] 
**inventory** | **Boolean** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] 
**tenantPermissions** | [**[ApiZonesZoneIdFoldersIdFolderTenantPermissions]**](ApiZonesZoneIdFoldersIdFolderTenantPermissions.md) |  | [optional] 
**resourcePermissions** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions.md) |  | [optional] 



## Enum: VisibilityEnum


* `public` (value: `"public"`)

* `private` (value: `"private"`)




