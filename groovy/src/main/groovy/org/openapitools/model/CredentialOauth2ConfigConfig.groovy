package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CredentialOauth2ConfigConfig {
    /* OAuth 2.0 grant type */
    String grantType
    /* Token endpoint */
    String accessTokenUrl
    /* Client ID */
    String clientId
    /* Client Secret */
    String clientSecret
    /* Scope */
    String scope
    /* Client Authentication Method */
    String clientAuth
}
