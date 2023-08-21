package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstancesConfigVMWare {
    /* Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. */
    Boolean noAgent = false
    /* id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`. */
    String resourcePoolId
    /* Specific host to deploy to if so desired. */
    String hostId
    /* Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. */
    String smbiosAssetTag
    /* Enable Nested Virtualization */
    String nestedVirtualization = NestedVirtualizationEnum.OFF
    /* VMWare Folder External ID (as a String) or ID (as an Integer or String) */
    String vmwareFolderId
}
