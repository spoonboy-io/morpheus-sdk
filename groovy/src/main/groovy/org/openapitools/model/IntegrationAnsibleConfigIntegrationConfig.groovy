package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationAnsibleConfigIntegrationConfig {
    /* default branch */
    String defaultBranch
    /* Playbooks path */
    String ansiblePlaybooks
    /* Roles path */
    String ansibleRoles
    /* Group variables path */
    String ansibleGroupVars
    /* Host variables path */
    String ansibleHostVars
    /* Use Ansible Galaxy */
    Boolean ansibleGalaxyEnabled
    /* Use verbose logging */
    Boolean ansibleVerbose
    /* Use Morpheus Agent Command Bus */
    Boolean ansibleCommandBus
    /* Enable Git repository caching */
    Boolean cacheEnabled
}
