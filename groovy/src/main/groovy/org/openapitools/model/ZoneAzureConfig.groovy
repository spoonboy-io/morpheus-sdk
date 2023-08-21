package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ZoneVcenterConfigNetworkServer;

@Canonical
class ZoneAzureConfig {
    
    String subscriberId
    
    String tenantId
    
    String clientId
    
    String clientSecret
    
    String resourceGroup
    
    String importExisting
    
    String importExisting
    
    String inventoryLevel
    
    String applianceUrl
    
    String datacenterName
    
    String networkServerId
    
    ZoneVcenterConfigNetworkServer networkServer
    
    String securityMode
    
    String certificateProvider
    
    String backupMode
    
    String replicationMode
    
    String dnsIntegrationId
    
    String configManagementId
    
    String configCmdbId
    
    String securityServer
    
    String accountType
    
    String serviceRegistryId
    
    String cloudType
    
    String rpcMode
    
    String diskEncryption
    
    String encryptionSet
    
    String cspTenantId
    
    String cspClientId
    
    String cspClientSecret
    
    String cspCustomer
    
    Boolean configCmdbDiscovery
    
    String azureCostingMode
    
    String clientSecretHash
    
    String cspClientSecretHash
}
