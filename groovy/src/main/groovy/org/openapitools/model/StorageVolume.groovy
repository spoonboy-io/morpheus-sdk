package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class StorageVolume {
    
    Long id
    
    String name
    
    String description
    
    Long controllerId
    
    String controllerMountPoint
    
    Boolean resizeable
    
    Boolean rootVolume
    
    String unitNumber
    
    String deviceName
    
    String deviceDisplayName
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
    
    Long typeId
    
    String category
    
    String status
    
    String statusMessage
    
    Boolean configurableIOPS
    
    Long maxStorage
    
    Long displayOrder
    
    String maxIOPS
    
    String uuid
    
    Boolean active
    
    InlineResponse20040AppDeployInstance zone
    
    Long zoneId
    
    InlineResponse20040AppDeployInstance datastore
    
    Long datastoreId
    
    String storageGroup
    
    String namespace
    
    String storageServer
    
    String source
    
    InlineResponse20040AppDeployInstance owner
}
