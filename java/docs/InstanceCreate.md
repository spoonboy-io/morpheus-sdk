

# InstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceCreateInstance**](InstanceCreateInstance.md) |  | 
**zoneId** | **Long** | The Cloud ID to provision the instance onto. |  [optional]
**evars** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. |  [optional]
**copies** | **Long** | Number of copies to provision. |  [optional]
**layoutSize** | **Long** | Apply a multiply factor of containers/vms within the instance. |  [optional]
**servicePlanOptions** | **Object** | Map of custom options depending on selected service plan. |  [optional]
**securityGroups** | **List&lt;Object&gt;** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. |  [optional]
**volumes** | [**List&lt;InstanceCreateVolume&gt;**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of |  [optional]
**networkInterfaces** | [**List&lt;InstanceCreateNetwork&gt;**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  |  [optional]
**config** | [**AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject**](AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject.md) |  | 
**labels** | **List&lt;String&gt;** | Array of strings (keywords). |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. |  [optional]
**metadata** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. |  [optional]
**ports** | [**List&lt;InstanceCreatePorts&gt;**](InstanceCreatePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  |  [optional]
**taskSetId** | **Long** | The Workflow ID to execute. |  [optional]
**taskSetName** | **String** | The Workflow Name to execute. |  [optional]



