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
    instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
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

  describe('InlineResponse20095NetworkRouterFirewall', function() {
    it('should create an instance of InlineResponse20095NetworkRouterFirewall', function() {
      // uncomment below and update the code to test InlineResponse20095NetworkRouterFirewall
      //var instane = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be.a(MorpheusApi.InlineResponse20095NetworkRouterFirewall);
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be();
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be();
    });

    it('should have the property defaultPolicy (base name: "defaultPolicy")', function() {
      // uncomment below and update the code to test the property defaultPolicy
      //var instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be();
    });

    it('should have the property global (base name: "global")', function() {
      // uncomment below and update the code to test the property global
      //var instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be();
    });

    it('should have the property ruleGroups (base name: "ruleGroups")', function() {
      // uncomment below and update the code to test the property ruleGroups
      //var instance = new MorpheusApi.InlineResponse20095NetworkRouterFirewall();
      //expect(instance).to.be();
    });

  });

}));