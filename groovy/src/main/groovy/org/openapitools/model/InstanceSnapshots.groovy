package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.Snapshot;

@Canonical
class InstanceSnapshots {
    /* List of snapshot objects */
    List<Snapshot> snapshots = new ArrayList<Snapshot>()
}
