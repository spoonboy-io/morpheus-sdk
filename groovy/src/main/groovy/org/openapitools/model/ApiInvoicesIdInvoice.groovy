package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiInvoicesIdInvoice {
    /* This adds or updates the specified Metadata tags and removes any tags not specified. Array of objects having a name and value.  */
    List<Object> tags = new ArrayList<Object>()
    /* Add or update value of Metadata tags. Array of objects having a name and value.  */
    List<Object> addTags = new ArrayList<Object>()
    /* This removes the specified Metadata tags matching name and optionally value (if provided). Array of objects having a name and value.  */
    List<Object> removeTags = new ArrayList<Object>()
}
