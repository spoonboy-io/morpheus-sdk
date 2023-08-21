package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class CheckSqlConfig {
    /* Hostname or IP address of the database */
    String dbHost
    /* Database Port (defaults to default port of DB type selected) */
    String dbPort
    /* Database username */
    String dbUser
    /* Database password, (all check data is encrypted inside the database) */
    String dbPassword
    /* Database password hash */
    String dbPasswordHash
    /* Database name you would like to connect to */
    String dbName
    /* Query to test */
    String dbQuery
    /* Can be set to `lt` (less than), `gt` (greater than), `equal` (Equal to) for comparison */
    String checkOperator
    
    BigDecimal checkResult
    
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
