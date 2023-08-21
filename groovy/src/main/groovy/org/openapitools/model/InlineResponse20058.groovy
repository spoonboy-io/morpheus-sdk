package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InstanceSchedule;
import org.openapitools.model.InstanceThreshold;

@Canonical
class InlineResponse20058 {
    
    InlineResponse20040AppDeployInstance instance
    
    InstanceThreshold instanceThreshold
    
    List<InstanceSchedule> instanceSchedules = new ArrayList<InstanceSchedule>()
}
