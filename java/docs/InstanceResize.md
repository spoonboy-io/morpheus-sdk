

# InstanceResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instance** | [**InstanceResizeInstance**](InstanceResizeInstance.md) |  |  [optional]
**servicePlanOptions** | [**ServicePlanOptions**](ServicePlanOptions.md) |  |  [optional]
**volumes** | [**List&lt;InstanceCreateVolume&gt;**](InstanceCreateVolume.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan |  [optional]
**deleteOriginalVolumes** | **Boolean** | Delete the original volumes after resizing. (Amazon only) |  [optional]
**networkInterfaces** | [**List&lt;InstanceCreateNetwork&gt;**](InstanceCreateNetwork.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. |  [optional]



