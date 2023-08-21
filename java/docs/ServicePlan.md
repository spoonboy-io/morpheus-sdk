

# ServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**name** | **String** |  |  [optional]
**code** | **String** |  |  [optional]
**active** | **Boolean** |  |  [optional]
**sortOrder** | **Long** |  |  [optional]
**description** | **String** |  |  [optional]
**maxStorage** | **BigDecimal** |  |  [optional]
**maxMemory** | **BigDecimal** |  |  [optional]
**maxCpu** | **BigDecimal** |  |  [optional]
**maxCores** | **BigDecimal** |  |  [optional]
**maxDisks** | **BigDecimal** |  |  [optional]
**coresPerSocket** | **BigDecimal** |  |  [optional]
**customCpu** | **Boolean** |  |  [optional]
**customCores** | **Boolean** |  |  [optional]
**customMaxStorage** | **Boolean** |  |  [optional]
**customMaxDataStorage** | **Boolean** |  |  [optional]
**customMaxMemory** | **Boolean** |  |  [optional]
**addVolumes** | **Boolean** |  |  [optional]
**memoryOptionSource** | **String** |  |  [optional]
**cpuOptionSource** | **String** |  |  [optional]
**dateCreated** | **OffsetDateTime** |  |  [optional]
**lastUpdated** | **OffsetDateTime** |  |  [optional]
**regionCode** | **String** |  |  [optional]
**visibility** | **String** |  |  [optional]
**editable** | **Boolean** |  |  [optional]
**provisionType** | [**GuidanceVmwareSizingPlanBeforeActionProvisionType**](GuidanceVmwareSizingPlanBeforeActionProvisionType.md) |  |  [optional]
**tenants** | **String** |  |  [optional]
**priceSets** | [**List&lt;GuidanceVmwareSizingPlanBeforeActionPriceSets&gt;**](GuidanceVmwareSizingPlanBeforeActionPriceSets.md) |  |  [optional]
**config** | [**ServicePlanConfig**](ServicePlanConfig.md) |  |  [optional]
**zones** | [**List&lt;InlineResponse20094Network&gt;**](InlineResponse20094Network.md) |  |  [optional]
**permissions** | [**ServicePlanPermissions**](ServicePlanPermissions.md) |  |  [optional]



