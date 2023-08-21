package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionDefaultReportType {
    /* `ReportTypes` is the code for Default Report Type Access */
    String permissionCode
    /* The new access level. */
    String access
}
