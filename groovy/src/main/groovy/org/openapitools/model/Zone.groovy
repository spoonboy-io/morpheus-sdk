package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfobjectobject;
import org.openapitools.model.AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.ZoneGroups;
import org.openapitools.model.ZoneStats;

@Canonical
class Zone {
    
    Long id
    
    String uuid
    
    String externalId
    
    String name
    
    String code
    
    String location
    
    InlineResponse20040AppDeployInstance owner
    
    Long accountId
    
    InlineResponse20040AppDeployInstance account
    
    String visibility
    
    Boolean enabled
    
    String status
    
    String statusMessage
    
    Date statusDate
    
    String costStatus
    
    String costStatusMessage
    
    Date costStatusDate
    
    Long costLastSyncDuration
    
    Date costLastSync
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType zoneType
    
    Long zoneTypeId
    
    String guidanceMode
    
    String storageMode
    
    String agentMode
    
    String userDataLinux
    
    String userDataWindows
    
    String consoleKeymap
    
    String containerMode
    
    String costingMode
    
    String serviceVersion
    
    String securityMode
    
    String inventoryLevel
    
    String timezone
    
    String apiProxy
    
    String provisioningProxy
    
    InlineResponse20082LoadBalancerInstanceSslCert networkDomain
    
    String domainName
    
    String regionCode
    
    Boolean autoRecoverPowerState
    
    Long scalePriority
    
    AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig config = null
    
    AnyOfobjectobject credential = null
    /* Logo image URL */
    String imagePath
    /* Dark logo image URL */
    String darkImagePath
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastSync
    
    Long lastSyncDuration
    
    Date nextRunDate
    
    List<ZoneGroups> groups = new ArrayList<ZoneGroups>()
    
    InlineResponse20082LoadBalancerInstanceSslCert securityServer
    
    InlineResponse20082LoadBalancerInstanceSslCert networkServer
    
    ZoneStats stats
    
    Long serverCount
}
