package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class MetaMeta {
    /* Offset records, the number of records to skip, can be used to paginate over results. */
    Long offset = 0l
    /* Max size, the maximum number of records to include in the response. */
    Long max = 25l
    /* Number of records returned in the response */
    Long size = 0l
    /* Total number of records found */
    Long total = 0l
}
