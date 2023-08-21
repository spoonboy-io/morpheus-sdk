package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InstanceBackupsInstance;

@Canonical
class InstanceBackups {
    
    InstanceBackupsInstance instance
    /* List of backup objects */
    List<Object> backups = new ArrayList<Object>()
}
