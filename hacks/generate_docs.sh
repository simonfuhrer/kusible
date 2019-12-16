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

PACKAGES="$(go list -f "{{.Name}}:{{.ImportPath}}" ./... | grep -v -E "main:|vendor/|examples" | cut -d ":" -f 2)"

# loop over all packages generating all their documentation
for pkg in $PACKAGES; do
  echo "godoc2md $pkg > src/$pkg/README.md"
  godoc2md $pkg > src/$pkg/README.md
done
