// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://f347242f5b8556a1b5662f3809e4b8e638b7a5ec063b61d68006c2736ee09f12bfb74df6bb6e72e79dd2e5ceed572f80f412e161e4f6ec5344adf2decceaf80d@35.163.165.80:31818",
	"enode://6773392b2218427a4efa7da736625e8018813443973103257402772336fbf3f08f57553d7aad39aef8a0993657f74f44815107a9ef1ad20f8cf9064aab93e26d@54.172.143.189:31818",
	"enode://7113b2433b987a09a5b46b3d4b86884eb797159e44864f8c75c61809182478017a7874cc3065e0599e656b834ef7f36a02e0db930367de8053cdc52d6269359a@3.64.146.10:31818",
	"enode://3db1081ad1d977914227ee959029415766a9065826958ba54a4366635dd04fa817a5b2178cd859976d6b5c1194085114521631e2b1a7dadf9a9430df7d76def8@54.255.57.25:31818",
	"enode://0f25489711ad39240cc56b96a3d0e0a470eae9e3e34c3e0531d075f393999ad8aa85d2018f1780149f5fecfd06fd6468058b2d782bfbd48b68eecd7bca5ced89@44.194.154.194:31818",
	"enode://9b54b04971c6f6178b3ab3b9726036014d4c305b8019ff0be92dacf0b80ea30d81455bcce459fc2d12ee1477e91c2608e93e6abe39ea5c3b2c4233bc2f370bb7@3.65.110.71:31818",
	"enode://451338ee184407e10ac5449e67de4f8b3a1d776e524dfbe12102ab5b473550e0d3f72132c10625873fb4bde9592b97fe9cf3ea022c8e037c84718929e1f725e1@18.138.63.212:31818",
	"enode://06ce0b5ac00e2a7e9b68e3e048c6c6603615aefcf7082391fbcede3895effd39e1b2076ae8710640729014c4f5a0af930ff9938563dde29ed86657b0401c6651@44.225.80.85:31818",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://056ff9068fec422fa2e81a318f5fc387e5365328b8b9a9fcc7784dcdf1da785ec5587450a29b25c77a696ad6342ee0d37a2daf293d6c70f43cedcb738fd11fdd@35.155.130.142:31818",
	"enode://3a1b6b7f4b9459e0f4ed302ceee7d30a2c6aa353818ee1f1dbd5ba80d8a8a118db113880977fad12085d1519da466fef9318c8d6381fc22101f9e05887923924@35.161.230.65:31818",
	"enode://bf25ac807ce0c430ddc375422e15218b568f3db5eebc670dc2092700e5d4e67d6746793a46db8bb407c3df4bb4dda048894695e3522bbf9097404159ff76b88d@35.155.152.134:31818",
	"enode://fd2606356799c0b5954d14913128596ce870a6111094294a05f3c7e852931afe06e8056303f5955ee271c04baf19820636456c48cb7ca4a0d9b3da57d01b5fff@44.228.225.190:31818",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
