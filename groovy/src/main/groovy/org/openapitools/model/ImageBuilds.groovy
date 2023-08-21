package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ImageBuildsBootScript;
import org.openapitools.model.ImageBuildsConfig;
import org.openapitools.model.ImageBuildsLastResult;
import org.openapitools.model.ImageBuildsScripts;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class ImageBuilds {
    
    Long id
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
    
    InlineResponse20040AppDeployInstance site
    
    InlineResponse20040AppDeployInstance zone
    
    String name
    
    String description
    
    ImageBuildsBootScript bootScript
    
    String bootCommand
    
    ImageBuildsBootScript preseedScript
    
    List<ImageBuildsScripts> scripts = new ArrayList<ImageBuildsScripts>()
    
    String sshUsername
    
    String sshPassword
    
    String storageProvider
    
    String buildOutputName
    
    String conversionFormats
    
    Boolean isCloudInit
    
    Boolean vmToolsInstalled
    
    Long keepResults
    
    ImageBuildsConfig config
    
    ImageBuildsLastResult lastResult
    
    Long executionCount
}
