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
    instance = new MorpheusApi.ClusterMastersInterfaces();
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

  describe('ClusterMastersInterfaces', function() {
    it('should create an instance of ClusterMastersInterfaces', function() {
      // uncomment below and update the code to test ClusterMastersInterfaces
      //var instane = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be.a(MorpheusApi.ClusterMastersInterfaces);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property uniqueId (base name: "uniqueId")', function() {
      // uncomment below and update the code to test the property uniqueId
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property publicIpAddress (base name: "publicIpAddress")', function() {
      // uncomment below and update the code to test the property publicIpAddress
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property publicIpv6Address (base name: "publicIpv6Address")', function() {
      // uncomment below and update the code to test the property publicIpv6Address
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property ipAddress (base name: "ipAddress")', function() {
      // uncomment below and update the code to test the property ipAddress
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property ipv6Address (base name: "ipv6Address")', function() {
      // uncomment below and update the code to test the property ipv6Address
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property ipSubnet (base name: "ipSubnet")', function() {
      // uncomment below and update the code to test the property ipSubnet
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property ipv6Subnet (base name: "ipv6Subnet")', function() {
      // uncomment below and update the code to test the property ipv6Subnet
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property dhcp (base name: "dhcp")', function() {
      // uncomment below and update the code to test the property dhcp
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property poolAssigned (base name: "poolAssigned")', function() {
      // uncomment below and update the code to test the property poolAssigned
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property primaryInterface (base name: "primaryInterface")', function() {
      // uncomment below and update the code to test the property primaryInterface
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property network (base name: "network")', function() {
      // uncomment below and update the code to test the property network
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property subnet (base name: "subnet")', function() {
      // uncomment below and update the code to test the property subnet
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property networkGroup (base name: "networkGroup")', function() {
      // uncomment below and update the code to test the property networkGroup
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property networkPosition (base name: "networkPosition")', function() {
      // uncomment below and update the code to test the property networkPosition
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property networkPool (base name: "networkPool")', function() {
      // uncomment below and update the code to test the property networkPool
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property networkDomain (base name: "networkDomain")', function() {
      // uncomment below and update the code to test the property networkDomain
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property ipMode (base name: "ipMode")', function() {
      // uncomment below and update the code to test the property ipMode
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

    it('should have the property macAddress (base name: "macAddress")', function() {
      // uncomment below and update the code to test the property macAddress
      //var instance = new MorpheusApi.ClusterMastersInterfaces();
      //expect(instance).to.be();
    });

  });

}));