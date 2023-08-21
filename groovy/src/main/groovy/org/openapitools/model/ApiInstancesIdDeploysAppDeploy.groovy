package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiInstancesIdDeploysAppDeploy {
    /* Deployment ID. */
    Long deploymentId
    /* Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version */
    String version
    /* Deployment Version ID. This can be passed instead of deploymentId and version. */
    Long versionId
    /* Map of configuration properties that vary by instance type. */
    Object config
    /* Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. */
    Boolean stageOnly = false
}
