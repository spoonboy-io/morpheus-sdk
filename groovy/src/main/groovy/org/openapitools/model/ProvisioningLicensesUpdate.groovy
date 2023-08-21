package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ProvisioningLicensesUpdate {
    /* Name */
    String name
    /* License Version */
    String licenseVersion
    /* Copies - The number of times the key can be used. */
    Long copies = 1l
    /* Description */
    String description
    /* Virtual Images - Array of Virtual Image IDs to associate with license. */
    List<Long> virtualImages = new ArrayList<Long>()
    /* Tenants - Array of tenants that are allowed to use the key. */
    List<Long> tenants = new ArrayList<Long>()
}
