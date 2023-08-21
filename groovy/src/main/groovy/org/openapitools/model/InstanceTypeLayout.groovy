package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ContainerType;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.InlineResponse20094Network;
import org.openapitools.model.InstanceTypeLayoutPermissions;
import org.openapitools.model.ProvisionType;

@Canonical
class InstanceTypeLayout {
    
    Long id
    
    InlineResponse20094Network instanceType
    
    InlineResponse20082LoadBalancerInstanceSslCert account
    
    String code
    
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    String instanceVersion
    
    String description
    
    Boolean creatable
    
    Long memoryRequirement
    
    Long sortOrder
    
    Boolean supportsConvertToManaged
    
    ProvisionType provisionType
    
    List<Object> taskSets = new ArrayList<Object>()
    
    List<ContainerType> containerTypes = new ArrayList<ContainerType>()
    
    List<Object> mounts = new ArrayList<Object>()
    
    List<Object> ports = new ArrayList<Object>()
    
    List<Object> optionTypes = new ArrayList<Object>()
    
    List<Object> environmentVariables = new ArrayList<Object>()
    
    List<Object> priceSets = new ArrayList<Object>()
    
    List<Object> specTemplates = new ArrayList<Object>()
    
    List<Object> tfvarSecret = new ArrayList<Object>()
    
    InstanceTypeLayoutPermissions permissions
}
