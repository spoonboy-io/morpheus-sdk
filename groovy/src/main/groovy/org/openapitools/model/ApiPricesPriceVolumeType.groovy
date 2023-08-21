package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiPricesPriceVolumeType {
    /* Volume type ID, required for `storage` price type. The endpoint /api/prices/volume-types provides a list of available volume type options.  */
    Long id
}
