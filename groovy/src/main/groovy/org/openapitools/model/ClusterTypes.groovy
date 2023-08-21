package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterTypesControllerTypes;
import org.openapitools.model.OptionType;

@Canonical
class ClusterTypes {
    
    Long id
    
    String deployTargetService
    
    String shortName
    
    String providerType
    
    String code
    
    String hostService
    
    Boolean managed
    
    Boolean hasMasters
    
    Boolean hasWorkers
    
    String viewSet
    
    String imageCode
    
    Boolean kubeCtlLocal
    
    Boolean hasDatastore
    
    Boolean supportsCloudScaling
    
    String name
    
    Boolean hasDefaultDataDisk
    
    Boolean canManage
    
    Boolean hasCluster
    
    String description
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<ClusterTypesControllerTypes> controllerTypes = new ArrayList<ClusterTypesControllerTypes>()
    
    List<ClusterTypesControllerTypes> workerTypes = new ArrayList<ClusterTypesControllerTypes>()
}
