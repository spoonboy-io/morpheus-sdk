package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.ExecuteSchedule;
import org.openapitools.model.MetaMeta;

@Canonical
class ListExecuteSchedules200Response {
    
    MetaMeta meta
    
    List<ExecuteSchedule> schedules
}
