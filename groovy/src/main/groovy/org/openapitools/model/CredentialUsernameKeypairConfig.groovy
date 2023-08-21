package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.CredentialAccessSecretKeyConfigIntegration;
import org.openapitools.model.CredentialEmailPrivateKeyConfigAuthKey;

@Canonical
class CredentialUsernameKeypairConfig {
    /* Credential Type Code */
    String type
    /* A unique name scoped to your account for the credential */
    String name
    /* Optional Description */
    String description
    /* Credential enabled */
    Boolean enabled = true
    
    CredentialAccessSecretKeyConfigIntegration integration
    /* Username */
    String username
    
    CredentialEmailPrivateKeyConfigAuthKey authKey
}
