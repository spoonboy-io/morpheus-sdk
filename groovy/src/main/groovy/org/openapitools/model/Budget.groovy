package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.model.BudgetStats;

@Canonical
class Budget {
    
    Long id
    
    String name
    
    String description
    
    ApplianceSettingsEnabledZoneTypesInner account
    
    Boolean enabled
    
    String refScope
    
    String refType
    
    Long refId
    
    String refName
    
    String period
    
    String year
    
    String resourceType
    
    String timezone
    
    Date startDate
    
    Date endDate
    
    String interval
    
    List<Long> costs
    
    Boolean isFiscal
    
    Long averageCost
    
    Long totalCost
    
    String currency
    
    Boolean rollover
    
    String warningLimit
    
    String overLimit
    
    String externalId
    
    String internalId
    
    Long createdById
    
    String createdByName
    
    String updatedById
    
    String updatedByName
    
    Date dateCreated
    
    Date lastUpdated
    
    BudgetStats stats
}
