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
    instance = new MorpheusApi.InstancesConfigVMWare();
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

  describe('InstancesConfigVMWare', function() {
    it('should create an instance of InstancesConfigVMWare', function() {
      // uncomment below and update the code to test InstancesConfigVMWare
      //var instane = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be.a(MorpheusApi.InstancesConfigVMWare);
    });

    it('should have the property noAgent (base name: "noAgent")', function() {
      // uncomment below and update the code to test the property noAgent
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

    it('should have the property resourcePoolId (base name: "resourcePoolId")', function() {
      // uncomment below and update the code to test the property resourcePoolId
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

    it('should have the property hostId (base name: "hostId")', function() {
      // uncomment below and update the code to test the property hostId
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

    it('should have the property smbiosAssetTag (base name: "smbiosAssetTag")', function() {
      // uncomment below and update the code to test the property smbiosAssetTag
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

    it('should have the property nestedVirtualization (base name: "nestedVirtualization")', function() {
      // uncomment below and update the code to test the property nestedVirtualization
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

    it('should have the property vmwareFolderId (base name: "vmwareFolderId")', function() {
      // uncomment below and update the code to test the property vmwareFolderId
      //var instance = new MorpheusApi.InstancesConfigVMWare();
      //expect(instance).to.be();
    });

  });

}));
