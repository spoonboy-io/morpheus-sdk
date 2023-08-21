# MorpheusApi.InstanceResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceResizeInstance**](InstanceResizeInstance.md) |  | [optional] 
**servicePlanOptions** | [**ServicePlanOptions**](ServicePlanOptions.md) |  | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan | [optional] 
**deleteOriginalVolumes** | **Boolean** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**networkInterfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. | [optional] 


