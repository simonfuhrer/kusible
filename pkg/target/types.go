/*
Copyright © 2019 Michael Gruener

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package target

import (
	inv "github.com/bedag/kusible/pkg/inventory"
	"github.com/bedag/kusible/pkg/values"
)

type Targets struct {
	limits     []string
	filter     string
	valuesPath string
	ejson      *values.EjsonSettings
	targets    map[string]*Target
}

type Target struct {
	entry  *inv.Entry
	values values.Values
}
