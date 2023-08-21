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
    instance = new MorpheusApi.InstancesConfigAzure();
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

  describe('InstancesConfigAzure', function() {
    it('should create an instance of InstancesConfigAzure', function() {
      // uncomment below and update the code to test InstancesConfigAzure
      //var instane = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be.a(MorpheusApi.InstancesConfigAzure);
    });

    it('should have the property resourcePoolId (base name: "resourcePoolId")', function() {
      // uncomment below and update the code to test the property resourcePoolId
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property availabilityOptions (base name: "availabilityOptions")', function() {
      // uncomment below and update the code to test the property availabilityOptions
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property availabilitySet (base name: "availabilitySet")', function() {
      // uncomment below and update the code to test the property availabilitySet
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property availabilityZone (base name: "availabilityZone")', function() {
      // uncomment below and update the code to test the property availabilityZone
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property azurefloatingIp (base name: "azurefloatingIp")', function() {
      // uncomment below and update the code to test the property azurefloatingIp
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property bootDiagnostics (base name: "bootDiagnostics")', function() {
      // uncomment below and update the code to test the property bootDiagnostics
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property osGuestDiagnostics (base name: "osGuestDiagnostics")', function() {
      // uncomment below and update the code to test the property osGuestDiagnostics
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property diagnosticsStorageAccount (base name: "diagnosticsStorageAccount")', function() {
      // uncomment below and update the code to test the property diagnosticsStorageAccount
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

    it('should have the property createUser (base name: "createUser")', function() {
      // uncomment below and update the code to test the property createUser
      //var instance = new MorpheusApi.InstancesConfigAzure();
      //expect(instance).to.be();
    });

  });

}));