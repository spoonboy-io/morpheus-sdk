package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.WhitelabelSettingsSupportMenuLinks;

@Canonical
class WhitelabelSettings {
    
    Boolean enabled
    
    String applianceName
    
    Boolean disableSupportMenu
    
    String headerLogo
    
    String footerLogo
    
    String loginLogo
    
    String favicon
    
    String headerBgColor
    
    String headerFgColor
    
    String navBgColor
    
    String navFgColor
    
    String navHoverColor
    
    String primaryButtonBgColor
    
    String primaryButtonFgColor
    
    String primaryButtonHoverBgColor
    
    String primaryButtonHoverFgColor
    
    String footerBgColor
    
    String footerFgColor
    
    String loginBgColor
    
    String overrideCss
    
    String copyrightString
    
    String termsOfUse
    
    String privacyPolicy
    
    List<WhitelabelSettingsSupportMenuLinks> supportMenuLinks = new ArrayList<WhitelabelSettingsSupportMenuLinks>()
}
