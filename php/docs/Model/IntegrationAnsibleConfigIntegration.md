# # IntegrationAnsibleConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name, a unique identifier for the integration |
**type** | **string** | Integration Type Code |
**enabled** | **bool** | Set &#x60;true&#x60; to enable integration | [optional]
**refresh** | **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type. | [optional] [default to true]
**service_url** | **string** | Ansible Git URL |
**service_username** | **string** | Git Username | [optional]
**service_password** | **string** | Git Password or Token depending on the Git host | [optional]
**service_token** | **string** | Git Token | [optional]
**service_key** | **int** | Keypair ID | [optional]
**config** | [**\OpenAPI\Client\Model\IntegrationAnsibleConfigIntegrationConfig**](IntegrationAnsibleConfigIntegrationConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
