# BackupTypeProvider
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationType** | **String** |  | 
**Name** | **String** | A name for the backup | 
**SourceProviderId** | **Int64** | Source Object Store. The ID of the storage provider (bucket) to be backed up. | [optional] 
**StorageProviderId** | **Int64** | Target Object Store. The ID of the storage provider (bucket) the backup will be copied to. | [optional] 
**BackupType** | **String** | The backup type code, options vary by the type of cloud and source | 
**JobAction** | **String** | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**JobId** | **Int64** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional] 
**JobName** | **String** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**JobSchedule** | **Int64** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**RetentionCount** | **Int64** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupTypeProvider = Initialize-PSOpenAPIToolsBackupTypeProvider  -LocationType null `
 -Name null `
 -SourceProviderId null `
 -StorageProviderId null `
 -BackupType null `
 -JobAction null `
 -JobId null `
 -JobName null `
 -JobSchedule null `
 -RetentionCount null
```

- Convert the resource to JSON
```powershell
$BackupTypeProvider | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

