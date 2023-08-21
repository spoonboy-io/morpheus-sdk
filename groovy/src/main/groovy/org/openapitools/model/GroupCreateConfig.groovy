package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class GroupCreateConfig {
    /* Optional DNS Integration ID */
    String dnsIntegrationId
    /* Optional CMDB Integration ID */
    String configCmdbId
    /* Optional Change Management Integration ID */
    String configCmId
    /* Optional Service Registry Integration ID */
    String serviceRegistryId
    /* Optional Configuration Management Integration ID */
    String configManagementId
    /* Enable or disable CMDB Discovery */
    Boolean configCmdbDiscovery
}
