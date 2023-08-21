package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class ClusterLayoutContainerType {
    
    Long id
    
    String account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String shortName
    
    String code
    
    String containerVersion
    
    InlineResponse20094Network provisionType
    
    String virtualImage
    
    String category
    
    Object config
    
    List<Object> containerPorts = new ArrayList<Object>()
    
    List<Object> containerScripts = new ArrayList<Object>()
    
    List<Object> containerTemplates = new ArrayList<Object>()
    
    List<Object> environmentVariables = new ArrayList<Object>()
}
