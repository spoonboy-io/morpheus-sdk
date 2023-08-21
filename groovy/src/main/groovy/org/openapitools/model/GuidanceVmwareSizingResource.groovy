package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutComputeServerType;
import org.openapitools.model.GuidanceVmwareSizingResourceControllers;
import org.openapitools.model.GuidanceVmwareSizingResourceInterfaces;
import org.openapitools.model.GuidanceVmwareSizingResourceServerOs;
import org.openapitools.model.GuidanceVmwareSizingResourceStats;
import org.openapitools.model.GuidanceVmwareSizingResourceVolumes;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class GuidanceVmwareSizingResource {
    
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
    
    InlineResponse200107NetworkPoolCreatedBy owner
    
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
    
    GuidanceVmwareSizingResourceStats stats
    
    String status
    
    String statusMessage
    
    String errorMessage
    
    Date statusDate
    
    String statusPercent
    
    String statusEta
    
    String powerState
    
    Boolean agentInstalled
    
    Date lastAgentUpdate
    
    String agentVersion
    
    Long maxCores
    
    Long coresPerSocket
    
    Long maxMemory
    
    Long maxStorage
    
    String maxCpu
    
    BigDecimal hourlyPrice
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType sourceImage
    
    GuidanceVmwareSizingResourceServerOs serverOs
    
    List<GuidanceVmwareSizingResourceVolumes> volumes = new ArrayList<GuidanceVmwareSizingResourceVolumes>()
    
    List<GuidanceVmwareSizingResourceControllers> controllers = new ArrayList<GuidanceVmwareSizingResourceControllers>()
    
    List<GuidanceVmwareSizingResourceInterfaces> interfaces = new ArrayList<GuidanceVmwareSizingResourceInterfaces>()
    
    List<Object> labels = new ArrayList<Object>()
    
    List<Object> tags = new ArrayList<Object>()
    
    Boolean enabled
    
    String tagCompliant
    
    List<Long> containers = new ArrayList<Long>()
}
