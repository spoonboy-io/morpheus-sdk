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
    instance = new MorpheusApi.IntegrationPuppet();
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

  describe('IntegrationPuppet', function() {
    it('should create an instance of IntegrationPuppet', function() {
      // uncomment below and update the code to test IntegrationPuppet
      //var instane = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be.a(MorpheusApi.IntegrationPuppet);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property integrationType (base name: "integrationType")', function() {
      // uncomment below and update the code to test the property integrationType
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property isPlugin (base name: "isPlugin")', function() {
      // uncomment below and update the code to test the property isPlugin
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property statusDate (base name: "statusDate")', function() {
      // uncomment below and update the code to test the property statusDate
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property statusMessage (base name: "statusMessage")', function() {
      // uncomment below and update the code to test the property statusMessage
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property lastSync (base name: "lastSync")', function() {
      // uncomment below and update the code to test the property lastSync
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property lastSyncDuration (base name: "lastSyncDuration")', function() {
      // uncomment below and update the code to test the property lastSyncDuration
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

    it('should have the property credential (base name: "credential")', function() {
      // uncomment below and update the code to test the property credential
      //var instance = new MorpheusApi.IntegrationPuppet();
      //expect(instance).to.be();
    });

  });

}));
