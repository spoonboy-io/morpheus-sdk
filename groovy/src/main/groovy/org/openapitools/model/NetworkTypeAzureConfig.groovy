package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkTypeAzureConfig {
    /* Resource Group Name */
    String resourceGroupId
    /* Subnet Name */
    String subnetName
    /* The subnet's address range in CIDR notation (e.g. 192.168.1.0/24). It must be contained by the address space of the virtual network. */
    String subnetCidr
}
