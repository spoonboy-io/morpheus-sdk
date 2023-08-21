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
    instance = new MorpheusApi.LoadBalancersApi();
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

  describe('LoadBalancersApi', function() {
    describe('createLoadBalancer', function() {
      it('should call createLoadBalancer successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancer
        //instance.createLoadBalancer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createLoadBalancerMonitor', function() {
      it('should call createLoadBalancerMonitor successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancerMonitor
        //instance.createLoadBalancerMonitor(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createLoadBalancerPool', function() {
      it('should call createLoadBalancerPool successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancerPool
        //instance.createLoadBalancerPool(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createLoadBalancerPoolNode', function() {
      it('should call createLoadBalancerPoolNode successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancerPoolNode
        //instance.createLoadBalancerPoolNode(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createLoadBalancerProfile', function() {
      it('should call createLoadBalancerProfile successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancerProfile
        //instance.createLoadBalancerProfile(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createLoadBalancerVirtualServer', function() {
      it('should call createLoadBalancerVirtualServer successfully', function(done) {
        //uncomment below and update the code to test createLoadBalancerVirtualServer
        //instance.createLoadBalancerVirtualServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancer', function() {
      it('should call deleteLoadBalancer successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancer
        //instance.deleteLoadBalancer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancerMonitor', function() {
      it('should call deleteLoadBalancerMonitor successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancerMonitor
        //instance.deleteLoadBalancerMonitor(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancerPool', function() {
      it('should call deleteLoadBalancerPool successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancerPool
        //instance.deleteLoadBalancerPool(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancerPoolNode', function() {
      it('should call deleteLoadBalancerPoolNode successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancerPoolNode
        //instance.deleteLoadBalancerPoolNode(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancerProfile', function() {
      it('should call deleteLoadBalancerProfile successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancerProfile
        //instance.deleteLoadBalancerProfile(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteLoadBalancerVirtualServer', function() {
      it('should call deleteLoadBalancerVirtualServer successfully', function(done) {
        //uncomment below and update the code to test deleteLoadBalancerVirtualServer
        //instance.deleteLoadBalancerVirtualServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancer', function() {
      it('should call getLoadBalancer successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancer
        //instance.getLoadBalancer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerMonitor', function() {
      it('should call getLoadBalancerMonitor successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerMonitor
        //instance.getLoadBalancerMonitor(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerPool', function() {
      it('should call getLoadBalancerPool successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerPool
        //instance.getLoadBalancerPool(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerPoolNode', function() {
      it('should call getLoadBalancerPoolNode successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerPoolNode
        //instance.getLoadBalancerPoolNode(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerProfile', function() {
      it('should call getLoadBalancerProfile successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerProfile
        //instance.getLoadBalancerProfile(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerType', function() {
      it('should call getLoadBalancerType successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerType
        //instance.getLoadBalancerType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getLoadBalancerVirtualServer', function() {
      it('should call getLoadBalancerVirtualServer successfully', function(done) {
        //uncomment below and update the code to test getLoadBalancerVirtualServer
        //instance.getLoadBalancerVirtualServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerMonitors', function() {
      it('should call listLoadBalancerMonitors successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerMonitors
        //instance.listLoadBalancerMonitors(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerPoolNodes', function() {
      it('should call listLoadBalancerPoolNodes successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerPoolNodes
        //instance.listLoadBalancerPoolNodes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerPools', function() {
      it('should call listLoadBalancerPools successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerPools
        //instance.listLoadBalancerPools(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerProfiles', function() {
      it('should call listLoadBalancerProfiles successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerProfiles
        //instance.listLoadBalancerProfiles(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerTypes', function() {
      it('should call listLoadBalancerTypes successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerTypes
        //instance.listLoadBalancerTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancerVirtualServers', function() {
      it('should call listLoadBalancerVirtualServers successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancerVirtualServers
        //instance.listLoadBalancerVirtualServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listLoadBalancers', function() {
      it('should call listLoadBalancers successfully', function(done) {
        //uncomment below and update the code to test listLoadBalancers
        //instance.listLoadBalancers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('refreshLoadBalancer', function() {
      it('should call refreshLoadBalancer successfully', function(done) {
        //uncomment below and update the code to test refreshLoadBalancer
        //instance.refreshLoadBalancer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancer', function() {
      it('should call updateLoadBalancer successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancer
        //instance.updateLoadBalancer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancerMonitor', function() {
      it('should call updateLoadBalancerMonitor successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancerMonitor
        //instance.updateLoadBalancerMonitor(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancerPool', function() {
      it('should call updateLoadBalancerPool successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancerPool
        //instance.updateLoadBalancerPool(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancerPoolNode', function() {
      it('should call updateLoadBalancerPoolNode successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancerPoolNode
        //instance.updateLoadBalancerPoolNode(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancerProfile', function() {
      it('should call updateLoadBalancerProfile successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancerProfile
        //instance.updateLoadBalancerProfile(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateLoadBalancerVirtualServer', function() {
      it('should call updateLoadBalancerVirtualServer successfully', function(done) {
        //uncomment below and update the code to test updateLoadBalancerVirtualServer
        //instance.updateLoadBalancerVirtualServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));