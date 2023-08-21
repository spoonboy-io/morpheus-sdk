package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CheckSocketConfig {
    /* Hostname or IP address of the socket server */
    String host
    /* TCP port */
    String port
    /* Connection string you might want to send to the service */
    String send
    /* Response from the service to match against */
    String responseMatch
    
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
