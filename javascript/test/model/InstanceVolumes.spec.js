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
    instance = new MorpheusApi.InstanceVolumes();
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

  describe('InstanceVolumes', function() {
    it('should create an instance of InstanceVolumes', function() {
      // uncomment below and update the code to test InstanceVolumes
      //var instane = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be.a(MorpheusApi.InstanceVolumes);
    });

    it('should have the property controllerId (base name: "controllerId")', function() {
      // uncomment below and update the code to test the property controllerId
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property datastoreId (base name: "datastoreId")', function() {
      // uncomment below and update the code to test the property datastoreId
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property displayOrder (base name: "displayOrder")', function() {
      // uncomment below and update the code to test the property displayOrder
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property uuid (base name: "uuid")', function() {
      // uncomment below and update the code to test the property uuid
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property maxIOPS (base name: "maxIOPS")', function() {
      // uncomment below and update the code to test the property maxIOPS
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property shortName (base name: "shortName")', function() {
      // uncomment below and update the code to test the property shortName
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property resizeable (base name: "resizeable")', function() {
      // uncomment below and update the code to test the property resizeable
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property planResizable (base name: "planResizable")', function() {
      // uncomment below and update the code to test the property planResizable
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property rootVolume (base name: "rootVolume")', function() {
      // uncomment below and update the code to test the property rootVolume
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property size (base name: "size")', function() {
      // uncomment below and update the code to test the property size
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property storageType (base name: "storageType")', function() {
      // uncomment below and update the code to test the property storageType
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property unitNumber (base name: "unitNumber")', function() {
      // uncomment below and update the code to test the property unitNumber
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

    it('should have the property controllerMountPoint (base name: "controllerMountPoint")', function() {
      // uncomment below and update the code to test the property controllerMountPoint
      //var instance = new MorpheusApi.InstanceVolumes();
      //expect(instance).to.be();
    });

  });

}));
