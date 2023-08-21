package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CheckWebConfig {
    /* HTTP method to use for testing */
    String webMethod
    /* Web URL you wish to use to run a check on */
    String webUrl
    /* Ignore SSL Errors */
    Boolean ignoreSSL = false
    /* If you want to use HTTP Basic Authentication, populate this field with the username */
    String checkUser
    /* If you want to use HTTP basic Authentication, populate this field with the password */
    String checkPassword
    /* Set value to `on` if you want to turn on text matching */
    String textCheckOn
    /* Set the string you want to look for in the page source */
    String webTextMatch
    /* Set to on to turn on tunneling */
    String tunnelOn
    /* Hostname or IP address of the proxy host */
    String sshHost
    /* Port for SSH on the proxy host, defaults to 22 */
    Long sshPort
    /* SSH user on the proxy host to login as */
    String sshUser
    /* Password for user, if not using key based authentication */
    String sshPassword
}
