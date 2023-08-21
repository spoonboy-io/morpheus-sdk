package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InstanceConfigBackup;
import org.openapitools.model.InstanceConfigLayout;
import org.openapitools.model.InstanceServicePlanAutoOptions;
import org.openapitools.model.VdiPoolConfigConfig;
import org.openapitools.model.VdiPoolConfigDisplayNetworks;
import org.openapitools.model.VdiPoolConfigInstance;
import org.openapitools.model.VdiPoolConfigNetworkInterfaces;
import org.openapitools.model.VdiPoolConfigStorageControllers;
import org.openapitools.model.VdiPoolConfigVolumes;
import org.openapitools.model.VdiPoolConfigVolumesDisplay;

@Canonical
class VdiPoolConfig {
    
    InstanceServicePlanAutoOptions group
    
    InlineResponse20040AppDeployInstance cloud
    
    String type
    
    VdiPoolConfigInstance instance
    
    String name
    
    String environment
    
    VdiPoolConfigConfig config
    
    List<VdiPoolConfigVolumes> volumes = new ArrayList<VdiPoolConfigVolumes>()
    
    String hostName
    
    InstanceConfigLayout layout
    
    List<VdiPoolConfigStorageControllers> storageControllers = new ArrayList<VdiPoolConfigStorageControllers>()
    
    InstanceConfigLayout plan
    
    String version
    
    List<VdiPoolConfigNetworkInterfaces> networkInterfaces = new ArrayList<VdiPoolConfigNetworkInterfaces>()
    
    String executionId
    
    InstanceConfigBackup backup
    
    List<Object> loadBalancer = new ArrayList<Object>()
    
    Boolean hideLock
    
    Boolean hasNetworks
    
    List<VdiPoolConfigDisplayNetworks> displayNetworks = new ArrayList<VdiPoolConfigDisplayNetworks>()
    
    Long copies
    
    Boolean showScale
    
    Boolean hasPreview
    
    List<VdiPoolConfigVolumesDisplay> volumesDisplay = new ArrayList<VdiPoolConfigVolumesDisplay>()
}
