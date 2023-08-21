package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class AccessToken {
    /* Token that grants API Access */
    String accessToken
    /* Token that can request an new API access_token */
    String refreshToken
    /* Epoch time when token expires */
    BigDecimal expiresIn
    /* Token type granted */
    String tokenType
    /* Scope granted */
    String scope
}
