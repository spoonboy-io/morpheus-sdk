package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkPoolServerCreateBluecatCredential;
import org.openapitools.model.NetworkPoolServerCreatePhpIpamConfig;

@Canonical
class NetworkPoolServerCreatePhpIpam {
    /* Type Code (phpIPAM) */
    String type
    /* Name */
    String name
    /* Can be used to enable / disable the network pool server. */
    Boolean enabled = true
    /* URL */
    String serviceUrl
    /* Username */
    String serviceUsername
    /* Password */
    String servicePassword
    /* Throttle Rate */
    Long serviceThrottleRate = 0l
    /* Disable SSL SNI Verification */
    Boolean ignoreSsl = true
    /* Network Filter */
    String networkFilter
    
    NetworkPoolServerCreatePhpIpamConfig config
    
    NetworkPoolServerCreateBluecatCredential credential
}
