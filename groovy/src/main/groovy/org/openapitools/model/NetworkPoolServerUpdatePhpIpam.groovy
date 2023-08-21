package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkPoolServerCreateBluecatCredential;
import org.openapitools.model.NetworkPoolServerUpdatePhpIpamConfig;

@Canonical
class NetworkPoolServerUpdatePhpIpam {
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
    Boolean ignoreSsl
    /* Network Filter */
    String networkFilter
    
    NetworkPoolServerUpdatePhpIpamConfig config
    
    NetworkPoolServerCreateBluecatCredential credential
}
