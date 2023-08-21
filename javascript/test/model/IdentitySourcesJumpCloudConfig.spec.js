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
    instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
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

  describe('IdentitySourcesJumpCloudConfig', function() {
    it('should create an instance of IdentitySourcesJumpCloudConfig', function() {
      // uncomment below and update the code to test IdentitySourcesJumpCloudConfig
      //var instane = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be.a(MorpheusApi.IdentitySourcesJumpCloudConfig);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property deleted (base name: "deleted")', function() {
      // uncomment below and update the code to test the property deleted
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property autoSyncOnLogin (base name: "autoSyncOnLogin")', function() {
      // uncomment below and update the code to test the property autoSyncOnLogin
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property externalLogin (base name: "externalLogin")', function() {
      // uncomment below and update the code to test the property externalLogin
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property allowCustomMappings (base name: "allowCustomMappings")', function() {
      // uncomment below and update the code to test the property allowCustomMappings
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property manualRoleAssignment (base name: "manualRoleAssignment")', function() {
      // uncomment below and update the code to test the property manualRoleAssignment
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property defaultAccountRole (base name: "defaultAccountRole")', function() {
      // uncomment below and update the code to test the property defaultAccountRole
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property roleMappings (base name: "roleMappings")', function() {
      // uncomment below and update the code to test the property roleMappings
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property subdomain (base name: "subdomain")', function() {
      // uncomment below and update the code to test the property subdomain
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property loginURL (base name: "loginURL")', function() {
      // uncomment below and update the code to test the property loginURL
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property providerSettings (base name: "providerSettings")', function() {
      // uncomment below and update the code to test the property providerSettings
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.IdentitySourcesJumpCloudConfig();
      //expect(instance).to.be();
    });

  });

}));