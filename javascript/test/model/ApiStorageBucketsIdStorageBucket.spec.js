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
    instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
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

  describe('ApiStorageBucketsIdStorageBucket', function() {
    it('should create an instance of ApiStorageBucketsIdStorageBucket', function() {
      // uncomment below and update the code to test ApiStorageBucketsIdStorageBucket
      //var instane = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be.a(MorpheusApi.ApiStorageBucketsIdStorageBucket);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property providerType (base name: "providerType")', function() {
      // uncomment below and update the code to test the property providerType
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property defaultBackupTarget (base name: "defaultBackupTarget")', function() {
      // uncomment below and update the code to test the property defaultBackupTarget
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property copyToStore (base name: "copyToStore")', function() {
      // uncomment below and update the code to test the property copyToStore
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property defaultDeploymentTarget (base name: "defaultDeploymentTarget")', function() {
      // uncomment below and update the code to test the property defaultDeploymentTarget
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property defaultVirtualImageTarget (base name: "defaultVirtualImageTarget")', function() {
      // uncomment below and update the code to test the property defaultVirtualImageTarget
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property retentionPolicyType (base name: "retentionPolicyType")', function() {
      // uncomment below and update the code to test the property retentionPolicyType
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property retentionPolicyDays (base name: "retentionPolicyDays")', function() {
      // uncomment below and update the code to test the property retentionPolicyDays
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property retentionProvider (base name: "retentionProvider")', function() {
      // uncomment below and update the code to test the property retentionProvider
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property bucketName (base name: "bucketName")', function() {
      // uncomment below and update the code to test the property bucketName
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property createBucket (base name: "createBucket")', function() {
      // uncomment below and update the code to test the property createBucket
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.ApiStorageBucketsIdStorageBucket();
      //expect(instance).to.be();
    });

  });

}));
