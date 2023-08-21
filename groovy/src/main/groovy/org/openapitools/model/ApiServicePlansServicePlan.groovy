package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServicePlansServicePlanConfig;
import org.openapitools.model.ApiServicePlansServicePlanProvisionType;

@Canonical
class ApiServicePlansServicePlan {
    /* Service plan name */
    String name
    /* Service plan code, must be unique */
    String code
    /* Service plan description */
    String description
    /* Can be used to enable / disable the editability of the service plan. */
    Boolean editable = true
    /* Max storage size in bytes */
    BigDecimal maxStorage
    /* Max memory size in bytes */
    BigDecimal maxMemory
    /* Max cores */
    BigDecimal maxCores
    /* Max disks allowed */
    BigDecimal maxDisks
    
    List<ApiServicePlansServicePlanProvisionType> provisionType = new ArrayList<ApiServicePlansServicePlanProvisionType>()
    /* Can be used to enable / disable customizable cores */
    Boolean customCores = false
    /* Can be used to enable / disable customizable storage */
    Boolean customMaxStorage = false
    /* Can be used to enable / disable customizable extra volumes. */
    Boolean customMaxDataStorage = false
    /* Can be used to enable / disable customizable memory. */
    Boolean customMaxMemory = false
    /* Can be used to enable / disable ability to add volumes */
    Boolean addVolumes = false
    /* Sort order */
    BigDecimal sortOrder
    /* List of price sets to include in service plan */
    List<Long> priceSets = new ArrayList<Long>()
    
    ApiServicePlansServicePlanConfig config
}
