

# CatalogItemTypeInstanceScribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group** | [**CatalogItemTypeInstanceScribeGroup**](CatalogItemTypeInstanceScribeGroup.md) |  | 
**cloud** | [**CatalogItemTypeInstanceScribeCloud**](CatalogItemTypeInstanceScribeCloud.md) |  | 
**type** | **String** | The type of instance by code we want to fetch. | 
**name** | **String** | Name of the instance to be created. | 
**config** | [**AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject**](AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject.md) |  | 
**volumes** | [**List&lt;InstanceCreateVolume&gt;**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**hostName** | **String** | Hostname of the instance to be created.  Can be the same as the instance name. |  [optional]
**environment** | **String** | Environment code |  [optional]
**layout** | [**CatalogItemTypeInstanceScribeLayout**](CatalogItemTypeInstanceScribeLayout.md) |  | 
**plan** | [**CatalogItemTypeInstanceScribePlan**](CatalogItemTypeInstanceScribePlan.md) |  | 
**version** | **String** | Version of the layout to create. |  [optional]
**evars** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. |  [optional]
**servicePlanOptions** | [**ServicePlanOptions**](ServicePlanOptions.md) |  |  [optional]
**securityGroups** | [**List&lt;CatalogItemTypeInstanceScribeSecurityGroups&gt;**](CatalogItemTypeInstanceScribeSecurityGroups.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. |  [optional]
**networkInterfaces** | [**List&lt;InstanceCreateNetwork&gt;**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  |  [optional]
**labels** | **List&lt;String&gt;** | Array of strings (keywords). |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. |  [optional]
**metadata** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. |  [optional]
**ports** | [**List&lt;CatalogItemTypeInstanceScribePorts&gt;**](CatalogItemTypeInstanceScribePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  |  [optional]
**taskSetId** | **Long** | The Workflow ID to execute. |  [optional]
**taskSetName** | **String** | The Workflow Name to execute. |  [optional]



