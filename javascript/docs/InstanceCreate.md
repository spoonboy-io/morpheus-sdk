# MorpheusApi.InstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceCreateInstance**](InstanceCreateInstance.md) |  | 
**zoneId** | **Number** | The Cloud ID to provision the instance onto. | [optional] 
**evars** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**copies** | **Number** | Number of copies to provision. | [optional] [default to 1]
**layoutSize** | **Number** | Apply a multiply factor of containers/vms within the instance. | [optional] [default to 1]
**servicePlanOptions** | **Object** | Map of custom options depending on selected service plan. | [optional] 
**securityGroups** | **[Object]** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | [optional] 
**networkInterfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**config** | [**AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject**](AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject.md) |  | 
**labels** | **[String]** | Array of strings (keywords). | [optional] 
**tags** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**metadata** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. | [optional] 
**ports** | [**[InstanceCreatePorts]**](InstanceCreatePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**taskSetId** | **Number** | The Workflow ID to execute. | [optional] 
**taskSetName** | **String** | The Workflow Name to execute. | [optional] 


