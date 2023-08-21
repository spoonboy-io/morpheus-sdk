package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiAccountsAccountRole;
import org.openapitools.model.CurrencyCode;

@Canonical
class ApiAccountsAccount {
    /* Name */
    String name
    /* Description */
    String description
    
    ApiAccountsAccountRole role
    /* The subdomain. This will be part of the login URL and username for sub tenant users. */
    String subdomain
    
    CurrencyCode currency = CurrencyCode.USD
}
