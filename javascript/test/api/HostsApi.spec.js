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
    instance = new MorpheusApi.HostsApi();
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

  describe('HostsApi', function() {
    describe('getHost', function() {
      it('should call getHost successfully', function(done) {
        //uncomment below and update the code to test getHost
        //instance.getHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getHostSnpshots', function() {
      it('should call getHostSnpshots successfully', function(done) {
        //uncomment below and update the code to test getHostSnpshots
        //instance.getHostSnpshots(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getHostType', function() {
      it('should call getHostType successfully', function(done) {
        //uncomment below and update the code to test getHostType
        //instance.getHostType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getWikiServer', function() {
      it('should call getWikiServer successfully', function(done) {
        //uncomment below and update the code to test getWikiServer
        //instance.getWikiServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listHostTypes', function() {
      it('should call listHostTypes successfully', function(done) {
        //uncomment below and update the code to test listHostTypes
        //instance.listHostTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listHosts', function() {
      it('should call listHosts successfully', function(done) {
        //uncomment below and update the code to test listHosts
        //instance.listHosts(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listServerServicePlans', function() {
      it('should call listServerServicePlans successfully', function(done) {
        //uncomment below and update the code to test listServerServicePlans
        //instance.listServerServicePlans(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeHost', function() {
      it('should call removeHost successfully', function(done) {
        //uncomment below and update the code to test removeHost
        //instance.removeHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restartHost', function() {
      it('should call restartHost successfully', function(done) {
        //uncomment below and update the code to test restartHost
        //instance.restartHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('startHost', function() {
      it('should call startHost successfully', function(done) {
        //uncomment below and update the code to test startHost
        //instance.startHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopHost', function() {
      it('should call stopHost successfully', function(done) {
        //uncomment below and update the code to test stopHost
        //instance.stopHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHost', function() {
      it('should call updateHost successfully', function(done) {
        //uncomment below and update the code to test updateHost
        //instance.updateHost(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostAssignTenant', function() {
      it('should call updateHostAssignTenant successfully', function(done) {
        //uncomment below and update the code to test updateHostAssignTenant
        //instance.updateHostAssignTenant(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostCloud', function() {
      it('should call updateHostCloud successfully', function(done) {
        //uncomment below and update the code to test updateHostCloud
        //instance.updateHostCloud(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostExecuteWorkflow', function() {
      it('should call updateHostExecuteWorkflow successfully', function(done) {
        //uncomment below and update the code to test updateHostExecuteWorkflow
        //instance.updateHostExecuteWorkflow(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostInstallAgent', function() {
      it('should call updateHostInstallAgent successfully', function(done) {
        //uncomment below and update the code to test updateHostInstallAgent
        //instance.updateHostInstallAgent(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostManaged', function() {
      it('should call updateHostManaged successfully', function(done) {
        //uncomment below and update the code to test updateHostManaged
        //instance.updateHostManaged(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostResize', function() {
      it('should call updateHostResize successfully', function(done) {
        //uncomment below and update the code to test updateHostResize
        //instance.updateHostResize(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateHostUpgradeAgent', function() {
      it('should call updateHostUpgradeAgent successfully', function(done) {
        //uncomment below and update the code to test updateHostUpgradeAgent
        //instance.updateHostUpgradeAgent(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateServerNetworkInterface', function() {
      it('should call updateServerNetworkInterface successfully', function(done) {
        //uncomment below and update the code to test updateServerNetworkInterface
        //instance.updateServerNetworkInterface(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateWikiServer', function() {
      it('should call updateWikiServer successfully', function(done) {
        //uncomment below and update the code to test updateWikiServer
        //instance.updateWikiServer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
