package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class RolePermissionReportType {
    /* `id` of the report type */
    Integer reportTypeId
    /* The new access level. */
    String access
}
