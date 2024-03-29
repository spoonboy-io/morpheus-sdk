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
from openapi_client.model.update_host_install_agent_request_server_server_os import UpdateHostInstallAgentRequestServerServerOs
from openapi_client.model.update_host_managed_request_server_account import UpdateHostManagedRequestServerAccount
from openapi_client.model.update_host_managed_request_server_config import UpdateHostManagedRequestServerConfig
from openapi_client.model.update_host_managed_request_server_plan import UpdateHostManagedRequestServerPlan
from openapi_client.model.update_host_managed_request_server_tags_inner import UpdateHostManagedRequestServerTagsInner
globals()['UpdateHostInstallAgentRequestServerServerOs'] = UpdateHostInstallAgentRequestServerServerOs
globals()['UpdateHostManagedRequestServerAccount'] = UpdateHostManagedRequestServerAccount
globals()['UpdateHostManagedRequestServerConfig'] = UpdateHostManagedRequestServerConfig
globals()['UpdateHostManagedRequestServerPlan'] = UpdateHostManagedRequestServerPlan
globals()['UpdateHostManagedRequestServerTagsInner'] = UpdateHostManagedRequestServerTagsInner
from openapi_client.model.update_host_managed_request_server import UpdateHostManagedRequestServer


class TestUpdateHostManagedRequestServer(unittest.TestCase):
    """UpdateHostManagedRequestServer unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUpdateHostManagedRequestServer(self):
        """Test UpdateHostManagedRequestServer"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UpdateHostManagedRequestServer()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
