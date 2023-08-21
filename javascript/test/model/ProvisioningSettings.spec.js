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
    instance = new MorpheusApi.ProvisioningSettings();
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

  describe('ProvisioningSettings', function() {
    it('should create an instance of ProvisioningSettings', function() {
      // uncomment below and update the code to test ProvisioningSettings
      //var instane = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be.a(MorpheusApi.ProvisioningSettings);
    });

    it('should have the property allowZoneSelection (base name: "allowZoneSelection")', function() {
      // uncomment below and update the code to test the property allowZoneSelection
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property allowServerSelection (base name: "allowServerSelection")', function() {
      // uncomment below and update the code to test the property allowServerSelection
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property requireEnvironments (base name: "requireEnvironments")', function() {
      // uncomment below and update the code to test the property requireEnvironments
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property showPricing (base name: "showPricing")', function() {
      // uncomment below and update the code to test the property showPricing
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property hideDatastoreStats (base name: "hideDatastoreStats")', function() {
      // uncomment below and update the code to test the property hideDatastoreStats
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property crossTenantNamingPolicies (base name: "crossTenantNamingPolicies")', function() {
      // uncomment below and update the code to test the property crossTenantNamingPolicies
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property reuseSequence (base name: "reuseSequence")', function() {
      // uncomment below and update the code to test the property reuseSequence
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property cloudInitUsername (base name: "cloudInitUsername")', function() {
      // uncomment below and update the code to test the property cloudInitUsername
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property cloudInitPassword (base name: "cloudInitPassword")', function() {
      // uncomment below and update the code to test the property cloudInitPassword
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property cloudInitKeyPair (base name: "cloudInitKeyPair")', function() {
      // uncomment below and update the code to test the property cloudInitKeyPair
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property windowsPassword (base name: "windowsPassword")', function() {
      // uncomment below and update the code to test the property windowsPassword
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property pxeRootPassword (base name: "pxeRootPassword")', function() {
      // uncomment below and update the code to test the property pxeRootPassword
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property defaultTemplateType (base name: "defaultTemplateType")', function() {
      // uncomment below and update the code to test the property defaultTemplateType
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

    it('should have the property deployStorageProvider (base name: "deployStorageProvider")', function() {
      // uncomment below and update the code to test the property deployStorageProvider
      //var instance = new MorpheusApi.ProvisioningSettings();
      //expect(instance).to.be();
    });

  });

}));