// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package test

import (
	"fmt"
	"web3go/common"
	"web3go/rpc"
)

// MockAdminAPI ...
type MockAdminAPI struct {
	rpc rpc.RPC
}

// NewMockNetAPI ...
func NewMockAdminAPI(rpc rpc.RPC) MockAPI {
	return &MockAdminAPI{rpc: rpc}
}

// Do ...
func (admin *MockAdminAPI) Do(request rpc.Request) (response rpc.Response, err error) {
	method := request.Get("method").(string)
	switch method {
	case "admin_nodeInfo":
		nodeInfo := common.NodeInfo{
			Enode: 			string("enode://d1a4eab40a2f846478e87915d9c13b3a5a68025df80a211f9af8db6e40b64c5301b44593540ef5745e8aaaacee58ea06f36029803fb5f80d991ff62edb703686@34.243.28.236:21000"),
			Enr: 			string("0xf89db8402d58b4aafeaceb7f6caff4dd83d3eda7ee08790cbafdf8cc0d914e7d7023996e4030502adaed4938943a141e6c399b7bb8c4e8c2632c5c663499918029ac9b5082019883636170cbca88736d696c6f626674408269648276348269708422f31cec89736563703235366b31a102d1a4eab40a2f846478e87915d9c13b3a5a68025df80a211f9af8db6e40b64c538374637082520883756470825208"),
			Id: 			string("12200bbb0ad607e5ee158d5a644097fadec1401f773d2407744006deeea5c12c"),
		}
		return generateResponse(admin.rpc, request, nodeInfo)
	}

	return nil, fmt.Errorf("Invalid method %s", method)
}
