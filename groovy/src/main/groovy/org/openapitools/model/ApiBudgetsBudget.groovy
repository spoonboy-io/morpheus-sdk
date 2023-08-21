package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiBudgetsBudget {
    
    String name
    
    String description
    
    String scope = ScopeEnum.ACCOUNT
    
    String period = "year"
    
    Long year
    
    Date startDate
    
    Date endDate
    
    String interval = IntervalEnum.YEAR
    /* The Tenant ID to scope to, for use with ``scope``=tenant  */
    Long scopeTenantId
    /* The Tenant ID to scope to, for use with ``scope``=group   */
    Long scopeGroupId
    /* The Tenant ID to scope to, for use with ``scope``=cloud  */
    Long scopeCloudId
    /* The Tenant ID to scope to, for use with ``scope``=user  */
    Long scopeUserId
    
    List<Long> costs = new ArrayList<Long>()
    
    Boolean enabled = true
}
