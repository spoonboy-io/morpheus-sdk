package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClustersLayout;
import org.openapitools.model.ClustersServers;
import org.openapitools.model.ClustersWorkerStats;
import org.openapitools.model.ClustersZone;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Clusters {
    
    Long id
    
    String name
    
    String code
    
    String category
    
    String visibility
    
    String description
    
    String location
    
    Boolean enabled
    
    String serviceUrl
    
    String serviceHost
    
    String servicePath
    
    String serviceHostname
    
    Long servicePort
    
    String serviceUsername
    
    String servicePassword
    
    String servicePasswordHash
    
    String serviceToken
    
    String serviceTokenHash
    
    String serviceAccess
    
    String serviceAccessHash
    
    String serviceCert
    
    String serviceCertHash
    
    String serviceVersion
    
    String searchDomains
    
    Boolean enableInternalDns
    
    String internalId
    
    String externalId
    
    String datacenterId
    
    String status
    
    Date statusDate
    
    String statusMessage
    
    String inventoryLevel
    
    Date lastSync
    
    Date nextRunDate
    
    Long lastSyncDuration
    
    Date dateCreated
    
    Date lastUpdated
    
    Boolean managed
    
    List<String> labels = new ArrayList<String>()
    
    String serviceEntry
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    String userGroup
    
    ClustersLayout layout
    
    InlineResponse20040AppDeployInstance owner
    
    List<ClustersServers> servers = new ArrayList<ClustersServers>()
    
    List<Object> accounts = new ArrayList<Object>()
    
    List<Object> integrations = new ArrayList<Object>()
    
    InlineResponse20040AppDeployInstance site
    
    InlineResponse20040AppDeployInstance type
    
    ClustersZone zone
    
    ClustersWorkerStats workerStats
    
    Object config
}
