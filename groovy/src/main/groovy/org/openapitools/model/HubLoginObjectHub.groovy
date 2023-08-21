package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HubLoginObjectHub {
    /* Email for existing Morpheus Hub user */
    String email
    /* Password for existing Morpheus Hub user */
    String password
}
