# # ApiServicePlansServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Service plan name |
**code** | **string** | Service plan code, must be unique |
**description** | **string** | Service plan description | [optional]
**editable** | **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**max_storage** | **float** | Max storage size in bytes |
**max_memory** | **float** | Max memory size in bytes |
**max_cores** | **float** | Max cores | [optional]
**max_disks** | **float** | Max disks allowed | [optional]
**provision_type** | [**\OpenAPI\Client\Model\ApiServicePlansServicePlanProvisionType[]**](ApiServicePlansServicePlanProvisionType.md) |  |
**custom_cores** | **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**custom_max_storage** | **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**custom_max_data_storage** | **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**custom_max_memory** | **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**add_volumes** | **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**sort_order** | **float** | Sort order | [optional]
**price_sets** | **int[]** | List of price sets to include in service plan | [optional]
**config** | [**\OpenAPI\Client\Model\ApiServicePlansServicePlanConfig**](ApiServicePlansServicePlanConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
