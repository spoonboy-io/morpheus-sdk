package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class CloudFoundryResourcePoolConfig {
    /* Array of manager usernames */
    List<String> managers = new ArrayList<String>()
    /* Array of developer usernames */
    List<String> developers = new ArrayList<String>()
    /* Array of auditor usernames */
    List<String> auditors = new ArrayList<String>()
}
