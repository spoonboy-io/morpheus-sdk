

# InlineResponse200108NetworkFloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**externalId** | **String** |  |  [optional]
**cloud** | [**InlineResponse200108NetworkFloatingIpCloud**](InlineResponse200108NetworkFloatingIpCloud.md) |  |  [optional]
**server** | [**InlineResponse200108NetworkFloatingIpServer**](InlineResponse200108NetworkFloatingIpServer.md) |  |  [optional]
**ipStatus** | [**IpStatusEnum**](#IpStatusEnum) |  |  [optional]
**ipAddress** | **String** | IP Address |  [optional]
**ipRange** | **String** |  |  [optional]
**ptrId** | **String** |  |  [optional]
**networkDomain** | [**InlineResponse200108NetworkFloatingIpNetworkDomain**](InlineResponse200108NetworkFloatingIpNetworkDomain.md) |  |  [optional]
**createdBy** | [**InlineResponse200108NetworkFloatingIpCreatedBy**](InlineResponse200108NetworkFloatingIpCreatedBy.md) |  |  [optional]
**dateCreated** | **OffsetDateTime** |  |  [optional]
**lastUpdated** | **OffsetDateTime** |  |  [optional]



## Enum: IpStatusEnum

Name | Value
---- | -----
ASSIGNED | &quot;assigned&quot;
FREE | &quot;free&quot;
PENDING | &quot;pending&quot;



