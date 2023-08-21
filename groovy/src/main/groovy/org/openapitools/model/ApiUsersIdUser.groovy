package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ReferenceObject;

@Canonical
class ApiUsersIdUser {
    /* First Name */
    String firstName
    /* Last Name */
    String lastName
    /* Username (unique per tenant). */
    String username
    /* Email address */
    String email
    /* Password */
    String password
    /* List of Roles */
    List<ReferenceObject> roles = new ArrayList<ReferenceObject>()
}
