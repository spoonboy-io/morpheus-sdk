# MorpheusApi.InlineObject224

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**server** | [**ApiServersIdResizeServer**](ApiServersIdResizeServer.md) |  | [optional] 
**servicePlanOptions** | [**ApiServersIdResizeServicePlanOptions**](ApiServersIdResizeServicePlanOptions.md) |  | [optional] 
**volumes** | [**[InstanceCreateVolume]**](InstanceCreateVolume.md) | List of volumes with their new sizes. | [optional] 
**deleteOriginalVolumes** | **Boolean** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**networkInterfaces** | [**[InstanceCreateNetwork]**](InstanceCreateNetwork.md) | Key for network configurations. Include id to update an existing interface. | [optional] 


