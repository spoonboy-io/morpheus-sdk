

# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**uuid** | **String** |  |  [optional]
**accountId** | **Long** |  |  [optional]
**tenant** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  |  [optional]
**instanceType** | [**InstanceInstanceType**](InstanceInstanceType.md) |  |  [optional]
**group** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  |  [optional]
**cloud** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  |  [optional]
**containers** | **List&lt;Long&gt;** |  |  [optional]
**servers** | **List&lt;Long&gt;** |  |  [optional]
**connectionInfo** | [**List&lt;InstanceConnectionInfo&gt;**](InstanceConnectionInfo.md) |  |  [optional]
**layout** | [**InstanceLayout**](InstanceLayout.md) |  |  [optional]
**plan** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  |  [optional]
**name** | **String** |  |  [optional]
**description** | **String** |  |  [optional]
**environment** | **String** |  |  [optional]
**config** | [**InstanceConfig**](InstanceConfig.md) |  |  [optional]
**configGroup** | **String** |  |  [optional]
**configId** | **String** |  |  [optional]
**configRole** | **String** |  |  [optional]
**volumes** | [**List&lt;InstanceVolumes&gt;**](InstanceVolumes.md) |  |  [optional]
**controllers** | [**List&lt;GuidanceVmwareSizingResourceControllers&gt;**](GuidanceVmwareSizingResourceControllers.md) |  |  [optional]
**interfaces** | [**List&lt;InstanceInterfaces&gt;**](InstanceInterfaces.md) |  |  [optional]
**customOptions** | **Object** |  |  [optional]
**instanceVersion** | **String** |  |  [optional]
**labels** | **List&lt;String&gt;** |  |  [optional]
**tags** | **List&lt;Object&gt;** |  |  [optional]
**evars** | **List&lt;Object&gt;** |  |  [optional]
**maxMemory** | **Long** |  |  [optional]
**maxStorage** | **Long** |  |  [optional]
**maxCores** | **Long** |  |  [optional]
**coresPerSocket** | **Long** |  |  [optional]
**maxCpu** | **String** |  |  [optional]
**hourlyCost** | **BigDecimal** |  |  [optional]
**hourlyPrice** | **BigDecimal** |  |  [optional]
**instancePrice** | [**InstanceInstancePrice**](InstanceInstancePrice.md) |  |  [optional]
**dateCreated** | **OffsetDateTime** |  |  [optional]
**lastUpdated** | **OffsetDateTime** |  |  [optional]
**hostName** | **String** |  |  [optional]
**domainName** | **String** |  |  [optional]
**environmentPrefix** | **String** |  |  [optional]
**firewallEnabled** | **Boolean** |  |  [optional]
**networkLevel** | **String** |  |  [optional]
**autoScale** | **Boolean** |  |  [optional]
**instanceContext** | **String** |  |  [optional]
**currentDeployId** | **String** |  |  [optional]
**locked** | **Boolean** |  |  [optional]
**status** | **String** |  |  [optional]
**statusMessage** | **String** |  |  [optional]
**errorMessage** | **String** |  |  [optional]
**statusDate** | **OffsetDateTime** |  |  [optional]
**statusPercent** | **String** |  |  [optional]
**statusEta** | **String** |  |  [optional]
**userStatus** | **String** |  |  [optional]
**expireDays** | **Long** |  |  [optional]
**renewDays** | **Long** |  |  [optional]
**expireCount** | **Long** |  |  [optional]
**expireDate** | **OffsetDateTime** |  |  [optional]
**expireWarningDate** | **OffsetDateTime** |  |  [optional]
**expireWarningSent** | **Boolean** |  |  [optional]
**shutdownDays** | **Long** |  |  [optional]
**shutdownRenewDays** | **Long** |  |  [optional]
**shutdownCount** | **Long** |  |  [optional]
**shutdownDate** | **OffsetDateTime** |  |  [optional]
**shutdownWarningDate** | **OffsetDateTime** |  |  [optional]
**shutdownWarningSent** | **Boolean** |  |  [optional]
**removalDate** | **OffsetDateTime** |  |  [optional]
**createdBy** | [**InlineResponse200107NetworkPoolCreatedBy**](InlineResponse200107NetworkPoolCreatedBy.md) |  |  [optional]
**owner** | [**InlineResponse200107NetworkPoolCreatedBy**](InlineResponse200107NetworkPoolCreatedBy.md) |  |  [optional]
**notes** | **String** |  |  [optional]
**stats** | [**InstanceStats**](InstanceStats.md) |  |  [optional]
**powerSchedule** | **String** |  |  [optional]
**isScalable** | **Boolean** |  |  [optional]
**instanceThreshold** | **Object** |  |  [optional]
**isBusy** | **Boolean** |  |  [optional]
**apps** | **List&lt;Object&gt;** |  |  [optional]


