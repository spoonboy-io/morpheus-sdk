package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.InstanceConfigBackup;
import org.openapitools.model.InstanceConfigInstanceType;
import org.openapitools.model.InstanceConfigLayout;
import org.openapitools.model.InstanceConfigReplicationGroup;
import org.openapitools.model.InstanceConfigSecurityGroups;

@Canonical
class InstanceConfig {
    
    Boolean createUser
    
    Boolean isEC2
    
    Boolean isVpcSelectable
    
    Boolean noAgent
    
    List<InstanceConfigSecurityGroups> securityGroups = new ArrayList<InstanceConfigSecurityGroups>()
    
    String smbiosAssetTag
    
    String nestedVirtualization
    
    String vmwareFolderId
    
    Object customOptions
    
    Long resourcePoolId
    
    String poolProviderType
    
    InstanceConfigSecurityGroups userGroup
    
    String expireDays
    
    String shutdownDays
    
    String name
    
    String hostName
    
    InstanceConfigInstanceType instanceType
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites site
    
    String environmentPrefix
    
    InstanceConfigLayout layout
    
    String type
    
    String instanceContext
    
    String memoryDisplay
    
    List<Object> expose = new ArrayList<Object>()
    
    Boolean createBackup
    
    InstanceConfigBackup backup
    
    InstanceConfigReplicationGroup replicationGroup
    
    Long layoutSize
    
    List<Object> lbInstances = new ArrayList<Object>()
}
