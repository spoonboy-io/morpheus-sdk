

# BackupStats

Backup Result Statistics
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**totalSize** | **Long** | Total size of all backups in bytes |  [optional]
**avgSize** | **Long** | Average size of each backup in bytes |  [optional]
**totalCompleted** | **Long** | Total completed count |  [optional]
**success** | **Long** | Successful backup count |  [optional]
**failed** | **Long** | Failed backup count |  [optional]
**successRate** | **Double** | Success rate 0-1 |  [optional]
**failRate** | **Double** | Failure rate 0-1 |  [optional]
**lastFiveResults** | [**List&lt;LastFiveResultsEnum&gt;**](#List&lt;LastFiveResultsEnum&gt;) | List of the last 5 backup result statuses |  [optional]



## Enum: List&lt;LastFiveResultsEnum&gt;

Name | Value
---- | -----
SUCCEEDED | &quot;SUCCEEDED&quot;
FAILED | &quot;FAILED&quot;
IN_PROGRESS | &quot;IN_PROGRESS&quot;
START_REQUESTED | &quot;START_REQUESTED&quot;



