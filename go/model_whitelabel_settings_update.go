/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WhitelabelSettingsUpdate struct for WhitelabelSettingsUpdate
type WhitelabelSettingsUpdate struct {
	// Can be used to enable / disable whitelabel feature
	Enabled *bool `json:"enabled,omitempty"`
	// Appliance name. Master account only
	ApplianceName *string `json:"applianceName,omitempty"`
	// Can be used to disable support menu
	DisableSupportMenu *bool `json:"disableSupportMenu,omitempty"`
	// Resets header logo to default header logo
	ResetHeaderLogo *bool `json:"resetHeaderLogo,omitempty"`
	// Resets footer logo to default footer logo
	ResetFooterLogo *bool `json:"resetFooterLogo,omitempty"`
	// Resets login logo to default login logo
	ResetLoginLogo *bool `json:"resetLoginLogo,omitempty"`
	// Resets favicon to default favicon
	ResetFavicon *bool `json:"resetFavicon,omitempty"`
	// Header background color
	HeaderBgColor *string `json:"headerBgColor,omitempty"`
	// Header foreground color
	HeaderFgColor *string `json:"headerFgColor,omitempty"`
	// Nav background color
	NavBgColor *string `json:"navBgColor,omitempty"`
	// Nav foreground color
	NavFgColor *string `json:"navFgColor,omitempty"`
	// Nav hover color
	NavHoverColor *string `json:"navHoverColor,omitempty"`
	// Primary button background color
	PrimaryButtonBgColor *string `json:"primaryButtonBgColor,omitempty"`
	// Primary button foreground color
	PrimaryButtonFgColor *string `json:"primaryButtonFgColor,omitempty"`
	// Primary button hover background color
	PrimaryButtonHoverBgColor *string `json:"primaryButtonHoverBgColor,omitempty"`
	// Primary button hover foreground color
	PrimaryButtonHoverFgColor *string `json:"primaryButtonHoverFgColor,omitempty"`
	// Footer background color
	FooterBgColor *string `json:"footerBgColor,omitempty"`
	// Footer foreground color
	FooterFgColor *string `json:"footerFgColor,omitempty"`
	// Login background color
	LoginBgColor *string `json:"loginBgColor,omitempty"`
	// Copyright String
	CopyrightString *string `json:"copyrightString,omitempty"`
	// Override CSS
	OverrideCss *string `json:"overrideCss,omitempty"`
	// Terms of use content
	TermsOfUse *string `json:"termsOfUse,omitempty"`
	// Privacy policy content
	PrivacyPolicy *string `json:"privacyPolicy,omitempty"`
	SupportMenuLinks *[]WhitelabelSettingsUpdateSupportMenuLinks `json:"supportMenuLinks,omitempty"`
}

// NewWhitelabelSettingsUpdate instantiates a new WhitelabelSettingsUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhitelabelSettingsUpdate() *WhitelabelSettingsUpdate {
	this := WhitelabelSettingsUpdate{}
	return &this
}

// NewWhitelabelSettingsUpdateWithDefaults instantiates a new WhitelabelSettingsUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhitelabelSettingsUpdateWithDefaults() *WhitelabelSettingsUpdate {
	this := WhitelabelSettingsUpdate{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WhitelabelSettingsUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetApplianceName returns the ApplianceName field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetApplianceName() string {
	if o == nil || o.ApplianceName == nil {
		var ret string
		return ret
	}
	return *o.ApplianceName
}

// GetApplianceNameOk returns a tuple with the ApplianceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetApplianceNameOk() (*string, bool) {
	if o == nil || o.ApplianceName == nil {
		return nil, false
	}
	return o.ApplianceName, true
}

// HasApplianceName returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasApplianceName() bool {
	if o != nil && o.ApplianceName != nil {
		return true
	}

	return false
}

// SetApplianceName gets a reference to the given string and assigns it to the ApplianceName field.
func (o *WhitelabelSettingsUpdate) SetApplianceName(v string) {
	o.ApplianceName = &v
}

// GetDisableSupportMenu returns the DisableSupportMenu field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetDisableSupportMenu() bool {
	if o == nil || o.DisableSupportMenu == nil {
		var ret bool
		return ret
	}
	return *o.DisableSupportMenu
}

// GetDisableSupportMenuOk returns a tuple with the DisableSupportMenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetDisableSupportMenuOk() (*bool, bool) {
	if o == nil || o.DisableSupportMenu == nil {
		return nil, false
	}
	return o.DisableSupportMenu, true
}

// HasDisableSupportMenu returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasDisableSupportMenu() bool {
	if o != nil && o.DisableSupportMenu != nil {
		return true
	}

	return false
}

// SetDisableSupportMenu gets a reference to the given bool and assigns it to the DisableSupportMenu field.
func (o *WhitelabelSettingsUpdate) SetDisableSupportMenu(v bool) {
	o.DisableSupportMenu = &v
}

// GetResetHeaderLogo returns the ResetHeaderLogo field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetResetHeaderLogo() bool {
	if o == nil || o.ResetHeaderLogo == nil {
		var ret bool
		return ret
	}
	return *o.ResetHeaderLogo
}

// GetResetHeaderLogoOk returns a tuple with the ResetHeaderLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetResetHeaderLogoOk() (*bool, bool) {
	if o == nil || o.ResetHeaderLogo == nil {
		return nil, false
	}
	return o.ResetHeaderLogo, true
}

// HasResetHeaderLogo returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasResetHeaderLogo() bool {
	if o != nil && o.ResetHeaderLogo != nil {
		return true
	}

	return false
}

// SetResetHeaderLogo gets a reference to the given bool and assigns it to the ResetHeaderLogo field.
func (o *WhitelabelSettingsUpdate) SetResetHeaderLogo(v bool) {
	o.ResetHeaderLogo = &v
}

// GetResetFooterLogo returns the ResetFooterLogo field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetResetFooterLogo() bool {
	if o == nil || o.ResetFooterLogo == nil {
		var ret bool
		return ret
	}
	return *o.ResetFooterLogo
}

// GetResetFooterLogoOk returns a tuple with the ResetFooterLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetResetFooterLogoOk() (*bool, bool) {
	if o == nil || o.ResetFooterLogo == nil {
		return nil, false
	}
	return o.ResetFooterLogo, true
}

// HasResetFooterLogo returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasResetFooterLogo() bool {
	if o != nil && o.ResetFooterLogo != nil {
		return true
	}

	return false
}

// SetResetFooterLogo gets a reference to the given bool and assigns it to the ResetFooterLogo field.
func (o *WhitelabelSettingsUpdate) SetResetFooterLogo(v bool) {
	o.ResetFooterLogo = &v
}

// GetResetLoginLogo returns the ResetLoginLogo field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetResetLoginLogo() bool {
	if o == nil || o.ResetLoginLogo == nil {
		var ret bool
		return ret
	}
	return *o.ResetLoginLogo
}

// GetResetLoginLogoOk returns a tuple with the ResetLoginLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetResetLoginLogoOk() (*bool, bool) {
	if o == nil || o.ResetLoginLogo == nil {
		return nil, false
	}
	return o.ResetLoginLogo, true
}

// HasResetLoginLogo returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasResetLoginLogo() bool {
	if o != nil && o.ResetLoginLogo != nil {
		return true
	}

	return false
}

// SetResetLoginLogo gets a reference to the given bool and assigns it to the ResetLoginLogo field.
func (o *WhitelabelSettingsUpdate) SetResetLoginLogo(v bool) {
	o.ResetLoginLogo = &v
}

// GetResetFavicon returns the ResetFavicon field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetResetFavicon() bool {
	if o == nil || o.ResetFavicon == nil {
		var ret bool
		return ret
	}
	return *o.ResetFavicon
}

// GetResetFaviconOk returns a tuple with the ResetFavicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetResetFaviconOk() (*bool, bool) {
	if o == nil || o.ResetFavicon == nil {
		return nil, false
	}
	return o.ResetFavicon, true
}

// HasResetFavicon returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasResetFavicon() bool {
	if o != nil && o.ResetFavicon != nil {
		return true
	}

	return false
}

// SetResetFavicon gets a reference to the given bool and assigns it to the ResetFavicon field.
func (o *WhitelabelSettingsUpdate) SetResetFavicon(v bool) {
	o.ResetFavicon = &v
}

// GetHeaderBgColor returns the HeaderBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetHeaderBgColor() string {
	if o == nil || o.HeaderBgColor == nil {
		var ret string
		return ret
	}
	return *o.HeaderBgColor
}

// GetHeaderBgColorOk returns a tuple with the HeaderBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetHeaderBgColorOk() (*string, bool) {
	if o == nil || o.HeaderBgColor == nil {
		return nil, false
	}
	return o.HeaderBgColor, true
}

// HasHeaderBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasHeaderBgColor() bool {
	if o != nil && o.HeaderBgColor != nil {
		return true
	}

	return false
}

// SetHeaderBgColor gets a reference to the given string and assigns it to the HeaderBgColor field.
func (o *WhitelabelSettingsUpdate) SetHeaderBgColor(v string) {
	o.HeaderBgColor = &v
}

// GetHeaderFgColor returns the HeaderFgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetHeaderFgColor() string {
	if o == nil || o.HeaderFgColor == nil {
		var ret string
		return ret
	}
	return *o.HeaderFgColor
}

// GetHeaderFgColorOk returns a tuple with the HeaderFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetHeaderFgColorOk() (*string, bool) {
	if o == nil || o.HeaderFgColor == nil {
		return nil, false
	}
	return o.HeaderFgColor, true
}

// HasHeaderFgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasHeaderFgColor() bool {
	if o != nil && o.HeaderFgColor != nil {
		return true
	}

	return false
}

// SetHeaderFgColor gets a reference to the given string and assigns it to the HeaderFgColor field.
func (o *WhitelabelSettingsUpdate) SetHeaderFgColor(v string) {
	o.HeaderFgColor = &v
}

// GetNavBgColor returns the NavBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetNavBgColor() string {
	if o == nil || o.NavBgColor == nil {
		var ret string
		return ret
	}
	return *o.NavBgColor
}

// GetNavBgColorOk returns a tuple with the NavBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetNavBgColorOk() (*string, bool) {
	if o == nil || o.NavBgColor == nil {
		return nil, false
	}
	return o.NavBgColor, true
}

// HasNavBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasNavBgColor() bool {
	if o != nil && o.NavBgColor != nil {
		return true
	}

	return false
}

// SetNavBgColor gets a reference to the given string and assigns it to the NavBgColor field.
func (o *WhitelabelSettingsUpdate) SetNavBgColor(v string) {
	o.NavBgColor = &v
}

// GetNavFgColor returns the NavFgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetNavFgColor() string {
	if o == nil || o.NavFgColor == nil {
		var ret string
		return ret
	}
	return *o.NavFgColor
}

// GetNavFgColorOk returns a tuple with the NavFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetNavFgColorOk() (*string, bool) {
	if o == nil || o.NavFgColor == nil {
		return nil, false
	}
	return o.NavFgColor, true
}

// HasNavFgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasNavFgColor() bool {
	if o != nil && o.NavFgColor != nil {
		return true
	}

	return false
}

// SetNavFgColor gets a reference to the given string and assigns it to the NavFgColor field.
func (o *WhitelabelSettingsUpdate) SetNavFgColor(v string) {
	o.NavFgColor = &v
}

// GetNavHoverColor returns the NavHoverColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetNavHoverColor() string {
	if o == nil || o.NavHoverColor == nil {
		var ret string
		return ret
	}
	return *o.NavHoverColor
}

// GetNavHoverColorOk returns a tuple with the NavHoverColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetNavHoverColorOk() (*string, bool) {
	if o == nil || o.NavHoverColor == nil {
		return nil, false
	}
	return o.NavHoverColor, true
}

// HasNavHoverColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasNavHoverColor() bool {
	if o != nil && o.NavHoverColor != nil {
		return true
	}

	return false
}

// SetNavHoverColor gets a reference to the given string and assigns it to the NavHoverColor field.
func (o *WhitelabelSettingsUpdate) SetNavHoverColor(v string) {
	o.NavHoverColor = &v
}

// GetPrimaryButtonBgColor returns the PrimaryButtonBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonBgColor() string {
	if o == nil || o.PrimaryButtonBgColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonBgColor
}

// GetPrimaryButtonBgColorOk returns a tuple with the PrimaryButtonBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonBgColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonBgColor == nil {
		return nil, false
	}
	return o.PrimaryButtonBgColor, true
}

// HasPrimaryButtonBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasPrimaryButtonBgColor() bool {
	if o != nil && o.PrimaryButtonBgColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonBgColor gets a reference to the given string and assigns it to the PrimaryButtonBgColor field.
func (o *WhitelabelSettingsUpdate) SetPrimaryButtonBgColor(v string) {
	o.PrimaryButtonBgColor = &v
}

// GetPrimaryButtonFgColor returns the PrimaryButtonFgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonFgColor() string {
	if o == nil || o.PrimaryButtonFgColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonFgColor
}

// GetPrimaryButtonFgColorOk returns a tuple with the PrimaryButtonFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonFgColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonFgColor == nil {
		return nil, false
	}
	return o.PrimaryButtonFgColor, true
}

// HasPrimaryButtonFgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasPrimaryButtonFgColor() bool {
	if o != nil && o.PrimaryButtonFgColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonFgColor gets a reference to the given string and assigns it to the PrimaryButtonFgColor field.
func (o *WhitelabelSettingsUpdate) SetPrimaryButtonFgColor(v string) {
	o.PrimaryButtonFgColor = &v
}

// GetPrimaryButtonHoverBgColor returns the PrimaryButtonHoverBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonHoverBgColor() string {
	if o == nil || o.PrimaryButtonHoverBgColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonHoverBgColor
}

// GetPrimaryButtonHoverBgColorOk returns a tuple with the PrimaryButtonHoverBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonHoverBgColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonHoverBgColor == nil {
		return nil, false
	}
	return o.PrimaryButtonHoverBgColor, true
}

// HasPrimaryButtonHoverBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasPrimaryButtonHoverBgColor() bool {
	if o != nil && o.PrimaryButtonHoverBgColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonHoverBgColor gets a reference to the given string and assigns it to the PrimaryButtonHoverBgColor field.
func (o *WhitelabelSettingsUpdate) SetPrimaryButtonHoverBgColor(v string) {
	o.PrimaryButtonHoverBgColor = &v
}

// GetPrimaryButtonHoverFgColor returns the PrimaryButtonHoverFgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonHoverFgColor() string {
	if o == nil || o.PrimaryButtonHoverFgColor == nil {
		var ret string
		return ret
	}
	return *o.PrimaryButtonHoverFgColor
}

// GetPrimaryButtonHoverFgColorOk returns a tuple with the PrimaryButtonHoverFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetPrimaryButtonHoverFgColorOk() (*string, bool) {
	if o == nil || o.PrimaryButtonHoverFgColor == nil {
		return nil, false
	}
	return o.PrimaryButtonHoverFgColor, true
}

// HasPrimaryButtonHoverFgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasPrimaryButtonHoverFgColor() bool {
	if o != nil && o.PrimaryButtonHoverFgColor != nil {
		return true
	}

	return false
}

// SetPrimaryButtonHoverFgColor gets a reference to the given string and assigns it to the PrimaryButtonHoverFgColor field.
func (o *WhitelabelSettingsUpdate) SetPrimaryButtonHoverFgColor(v string) {
	o.PrimaryButtonHoverFgColor = &v
}

// GetFooterBgColor returns the FooterBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetFooterBgColor() string {
	if o == nil || o.FooterBgColor == nil {
		var ret string
		return ret
	}
	return *o.FooterBgColor
}

// GetFooterBgColorOk returns a tuple with the FooterBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetFooterBgColorOk() (*string, bool) {
	if o == nil || o.FooterBgColor == nil {
		return nil, false
	}
	return o.FooterBgColor, true
}

// HasFooterBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasFooterBgColor() bool {
	if o != nil && o.FooterBgColor != nil {
		return true
	}

	return false
}

// SetFooterBgColor gets a reference to the given string and assigns it to the FooterBgColor field.
func (o *WhitelabelSettingsUpdate) SetFooterBgColor(v string) {
	o.FooterBgColor = &v
}

// GetFooterFgColor returns the FooterFgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetFooterFgColor() string {
	if o == nil || o.FooterFgColor == nil {
		var ret string
		return ret
	}
	return *o.FooterFgColor
}

// GetFooterFgColorOk returns a tuple with the FooterFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetFooterFgColorOk() (*string, bool) {
	if o == nil || o.FooterFgColor == nil {
		return nil, false
	}
	return o.FooterFgColor, true
}

// HasFooterFgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasFooterFgColor() bool {
	if o != nil && o.FooterFgColor != nil {
		return true
	}

	return false
}

// SetFooterFgColor gets a reference to the given string and assigns it to the FooterFgColor field.
func (o *WhitelabelSettingsUpdate) SetFooterFgColor(v string) {
	o.FooterFgColor = &v
}

// GetLoginBgColor returns the LoginBgColor field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetLoginBgColor() string {
	if o == nil || o.LoginBgColor == nil {
		var ret string
		return ret
	}
	return *o.LoginBgColor
}

// GetLoginBgColorOk returns a tuple with the LoginBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetLoginBgColorOk() (*string, bool) {
	if o == nil || o.LoginBgColor == nil {
		return nil, false
	}
	return o.LoginBgColor, true
}

// HasLoginBgColor returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasLoginBgColor() bool {
	if o != nil && o.LoginBgColor != nil {
		return true
	}

	return false
}

// SetLoginBgColor gets a reference to the given string and assigns it to the LoginBgColor field.
func (o *WhitelabelSettingsUpdate) SetLoginBgColor(v string) {
	o.LoginBgColor = &v
}

// GetCopyrightString returns the CopyrightString field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetCopyrightString() string {
	if o == nil || o.CopyrightString == nil {
		var ret string
		return ret
	}
	return *o.CopyrightString
}

// GetCopyrightStringOk returns a tuple with the CopyrightString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetCopyrightStringOk() (*string, bool) {
	if o == nil || o.CopyrightString == nil {
		return nil, false
	}
	return o.CopyrightString, true
}

// HasCopyrightString returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasCopyrightString() bool {
	if o != nil && o.CopyrightString != nil {
		return true
	}

	return false
}

// SetCopyrightString gets a reference to the given string and assigns it to the CopyrightString field.
func (o *WhitelabelSettingsUpdate) SetCopyrightString(v string) {
	o.CopyrightString = &v
}

// GetOverrideCss returns the OverrideCss field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetOverrideCss() string {
	if o == nil || o.OverrideCss == nil {
		var ret string
		return ret
	}
	return *o.OverrideCss
}

// GetOverrideCssOk returns a tuple with the OverrideCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetOverrideCssOk() (*string, bool) {
	if o == nil || o.OverrideCss == nil {
		return nil, false
	}
	return o.OverrideCss, true
}

// HasOverrideCss returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasOverrideCss() bool {
	if o != nil && o.OverrideCss != nil {
		return true
	}

	return false
}

// SetOverrideCss gets a reference to the given string and assigns it to the OverrideCss field.
func (o *WhitelabelSettingsUpdate) SetOverrideCss(v string) {
	o.OverrideCss = &v
}

// GetTermsOfUse returns the TermsOfUse field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetTermsOfUse() string {
	if o == nil || o.TermsOfUse == nil {
		var ret string
		return ret
	}
	return *o.TermsOfUse
}

// GetTermsOfUseOk returns a tuple with the TermsOfUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetTermsOfUseOk() (*string, bool) {
	if o == nil || o.TermsOfUse == nil {
		return nil, false
	}
	return o.TermsOfUse, true
}

// HasTermsOfUse returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasTermsOfUse() bool {
	if o != nil && o.TermsOfUse != nil {
		return true
	}

	return false
}

// SetTermsOfUse gets a reference to the given string and assigns it to the TermsOfUse field.
func (o *WhitelabelSettingsUpdate) SetTermsOfUse(v string) {
	o.TermsOfUse = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetPrivacyPolicy() string {
	if o == nil || o.PrivacyPolicy == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetPrivacyPolicyOk() (*string, bool) {
	if o == nil || o.PrivacyPolicy == nil {
		return nil, false
	}
	return o.PrivacyPolicy, true
}

// HasPrivacyPolicy returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasPrivacyPolicy() bool {
	if o != nil && o.PrivacyPolicy != nil {
		return true
	}

	return false
}

// SetPrivacyPolicy gets a reference to the given string and assigns it to the PrivacyPolicy field.
func (o *WhitelabelSettingsUpdate) SetPrivacyPolicy(v string) {
	o.PrivacyPolicy = &v
}

// GetSupportMenuLinks returns the SupportMenuLinks field value if set, zero value otherwise.
func (o *WhitelabelSettingsUpdate) GetSupportMenuLinks() []WhitelabelSettingsUpdateSupportMenuLinks {
	if o == nil || o.SupportMenuLinks == nil {
		var ret []WhitelabelSettingsUpdateSupportMenuLinks
		return ret
	}
	return *o.SupportMenuLinks
}

// GetSupportMenuLinksOk returns a tuple with the SupportMenuLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelabelSettingsUpdate) GetSupportMenuLinksOk() (*[]WhitelabelSettingsUpdateSupportMenuLinks, bool) {
	if o == nil || o.SupportMenuLinks == nil {
		return nil, false
	}
	return o.SupportMenuLinks, true
}

// HasSupportMenuLinks returns a boolean if a field has been set.
func (o *WhitelabelSettingsUpdate) HasSupportMenuLinks() bool {
	if o != nil && o.SupportMenuLinks != nil {
		return true
	}

	return false
}

// SetSupportMenuLinks gets a reference to the given []WhitelabelSettingsUpdateSupportMenuLinks and assigns it to the SupportMenuLinks field.
func (o *WhitelabelSettingsUpdate) SetSupportMenuLinks(v []WhitelabelSettingsUpdateSupportMenuLinks) {
	o.SupportMenuLinks = &v
}

func (o WhitelabelSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ApplianceName != nil {
		toSerialize["applianceName"] = o.ApplianceName
	}
	if o.DisableSupportMenu != nil {
		toSerialize["disableSupportMenu"] = o.DisableSupportMenu
	}
	if o.ResetHeaderLogo != nil {
		toSerialize["resetHeaderLogo"] = o.ResetHeaderLogo
	}
	if o.ResetFooterLogo != nil {
		toSerialize["resetFooterLogo"] = o.ResetFooterLogo
	}
	if o.ResetLoginLogo != nil {
		toSerialize["resetLoginLogo"] = o.ResetLoginLogo
	}
	if o.ResetFavicon != nil {
		toSerialize["resetFavicon"] = o.ResetFavicon
	}
	if o.HeaderBgColor != nil {
		toSerialize["headerBgColor"] = o.HeaderBgColor
	}
	if o.HeaderFgColor != nil {
		toSerialize["headerFgColor"] = o.HeaderFgColor
	}
	if o.NavBgColor != nil {
		toSerialize["navBgColor"] = o.NavBgColor
	}
	if o.NavFgColor != nil {
		toSerialize["navFgColor"] = o.NavFgColor
	}
	if o.NavHoverColor != nil {
		toSerialize["navHoverColor"] = o.NavHoverColor
	}
	if o.PrimaryButtonBgColor != nil {
		toSerialize["primaryButtonBgColor"] = o.PrimaryButtonBgColor
	}
	if o.PrimaryButtonFgColor != nil {
		toSerialize["primaryButtonFgColor"] = o.PrimaryButtonFgColor
	}
	if o.PrimaryButtonHoverBgColor != nil {
		toSerialize["primaryButtonHoverBgColor"] = o.PrimaryButtonHoverBgColor
	}
	if o.PrimaryButtonHoverFgColor != nil {
		toSerialize["primaryButtonHoverFgColor"] = o.PrimaryButtonHoverFgColor
	}
	if o.FooterBgColor != nil {
		toSerialize["footerBgColor"] = o.FooterBgColor
	}
	if o.FooterFgColor != nil {
		toSerialize["footerFgColor"] = o.FooterFgColor
	}
	if o.LoginBgColor != nil {
		toSerialize["loginBgColor"] = o.LoginBgColor
	}
	if o.CopyrightString != nil {
		toSerialize["copyrightString"] = o.CopyrightString
	}
	if o.OverrideCss != nil {
		toSerialize["overrideCss"] = o.OverrideCss
	}
	if o.TermsOfUse != nil {
		toSerialize["termsOfUse"] = o.TermsOfUse
	}
	if o.PrivacyPolicy != nil {
		toSerialize["privacyPolicy"] = o.PrivacyPolicy
	}
	if o.SupportMenuLinks != nil {
		toSerialize["supportMenuLinks"] = o.SupportMenuLinks
	}
	return json.Marshal(toSerialize)
}

type NullableWhitelabelSettingsUpdate struct {
	value *WhitelabelSettingsUpdate
	isSet bool
}

func (v NullableWhitelabelSettingsUpdate) Get() *WhitelabelSettingsUpdate {
	return v.value
}

func (v *NullableWhitelabelSettingsUpdate) Set(val *WhitelabelSettingsUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWhitelabelSettingsUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWhitelabelSettingsUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhitelabelSettingsUpdate(val *WhitelabelSettingsUpdate) *NullableWhitelabelSettingsUpdate {
	return &NullableWhitelabelSettingsUpdate{value: val, isSet: true}
}

func (v NullableWhitelabelSettingsUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhitelabelSettingsUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

