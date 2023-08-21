package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstancesConfigAzure {
    /* id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`. */
    String resourcePoolId
    /* Availability Options */
    String availabilityOptions
    /* Availability Set */
    String availabilitySet
    /* Availability Zone */
    Long availabilityZone
    /* Assign Public IP */
    String azurefloatingIp
    /* Boot Diagnostics */
    String bootDiagnostics
    /* OS Guest Diagnostics */
    String osGuestDiagnostics
    /* Diagnostics Storage Account */
    String diagnosticsStorageAccount
    /* Create User */
    Boolean createUser = true
}
