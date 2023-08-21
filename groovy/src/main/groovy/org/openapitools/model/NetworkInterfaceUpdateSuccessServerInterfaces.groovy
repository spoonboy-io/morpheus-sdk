package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class NetworkInterfaceUpdateSuccessServerInterfaces {
    
    Long id
    
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> addresses = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    
    String internalId
    
    String interfaceId
    
    Long displayOrder
    
    Object networkPool
    
    Boolean dhcp
    
    String uuid
    
    Boolean active
    
    String uniqueId
    
    String subnet
    
    Boolean replaceHostRecord
    
    String ipMode
    
    String version
    
    String ipSubnet
    
    String config
    
    String publicIpAddress
    
    String fabricId
    
    String ipv6Subnet
    
    String macAddress
    
    String publicIpv6Address
    
    String refType
    
    String networkGroup
    
    String refId
    
    String networkDomain
    
    String name
    
    Boolean primaryInterface
    
    Object networkPoolIPv6
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites network
    
    String vlanId
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites type
    
    String networkPosition
    
    Boolean poolAssigned
    
    String description
    
    String externalType
    
    String externalId
}
