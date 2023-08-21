# # IntegrationSaltMasterConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name, a unique identifier for the integration |
**type** | **string** | Integration Type Code |
**service_mode** | **string** | Topology | [optional] [default to 'single']
**service_url** | **string** | Salt Master |
**secondary** | **string** | Salt Syndic | [optional]
**service_port** | **int** | SSH Port | [optional] [default to 22]
**service_username** | **string** | Username |
**service_password** | **string** | Password | [optional]
**service_key** | **string** | Master Key Pair | [optional]
**auth_key** | **string** | Signing Key | [optional]
**service_path** | **string** | Working Directory | [optional]
**service_version** | **string** | Salt Version | [optional]
**service_windows_version** | **string** | Salt Version (Windows) | [optional]
**repo_url** | **string** | Repo URL | [optional]
**service_config** | **string** | Minion Config | [optional]
**service_command** | **string** | Post Provision Commands | [optional]
**config** | [**\OpenAPI\Client\Model\IntegrationSaltMasterConfigIntegrationConfig**](IntegrationSaltMasterConfigIntegrationConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
