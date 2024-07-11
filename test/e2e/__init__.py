# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

import pytest
import yaml
from typing import Dict, Any
from pathlib import Path

from acktest.resources import load_resource_file, _replace_placeholder_values, default_placeholder_values

SERVICE_NAME = "acmpca"
CRD_GROUP = "acmpca.services.k8s.aws"
CRD_VERSION = "v1alpha1"

# PyTest marker for the current service
service_marker = pytest.mark.service(arg=SERVICE_NAME)

bootstrap_directory = Path(__file__).parent
resource_directory = Path(__file__).parent / "resources"

def load_acmpca_resource(resource_name: str, additional_replacements: Dict[str, Any] = {}):
    """ Overrides the default `load_resource_file` to access the specific resources
    directory for the current service.
    """
    return load_resource_file(resource_directory, resource_name, additional_replacements=additional_replacements)

def load_all_acmpca_resources(resource_name: str, additional_replacements: Dict[str, Any] = {}):
    with open(resource_directory / f"{resource_name}.yaml", "r") as stream:
        resource_contents = stream.read()
        injected_contents = _replace_placeholder_values(
            resource_contents, default_placeholder_values())
        injected_contents = _replace_placeholder_values(
            injected_contents, additional_replacements)
        return yaml.safe_load_all(injected_contents)