

# IntegrationSNOWConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | [**TypeEnum**](#TypeEnum) | Integration Type Code | 
**enabled** | **Boolean** | Set &#x60;true&#x60; to enable integration |  [optional]
**refresh** | **Boolean** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  |  [optional]
**serviceUrl** | **String** | ServiceNow Host | 
**serviceUsername** | **String** | ServiceNow Username | 
**servicePassword** | **String** | ServiceNow Password | 
**config** | [**IntegrationSNOWConfigIntegrationConfig**](IntegrationSNOWConfigIntegrationConfig.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
SERVICENOW | &quot;serviceNow&quot;



