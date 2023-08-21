package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AppBlueprint;
import org.openapitools.model.CatalogItemInstance;

@Canonical
class CatalogItem {
    
    Long id
    
    String name
    
    AppBlueprint type
    
    Long quantity
    
    String status
    
    String statusMessage
    
    String refType
    
    CatalogItemInstance instance
    
    Date orderDate
    
    Date dateCreated
    
    Date lastUpdated
}
