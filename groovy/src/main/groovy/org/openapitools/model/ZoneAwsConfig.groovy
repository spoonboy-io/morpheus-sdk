package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ZoneVcenterConfigNetworkServer;

@Canonical
class ZoneAwsConfig {
    
    String endpoint
    
    String accessKey
    
    String secretKey
    
    String useHostCredentials
    
    String stsAssumeRole
    
    String isVpc
    
    String vpc
    
    String imageStoreId
    
    String ebsEncryption
    
    String costingReport
    
    String costingFolder
    
    String costingBucket
    
    String costingBucketName
    
    String costingRegion
    
    String costingAccessKey
    
    String costingSecretKey
    
    String costingReportName
    
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
    
    Boolean configCmdbDiscovery
    
    String secretKeyHash
    
    String costingSecretKeyHash
}
