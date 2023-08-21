

# CheckElastic

Elasticsearch check is capable of connecting to your Elasticsearch, cluster or node, verifying its health. In addition, Morpheus will also pull statistical information such as: document size, capacity, and cpu usage. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check |  [optional]
**description** | **String** | Optional description field |  [optional]
**checkType** | [**CheckElasticCheckType**](CheckElasticCheckType.md) |  |  [optional]
**checkInterval** | **Integer** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**active** | **Boolean** | Used to determine if check should be scheduled to execute |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level threshold for sending notifications. |  [optional]
**config** | [**CheckElasticsearchConfig**](CheckElasticsearchConfig.md) |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



