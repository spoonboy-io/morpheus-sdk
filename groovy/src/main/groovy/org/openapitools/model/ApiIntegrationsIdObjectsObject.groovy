package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiIntegrationsIdObjectsObject {
    /* Name to display */
    String name
    /* Integration Object Type Code */
    String type
    /* Catalog Item Type ID */
    Long catalogItemType
}
