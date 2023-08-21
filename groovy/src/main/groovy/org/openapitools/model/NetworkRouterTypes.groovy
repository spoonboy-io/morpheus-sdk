package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.NetworkRouterTypesOptionTypes;

@Canonical
class NetworkRouterTypes {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Boolean enabled
    
    Boolean creatable
    
    Boolean selectable
    
    Boolean hasFirewall
    
    Boolean hasDhcp
    
    Boolean hasRouting
    
    Boolean hasNetworkServer
    
    List<NetworkRouterTypesOptionTypes> optionTypes = new ArrayList<NetworkRouterTypesOptionTypes>()
    
    List<Object> ruleOptionTypes = new ArrayList<Object>()
    
    List<Object> natOptionTypes = new ArrayList<Object>()
    
    List<Object> ruleGroupOptionTypes = new ArrayList<Object>()
}
