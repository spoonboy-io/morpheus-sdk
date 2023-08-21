# # CheckElastic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the check | [optional]
**description** | **string** | Optional description field | [optional]
**check_type** | [**\OpenAPI\Client\Model\CheckElasticCheckType**](CheckElasticCheckType.md) |  | [optional]
**check_interval** | **int** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) | [optional] [default to 300]
**in_uptime** | **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**active** | **bool** | Used to determine if check should be scheduled to execute | [optional] [default to true]
**severity** | **string** | Severity level threshold for sending notifications. | [optional] [default to 'critical']
**config** | [**CheckElasticsearchConfig**](CheckElasticsearchConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
