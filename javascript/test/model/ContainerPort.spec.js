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
    instance = new MorpheusApi.ContainerPort();
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

  describe('ContainerPort', function() {
    it('should create an instance of ContainerPort', function() {
      // uncomment below and update the code to test ContainerPort
      //var instane = new MorpheusApi.ContainerPort();
      //expect(instance).to.be.a(MorpheusApi.ContainerPort);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property index (base name: "index")', function() {
      // uncomment below and update the code to test the property index
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property external (base name: "external")', function() {
      // uncomment below and update the code to test the property external
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property internal (base name: "internal")', function() {
      // uncomment below and update the code to test the property internal
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property primaryPort (base name: "primaryPort")', function() {
      // uncomment below and update the code to test the property primaryPort
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property _export (base name: "export")', function() {
      // uncomment below and update the code to test the property _export
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property visible (base name: "visible")', function() {
      // uncomment below and update the code to test the property visible
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property exportName (base name: "exportName")', function() {
      // uncomment below and update the code to test the property exportName
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property loadBalanceProtocol (base name: "loadBalanceProtocol")', function() {
      // uncomment below and update the code to test the property loadBalanceProtocol
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property loadBalance (base name: "loadBalance")', function() {
      // uncomment below and update the code to test the property loadBalance
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property protocol (base name: "protocol")', function() {
      // uncomment below and update the code to test the property protocol
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property link (base name: "link")', function() {
      // uncomment below and update the code to test the property link
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property externalIp (base name: "externalIp")', function() {
      // uncomment below and update the code to test the property externalIp
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

    it('should have the property internalIp (base name: "internalIp")', function() {
      // uncomment below and update the code to test the property internalIp
      //var instance = new MorpheusApi.ContainerPort();
      //expect(instance).to.be();
    });

  });

}));