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
    instance = new MorpheusApi.ClusterDatastore();
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

  describe('ClusterDatastore', function() {
    it('should create an instance of ClusterDatastore', function() {
      // uncomment below and update the code to test ClusterDatastore
      //var instane = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be.a(MorpheusApi.ClusterDatastore);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property storageSize (base name: "storageSize")', function() {
      // uncomment below and update the code to test the property storageSize
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property freeSpace (base name: "freeSpace")', function() {
      // uncomment below and update the code to test the property freeSpace
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property drsEnabled (base name: "drsEnabled")', function() {
      // uncomment below and update the code to test the property drsEnabled
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property allowWrite (base name: "allowWrite")', function() {
      // uncomment below and update the code to test the property allowWrite
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property defaultStore (base name: "defaultStore")', function() {
      // uncomment below and update the code to test the property defaultStore
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property online (base name: "online")', function() {
      // uncomment below and update the code to test the property online
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property allowRead (base name: "allowRead")', function() {
      // uncomment below and update the code to test the property allowRead
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property allowProvision (base name: "allowProvision")', function() {
      // uncomment below and update the code to test the property allowProvision
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property zone (base name: "zone")', function() {
      // uncomment below and update the code to test the property zone
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property zonePool (base name: "zonePool")', function() {
      // uncomment below and update the code to test the property zonePool
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property owner (base name: "owner")', function() {
      // uncomment below and update the code to test the property owner
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property tenants (base name: "tenants")', function() {
      // uncomment below and update the code to test the property tenants
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property permissions (base name: "permissions")', function() {
      // uncomment below and update the code to test the property permissions
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

    it('should have the property datastores (base name: "datastores")', function() {
      // uncomment below and update the code to test the property datastores
      //var instance = new MorpheusApi.ClusterDatastore();
      //expect(instance).to.be();
    });

  });

}));