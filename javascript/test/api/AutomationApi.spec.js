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
    instance = new MorpheusApi.AutomationApi();
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

  describe('AutomationApi', function() {
    describe('addExecuteSchedules', function() {
      it('should call addExecuteSchedules successfully', function(done) {
        //uncomment below and update the code to test addExecuteSchedules
        //instance.addExecuteSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addPowerScheduleInstances', function() {
      it('should call addPowerScheduleInstances successfully', function(done) {
        //uncomment below and update the code to test addPowerScheduleInstances
        //instance.addPowerScheduleInstances(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addPowerScheduleServers', function() {
      it('should call addPowerScheduleServers successfully', function(done) {
        //uncomment below and update the code to test addPowerScheduleServers
        //instance.addPowerScheduleServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addPowerSchedules', function() {
      it('should call addPowerSchedules successfully', function(done) {
        //uncomment below and update the code to test addPowerSchedules
        //instance.addPowerSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addScaleThresholds', function() {
      it('should call addScaleThresholds successfully', function(done) {
        //uncomment below and update the code to test addScaleThresholds
        //instance.addScaleThresholds(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addTasks', function() {
      it('should call addTasks successfully', function(done) {
        //uncomment below and update the code to test addTasks
        //instance.addTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addWorkflows', function() {
      it('should call addWorkflows successfully', function(done) {
        //uncomment below and update the code to test addWorkflows
        //instance.addWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deletePowerScheduleInstances', function() {
      it('should call deletePowerScheduleInstances successfully', function(done) {
        //uncomment below and update the code to test deletePowerScheduleInstances
        //instance.deletePowerScheduleInstances(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deletePowerScheduleServers', function() {
      it('should call deletePowerScheduleServers successfully', function(done) {
        //uncomment below and update the code to test deletePowerScheduleServers
        //instance.deletePowerScheduleServers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('executeExecutionRequest', function() {
      it('should call executeExecutionRequest successfully', function(done) {
        //uncomment below and update the code to test executeExecutionRequest
        //instance.executeExecutionRequest(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('executeTasks', function() {
      it('should call executeTasks successfully', function(done) {
        //uncomment below and update the code to test executeTasks
        //instance.executeTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('executeWorkflows', function() {
      it('should call executeWorkflows successfully', function(done) {
        //uncomment below and update the code to test executeWorkflows
        //instance.executeWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getExecuteSchedules', function() {
      it('should call getExecuteSchedules successfully', function(done) {
        //uncomment below and update the code to test getExecuteSchedules
        //instance.getExecuteSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getExecutionRequest', function() {
      it('should call getExecutionRequest successfully', function(done) {
        //uncomment below and update the code to test getExecutionRequest
        //instance.getExecutionRequest(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getPowerSchedules', function() {
      it('should call getPowerSchedules successfully', function(done) {
        //uncomment below and update the code to test getPowerSchedules
        //instance.getPowerSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getScaleThresholds', function() {
      it('should call getScaleThresholds successfully', function(done) {
        //uncomment below and update the code to test getScaleThresholds
        //instance.getScaleThresholds(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getTaskTypes', function() {
      it('should call getTaskTypes successfully', function(done) {
        //uncomment below and update the code to test getTaskTypes
        //instance.getTaskTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getTasks', function() {
      it('should call getTasks successfully', function(done) {
        //uncomment below and update the code to test getTasks
        //instance.getTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getWorkflows', function() {
      it('should call getWorkflows successfully', function(done) {
        //uncomment below and update the code to test getWorkflows
        //instance.getWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listExecuteSchedules', function() {
      it('should call listExecuteSchedules successfully', function(done) {
        //uncomment below and update the code to test listExecuteSchedules
        //instance.listExecuteSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listPowerSchedules', function() {
      it('should call listPowerSchedules successfully', function(done) {
        //uncomment below and update the code to test listPowerSchedules
        //instance.listPowerSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listScaleThresholds', function() {
      it('should call listScaleThresholds successfully', function(done) {
        //uncomment below and update the code to test listScaleThresholds
        //instance.listScaleThresholds(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listTaskTypes', function() {
      it('should call listTaskTypes successfully', function(done) {
        //uncomment below and update the code to test listTaskTypes
        //instance.listTaskTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listTasks', function() {
      it('should call listTasks successfully', function(done) {
        //uncomment below and update the code to test listTasks
        //instance.listTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listWorkflows', function() {
      it('should call listWorkflows successfully', function(done) {
        //uncomment below and update the code to test listWorkflows
        //instance.listWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeExecuteSchedules', function() {
      it('should call removeExecuteSchedules successfully', function(done) {
        //uncomment below and update the code to test removeExecuteSchedules
        //instance.removeExecuteSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removePowerSchedules', function() {
      it('should call removePowerSchedules successfully', function(done) {
        //uncomment below and update the code to test removePowerSchedules
        //instance.removePowerSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeScaleThresholds', function() {
      it('should call removeScaleThresholds successfully', function(done) {
        //uncomment below and update the code to test removeScaleThresholds
        //instance.removeScaleThresholds(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeTasks', function() {
      it('should call removeTasks successfully', function(done) {
        //uncomment below and update the code to test removeTasks
        //instance.removeTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeWorkflows', function() {
      it('should call removeWorkflows successfully', function(done) {
        //uncomment below and update the code to test removeWorkflows
        //instance.removeWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateExecuteSchedules', function() {
      it('should call updateExecuteSchedules successfully', function(done) {
        //uncomment below and update the code to test updateExecuteSchedules
        //instance.updateExecuteSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updatePowerSchedules', function() {
      it('should call updatePowerSchedules successfully', function(done) {
        //uncomment below and update the code to test updatePowerSchedules
        //instance.updatePowerSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateScaleThresholds', function() {
      it('should call updateScaleThresholds successfully', function(done) {
        //uncomment below and update the code to test updateScaleThresholds
        //instance.updateScaleThresholds(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateTasks', function() {
      it('should call updateTasks successfully', function(done) {
        //uncomment below and update the code to test updateTasks
        //instance.updateTasks(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateWorkflows', function() {
      it('should call updateWorkflows successfully', function(done) {
        //uncomment below and update the code to test updateWorkflows
        //instance.updateWorkflows(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
