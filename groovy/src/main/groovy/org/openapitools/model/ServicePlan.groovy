package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionPriceSets;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionProvisionType;
import org.openapitools.model.InlineResponse20094Network;
import org.openapitools.model.ServicePlanConfig;
import org.openapitools.model.ServicePlanPermissions;

@Canonical
class ServicePlan {
    
    Long id
    
    String name
    
    String code
    
    Boolean active
    
    Long sortOrder
    
    String description
    
    BigDecimal maxStorage
    
    BigDecimal maxMemory
    
    BigDecimal maxCpu
    
    BigDecimal maxCores
    
    BigDecimal maxDisks
    
    BigDecimal coresPerSocket
    
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
    
    ServicePlanConfig config
    
    List<InlineResponse20094Network> zones = new ArrayList<InlineResponse20094Network>()
    
    ServicePlanPermissions permissions
}
