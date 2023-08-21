package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutComputeServerType;
import org.openapitools.model.GuidanceVmwareSizingResourceControllers;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.Owner;
import org.openapitools.model.PriceSetVolumeType;
import org.openapitools.model.ServerConfig;
import org.openapitools.model.ServerInterfaces;
import org.openapitools.model.ServerServerOs;
import org.openapitools.model.ServerStats;
import org.openapitools.model.ServerVolumes;

@Canonical
class Server {
    
    Long id
    
    String uuid
    
    String externalId
    
    String internalId
    
    String externalUniqueId
    
    String name
    
    String externalName
    
    String hostname
    
    InlineResponse20040AppDeployInstance parentServer
    
    Long accountId
    
    InlineResponse20040AppDeployInstance account
    
    Owner owner
    
    InlineResponse20040AppDeployInstance zone
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType plan
    
    ClusterLayoutComputeServerType computeServerType
    
    String visibility
    
    String description
    
    Long zoneId
    
    Long siteId
    
    Long resourcePoolId
    
    Long folderId
    
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
    
    ServerStats stats
    
    String status
    
    String statusMessage
    
    String errorMessage
    
    Date statusDate
    
    String statusPercent
    
    String statusEta
    
    String powerState
    
    Boolean agentInstalled
    
    String lastAgentUpdate
    
    String agentVersion
    
    Long maxCores
    
    Long coresPerSocket
    
    Long maxMemory
    
    Long maxStorage
    
    Long maxCpu
    
    BigDecimal hourlyCost
    
    BigDecimal hourlyPrice
    
    PriceSetVolumeType sourceImage
    
    ServerServerOs serverOs
    
    List<ServerVolumes> volumes = new ArrayList<ServerVolumes>()
    
    List<GuidanceVmwareSizingResourceControllers> controllers = new ArrayList<GuidanceVmwareSizingResourceControllers>()
    
    List<ServerInterfaces> interfaces = new ArrayList<ServerInterfaces>()
    
    List<Object> labels = new ArrayList<Object>()
    
    List<Object> tags = new ArrayList<Object>()
    
    Boolean enabled
    
    String tagCompliant
    
    List<Long> containers = new ArrayList<Long>()
    
    ServerConfig config
    
    Boolean guestConsolePreferred
    
    String guestConsoleType
    
    String guestConsoleUsername
    
    String guestConsolePassword
    
    String guestConsolePasswordHash
    
    String guestConsolePort
}
