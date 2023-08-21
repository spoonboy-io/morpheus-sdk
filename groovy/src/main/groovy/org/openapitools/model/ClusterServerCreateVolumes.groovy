package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterServerCreateVolumes {
    /* The id for the LV configuration being created */
    Long id = -1l
    /* If set to false then a non-root LV will be created */
    Boolean rootVolume = true
    /* Name/type of the LV being created */
    String name = "root"
    /* Size of the LV to be created in GBs  Default is from the service plan  */
    Long size
    /* Can be used to select pre-existing LV choices from Morpheus */
    String sizeId
    /* Identifier for LV type */
    Long storageType
    /* The ID of the specific datastore. Auto selection can be specified as auto or `autoCluster` (for clusters). */
    String datastoreId
}
