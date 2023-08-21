package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiSecurityPackagesSecurityPackage {
    /* A name for the security package */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Security Package Type Code */
    String type = "scap-package"
    /* A description for the security package */
    String description
    /* URL to download the security package zip file from */
    String url
    /* Can be used to disable the security package */
    Boolean enabled = true
}
