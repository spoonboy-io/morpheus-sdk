package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GuidanceVmwareSizingResourceControllers;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InstanceConfig;
import org.openapitools.model.InstanceConnectionInfo;
import org.openapitools.model.InstanceInstancePrice;
import org.openapitools.model.InstanceInstanceType;
import org.openapitools.model.InstanceInterfaces;
import org.openapitools.model.InstanceLayout;
import org.openapitools.model.InstanceStats;
import org.openapitools.model.InstanceVolumes;

@Canonical
class Instance {
    
    Long id
    
    String uuid
    
    Long accountId
    
    InlineResponse20040AppDeployInstance tenant
    
    InstanceInstanceType instanceType
    
    InlineResponse20040AppDeployInstance group
    
    InlineResponse20040AppDeployInstance cloud
    
    List<Long> containers = new ArrayList<Long>()
    
    List<Long> servers = new ArrayList<Long>()
    
    List<InstanceConnectionInfo> connectionInfo = new ArrayList<InstanceConnectionInfo>()
    
    InstanceLayout layout
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType plan
    
    String name
    
    String description
    
    String environment
    
    InstanceConfig config
    
    String configGroup
    
    String configId
    
    String configRole
    
    List<InstanceVolumes> volumes = new ArrayList<InstanceVolumes>()
    
    List<GuidanceVmwareSizingResourceControllers> controllers = new ArrayList<GuidanceVmwareSizingResourceControllers>()
    
    List<InstanceInterfaces> interfaces = new ArrayList<InstanceInterfaces>()
    
    Object customOptions
    
    String instanceVersion
    
    List<String> labels = new ArrayList<String>()
    
    List<Object> tags = new ArrayList<Object>()
    
    List<Object> evars = new ArrayList<Object>()
    
    Long maxMemory
    
    Long maxStorage
    
    Long maxCores
    
    Long coresPerSocket
    
    String maxCpu
    
    BigDecimal hourlyCost
    
    BigDecimal hourlyPrice
    
    InstanceInstancePrice instancePrice
    
    Date dateCreated
    
    Date lastUpdated
    
    String hostName
    
    String domainName
    
    String environmentPrefix
    
    Boolean firewallEnabled
    
    String networkLevel
    
    Boolean autoScale
    
    String instanceContext
    
    String currentDeployId
    
    Boolean locked
    
    String status
    
    String statusMessage
    
    String errorMessage
    
    Date statusDate
    
    String statusPercent
    
    String statusEta
    
    String userStatus
    
    Long expireDays
    
    Long renewDays
    
    Long expireCount
    
    Date expireDate
    
    Date expireWarningDate
    
    Boolean expireWarningSent
    
    Long shutdownDays
    
    Long shutdownRenewDays
    
    Long shutdownCount
    
    Date shutdownDate
    
    Date shutdownWarningDate
    
    Boolean shutdownWarningSent
    
    Date removalDate
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    InlineResponse200107NetworkPoolCreatedBy owner
    
    String notes
    
    InstanceStats stats
    
    String powerSchedule
    
    Boolean isScalable
    
    Object instanceThreshold
    
    Boolean isBusy
    
    List<Object> apps = new ArrayList<Object>()
}
