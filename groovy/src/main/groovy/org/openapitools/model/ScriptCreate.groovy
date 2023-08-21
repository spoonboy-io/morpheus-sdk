package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ScriptCreate {
    /* Script name */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Script category */
    String category
    /* Version of the script */
    String scriptVersion = "1"
    /* Phase for the script, provision, start, etc. */
    String scriptPhase
    /* Type for the script */
    String scriptType = ScriptTypeEnum.BASH
    /* Script content, that is, the code itself. */
    String script
    /* Run as a specific user. */
    String runAsUser
    /* Sudo, whether or not to run with sudo. */
    Boolean sudoUser = false
}
