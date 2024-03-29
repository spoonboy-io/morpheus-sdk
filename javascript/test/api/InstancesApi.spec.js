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
    instance = new MorpheusApi.InstancesApi();
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

  describe('InstancesApi', function() {
    describe('addInstance', function() {
      it('should call addInstance successfully', function(done) {
        //uncomment below and update the code to test addInstance
        //instance.addInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('backupInstance', function() {
      it('should call backupInstance successfully', function(done) {
        //uncomment below and update the code to test backupInstance
        //instance.backupInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('backupsInstance', function() {
      it('should call backupsInstance successfully', function(done) {
        //uncomment below and update the code to test backupsInstance
        //instance.backupsInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('cancelExpirationInstance', function() {
      it('should call cancelExpirationInstance successfully', function(done) {
        //uncomment below and update the code to test cancelExpirationInstance
        //instance.cancelExpirationInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('cancelRemovalInstance', function() {
      it('should call cancelRemovalInstance successfully', function(done) {
        //uncomment below and update the code to test cancelRemovalInstance
        //instance.cancelRemovalInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('cancelShutdownInstance', function() {
      it('should call cancelShutdownInstance successfully', function(done) {
        //uncomment below and update the code to test cancelShutdownInstance
        //instance.cancelShutdownInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('cloneImageInstance', function() {
      it('should call cloneImageInstance successfully', function(done) {
        //uncomment below and update the code to test cloneImageInstance
        //instance.cloneImageInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('cloneInstance', function() {
      it('should call cloneInstance successfully', function(done) {
        //uncomment below and update the code to test cloneInstance
        //instance.cloneInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createInstanceSchedule', function() {
      it('should call createInstanceSchedule successfully', function(done) {
        //uncomment below and update the code to test createInstanceSchedule
        //instance.createInstanceSchedule(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteAllSnapshotsInstance', function() {
      it('should call deleteAllSnapshotsInstance successfully', function(done) {
        //uncomment below and update the code to test deleteAllSnapshotsInstance
        //instance.deleteAllSnapshotsInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteAllSnapshotsInstanceContainer', function() {
      it('should call deleteAllSnapshotsInstanceContainer successfully', function(done) {
        //uncomment below and update the code to test deleteAllSnapshotsInstanceContainer
        //instance.deleteAllSnapshotsInstanceContainer(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteInstance', function() {
      it('should call deleteInstance successfully', function(done) {
        //uncomment below and update the code to test deleteInstance
        //instance.deleteInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteInstanceSchedule', function() {
      it('should call deleteInstanceSchedule successfully', function(done) {
        //uncomment below and update the code to test deleteInstanceSchedule
        //instance.deleteInstanceSchedule(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteSnapshotInstance', function() {
      it('should call deleteSnapshotInstance successfully', function(done) {
        //uncomment below and update the code to test deleteSnapshotInstance
        //instance.deleteSnapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('ejectInstance', function() {
      it('should call ejectInstance successfully', function(done) {
        //uncomment below and update the code to test ejectInstance
        //instance.ejectInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('extendExpirationInstance', function() {
      it('should call extendExpirationInstance successfully', function(done) {
        //uncomment below and update the code to test extendExpirationInstance
        //instance.extendExpirationInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('extendShutdownInstance', function() {
      it('should call extendShutdownInstance successfully', function(done) {
        //uncomment below and update the code to test extendShutdownInstance
        //instance.extendShutdownInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getEnvVariables', function() {
      it('should call getEnvVariables successfully', function(done) {
        //uncomment below and update the code to test getEnvVariables
        //instance.getEnvVariables(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstance', function() {
      it('should call getInstance successfully', function(done) {
        //uncomment below and update the code to test getInstance
        //instance.getInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceContainers', function() {
      it('should call getInstanceContainers successfully', function(done) {
        //uncomment below and update the code to test getInstanceContainers
        //instance.getInstanceContainers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceHistory', function() {
      it('should call getInstanceHistory successfully', function(done) {
        //uncomment below and update the code to test getInstanceHistory
        //instance.getInstanceHistory(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceSchedule', function() {
      it('should call getInstanceSchedule successfully', function(done) {
        //uncomment below and update the code to test getInstanceSchedule
        //instance.getInstanceSchedule(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceSchedules', function() {
      it('should call getInstanceSchedules successfully', function(done) {
        //uncomment below and update the code to test getInstanceSchedules
        //instance.getInstanceSchedules(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceThreshold', function() {
      it('should call getInstanceThreshold successfully', function(done) {
        //uncomment below and update the code to test getInstanceThreshold
        //instance.getInstanceThreshold(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getInstanceTypeProvisioning', function() {
      it('should call getInstanceTypeProvisioning successfully', function(done) {
        //uncomment below and update the code to test getInstanceTypeProvisioning
        //instance.getInstanceTypeProvisioning(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getPrepareApplyInstance', function() {
      it('should call getPrepareApplyInstance successfully', function(done) {
        //uncomment below and update the code to test getPrepareApplyInstance
        //instance.getPrepareApplyInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getSnapshotInstance', function() {
      it('should call getSnapshotInstance successfully', function(done) {
        //uncomment below and update the code to test getSnapshotInstance
        //instance.getSnapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getStateInstance', function() {
      it('should call getStateInstance successfully', function(done) {
        //uncomment below and update the code to test getStateInstance
        //instance.getStateInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getValidateApplyInstance', function() {
      it('should call getValidateApplyInstance successfully', function(done) {
        //uncomment below and update the code to test getValidateApplyInstance
        //instance.getValidateApplyInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getWikiInstance', function() {
      it('should call getWikiInstance successfully', function(done) {
        //uncomment below and update the code to test getWikiInstance
        //instance.getWikiInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('importSnapshotInstance', function() {
      it('should call importSnapshotInstance successfully', function(done) {
        //uncomment below and update the code to test importSnapshotInstance
        //instance.importSnapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('linkedCloneSnapshotInstance', function() {
      it('should call linkedCloneSnapshotInstance successfully', function(done) {
        //uncomment below and update the code to test linkedCloneSnapshotInstance
        //instance.linkedCloneSnapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listInstanceServicePlans', function() {
      it('should call listInstanceServicePlans successfully', function(done) {
        //uncomment below and update the code to test listInstanceServicePlans
        //instance.listInstanceServicePlans(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listInstanceTypesProvisioning', function() {
      it('should call listInstanceTypesProvisioning successfully', function(done) {
        //uncomment below and update the code to test listInstanceTypesProvisioning
        //instance.listInstanceTypesProvisioning(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listInstances', function() {
      it('should call listInstances successfully', function(done) {
        //uncomment below and update the code to test listInstances
        //instance.listInstances(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listSecurityGroupsInstance', function() {
      it('should call listSecurityGroupsInstance successfully', function(done) {
        //uncomment below and update the code to test listSecurityGroupsInstance
        //instance.listSecurityGroupsInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('lockInstance', function() {
      it('should call lockInstance successfully', function(done) {
        //uncomment below and update the code to test lockInstance
        //instance.lockInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('refreshStateInstance', function() {
      it('should call refreshStateInstance successfully', function(done) {
        //uncomment below and update the code to test refreshStateInstance
        //instance.refreshStateInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeInstancesFromControl', function() {
      it('should call removeInstancesFromControl successfully', function(done) {
        //uncomment below and update the code to test removeInstancesFromControl
        //instance.removeInstancesFromControl(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('resizeInstance', function() {
      it('should call resizeInstance successfully', function(done) {
        //uncomment below and update the code to test resizeInstance
        //instance.resizeInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restartInstance', function() {
      it('should call restartInstance successfully', function(done) {
        //uncomment below and update the code to test restartInstance
        //instance.restartInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('revertSnapshotInstance', function() {
      it('should call revertSnapshotInstance successfully', function(done) {
        //uncomment below and update the code to test revertSnapshotInstance
        //instance.revertSnapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('runWorkflowInstance', function() {
      it('should call runWorkflowInstance successfully', function(done) {
        //uncomment below and update the code to test runWorkflowInstance
        //instance.runWorkflowInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('setApplyInstance', function() {
      it('should call setApplyInstance successfully', function(done) {
        //uncomment below and update the code to test setApplyInstance
        //instance.setApplyInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('setInstanceSecurityGroups', function() {
      it('should call setInstanceSecurityGroups successfully', function(done) {
        //uncomment below and update the code to test setInstanceSecurityGroups
        //instance.setInstanceSecurityGroups(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('snapshotInstance', function() {
      it('should call snapshotInstance successfully', function(done) {
        //uncomment below and update the code to test snapshotInstance
        //instance.snapshotInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('snapshotsInstance', function() {
      it('should call snapshotsInstance successfully', function(done) {
        //uncomment below and update the code to test snapshotsInstance
        //instance.snapshotsInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('startInstance', function() {
      it('should call startInstance successfully', function(done) {
        //uncomment below and update the code to test startInstance
        //instance.startInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopInstance', function() {
      it('should call stopInstance successfully', function(done) {
        //uncomment below and update the code to test stopInstance
        //instance.stopInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('suspendInstance', function() {
      it('should call suspendInstance successfully', function(done) {
        //uncomment below and update the code to test suspendInstance
        //instance.suspendInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('unlockInstance', function() {
      it('should call unlockInstance successfully', function(done) {
        //uncomment below and update the code to test unlockInstance
        //instance.unlockInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateInstance', function() {
      it('should call updateInstance successfully', function(done) {
        //uncomment below and update the code to test updateInstance
        //instance.updateInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateInstanceNetworkInterface', function() {
      it('should call updateInstanceNetworkInterface successfully', function(done) {
        //uncomment below and update the code to test updateInstanceNetworkInterface
        //instance.updateInstanceNetworkInterface(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateInstanceSchedule', function() {
      it('should call updateInstanceSchedule successfully', function(done) {
        //uncomment below and update the code to test updateInstanceSchedule
        //instance.updateInstanceSchedule(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateInstanceThreshold', function() {
      it('should call updateInstanceThreshold successfully', function(done) {
        //uncomment below and update the code to test updateInstanceThreshold
        //instance.updateInstanceThreshold(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateWikiInstance', function() {
      it('should call updateWikiInstance successfully', function(done) {
        //uncomment below and update the code to test updateWikiInstance
        //instance.updateWikiInstance(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
