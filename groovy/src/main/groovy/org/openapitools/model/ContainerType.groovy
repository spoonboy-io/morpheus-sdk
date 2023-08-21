package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ContainerTypeAccount;
import org.openapitools.model.ContainerTypeContainerPorts;
import org.openapitools.model.ContainerTypeProvisionType;

@Canonical
class ContainerType {
    
    Integer id
    
    ContainerTypeAccount account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String shortName
    
    String code
    
    String containerVersion
    
    ContainerTypeProvisionType provisionType
    
    ContainerTypeAccount virtualImage
    
    String category
    
    Object config
    
    List<ContainerTypeContainerPorts> containerPorts = new ArrayList<ContainerTypeContainerPorts>()
    
    List<Object> containerScripts = new ArrayList<Object>()
    
    List<Object> containerTemplates = new ArrayList<Object>()
    
    List<Object> environmentVariables = new ArrayList<Object>()
}
