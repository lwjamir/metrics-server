// Copyright 2018 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package summary

import (
	"fmt"

	"k8s.io/client-go/rest"
)

// GetKubeletConfig fetches connection config for connecting to the Kubelet.
func GetKubeletConfig(baseKubeConfig *rest.Config, port int, portIsInsecure bool) *KubeletClientConfig {
	kubeletConfig := &KubeletClientConfig{
		// TODO: deprecate and remove this option
		PortIsInsecure: portIsInsecure,
		Port:           uint(port),
		RESTConfig:     baseKubeConfig,
	}

	return kubeletConfig
}

type KubeletClientConfig struct {
	PortIsInsecure bool
	Port           uint
	RESTConfig     *rest.Config
}

func KubeletClientFor(config *KubeletClientConfig) (KubeletInterface, error) {
	transport, err := rest.TransportFor(config.RESTConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to construct transport: %v", err)
	}

	return NewKubeletClient(transport, config.Port, config.PortIsInsecure)
}
