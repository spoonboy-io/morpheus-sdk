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
from openapi_client.model.task_ansible_playbook_config_file import TaskAnsiblePlaybookConfigFile
from openapi_client.model.task_library_script_config_task_options import TaskLibraryScriptConfigTaskOptions
from openapi_client.model.task_library_script_config_task_type import TaskLibraryScriptConfigTaskType
from openapi_client.model.zone_credential_any_of import ZoneCredentialAnyOf
globals()['TaskAnsiblePlaybookConfigFile'] = TaskAnsiblePlaybookConfigFile
globals()['TaskLibraryScriptConfigTaskOptions'] = TaskLibraryScriptConfigTaskOptions
globals()['TaskLibraryScriptConfigTaskType'] = TaskLibraryScriptConfigTaskType
globals()['ZoneCredentialAnyOf'] = ZoneCredentialAnyOf
from openapi_client.model.task_library_script_config import TaskLibraryScriptConfig


class TestTaskLibraryScriptConfig(unittest.TestCase):
    """TaskLibraryScriptConfig unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testTaskLibraryScriptConfig(self):
        """Test TaskLibraryScriptConfig"""
        # FIXME: construct object with mandatory attributes with example values
        # model = TaskLibraryScriptConfig()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
