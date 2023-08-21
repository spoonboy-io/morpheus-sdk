package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerZoneNetworkServer;

@Canonical
class NetworkInterfaceUpdateSuccessServerZone {
    
    Long id
    
    Long accountId
    
    List<Long> groups = new ArrayList<Long>()
    
    String name
    
    String description
    
    String location
    
    String visibility
    
    Long zoneTypeId
    
    NetworkInterfaceUpdateSuccessServerZoneNetworkServer networkServer
    
    String securityServer
}
