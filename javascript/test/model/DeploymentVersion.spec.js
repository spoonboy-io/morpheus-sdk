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
    instance = new MorpheusApi.DeploymentVersion();
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

  describe('DeploymentVersion', function() {
    it('should create an instance of DeploymentVersion', function() {
      // uncomment below and update the code to test DeploymentVersion
      //var instane = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be.a(MorpheusApi.DeploymentVersion);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property deployType (base name: "deployType")', function() {
      // uncomment below and update the code to test the property deployType
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property deploymentId (base name: "deploymentId")', function() {
      // uncomment below and update the code to test the property deploymentId
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property fetchUrl (base name: "fetchUrl")', function() {
      // uncomment below and update the code to test the property fetchUrl
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property gitUrl (base name: "gitUrl")', function() {
      // uncomment below and update the code to test the property gitUrl
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property gitRef (base name: "gitRef")', function() {
      // uncomment below and update the code to test the property gitRef
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property userVersion (base name: "userVersion")', function() {
      // uncomment below and update the code to test the property userVersion
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.DeploymentVersion();
      //expect(instance).to.be();
    });

  });

}));