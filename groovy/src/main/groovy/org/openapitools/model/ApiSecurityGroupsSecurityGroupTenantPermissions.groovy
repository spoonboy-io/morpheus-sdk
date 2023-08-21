package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiSecurityGroupsSecurityGroupTenantPermissions {
    /* Array of tenant account ids that are allowed access */
    List<Long> accounts = new ArrayList<Long>()
    /* Array of tenant account ids that can manage */
    List<Long> canManageAccounts = new ArrayList<Long>()
}
