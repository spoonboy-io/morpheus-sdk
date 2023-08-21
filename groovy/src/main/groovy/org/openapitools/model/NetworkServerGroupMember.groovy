package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkServerGroupMember {
    
    Long id
    
    String category
    
    String type
    
    String memberName
    
    String memberType
    
    String memberValue
    
    String memberExpression
    
    Long displayOrder
    
    String internalId
    
    String externalId
    
    List<Object> members = new ArrayList<Object>()
}
