package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateOneLogin {
    /* OneLogin Subdomain */
    String subdomain
    /* OneLogin Region */
    String region
    /* API Client Secret */
    String clientSecret
    /* API Client ID */
    String clientId
    /* Required Role */
    String requiredRole
}
