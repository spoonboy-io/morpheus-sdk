package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiClientsClient {
    /* ClientId */
    String clientId
    /* ClientSecret */
    String clientSecret
    /* Length of time accessToken is valid in seconds. */
    Integer accessTokenValiditySeconds
    /* Length of time refreshToken is valid in seconds. */
    Integer refreshTokenValiditySeconds
}
