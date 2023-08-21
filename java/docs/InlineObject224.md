

# InlineObject224

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**server** | [**ApiServersIdResizeServer**](ApiServersIdResizeServer.md) |  |  [optional]
**servicePlanOptions** | [**ApiServersIdResizeServicePlanOptions**](ApiServersIdResizeServicePlanOptions.md) |  |  [optional]
**volumes** | [**List&lt;InstanceCreateVolume&gt;**](InstanceCreateVolume.md) | List of volumes with their new sizes. |  [optional]
**deleteOriginalVolumes** | **Boolean** | Delete the original volumes after resizing. (Amazon only) |  [optional]
**networkInterfaces** | [**List&lt;InstanceCreateNetwork&gt;**](InstanceCreateNetwork.md) | Key for network configurations. Include id to update an existing interface. |  [optional]



