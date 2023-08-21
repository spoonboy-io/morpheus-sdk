package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.SnapshotSnapshotZone;

@Canonical
class SnapshotSnapshot {
    
    Integer id
    
    String name
    
    String description
    
    String externalId
    
    String status
    
    String state
    
    String snapshotType
    
    Date snapshotCreated
    
    SnapshotSnapshotZone zone
    
    String datastore
    
    String parentSnapshot
    
    Boolean currentlyActive
    
    Date dateCreated
}
