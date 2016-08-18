/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"net"
	"net/http"
	"net/rpc"
	"os"
	"testing"
)

// rpc path for the tests
const testRPCPath = "/Echo"
const testDebugPath = "/Debug/Echo"

// Simple echo server
type testServer struct {
}

// Echo - returns the received message without modifying.
func (server *testServer) Echo(msg *string, reply *string) error {
	*reply = *msg
	return nil
}

// testRPCServer - starts a simple Echo server on port 1234
func testRPCServer() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterName("Test", &testServer{})
	rpcServer.HandleHTTP(testRPCPath, testDebugPath)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		new(testing.T).Errorf("listen error:", e)
		return
	}
	go http.Serve(l, nil)
}

// Setup function.
func TestMain(m *testing.M) {
	testRPCServer()
	os.Exit(m.Run())
}

func TestBasicRPCClient(t *testing.T) {
	clnt := newClient("localhost:1234", testRPCPath)
	var reply string
	err := clnt.Call("Test.Echo", "Hello", &reply)
	if err != nil {
		t.Errorf("Test.Echo RPC failed with ", err)
	}

}
