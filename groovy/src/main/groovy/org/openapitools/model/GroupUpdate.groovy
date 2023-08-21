package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.GroupCreateConfig;

@Canonical
class GroupUpdate {
    /* A unique name scoped to your account for the group */
    String name
    /* Optional code for use with policies */
    String code
    /* Optional location argument for your group */
    String location
    
    GroupCreateConfig config
}
