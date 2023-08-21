package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterApiConfig {
    
    String serviceUrl
    
    String serviceHost
    
    String servicePath
    
    String serviceHostname
    
    Long servicePort
    
    String serviceUsername
    
    String servicePassword
    
    String servicePasswordHash
    /* API Token */
    String serviceToken
    /* Kube Config */
    String serviceAccess
    
    String serviceCert
    
    String serviceVersion
}
