# # IntegrationSNOWConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name, a unique identifier for the integration |
**type** | **string** | Integration Type Code |
**enabled** | **bool** | Set &#x60;true&#x60; to enable integration | [optional]
**refresh** | **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type. | [optional] [default to true]
**service_url** | **string** | ServiceNow Host |
**service_username** | **string** | ServiceNow Username |
**service_password** | **string** | ServiceNow Password |
**config** | [**\OpenAPI\Client\Model\IntegrationSNOWConfigIntegrationConfig**](IntegrationSNOWConfigIntegrationConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
