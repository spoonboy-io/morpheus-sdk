package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.VirtualImageLocationVirtualImage;

@Canonical
class VirtualImageLocation {
    
    Long id
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType cloud
    
    String code
    
    String internalId
    
    String externalId
    
    String externalDiskId
    
    String remotePath
    
    String imagePath
    
    String imageName
    
    String imageRegion
    
    String imageFolder
    
    String refType
    
    Long refId
    
    String nodeRefType
    
    String nodeRefId
    
    String subRefType
    
    String subRefId
    
    Boolean isPublic
    
    Boolean systemImage
    
    Long diskIndex
    
    String pricePlan
    
    List<Object> volumes = new ArrayList<Object>()
    
    List<Object> storageControllers = new ArrayList<Object>()
    
    List<Object> networkInterfaces = new ArrayList<Object>()
    
    VirtualImageLocationVirtualImage virtualImage
}
