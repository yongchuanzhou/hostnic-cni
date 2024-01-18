/*
Copyright 2020 KubeSphere Authors

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

package ipam

type PoolInfo struct {
	IPPool string
	Block  string
}

// AutoAssignArgs defines the set of arguments for assigning one or more
// IP addresses.
type AutoAssignArgs struct {
	HandleID string

	// A key/value mapping of metadata to store with the allocations.
	Attrs map[string]string

	Pool string

	// extarnal
	Pools  []string
	Blocks []string
	Info   *PoolInfo
}

// GetUtilizationArgs defines the set of arguments for requesting IP utilization.
type GetUtilizationArgs struct {
	// If specified, the pools whose utilization should be reported.  Each string here
	// can be a pool name or CIDR.  If not specified, this defaults to all pools.
	Pools []string
	// NSAndBlocks map[string][]string //ns and it's ipamblocks
}

// PoolUtilization reports IP utilization for a single IP pool.
type PoolUtilization struct {
	// This pool's name.
	Name string

	// Number of possible IPs in this block.
	Capacity int

	// Number of available IPs in this block.
	Unallocated int

	Allocate int

	Reserved int
}

type BlockUtilization struct {
	// This block's name.
	Name string

	// Number of possible IPs in this block.
	Capacity int

	// Number of available IPs in this block.
	Unallocated int

	Allocate int

	Reserved int
}

type PoolBlocksUtilization struct {
	// This pool's name.
	Name string

	// Number of possible IPs in this block.
	Capacity int

	// Number of available IPs in this block.
	Unallocated int

	Allocate int

	Reserved int

	// This blocks' util which belong to pool.
	Blocks []*BlockUtilization

	// This blocks' util which belong to pool.
	BrokenBlocks          []*BrokenBlockUtilization
	BrokenBlockNames      []string
	BrokenBlockFixSucceed []string
}

type BrokenBlockUtilization struct {
	// This block's name.
	Name string

	IpToPods                map[string][]string
	IpWithoutRecord         map[string][]string // maybe allocated to more than one pod
	IpToPodsWithWrongRecord map[string][]string
}

// FixIpArgs assign from config annotation
type FixIpArgs struct {
	AutoAssignArgs
	//annotation ips
	TarGetIpList []string
}
