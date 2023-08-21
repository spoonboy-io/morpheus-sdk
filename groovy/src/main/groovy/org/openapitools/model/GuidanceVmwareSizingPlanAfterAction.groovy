package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionConfig;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionPriceSets;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionProvisionType;

@Canonical
class GuidanceVmwareSizingPlanAfterAction {
    
    Long id
    
    String name
    
    String code
    
    Boolean active
    
    Long sortOrder
    
    String description
    
    Long maxStorage
    
    Long maxMemory
    
    Long maxCpu
    
    Long maxCores
    
    String maxDisks
    
    Long coresPerSocket
    
    Boolean customCpu
    
    Boolean customCores
    
    Boolean customMaxStorage
    
    Boolean customMaxDataStorage
    
    Boolean customMaxMemory
    
    Boolean addVolumes
    
    String memoryOptionSource
    
    String cpuOptionSource
    
    Date dateCreated
    
    Date lastUpdated
    
    String regionCode
    
    String visibility
    
    Boolean editable
    
    GuidanceVmwareSizingPlanBeforeActionProvisionType provisionType
    
    String tenants
    
    List<GuidanceVmwareSizingPlanBeforeActionPriceSets> priceSets = new ArrayList<GuidanceVmwareSizingPlanBeforeActionPriceSets>()
    
    GuidanceVmwareSizingPlanBeforeActionConfig config
}
