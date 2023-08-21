package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200131ResourcePermissionSites;

@Canonical
class InlineResponse200131ResourcePermission {
    /* Pass `true` to allow access all groups */
    Boolean all = true
    /* Array of groups that are allowed access */
    List<InlineResponse200131ResourcePermissionSites> sites = new ArrayList<InlineResponse200131ResourcePermissionSites>()
}
