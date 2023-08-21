package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CatalogItemTypeInstanceScribePorts {
    /* Port number. */
    Long port
    /* A name for the port. */
    String name
    /* The load balancer protocol. HTTP, HTTPS, or TCP. */
    String lb
}
