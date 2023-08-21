package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerCapacityInfo;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerComputeServerType;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerInterfaces;
import org.openapitools.model.NetworkInterfaceUpdateSuccessServerZone;

@Canonical
class NetworkInterfaceUpdateSuccessServer {
    
    Long id
    
    String externalId
    
    Long accountId
    
    String name
    
    String displayName
    
    String visibility
    
    String description
    
    Long zoneId
    
    Long siteId
    
    String sshHost
    
    Long sshPort
    
    String externalIp
    
    String internalIp
    
    String volumeId
    
    String platform
    
    String platformVersion
    
    String sshUsername
    
    String sshPassword
    
    String osDevice
    
    String dataDevice
    
    Boolean lvmEnabled
    
    String apiKey
    
    Boolean softwareRaid
    
    String config
    
    NetworkInterfaceUpdateSuccessServerCapacityInfo capacityInfo
    
    Date dateCreated
    
    Date lastUpdated
    
    String lastStats
    
    String status
    
    NetworkInterfaceUpdateSuccessServerComputeServerType computeServerType
    
    List<NetworkInterfaceUpdateSuccessServerInterfaces> interfaces = new ArrayList<NetworkInterfaceUpdateSuccessServerInterfaces>()
    
    NetworkInterfaceUpdateSuccessServerZone zone
}
