# # BackupTypeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**location_type** | **string** |  |
**name** | **string** | A name for the backup |
**instance_id** | **int** | The ID of the instance to backup |
**container_id** | **int** | The ID of the container to backup |
**backup_type** | **string** | The backup type code, options vary by the type of cloud and source |
**job_action** | **string** | Create a new backup job, clone an existing job or add the new backup to an existing job |
**job_id** | **int** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional]
**job_name** | **string** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional]
**job_schedule** | **int** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional]
**retention_count** | **int** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
