// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package padmin

import "strconv"

type Config struct {
	Host      string
	Port      int
	urlPrefix string
}

type PulsarAdmin struct {
	Clusters            *Clusters
	Tenants             *Tenants
	Namespaces          *Namespaces
	PersistentTopics    *PersistentTopics
	NonPersistentTopics *NonPersistentTopics
}

func NewDefaultPulsarAdmin() (*PulsarAdmin, error) {
	return NewPulsarAdmin(Config{})
}

func NewPulsarAdmin(config Config) (*PulsarAdmin, error) {
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 8080
	}
	config.urlPrefix = "http://" + config.Host + ":" + strconv.Itoa(config.Port)
	client, err := newHttpClient(config)
	if err != nil {
		return nil, err
	}
	return &PulsarAdmin{
		Clusters:            newClusters(client),
		Tenants:             newTenants(client),
		Namespaces:          newNamespaces(client),
		PersistentTopics:    newPersistentTopics(client),
		NonPersistentTopics: newNonPersistentTopics(client),
	}, nil
}
