package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.HealthElasticMaster;
import org.openapitools.model.HealthElasticNodes;
import org.openapitools.model.HealthElasticStats;

@Canonical
class HealthElastic {
    
    Boolean success
    
    String status
    
    HealthElasticMaster master
    
    List<HealthElasticNodes> nodes = new ArrayList<HealthElasticNodes>()
    
    HealthElasticStats stats
    
    List<Object> indices = new ArrayList<Object>()
    
    List<Object> badIndices = new ArrayList<Object>()
}
