package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.ClusterLayoutComputeServerType;
import org.openapitools.model.ClusterMastersInterfaces;
import org.openapitools.model.ClusterMastersStats;
import org.openapitools.model.ClusterMastersVolumes;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class ClusterMasters {
    
    Long id
    
    String uuid
    
    String externalId
    
    String internalId
    
    String externalUniqueId
    
    String name
    
    String externalName
    
    String hostname
    
    Long accountId
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse200107NetworkPoolCreatedBy owner
    
    InlineResponse20040AppDeployInstance zone
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType plan
    
    ClusterLayoutComputeServerType computeServerType
    
    String visibility
    
    String description
    
    Long zoneId
    
    Long siteId
    
    Long resourcePoolId
    
    String folderId
    
    String sshHost
    
    Long sshPort
    
    String externalIp
    
    String internalIp
    
    String volumeId
    
    String platform
    
    String platformVersion
    
    String sshUsername
    
    String sshPassword
    
    String sshPasswordHash
    
    String osDevice
    
    String osType
    
    String dataDevice
    
    Boolean lvmEnabled
    
    String apiKey
    
    Boolean softwareRaid
    
    Date dateCreated
    
    Date lastUpdated
    
    ClusterMastersStats stats
    
    String status
    
    String statusMessage
    
    String errorMessage
    
    String statusDate
    
    String statusPercent
    
    String statusEta
    
    String powerState
    
    Boolean agentInstalled
    
    Date lastAgentUpdate
    
    String agentVersion
    
    Long maxCores
    
    String coresPerSocket
    
    Long maxMemory
    
    Long maxStorage
    
    Long maxCpu
    
    BigDecimal hourlyPrice
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType sourceImage
    
    String serverOs
    
    List<ClusterMastersVolumes> volumes = new ArrayList<ClusterMastersVolumes>()
    
    List<Object> controllers = new ArrayList<Object>()
    
    List<ClusterMastersInterfaces> interfaces = new ArrayList<ClusterMastersInterfaces>()
    
    List<String> labels = new ArrayList<String>()
    
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    
    Boolean enabled
    
    String tagCompliant
    
    List<Long> containers = new ArrayList<Long>()
    
    Boolean guestConsolePreferred
    
    String guestConsoleType
    
    String guestConsoleUsername
    
    String guestConsolePassword
    
    String guestConsolePasswordHash
    
    String guestConsolePort
}
