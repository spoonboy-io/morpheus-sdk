package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CheckElasticsearchConfig {
    /* Hostname or IP address of the Elasticsearch server */
    String esHost
    /* Port to connect to the HTTP service */
    String esPort
    
    String checkUser
    
    String textCheckOn
    
    String checkPassword
    
    String webTextMatch
    
    String checkPasswordHash
    /* Set to on to turn on tunneling */
    String tunnelOn = TunnelOnEnum.OFF
    /* Hostname or IP address of the proxy host */
    String sshHost
    /* Port for SSH on the proxy host, defaults to 22 */
    Long sshPort
    /* SSH user on the proxy host to login as */
    String sshUser
    /* Password for user, if not using key based authentication */
    String sshPassword
}
