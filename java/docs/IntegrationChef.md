

# IntegrationChef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**name** | **String** |  |  [optional]
**enabled** | **Boolean** |  |  [optional]
**type** | [**TypeEnum**](#TypeEnum) |  |  [optional]
**integrationType** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  |  [optional]
**url** | **String** |  |  [optional]
**isPlugin** | **Boolean** |  |  [optional]
**config** | [**IntegrationChefConfig**](IntegrationChefConfig.md) |  |  [optional]
**status** | **String** |  |  [optional]
**statusDate** | **OffsetDateTime** |  |  [optional]
**statusMessage** | **String** |  |  [optional]
**lastSync** | **String** |  |  [optional]
**lastSyncDuration** | **String** |  |  [optional]
**credential** | [**Creds**](Creds.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
CHEF | &quot;chef&quot;



