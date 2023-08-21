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
    instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
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

  describe('InlineResponse20080LoadBalancerPool', function() {
    it('should create an instance of InlineResponse20080LoadBalancerPool', function() {
      // uncomment below and update the code to test InlineResponse20080LoadBalancerPool
      //var instane = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be.a(MorpheusApi.InlineResponse20080LoadBalancerPool);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property loadBalancer (base name: "loadBalancer")', function() {
      // uncomment below and update the code to test the property loadBalancer
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property vipSticky (base name: "vipSticky")', function() {
      // uncomment below and update the code to test the property vipSticky
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property vipBalance (base name: "vipBalance")', function() {
      // uncomment below and update the code to test the property vipBalance
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property allowNat (base name: "allowNat")', function() {
      // uncomment below and update the code to test the property allowNat
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property allowSnat (base name: "allowSnat")', function() {
      // uncomment below and update the code to test the property allowSnat
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property vipClientIpMode (base name: "vipClientIpMode")', function() {
      // uncomment below and update the code to test the property vipClientIpMode
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property vipServerIpMode (base name: "vipServerIpMode")', function() {
      // uncomment below and update the code to test the property vipServerIpMode
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property minActive (base name: "minActive")', function() {
      // uncomment below and update the code to test the property minActive
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property minInService (base name: "minInService")', function() {
      // uncomment below and update the code to test the property minInService
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property minUpMonitor (base name: "minUpMonitor")', function() {
      // uncomment below and update the code to test the property minUpMonitor
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property minUpAction (base name: "minUpAction")', function() {
      // uncomment below and update the code to test the property minUpAction
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property maxQueueDepth (base name: "maxQueueDepth")', function() {
      // uncomment below and update the code to test the property maxQueueDepth
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property maxQueueTime (base name: "maxQueueTime")', function() {
      // uncomment below and update the code to test the property maxQueueTime
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property numberActive (base name: "numberActive")', function() {
      // uncomment below and update the code to test the property numberActive
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property numberInService (base name: "numberInService")', function() {
      // uncomment below and update the code to test the property numberInService
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property healthScore (base name: "healthScore")', function() {
      // uncomment below and update the code to test the property healthScore
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property performanceScore (base name: "performanceScore")', function() {
      // uncomment below and update the code to test the property performanceScore
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property healthPenalty (base name: "healthPenalty")', function() {
      // uncomment below and update the code to test the property healthPenalty
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property securityPenalty (base name: "securityPenalty")', function() {
      // uncomment below and update the code to test the property securityPenalty
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property errorPenalty (base name: "errorPenalty")', function() {
      // uncomment below and update the code to test the property errorPenalty
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property downAction (base name: "downAction")', function() {
      // uncomment below and update the code to test the property downAction
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property rampTime (base name: "rampTime")', function() {
      // uncomment below and update the code to test the property rampTime
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property port (base name: "port")', function() {
      // uncomment below and update the code to test the property port
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property portType (base name: "portType")', function() {
      // uncomment below and update the code to test the property portType
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property nodes (base name: "nodes")', function() {
      // uncomment below and update the code to test the property nodes
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property monitors (base name: "monitors")', function() {
      // uncomment below and update the code to test the property monitors
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property members (base name: "members")', function() {
      // uncomment below and update the code to test the property members
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.InlineResponse20080LoadBalancerPool();
      //expect(instance).to.be();
    });

  });

}));
