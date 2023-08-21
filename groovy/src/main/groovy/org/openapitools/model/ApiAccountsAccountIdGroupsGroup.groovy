package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiAccountsAccountIdGroupsGroup {
    /* A unique name scoped to the subtenant for the group */
    String name
    /* Optional description field if you want to put more info there */
    String description
    /* Optional code for use with policies */
    String code
    /* location */
    String location
}
