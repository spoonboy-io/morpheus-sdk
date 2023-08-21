package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class InlineObject94 {
    /* List of all security groups ids which should be applied. If no security groups should apply, pass '[]' */
    List<Long> securityGroupIds = new ArrayList<Long>()
}
