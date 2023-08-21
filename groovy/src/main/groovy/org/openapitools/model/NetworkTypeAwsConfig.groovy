package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkTypeAwsConfigZonePool;

@Canonical
class NetworkTypeAwsConfig {
    /* Availability Zone Name */
    String availabilityZone
    /* Network CIDR */
    String cidr
    /* Assign public IPs by default. */
    Boolean assignPublicIp
    
    NetworkTypeAwsConfigZonePool zonePool
}
