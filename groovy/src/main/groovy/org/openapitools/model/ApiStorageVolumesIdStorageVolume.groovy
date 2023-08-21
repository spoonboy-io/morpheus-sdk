package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiStorageVolumesStorageVolumeStorageServer;

@Canonical
class ApiStorageVolumesIdStorageVolume {
    /* A unique name scoped to your account for the storage volume */
    String name
    /* Storage Type Code or ID */
    String type
    /* Configuration object with parameters that vary by `type`. */
    Object config
    
    ApiStorageVolumesStorageVolumeStorageServer storageServer
    
    ApiStorageVolumesStorageVolumeStorageServer storageGroup
}
