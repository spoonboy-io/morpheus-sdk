# ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to $true]
**Sites** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites[]**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | **Boolean** | Pass true to allow access all plans | [optional] [default to $true]
**Plans** | [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites[]**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites.md) | Array of plans that are allowed access | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions = Initialize-PSOpenAPIToolsApiZonesZoneIdDataStoresIdDatastoreResourcePermissions  -All null `
 -Sites null `
 -AllPlans null `
 -Plans null
```

- Convert the resource to JSON
```powershell
$ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

