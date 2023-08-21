package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkTypeGcpConfigZonePool;

@Canonical
class NetworkTypeGcpConfig {
    /* GCP MTU */
    String mtu = MtuEnum._1460
    
    NetworkTypeGcpConfigZonePool zonePool
    /* Auto create subnets */
    Boolean autoCreate = true
}
