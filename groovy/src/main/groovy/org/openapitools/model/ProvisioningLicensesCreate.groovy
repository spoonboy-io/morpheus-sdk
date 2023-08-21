package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ProvisioningLicensesCreate {
    /* Name */
    String name
    /* License Type - The license type code. */
    String licenseType
    /* License Key - The license key, to be kept a secret. */
    String licenseKey
    /* Org Name - The Organization Name (if applicable) related to the license key */
    String orgName
    /* Full Name - The Full Name (if applicable) related to the license key */
    String fullName
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
