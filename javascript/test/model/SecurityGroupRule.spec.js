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
    instance = new MorpheusApi.SecurityGroupRule();
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

  describe('SecurityGroupRule', function() {
    it('should create an instance of SecurityGroupRule', function() {
      // uncomment below and update the code to test SecurityGroupRule
      //var instane = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be.a(MorpheusApi.SecurityGroupRule);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property ruleType (base name: "ruleType")', function() {
      // uncomment below and update the code to test the property ruleType
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property customRule (base name: "customRule")', function() {
      // uncomment below and update the code to test the property customRule
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property instanceTypeId (base name: "instanceTypeId")', function() {
      // uncomment below and update the code to test the property instanceTypeId
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property direction (base name: "direction")', function() {
      // uncomment below and update the code to test the property direction
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property policy (base name: "policy")', function() {
      // uncomment below and update the code to test the property policy
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property sourceType (base name: "sourceType")', function() {
      // uncomment below and update the code to test the property sourceType
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property source (base name: "source")', function() {
      // uncomment below and update the code to test the property source
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property sourceGroup (base name: "sourceGroup")', function() {
      // uncomment below and update the code to test the property sourceGroup
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property sourceTier (base name: "sourceTier")', function() {
      // uncomment below and update the code to test the property sourceTier
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property portRange (base name: "portRange")', function() {
      // uncomment below and update the code to test the property portRange
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property protocol (base name: "protocol")', function() {
      // uncomment below and update the code to test the property protocol
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property destinationType (base name: "destinationType")', function() {
      // uncomment below and update the code to test the property destinationType
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property destination (base name: "destination")', function() {
      // uncomment below and update the code to test the property destination
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property destinationGroup (base name: "destinationGroup")', function() {
      // uncomment below and update the code to test the property destinationGroup
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property destinationTier (base name: "destinationTier")', function() {
      // uncomment below and update the code to test the property destinationTier
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.SecurityGroupRule();
      //expect(instance).to.be();
    });

  });

}));