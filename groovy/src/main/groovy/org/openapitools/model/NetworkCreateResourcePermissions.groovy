package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkCreateResourcePermissions {
    /* Pass true to allow access all groups */
    Boolean all
    /* Array of groups that are allowed access */
    List<Long> sites = new ArrayList<Long>()
}
