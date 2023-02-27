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

type Bookies interface {
	// AllBookies Gets raw information for all the bookies in the cluster
	AllBookies() (*AllBookiesResp, error)
	// ListRacksInfo Gets the rack placement information for all the bookies in the cluster
	ListRacksInfo() (*ListRacksInfoResp, error)
	// RemoveRacksInfo Removed the rack placement information for a specific bookie in the cluster
	RemoveRacksInfo(string) error
	// GetRacksInfo Gets the rack placement information for a specific bookie in the cluster
	GetRacksInfo(string) (*BookieInfo, error)
	// UpdateRacksInfo Updates the rack placement information for a specific bookie in the cluster
	// (note. bookie address format:`address:port`)
	UpdateRacksInfo(string, *BookieInfo) error
}

type RawBookieInfo struct {
	BookieId string `json:"bookieId"`
}

type BookieInfo struct {
	Rack     string `json:"rack"`
	Hostname string `json:"hostname"`
}

type AllBookiesResp struct {
	Bookies []RawBookieInfo `json:"bookies"`
}

type ListRacksInfoResp map[string]map[string]BookieInfo
