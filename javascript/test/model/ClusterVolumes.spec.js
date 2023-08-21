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
    instance = new MorpheusApi.ClusterVolumes();
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

  describe('ClusterVolumes', function() {
    it('should create an instance of ClusterVolumes', function() {
      // uncomment below and update the code to test ClusterVolumes
      //var instane = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be.a(MorpheusApi.ClusterVolumes);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property displayOrder (base name: "displayOrder")', function() {
      // uncomment below and update the code to test the property displayOrder
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property usedStorage (base name: "usedStorage")', function() {
      // uncomment below and update the code to test the property usedStorage
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property provisionType (base name: "provisionType")', function() {
      // uncomment below and update the code to test the property provisionType
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property resizeable (base name: "resizeable")', function() {
      // uncomment below and update the code to test the property resizeable
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property online (base name: "online")', function() {
      // uncomment below and update the code to test the property online
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property deviceDisplayName (base name: "deviceDisplayName")', function() {
      // uncomment below and update the code to test the property deviceDisplayName
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property datastoreOption (base name: "datastoreOption")', function() {
      // uncomment below and update the code to test the property datastoreOption
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property claimName (base name: "claimName")', function() {
      // uncomment below and update the code to test the property claimName
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property volumeType (base name: "volumeType")', function() {
      // uncomment below and update the code to test the property volumeType
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property deviceName (base name: "deviceName")', function() {
      // uncomment below and update the code to test the property deviceName
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property removable (base name: "removable")', function() {
      // uncomment below and update the code to test the property removable
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property poolName (base name: "poolName")', function() {
      // uncomment below and update the code to test the property poolName
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property readOnly (base name: "readOnly")', function() {
      // uncomment below and update the code to test the property readOnly
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property sourceId (base name: "sourceId")', function() {
      // uncomment below and update the code to test the property sourceId
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property zoneId (base name: "zoneId")', function() {
      // uncomment below and update the code to test the property zoneId
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property rootVolume (base name: "rootVolume")', function() {
      // uncomment below and update the code to test the property rootVolume
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property rawData (base name: "rawData")', function() {
      // uncomment below and update the code to test the property rawData
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

    it('should have the property datastore (base name: "datastore")', function() {
      // uncomment below and update the code to test the property datastore
      //var instance = new MorpheusApi.ClusterVolumes();
      //expect(instance).to.be();
    });

  });

}));
