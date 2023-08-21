package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AnyOfstringlong;

@Canonical
class InstanceCreateVolume {
    /* The id for the LV configuration being created. */
    Long id = -1l
    /* If set to false then a non-root LV will be created. */
    Boolean rootVolume = true
    /* Name/type of the LV being created. */
    String name = "root"
    /* Size of the LV to be created in GBs.  Uses default from service plan. */
    Long size
    /* Can be used to select pre-existing LV choices from Morpheus. */
    Long sizeId
    /* Identifier for LV type */
    Long storageType
    /* The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters). */
    AnyOfstringlong datastoreId = null
}
