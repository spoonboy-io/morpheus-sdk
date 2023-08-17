package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;

@Canonical
class AddBudgetsRequestBudget {
    
    String name
    
    String description

    enum ScopeEnum {
    
        ACCOUNT("account"),
        
        GROUP("group"),
        
        CLOUD("cloud"),
        
        USER("user")
    
        private final String value
    
        ScopeEnum(String value) {
            this.value = value
        }
    
        String getValue() {
            value
        }
    
        @Override
        String toString() {
            String.valueOf(value)
        }
    }

    
    ScopeEnum scope = ScopeEnum.ACCOUNT
    
    String period = "year"
    
    Long year
    
    Date startDate
    
    Date endDate

    enum IntervalEnum {
    
        YEAR("year"),
        
        QUARTER("quarter"),
        
        MONTH("month")
    
        private final String value
    
        IntervalEnum(String value) {
            this.value = value
        }
    
        String getValue() {
            value
        }
    
        @Override
        String toString() {
            String.valueOf(value)
        }
    }

    
    IntervalEnum interval = IntervalEnum.YEAR
    /* The Tenant ID to scope to, for use with ``scope``=tenant  */
    Long scopeTenantId
    /* The Tenant ID to scope to, for use with ``scope``=group   */
    Long scopeGroupId
    /* The Tenant ID to scope to, for use with ``scope``=cloud  */
    Long scopeCloudId
    /* The Tenant ID to scope to, for use with ``scope``=user  */
    Long scopeUserId
    
    List<Long> costs
    
    Boolean enabled = true
}
