package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject;

@Canonical
class PolicyUpdate {
    /* A name for the policy */
    String name
    /* A description for the policy */
    String description
    /* A map of config values. The expected values vary by policy type. See `Retrieves all Policy Types` endpoint for `fieldName`(s) of required options. */
    AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config = null
    /* Set to false to disable */
    Boolean enabled = true
    /* Scope object type */
    String refType
    /* Scope object ID (`group`,`cloud`,`user`, etc) */
    Long refId
    /* Array of tenants to scope the policy to */
    List<Long> accounts = new ArrayList<Long>()
    /* Apply individually to each user in role.  Only when `refType` equals `Role` */
    Boolean eachUser
}
