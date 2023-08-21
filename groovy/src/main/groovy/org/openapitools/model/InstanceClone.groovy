package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceCloneGroup;

@Canonical
class InstanceClone {
    /* A name for the new cloned instance. If none is specified the existing name will be duplicated with the 'clone' suffix added. */
    String name
    
    InstanceCloneGroup group
}
