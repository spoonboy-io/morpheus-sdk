package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HubRegisterObjectHub {
    /* Company Name for new Morpheus Hub organization */
    String companyName
    /* First Name for new Morpheus Hub user */
    String firstName
    /* Last Name for new Morpheus Hub user */
    String lastName
    /* Email for new Morpheus Hub user */
    String email
    /* Password for new Morpheus Hub user */
    String password
    /* Job title of new Morpheus Hub user */
    String jobTitle
}
