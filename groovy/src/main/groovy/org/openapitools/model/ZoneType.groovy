package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ZoneTypeOptionTypes1;
import org.openapitools.model.ZoneTypeServerTypes;

@Canonical
class ZoneType {
    
    Long id
    
    String name
    
    String code
    
    Boolean enabled
    
    Boolean provision
    
    Boolean autoCapacity
    
    Boolean migrationTarget
    
    Boolean hasDatastores
    
    Boolean hasNetworks
    
    Boolean hasResourcePools
    
    Boolean hasSecurityGroups
    
    Boolean hasContainers
    
    Boolean hasBareMetal
    
    Boolean hasServices
    
    Boolean hasFunctions
    
    Boolean hasJobs
    
    Boolean hasDiscovery
    
    Boolean hasCloudInit
    
    Boolean hasFolders
    
    Boolean hasFloatingIps
    
    Boolean hasMarketplace
    
    Boolean canCreateResourcePools
    
    Boolean canDeleteResourcePools
    
    Boolean canCreateDatastores
    
    Boolean canCreateNetworks
    
    Boolean canChooseContainerMode
    
    Boolean provisionRequiresResourcePool
    
    Boolean supportsDistributedWorker
    
    String cloud
    
    List<Long> provisionTypes = new ArrayList<Long>()
    
    Long zoneInstanceTypeLayoutId
    
    List<ZoneTypeServerTypes> serverTypes = new ArrayList<ZoneTypeServerTypes>()
    
    List<ZoneTypeOptionTypes1> optionTypes = new ArrayList<ZoneTypeOptionTypes1>()
}
