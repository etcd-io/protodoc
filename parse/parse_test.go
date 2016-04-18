// Copyright 2016 CoreOS, Inc.
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

package parse

import "testing"

func TestReadDir(t *testing.T) {
	proto, err := ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	proto.Sort()

	if proto.Messages[51].Name != "MemberListRequest" {
		t.Fatalf("expected 'MemberListRequest', got %+v", proto.Messages[51])
	}
	if proto.Messages[52].Name != "MemberListResponse" {
		t.Fatalf("expected 'MemberListResponse', got %+v", proto.Messages[52])
	}
	if proto.Messages[52].Fields[0].Name != "header" {
		t.Fatalf("expected 'header', got %+v", proto.Messages[52])
	}
	if proto.Messages[52].Fields[0].UserDefinedProtoType != "ResponseHeader" {
		t.Fatalf("expected 'ResponseHeader', got %+v", proto.Messages[52])
	}

	if proto.Messages[75].Name != "WatchResponse" {
		t.Fatalf("expected 'WatchResponse', got %+v", proto.Messages[75])
	}
	if proto.Messages[75].Fields[0].Name != "header" {
		t.Fatalf("expected 'header', got %+v", proto.Messages[75])
	}
	if proto.Messages[75].Fields[1].ProtoType != Int64 {
		t.Fatalf("expected 'Int64', got %+v", proto.Messages[75])
	}
	if proto.Messages[75].Fields[1].Name != "watch_id" {
		t.Fatalf("expected 'watch_id', got %+v", proto.Messages[75])
	}
	if proto.Messages[75].Fields[0].UserDefinedProtoType != "ResponseHeader" {
		t.Fatalf("expected 'ResponseHeader', got %+v", proto.Messages[75])
	}

	if len(proto.Services[3].Methods) != 3 {
		t.Fatalf("expected 3 methods, got %+v", proto.Services[3].Methods)
	}
	if proto.Services[3].Methods[0].Name != "LeaseGrant" {
		t.Fatalf("expected 'LeaseGrant', got %+v", proto.Services[3].Methods[0])
	}
	if proto.Services[3].Methods[1].Name != "LeaseRevoke" {
		t.Fatalf("expected 'LeaseRevoke', got %+v", proto.Services[3].Methods[1])
	}
	if proto.Services[3].Methods[2].Name != "LeaseKeepAlive" {
		t.Fatalf("expected 'LeaseKeepAlive', got %+v", proto.Services[3].Methods[2])
	}

	if proto.Services[5].Name != "Watch" {
		t.Fatalf("expected 'Watch', got %+v", proto.Services[5])
	}
	if proto.Services[5].Methods[0].Name != "Watch" {
		t.Fatalf("expected 'Watch', got %+v", proto.Services[5])
	}
	if proto.Services[5].Methods[0].RequestType != "WatchRequest" {
		t.Fatalf("expected 'WatchRequest', got %+v", proto.Services[5])
	}
	if proto.Services[5].Methods[0].ResponseType != "WatchResponse" {
		t.Fatalf("expected 'WatchRequest', got %+v", proto.Services[5])
	}
}
