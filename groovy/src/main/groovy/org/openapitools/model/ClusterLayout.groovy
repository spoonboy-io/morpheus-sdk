package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutComputeServers;
import org.openapitools.model.ClusterLayoutSpecTemplates;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20094Network;
import org.openapitools.model.OptionType;

@Canonical
class ClusterLayout {
    
    Long id
    
    String internalId
    
    Long serverCount
    
    Date dateCreated
    
    String code
    
    Date lastUpdated
    
    Boolean hasAutoScale
    
    Long memoryRequirement
    
    String clusterVersion
    
    String computeVersion
    
    Boolean hasSettings
    
    Long sortOrder
    
    Boolean hasConfig
    
    String name
    
    Boolean creatable
    
    Boolean enabled
    
    String description
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType groupType
    
    List<String> labels = new ArrayList<String>()
    
    List<Object> environmentVariables = new ArrayList<Object>()
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<Object> actions = new ArrayList<Object>()
    
    List<ClusterLayoutComputeServers> computeServers = new ArrayList<ClusterLayoutComputeServers>()
    
    Boolean installContainerRuntime
    
    InlineResponse20094Network provisionType
    
    List<ClusterLayoutSpecTemplates> specTemplates = new ArrayList<ClusterLayoutSpecTemplates>()
    
    List<Object> taskSets = new ArrayList<Object>()
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
}
