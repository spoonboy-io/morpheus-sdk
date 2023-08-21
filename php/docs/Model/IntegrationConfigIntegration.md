# # IntegrationConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name, a unique identifier for the integration |
**type** | **string** | Integration Type Code |
**enabled** | **bool** | Set &#x60;true&#x60; to enable integration | [optional]
**refresh** | **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type. | [optional] [default to true]
**credential** | [**OneOfObjectObject**](OneOfObjectObject.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
