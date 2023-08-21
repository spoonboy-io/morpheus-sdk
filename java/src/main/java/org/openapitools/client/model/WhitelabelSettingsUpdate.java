/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.WhitelabelSettingsUpdateSupportMenuLinks;

/**
 * WhitelabelSettingsUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class WhitelabelSettingsUpdate {
  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_APPLIANCE_NAME = "applianceName";
  @SerializedName(SERIALIZED_NAME_APPLIANCE_NAME)
  private String applianceName;

  public static final String SERIALIZED_NAME_DISABLE_SUPPORT_MENU = "disableSupportMenu";
  @SerializedName(SERIALIZED_NAME_DISABLE_SUPPORT_MENU)
  private Boolean disableSupportMenu;

  public static final String SERIALIZED_NAME_RESET_HEADER_LOGO = "resetHeaderLogo";
  @SerializedName(SERIALIZED_NAME_RESET_HEADER_LOGO)
  private Boolean resetHeaderLogo;

  public static final String SERIALIZED_NAME_RESET_FOOTER_LOGO = "resetFooterLogo";
  @SerializedName(SERIALIZED_NAME_RESET_FOOTER_LOGO)
  private Boolean resetFooterLogo;

  public static final String SERIALIZED_NAME_RESET_LOGIN_LOGO = "resetLoginLogo";
  @SerializedName(SERIALIZED_NAME_RESET_LOGIN_LOGO)
  private Boolean resetLoginLogo;

  public static final String SERIALIZED_NAME_RESET_FAVICON = "resetFavicon";
  @SerializedName(SERIALIZED_NAME_RESET_FAVICON)
  private Boolean resetFavicon;

  public static final String SERIALIZED_NAME_HEADER_BG_COLOR = "headerBgColor";
  @SerializedName(SERIALIZED_NAME_HEADER_BG_COLOR)
  private String headerBgColor;

  public static final String SERIALIZED_NAME_HEADER_FG_COLOR = "headerFgColor";
  @SerializedName(SERIALIZED_NAME_HEADER_FG_COLOR)
  private String headerFgColor;

  public static final String SERIALIZED_NAME_NAV_BG_COLOR = "navBgColor";
  @SerializedName(SERIALIZED_NAME_NAV_BG_COLOR)
  private String navBgColor;

  public static final String SERIALIZED_NAME_NAV_FG_COLOR = "navFgColor";
  @SerializedName(SERIALIZED_NAME_NAV_FG_COLOR)
  private String navFgColor;

  public static final String SERIALIZED_NAME_NAV_HOVER_COLOR = "navHoverColor";
  @SerializedName(SERIALIZED_NAME_NAV_HOVER_COLOR)
  private String navHoverColor;

  public static final String SERIALIZED_NAME_PRIMARY_BUTTON_BG_COLOR = "primaryButtonBgColor";
  @SerializedName(SERIALIZED_NAME_PRIMARY_BUTTON_BG_COLOR)
  private String primaryButtonBgColor;

  public static final String SERIALIZED_NAME_PRIMARY_BUTTON_FG_COLOR = "primaryButtonFgColor";
  @SerializedName(SERIALIZED_NAME_PRIMARY_BUTTON_FG_COLOR)
  private String primaryButtonFgColor;

  public static final String SERIALIZED_NAME_PRIMARY_BUTTON_HOVER_BG_COLOR = "primaryButtonHoverBgColor";
  @SerializedName(SERIALIZED_NAME_PRIMARY_BUTTON_HOVER_BG_COLOR)
  private String primaryButtonHoverBgColor;

  public static final String SERIALIZED_NAME_PRIMARY_BUTTON_HOVER_FG_COLOR = "primaryButtonHoverFgColor";
  @SerializedName(SERIALIZED_NAME_PRIMARY_BUTTON_HOVER_FG_COLOR)
  private String primaryButtonHoverFgColor;

  public static final String SERIALIZED_NAME_FOOTER_BG_COLOR = "footerBgColor";
  @SerializedName(SERIALIZED_NAME_FOOTER_BG_COLOR)
  private String footerBgColor;

  public static final String SERIALIZED_NAME_FOOTER_FG_COLOR = "footerFgColor";
  @SerializedName(SERIALIZED_NAME_FOOTER_FG_COLOR)
  private String footerFgColor;

  public static final String SERIALIZED_NAME_LOGIN_BG_COLOR = "loginBgColor";
  @SerializedName(SERIALIZED_NAME_LOGIN_BG_COLOR)
  private String loginBgColor;

  public static final String SERIALIZED_NAME_COPYRIGHT_STRING = "copyrightString";
  @SerializedName(SERIALIZED_NAME_COPYRIGHT_STRING)
  private String copyrightString;

  public static final String SERIALIZED_NAME_OVERRIDE_CSS = "overrideCss";
  @SerializedName(SERIALIZED_NAME_OVERRIDE_CSS)
  private String overrideCss;

  public static final String SERIALIZED_NAME_TERMS_OF_USE = "termsOfUse";
  @SerializedName(SERIALIZED_NAME_TERMS_OF_USE)
  private String termsOfUse;

  public static final String SERIALIZED_NAME_PRIVACY_POLICY = "privacyPolicy";
  @SerializedName(SERIALIZED_NAME_PRIVACY_POLICY)
  private String privacyPolicy;

  public static final String SERIALIZED_NAME_SUPPORT_MENU_LINKS = "supportMenuLinks";
  @SerializedName(SERIALIZED_NAME_SUPPORT_MENU_LINKS)
  private List<WhitelabelSettingsUpdateSupportMenuLinks> supportMenuLinks = null;


  public WhitelabelSettingsUpdate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to enable / disable whitelabel feature
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to enable / disable whitelabel feature")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public WhitelabelSettingsUpdate applianceName(String applianceName) {
    
    this.applianceName = applianceName;
    return this;
  }

   /**
   * Appliance name. Master account only
   * @return applianceName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Appliance name. Master account only")

  public String getApplianceName() {
    return applianceName;
  }


  public void setApplianceName(String applianceName) {
    this.applianceName = applianceName;
  }


  public WhitelabelSettingsUpdate disableSupportMenu(Boolean disableSupportMenu) {
    
    this.disableSupportMenu = disableSupportMenu;
    return this;
  }

   /**
   * Can be used to disable support menu
   * @return disableSupportMenu
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to disable support menu")

  public Boolean getDisableSupportMenu() {
    return disableSupportMenu;
  }


  public void setDisableSupportMenu(Boolean disableSupportMenu) {
    this.disableSupportMenu = disableSupportMenu;
  }


  public WhitelabelSettingsUpdate resetHeaderLogo(Boolean resetHeaderLogo) {
    
    this.resetHeaderLogo = resetHeaderLogo;
    return this;
  }

   /**
   * Resets header logo to default header logo
   * @return resetHeaderLogo
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Resets header logo to default header logo")

  public Boolean getResetHeaderLogo() {
    return resetHeaderLogo;
  }


  public void setResetHeaderLogo(Boolean resetHeaderLogo) {
    this.resetHeaderLogo = resetHeaderLogo;
  }


  public WhitelabelSettingsUpdate resetFooterLogo(Boolean resetFooterLogo) {
    
    this.resetFooterLogo = resetFooterLogo;
    return this;
  }

   /**
   * Resets footer logo to default footer logo
   * @return resetFooterLogo
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Resets footer logo to default footer logo")

  public Boolean getResetFooterLogo() {
    return resetFooterLogo;
  }


  public void setResetFooterLogo(Boolean resetFooterLogo) {
    this.resetFooterLogo = resetFooterLogo;
  }


  public WhitelabelSettingsUpdate resetLoginLogo(Boolean resetLoginLogo) {
    
    this.resetLoginLogo = resetLoginLogo;
    return this;
  }

   /**
   * Resets login logo to default login logo
   * @return resetLoginLogo
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Resets login logo to default login logo")

  public Boolean getResetLoginLogo() {
    return resetLoginLogo;
  }


  public void setResetLoginLogo(Boolean resetLoginLogo) {
    this.resetLoginLogo = resetLoginLogo;
  }


  public WhitelabelSettingsUpdate resetFavicon(Boolean resetFavicon) {
    
    this.resetFavicon = resetFavicon;
    return this;
  }

   /**
   * Resets favicon to default favicon
   * @return resetFavicon
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Resets favicon to default favicon")

  public Boolean getResetFavicon() {
    return resetFavicon;
  }


  public void setResetFavicon(Boolean resetFavicon) {
    this.resetFavicon = resetFavicon;
  }


  public WhitelabelSettingsUpdate headerBgColor(String headerBgColor) {
    
    this.headerBgColor = headerBgColor;
    return this;
  }

   /**
   * Header background color
   * @return headerBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Header background color")

  public String getHeaderBgColor() {
    return headerBgColor;
  }


  public void setHeaderBgColor(String headerBgColor) {
    this.headerBgColor = headerBgColor;
  }


  public WhitelabelSettingsUpdate headerFgColor(String headerFgColor) {
    
    this.headerFgColor = headerFgColor;
    return this;
  }

   /**
   * Header foreground color
   * @return headerFgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Header foreground color")

  public String getHeaderFgColor() {
    return headerFgColor;
  }


  public void setHeaderFgColor(String headerFgColor) {
    this.headerFgColor = headerFgColor;
  }


  public WhitelabelSettingsUpdate navBgColor(String navBgColor) {
    
    this.navBgColor = navBgColor;
    return this;
  }

   /**
   * Nav background color
   * @return navBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Nav background color")

  public String getNavBgColor() {
    return navBgColor;
  }


  public void setNavBgColor(String navBgColor) {
    this.navBgColor = navBgColor;
  }


  public WhitelabelSettingsUpdate navFgColor(String navFgColor) {
    
    this.navFgColor = navFgColor;
    return this;
  }

   /**
   * Nav foreground color
   * @return navFgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Nav foreground color")

  public String getNavFgColor() {
    return navFgColor;
  }


  public void setNavFgColor(String navFgColor) {
    this.navFgColor = navFgColor;
  }


  public WhitelabelSettingsUpdate navHoverColor(String navHoverColor) {
    
    this.navHoverColor = navHoverColor;
    return this;
  }

   /**
   * Nav hover color
   * @return navHoverColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Nav hover color")

  public String getNavHoverColor() {
    return navHoverColor;
  }


  public void setNavHoverColor(String navHoverColor) {
    this.navHoverColor = navHoverColor;
  }


  public WhitelabelSettingsUpdate primaryButtonBgColor(String primaryButtonBgColor) {
    
    this.primaryButtonBgColor = primaryButtonBgColor;
    return this;
  }

   /**
   * Primary button background color
   * @return primaryButtonBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Primary button background color")

  public String getPrimaryButtonBgColor() {
    return primaryButtonBgColor;
  }


  public void setPrimaryButtonBgColor(String primaryButtonBgColor) {
    this.primaryButtonBgColor = primaryButtonBgColor;
  }


  public WhitelabelSettingsUpdate primaryButtonFgColor(String primaryButtonFgColor) {
    
    this.primaryButtonFgColor = primaryButtonFgColor;
    return this;
  }

   /**
   * Primary button foreground color
   * @return primaryButtonFgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Primary button foreground color")

  public String getPrimaryButtonFgColor() {
    return primaryButtonFgColor;
  }


  public void setPrimaryButtonFgColor(String primaryButtonFgColor) {
    this.primaryButtonFgColor = primaryButtonFgColor;
  }


  public WhitelabelSettingsUpdate primaryButtonHoverBgColor(String primaryButtonHoverBgColor) {
    
    this.primaryButtonHoverBgColor = primaryButtonHoverBgColor;
    return this;
  }

   /**
   * Primary button hover background color
   * @return primaryButtonHoverBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Primary button hover background color")

  public String getPrimaryButtonHoverBgColor() {
    return primaryButtonHoverBgColor;
  }


  public void setPrimaryButtonHoverBgColor(String primaryButtonHoverBgColor) {
    this.primaryButtonHoverBgColor = primaryButtonHoverBgColor;
  }


  public WhitelabelSettingsUpdate primaryButtonHoverFgColor(String primaryButtonHoverFgColor) {
    
    this.primaryButtonHoverFgColor = primaryButtonHoverFgColor;
    return this;
  }

   /**
   * Primary button hover foreground color
   * @return primaryButtonHoverFgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Primary button hover foreground color")

  public String getPrimaryButtonHoverFgColor() {
    return primaryButtonHoverFgColor;
  }


  public void setPrimaryButtonHoverFgColor(String primaryButtonHoverFgColor) {
    this.primaryButtonHoverFgColor = primaryButtonHoverFgColor;
  }


  public WhitelabelSettingsUpdate footerBgColor(String footerBgColor) {
    
    this.footerBgColor = footerBgColor;
    return this;
  }

   /**
   * Footer background color
   * @return footerBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Footer background color")

  public String getFooterBgColor() {
    return footerBgColor;
  }


  public void setFooterBgColor(String footerBgColor) {
    this.footerBgColor = footerBgColor;
  }


  public WhitelabelSettingsUpdate footerFgColor(String footerFgColor) {
    
    this.footerFgColor = footerFgColor;
    return this;
  }

   /**
   * Footer foreground color
   * @return footerFgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Footer foreground color")

  public String getFooterFgColor() {
    return footerFgColor;
  }


  public void setFooterFgColor(String footerFgColor) {
    this.footerFgColor = footerFgColor;
  }


  public WhitelabelSettingsUpdate loginBgColor(String loginBgColor) {
    
    this.loginBgColor = loginBgColor;
    return this;
  }

   /**
   * Login background color
   * @return loginBgColor
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Login background color")

  public String getLoginBgColor() {
    return loginBgColor;
  }


  public void setLoginBgColor(String loginBgColor) {
    this.loginBgColor = loginBgColor;
  }


  public WhitelabelSettingsUpdate copyrightString(String copyrightString) {
    
    this.copyrightString = copyrightString;
    return this;
  }

   /**
   * Copyright String
   * @return copyrightString
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Copyright String")

  public String getCopyrightString() {
    return copyrightString;
  }


  public void setCopyrightString(String copyrightString) {
    this.copyrightString = copyrightString;
  }


  public WhitelabelSettingsUpdate overrideCss(String overrideCss) {
    
    this.overrideCss = overrideCss;
    return this;
  }

   /**
   * Override CSS
   * @return overrideCss
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Override CSS")

  public String getOverrideCss() {
    return overrideCss;
  }


  public void setOverrideCss(String overrideCss) {
    this.overrideCss = overrideCss;
  }


  public WhitelabelSettingsUpdate termsOfUse(String termsOfUse) {
    
    this.termsOfUse = termsOfUse;
    return this;
  }

   /**
   * Terms of use content
   * @return termsOfUse
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Terms of use content")

  public String getTermsOfUse() {
    return termsOfUse;
  }


  public void setTermsOfUse(String termsOfUse) {
    this.termsOfUse = termsOfUse;
  }


  public WhitelabelSettingsUpdate privacyPolicy(String privacyPolicy) {
    
    this.privacyPolicy = privacyPolicy;
    return this;
  }

   /**
   * Privacy policy content
   * @return privacyPolicy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Privacy policy content")

  public String getPrivacyPolicy() {
    return privacyPolicy;
  }


  public void setPrivacyPolicy(String privacyPolicy) {
    this.privacyPolicy = privacyPolicy;
  }


  public WhitelabelSettingsUpdate supportMenuLinks(List<WhitelabelSettingsUpdateSupportMenuLinks> supportMenuLinks) {
    
    this.supportMenuLinks = supportMenuLinks;
    return this;
  }

  public WhitelabelSettingsUpdate addSupportMenuLinksItem(WhitelabelSettingsUpdateSupportMenuLinks supportMenuLinksItem) {
    if (this.supportMenuLinks == null) {
      this.supportMenuLinks = new ArrayList<WhitelabelSettingsUpdateSupportMenuLinks>();
    }
    this.supportMenuLinks.add(supportMenuLinksItem);
    return this;
  }

   /**
   * Get supportMenuLinks
   * @return supportMenuLinks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<WhitelabelSettingsUpdateSupportMenuLinks> getSupportMenuLinks() {
    return supportMenuLinks;
  }


  public void setSupportMenuLinks(List<WhitelabelSettingsUpdateSupportMenuLinks> supportMenuLinks) {
    this.supportMenuLinks = supportMenuLinks;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WhitelabelSettingsUpdate whitelabelSettingsUpdate = (WhitelabelSettingsUpdate) o;
    return Objects.equals(this.enabled, whitelabelSettingsUpdate.enabled) &&
        Objects.equals(this.applianceName, whitelabelSettingsUpdate.applianceName) &&
        Objects.equals(this.disableSupportMenu, whitelabelSettingsUpdate.disableSupportMenu) &&
        Objects.equals(this.resetHeaderLogo, whitelabelSettingsUpdate.resetHeaderLogo) &&
        Objects.equals(this.resetFooterLogo, whitelabelSettingsUpdate.resetFooterLogo) &&
        Objects.equals(this.resetLoginLogo, whitelabelSettingsUpdate.resetLoginLogo) &&
        Objects.equals(this.resetFavicon, whitelabelSettingsUpdate.resetFavicon) &&
        Objects.equals(this.headerBgColor, whitelabelSettingsUpdate.headerBgColor) &&
        Objects.equals(this.headerFgColor, whitelabelSettingsUpdate.headerFgColor) &&
        Objects.equals(this.navBgColor, whitelabelSettingsUpdate.navBgColor) &&
        Objects.equals(this.navFgColor, whitelabelSettingsUpdate.navFgColor) &&
        Objects.equals(this.navHoverColor, whitelabelSettingsUpdate.navHoverColor) &&
        Objects.equals(this.primaryButtonBgColor, whitelabelSettingsUpdate.primaryButtonBgColor) &&
        Objects.equals(this.primaryButtonFgColor, whitelabelSettingsUpdate.primaryButtonFgColor) &&
        Objects.equals(this.primaryButtonHoverBgColor, whitelabelSettingsUpdate.primaryButtonHoverBgColor) &&
        Objects.equals(this.primaryButtonHoverFgColor, whitelabelSettingsUpdate.primaryButtonHoverFgColor) &&
        Objects.equals(this.footerBgColor, whitelabelSettingsUpdate.footerBgColor) &&
        Objects.equals(this.footerFgColor, whitelabelSettingsUpdate.footerFgColor) &&
        Objects.equals(this.loginBgColor, whitelabelSettingsUpdate.loginBgColor) &&
        Objects.equals(this.copyrightString, whitelabelSettingsUpdate.copyrightString) &&
        Objects.equals(this.overrideCss, whitelabelSettingsUpdate.overrideCss) &&
        Objects.equals(this.termsOfUse, whitelabelSettingsUpdate.termsOfUse) &&
        Objects.equals(this.privacyPolicy, whitelabelSettingsUpdate.privacyPolicy) &&
        Objects.equals(this.supportMenuLinks, whitelabelSettingsUpdate.supportMenuLinks);
  }

  @Override
  public int hashCode() {
    return Objects.hash(enabled, applianceName, disableSupportMenu, resetHeaderLogo, resetFooterLogo, resetLoginLogo, resetFavicon, headerBgColor, headerFgColor, navBgColor, navFgColor, navHoverColor, primaryButtonBgColor, primaryButtonFgColor, primaryButtonHoverBgColor, primaryButtonHoverFgColor, footerBgColor, footerFgColor, loginBgColor, copyrightString, overrideCss, termsOfUse, privacyPolicy, supportMenuLinks);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WhitelabelSettingsUpdate {\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    applianceName: ").append(toIndentedString(applianceName)).append("\n");
    sb.append("    disableSupportMenu: ").append(toIndentedString(disableSupportMenu)).append("\n");
    sb.append("    resetHeaderLogo: ").append(toIndentedString(resetHeaderLogo)).append("\n");
    sb.append("    resetFooterLogo: ").append(toIndentedString(resetFooterLogo)).append("\n");
    sb.append("    resetLoginLogo: ").append(toIndentedString(resetLoginLogo)).append("\n");
    sb.append("    resetFavicon: ").append(toIndentedString(resetFavicon)).append("\n");
    sb.append("    headerBgColor: ").append(toIndentedString(headerBgColor)).append("\n");
    sb.append("    headerFgColor: ").append(toIndentedString(headerFgColor)).append("\n");
    sb.append("    navBgColor: ").append(toIndentedString(navBgColor)).append("\n");
    sb.append("    navFgColor: ").append(toIndentedString(navFgColor)).append("\n");
    sb.append("    navHoverColor: ").append(toIndentedString(navHoverColor)).append("\n");
    sb.append("    primaryButtonBgColor: ").append(toIndentedString(primaryButtonBgColor)).append("\n");
    sb.append("    primaryButtonFgColor: ").append(toIndentedString(primaryButtonFgColor)).append("\n");
    sb.append("    primaryButtonHoverBgColor: ").append(toIndentedString(primaryButtonHoverBgColor)).append("\n");
    sb.append("    primaryButtonHoverFgColor: ").append(toIndentedString(primaryButtonHoverFgColor)).append("\n");
    sb.append("    footerBgColor: ").append(toIndentedString(footerBgColor)).append("\n");
    sb.append("    footerFgColor: ").append(toIndentedString(footerFgColor)).append("\n");
    sb.append("    loginBgColor: ").append(toIndentedString(loginBgColor)).append("\n");
    sb.append("    copyrightString: ").append(toIndentedString(copyrightString)).append("\n");
    sb.append("    overrideCss: ").append(toIndentedString(overrideCss)).append("\n");
    sb.append("    termsOfUse: ").append(toIndentedString(termsOfUse)).append("\n");
    sb.append("    privacyPolicy: ").append(toIndentedString(privacyPolicy)).append("\n");
    sb.append("    supportMenuLinks: ").append(toIndentedString(supportMenuLinks)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

