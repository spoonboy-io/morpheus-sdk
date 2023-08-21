package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateOkta {
    /* Okta URL */
    String url
    /* Administrator API Token */
    String administratorAPIToken
    /* Required Group */
    String requiredGroup
}
