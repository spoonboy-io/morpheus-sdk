package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.AppSecurityGroups;

@Canonical
class SetAppSecurityGroups200Response {
    
    Boolean success
    
    List<AppSecurityGroups> securityGroups
}
