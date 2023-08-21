package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ContainerContainerType;
import org.openapitools.model.ContainerContainerTypeSet;
import org.openapitools.model.ContainerInstance;
import org.openapitools.model.ContainerPlan;
import org.openapitools.model.ContainerPort;
import org.openapitools.model.ContainerStats;

@Canonical
class Container {
    
    Integer id
    
    String uuid
    
    Integer accountId
    
    ContainerInstance instance
    
    ContainerContainerType containerType
    
    ContainerContainerTypeSet containerTypeSet
    
    ContainerInstance server
    
    ContainerInstance cloud
    
    String name
    
    String ip
    
    String internalIp
    
    String internalHostname
    
    String externalHostname
    
    String externalDomain
    
    String externalFqdn
    
    List<ContainerPort> ports = new ArrayList<ContainerPort>()
    
    ContainerPlan plan
    
    Date dateCreated
    
    Date lastUpdated
    
    Boolean statsEnabled
    
    String status
    
    String userStatus
    
    String environmentPrefix
    
    ContainerStats stats
    
    Object runtimeInfo
    
    String containerVersion
    
    String repositoryImage
    
    String planCategory
    
    String hostname
    
    String domainName
    
    Boolean volumeCreated
    
    Boolean containerCreated
    
    Integer maxStorage
    
    Integer maxMemory
    
    Integer maxCores
    
    Integer maxCpu
    
    List<Object> availableActions = new ArrayList<Object>()
    
    String configGroup
    
    String configId
    
    String configRole
    
    Double hourlyCost
    
    Double hourlyPrice
}
