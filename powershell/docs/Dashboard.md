# Dashboard
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**InstanceStats** | [**DashboardInstanceStats**](DashboardInstanceStats.md) |  | [optional] 
**Provisioning** | [**DashboardProvisioning**](DashboardProvisioning.md) |  | [optional] 
**Monitoring** | [**DashboardMonitoring**](DashboardMonitoring.md) |  | [optional] 
**Backups** | [**DashboardBackups**](DashboardBackups.md) |  | [optional] 
**Activity** | [**DashboardActivity[]**](DashboardActivity.md) |  | [optional] 
**LogStats** | [**DashboardLogStats**](DashboardLogStats.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Dashboard = Initialize-PSOpenAPIToolsDashboard  -Success null `
 -InstanceStats null `
 -Provisioning null `
 -Monitoring null `
 -Backups null `
 -Activity null `
 -LogStats null
```

- Convert the resource to JSON
```powershell
$Dashboard | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

