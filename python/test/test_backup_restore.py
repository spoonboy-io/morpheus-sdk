"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.backup_instance import BackupInstance
from openapi_client.model.backup_restore_backup import BackupRestoreBackup
from openapi_client.model.backup_restore_container import BackupRestoreContainer
globals()['BackupInstance'] = BackupInstance
globals()['BackupRestoreBackup'] = BackupRestoreBackup
globals()['BackupRestoreContainer'] = BackupRestoreContainer
from openapi_client.model.backup_restore import BackupRestore


class TestBackupRestore(unittest.TestCase):
    """BackupRestore unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testBackupRestore(self):
        """Test BackupRestore"""
        # FIXME: construct object with mandatory attributes with example values
        # model = BackupRestore()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
