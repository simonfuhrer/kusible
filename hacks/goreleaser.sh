#!/bin/bash

# Copyright 2019 the Bedag contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail


GIT_DIRTY=$(git status --porcelain 2> /dev/null)
if [[ -z "${GIT_DIRTY}" ]]; then
    export GIT_TREE_STATE=clean
else
    export GIT_TREE_STATE=dirty
fi

# $PUBLISH must explicitly be set to 'true' for goreleaser
# to publish the release to GitHub.
if [[ "${PUBLISH:-}" != "true" ]]; then
    goreleaser release \
        --rm-dist \
        --skip-publish \
        --snapshot
else
    goreleaser release \
        --rm-dist
fi
