package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.WhitelabelSettingsUpdateSupportMenuLinks;

@Canonical
class WhitelabelSettingsUpdate {
    /* Can be used to enable / disable whitelabel feature */
    Boolean enabled
    /* Appliance name. Master account only */
    String applianceName
    /* Can be used to disable support menu */
    Boolean disableSupportMenu
    /* Resets header logo to default header logo */
    Boolean resetHeaderLogo
    /* Resets footer logo to default footer logo */
    Boolean resetFooterLogo
    /* Resets login logo to default login logo */
    Boolean resetLoginLogo
    /* Resets favicon to default favicon */
    Boolean resetFavicon
    /* Header background color */
    String headerBgColor
    /* Header foreground color */
    String headerFgColor
    /* Nav background color */
    String navBgColor
    /* Nav foreground color */
    String navFgColor
    /* Nav hover color */
    String navHoverColor
    /* Primary button background color */
    String primaryButtonBgColor
    /* Primary button foreground color */
    String primaryButtonFgColor
    /* Primary button hover background color */
    String primaryButtonHoverBgColor
    /* Primary button hover foreground color */
    String primaryButtonHoverFgColor
    /* Footer background color */
    String footerBgColor
    /* Footer foreground color */
    String footerFgColor
    /* Login background color */
    String loginBgColor
    /* Copyright String */
    String copyrightString
    /* Override CSS */
    String overrideCss
    /* Terms of use content */
    String termsOfUse
    /* Privacy policy content */
    String privacyPolicy
    
    List<WhitelabelSettingsUpdateSupportMenuLinks> supportMenuLinks = new ArrayList<WhitelabelSettingsUpdateSupportMenuLinks>()
}
