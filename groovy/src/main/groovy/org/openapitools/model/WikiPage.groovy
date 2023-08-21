package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;

@Canonical
class WikiPage {
    
    Long id
    
    String name
    
    String urlName
    
    String category
    
    String refId
    
    String refType
    
    String format
    
    String content
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    InlineResponse200107NetworkPoolCreatedBy updatedBy
    
    Date dateCreated
    
    Date lastUpdated
}
