package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionReportTypeAll {
    /* Apply to all report types */
    Boolean allReportTypes
    /* The new access level. */
    String access
}
