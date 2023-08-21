package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.UserSettingsAccessTokens;
import org.openapitools.model.UserSettingsUser;

@Canonical
class UserSettings {
    
    UserSettingsUser user
    
    List<UserSettingsAccessTokens> accessTokens = new ArrayList<UserSettingsAccessTokens>()
}
