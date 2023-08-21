package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HealthDatabaseSlowQueries {
    
    Long count
    
    Long time
    
    String query
}
