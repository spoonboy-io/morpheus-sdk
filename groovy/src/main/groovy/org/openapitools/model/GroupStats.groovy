package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.GroupStatsInstanceCounts;
import org.openapitools.model.ZoneStatsServerCounts;

@Canonical
class GroupStats {
    
    GroupStatsInstanceCounts instanceCounts
    
    ZoneStatsServerCounts serverCounts
}
