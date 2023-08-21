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
    instance = new MorpheusApi.ClusterTypes();
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

  describe('ClusterTypes', function() {
    it('should create an instance of ClusterTypes', function() {
      // uncomment below and update the code to test ClusterTypes
      //var instane = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be.a(MorpheusApi.ClusterTypes);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property deployTargetService (base name: "deployTargetService")', function() {
      // uncomment below and update the code to test the property deployTargetService
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property shortName (base name: "shortName")', function() {
      // uncomment below and update the code to test the property shortName
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property providerType (base name: "providerType")', function() {
      // uncomment below and update the code to test the property providerType
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hostService (base name: "hostService")', function() {
      // uncomment below and update the code to test the property hostService
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property managed (base name: "managed")', function() {
      // uncomment below and update the code to test the property managed
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hasMasters (base name: "hasMasters")', function() {
      // uncomment below and update the code to test the property hasMasters
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hasWorkers (base name: "hasWorkers")', function() {
      // uncomment below and update the code to test the property hasWorkers
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property viewSet (base name: "viewSet")', function() {
      // uncomment below and update the code to test the property viewSet
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property imageCode (base name: "imageCode")', function() {
      // uncomment below and update the code to test the property imageCode
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property kubeCtlLocal (base name: "kubeCtlLocal")', function() {
      // uncomment below and update the code to test the property kubeCtlLocal
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hasDatastore (base name: "hasDatastore")', function() {
      // uncomment below and update the code to test the property hasDatastore
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property supportsCloudScaling (base name: "supportsCloudScaling")', function() {
      // uncomment below and update the code to test the property supportsCloudScaling
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hasDefaultDataDisk (base name: "hasDefaultDataDisk")', function() {
      // uncomment below and update the code to test the property hasDefaultDataDisk
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property canManage (base name: "canManage")', function() {
      // uncomment below and update the code to test the property canManage
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property hasCluster (base name: "hasCluster")', function() {
      // uncomment below and update the code to test the property hasCluster
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property optionTypes (base name: "optionTypes")', function() {
      // uncomment below and update the code to test the property optionTypes
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property controllerTypes (base name: "controllerTypes")', function() {
      // uncomment below and update the code to test the property controllerTypes
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

    it('should have the property workerTypes (base name: "workerTypes")', function() {
      // uncomment below and update the code to test the property workerTypes
      //var instance = new MorpheusApi.ClusterTypes();
      //expect(instance).to.be();
    });

  });

}));