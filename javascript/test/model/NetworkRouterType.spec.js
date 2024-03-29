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
    instance = new MorpheusApi.NetworkRouterType();
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

  describe('NetworkRouterType', function() {
    it('should create an instance of NetworkRouterType', function() {
      // uncomment below and update the code to test NetworkRouterType
      //var instane = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be.a(MorpheusApi.NetworkRouterType);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property creatable (base name: "creatable")', function() {
      // uncomment below and update the code to test the property creatable
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property selectable (base name: "selectable")', function() {
      // uncomment below and update the code to test the property selectable
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property hasFirewall (base name: "hasFirewall")', function() {
      // uncomment below and update the code to test the property hasFirewall
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property hasDhcp (base name: "hasDhcp")', function() {
      // uncomment below and update the code to test the property hasDhcp
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property hasRouting (base name: "hasRouting")', function() {
      // uncomment below and update the code to test the property hasRouting
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property hasNetworkServer (base name: "hasNetworkServer")', function() {
      // uncomment below and update the code to test the property hasNetworkServer
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property optionTypes (base name: "optionTypes")', function() {
      // uncomment below and update the code to test the property optionTypes
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property ruleOptionTypes (base name: "ruleOptionTypes")', function() {
      // uncomment below and update the code to test the property ruleOptionTypes
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property ruleGroupOptionTypes (base name: "ruleGroupOptionTypes")', function() {
      // uncomment below and update the code to test the property ruleGroupOptionTypes
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property natOptionTypes (base name: "natOptionTypes")', function() {
      // uncomment below and update the code to test the property natOptionTypes
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

    it('should have the property bgpOptionTypes (base name: "bgpOptionTypes")', function() {
      // uncomment below and update the code to test the property bgpOptionTypes
      //var instance = new MorpheusApi.NetworkRouterType();
      //expect(instance).to.be();
    });

  });

}));
