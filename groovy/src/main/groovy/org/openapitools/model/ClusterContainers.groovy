package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterContainersAvailableActions;
import org.openapitools.model.ClusterContainersContainerType;
import org.openapitools.model.ClusterContainersContainerTypeSet;
import org.openapitools.model.ClusterContainersPlan;
import org.openapitools.model.ClusterContainersStats;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ClusterContainers {
    
    Long id
    
    String uuid
    
    Long accountId
    
    String instance
    
    ClusterContainersContainerType containerType
    
    ClusterContainersContainerTypeSet containerTypeSet
    
    InlineResponse20040AppDeployInstance server
    
    InlineResponse20040AppDeployInstance cloud
    
    String name
    
    String ip
    
    String internalIp
    
    String internalHostname
    
    String externalHostname
    
    String externalDomain
    
    String externalFqdn
    
    List<Object> ports = new ArrayList<Object>()
    
    ClusterContainersPlan plan
    
    Date dateCreated
    
    Date lastUpdated
    
    Boolean statsEnabled
    
    String status
    
    String userStatus
    
    String environmentPrefix
    
    String configGroup
    
    String configId
    
    String configRole
    
    ClusterContainersStats stats
    
    Object runtimeInfo
    
    String containerVersion
    
    String repositoryImage
    
    String planCategory
    
    String hostname
    
    String domainName
    
    Boolean volumeCreated
    
    Boolean containerCreated
    
    String maxStorage
    
    String maxMemory
    
    String maxCores
    
    String maxCpu
    
    BigDecimal hourlyPrice
    
    List<ClusterContainersAvailableActions> availableActions = new ArrayList<ClusterContainersAvailableActions>()
}
