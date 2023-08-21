package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class VdiPoolConfigStorageControllers {
    
    Long id
    
    String name
    
    Boolean active
    
    Long typeId
    
    String typeName
    
    String unitNumber
    
    String busNumber
    
    BigDecimal maxDevices
    
    Boolean removable
    
    Boolean editable
    
    BigDecimal reservedUnitNumber
    
    String category
    
    BigDecimal displayOrder
}
