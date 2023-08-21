package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ZoneVcenterConfigNetworkServer;

@Canonical
class ZoneGcpConfig {
    
    String privateKey
    
    String clientEmail
    
    String projectId
    
    String googleRegionId
    
    String importExisting
    
    String importExisting
    
    String applianceUrl
    
    String datacenterName
    
    String networkServerId
    
    ZoneVcenterConfigNetworkServer networkServer
    
    String securityServer
    
    String certificateProvider
    
    String backupMode
    
    String replicationMode
    
    String dnsIntegrationId
    
    String serviceRegistryId
    
    String configManagementId
    
    String privateKeyHash
}
