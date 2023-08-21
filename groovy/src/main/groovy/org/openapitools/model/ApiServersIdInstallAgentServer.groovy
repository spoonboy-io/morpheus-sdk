package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiServersIdInstallAgentServerServerOs;

@Canonical
class ApiServersIdInstallAgentServer {
    /* SSH username to use when provisioning */
    String sshUsername
    /* SSH password to use, if not specified the account public key can be used */
    String sshPassword
    
    ApiServersIdInstallAgentServerServerOs serverOs
}
