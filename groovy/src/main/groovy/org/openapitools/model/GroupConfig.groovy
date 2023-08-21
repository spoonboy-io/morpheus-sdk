package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class GroupConfig {
    
    String dnsIntegrationId
    
    String configCmdbId
    
    String configCmId
    
    String serviceRegistryId
    
    String configManagementId
    
    Boolean configCmdbDiscovery
}
