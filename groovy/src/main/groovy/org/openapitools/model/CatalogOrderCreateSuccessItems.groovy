package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.AppBlueprint;

@Canonical
class CatalogOrderCreateSuccessItems {
    
    Long id
    
    AppBlueprint type
    
    Long quantity
    
    BigDecimal price
    
    String currency
    
    String unit
    
    Boolean valid
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
}
