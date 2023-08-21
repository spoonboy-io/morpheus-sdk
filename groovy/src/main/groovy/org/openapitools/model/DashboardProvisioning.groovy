package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class DashboardProvisioning {
    
    Long instanceCount
    
    List<Object> favoriteInstances = new ArrayList<Object>()
}
