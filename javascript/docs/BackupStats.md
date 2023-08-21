# MorpheusApi.BackupStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**totalSize** | **Number** | Total size of all backups in bytes | [optional] 
**avgSize** | **Number** | Average size of each backup in bytes | [optional] 
**totalCompleted** | **Number** | Total completed count | [optional] 
**success** | **Number** | Successful backup count | [optional] 
**failed** | **Number** | Failed backup count | [optional] 
**successRate** | **Number** | Success rate 0-1 | [optional] 
**failRate** | **Number** | Failure rate 0-1 | [optional] 
**lastFiveResults** | **[String]** | List of the last 5 backup result statuses | [optional] 



## Enum: [LastFiveResultsEnum]


* `SUCCEEDED` (value: `"SUCCEEDED"`)

* `FAILED` (value: `"FAILED"`)

* `IN_PROGRESS` (value: `"IN_PROGRESS"`)

* `START_REQUESTED` (value: `"START_REQUESTED"`)




