package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ClusterVolumes {
    
    Long id
    
    String internalId
    
    Long displayOrder
    
    Boolean active
    
    Long usedStorage
    
    String provisionType
    
    Boolean resizeable
    
    Boolean online
    
    String deviceDisplayName
    
    String refType
    
    String name
    
    String externalId
    
    String datastoreOption
    
    String claimName
    
    String volumeType
    
    String deviceName
    
    Boolean removable
    
    String poolName
    
    Boolean readOnly
    
    String sourceId
    
    Long zoneId
    
    Boolean rootVolume
    
    Long refId
    
    String category
    
    String status
    
    String rawData
    
    Long maxStorage
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites type
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites datastore
}
