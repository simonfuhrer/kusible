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

package inventory

import (
	"github.com/bedag/kusible/pkg/loader"
	"github.com/bedag/kusible/pkg/values"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Inventory struct {
	entries map[string]*Entry
	ejson   *values.EjsonSettings
}

type Entry struct {
	name            string
	groups          []string
	configNamespace string
	kubeconfig      *Kubeconfig
}

type Kubeconfig struct {
	loader loader.Loader
	config *clientcmdapi.Config
}
