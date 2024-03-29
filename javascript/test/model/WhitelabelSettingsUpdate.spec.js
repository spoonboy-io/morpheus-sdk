/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MorpheusApi);
  }
}(this, function(expect, MorpheusApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MorpheusApi.WhitelabelSettingsUpdate();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('WhitelabelSettingsUpdate', function() {
    it('should create an instance of WhitelabelSettingsUpdate', function() {
      // uncomment below and update the code to test WhitelabelSettingsUpdate
      //var instane = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be.a(MorpheusApi.WhitelabelSettingsUpdate);
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property applianceName (base name: "applianceName")', function() {
      // uncomment below and update the code to test the property applianceName
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property disableSupportMenu (base name: "disableSupportMenu")', function() {
      // uncomment below and update the code to test the property disableSupportMenu
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property resetHeaderLogo (base name: "resetHeaderLogo")', function() {
      // uncomment below and update the code to test the property resetHeaderLogo
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property resetFooterLogo (base name: "resetFooterLogo")', function() {
      // uncomment below and update the code to test the property resetFooterLogo
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property resetLoginLogo (base name: "resetLoginLogo")', function() {
      // uncomment below and update the code to test the property resetLoginLogo
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property resetFavicon (base name: "resetFavicon")', function() {
      // uncomment below and update the code to test the property resetFavicon
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property headerBgColor (base name: "headerBgColor")', function() {
      // uncomment below and update the code to test the property headerBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property headerFgColor (base name: "headerFgColor")', function() {
      // uncomment below and update the code to test the property headerFgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property navBgColor (base name: "navBgColor")', function() {
      // uncomment below and update the code to test the property navBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property navFgColor (base name: "navFgColor")', function() {
      // uncomment below and update the code to test the property navFgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property navHoverColor (base name: "navHoverColor")', function() {
      // uncomment below and update the code to test the property navHoverColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property primaryButtonBgColor (base name: "primaryButtonBgColor")', function() {
      // uncomment below and update the code to test the property primaryButtonBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property primaryButtonFgColor (base name: "primaryButtonFgColor")', function() {
      // uncomment below and update the code to test the property primaryButtonFgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property primaryButtonHoverBgColor (base name: "primaryButtonHoverBgColor")', function() {
      // uncomment below and update the code to test the property primaryButtonHoverBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property primaryButtonHoverFgColor (base name: "primaryButtonHoverFgColor")', function() {
      // uncomment below and update the code to test the property primaryButtonHoverFgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property footerBgColor (base name: "footerBgColor")', function() {
      // uncomment below and update the code to test the property footerBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property footerFgColor (base name: "footerFgColor")', function() {
      // uncomment below and update the code to test the property footerFgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property loginBgColor (base name: "loginBgColor")', function() {
      // uncomment below and update the code to test the property loginBgColor
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property copyrightString (base name: "copyrightString")', function() {
      // uncomment below and update the code to test the property copyrightString
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property overrideCss (base name: "overrideCss")', function() {
      // uncomment below and update the code to test the property overrideCss
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property termsOfUse (base name: "termsOfUse")', function() {
      // uncomment below and update the code to test the property termsOfUse
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property privacyPolicy (base name: "privacyPolicy")', function() {
      // uncomment below and update the code to test the property privacyPolicy
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

    it('should have the property supportMenuLinks (base name: "supportMenuLinks")', function() {
      // uncomment below and update the code to test the property supportMenuLinks
      //var instance = new MorpheusApi.WhitelabelSettingsUpdate();
      //expect(instance).to.be();
    });

  });

}));
