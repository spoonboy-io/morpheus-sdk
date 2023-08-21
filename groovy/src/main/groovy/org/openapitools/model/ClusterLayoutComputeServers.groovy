package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterLayoutComputeServerType;
import org.openapitools.model.ClusterLayoutContainerType;

@Canonical
class ClusterLayoutComputeServers {
    
    Long id
    
    Long priorityOrder
    
    Long nodeCount
    
    String nodeType
    
    Long minNodeCount
    
    String maxNodeCount
    
    Boolean dynamicCount
    
    Boolean installContainerRuntime
    
    Boolean installStorageRuntime
    
    String name
    
    String code
    
    String category
    
    String config
    
    ClusterLayoutContainerType containerType
    
    ClusterLayoutComputeServerType computeServerType
    
    String provisionService
    
    String planCategory
    
    String namePrefix
    
    String nameSuffix
    
    Boolean forceNameIndex
    
    Boolean loadBalance
}
