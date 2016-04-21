### etcdserverpb


##### service `Auth` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| AuthEnable | AuthEnableRequest | AuthEnableResponse | AuthEnable enables authentication. |
| AuthDisable | AuthDisableRequest | AuthDisableResponse | AuthDisable disables authentication. |
| Authenticate | AuthenticateRequest | AuthenticateResponse | Authenticate processes authenticate request. |
| UserAdd | AuthUserAddRequest | AuthUserAddResponse | UserAdd adds a new user. |
| UserGet | AuthUserGetRequest | AuthUserGetResponse | UserGet gets a detailed information of a user or lists entire users. |
| UserDelete | AuthUserDeleteRequest | AuthUserDeleteResponse | UserDelete deletes a specified user. |
| UserChangePassword | AuthUserChangePasswordRequest | AuthUserChangePasswordResponse | UserChangePassword changes password of a specified user. |
| UserGrant | AuthUserGrantRequest | AuthUserGrantResponse | UserGrant grants a role to a specified user. |
| UserRevoke | AuthUserRevokeRequest | AuthUserRevokeResponse | UserRevoke revokes a role of specified user. |
| RoleAdd | AuthRoleAddRequest | AuthRoleAddResponse | RoleAdd adds a new role. |
| RoleGet | AuthRoleGetRequest | AuthRoleGetResponse | RoleGet gets a detailed information of a role or lists entire roles. |
| RoleDelete | AuthRoleDeleteRequest | AuthRoleDeleteResponse | RoleDelete deletes a specified role. |
| RoleGrant | AuthRoleGrantRequest | AuthRoleGrantResponse | RoleGrant grants a permission of a specified key or range to a specified role. |
| RoleRevoke | AuthRoleRevokeRequest | AuthRoleRevokeResponse | RoleRevoke revokes a key or range permission of a specified role. |



##### service `Cluster` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| MemberAdd | MemberAddRequest | MemberAddResponse | MemberAdd adds a member into the cluster. |
| MemberRemove | MemberRemoveRequest | MemberRemoveResponse | MemberRemove removes an existing member from the cluster. |
| MemberUpdate | MemberUpdateRequest | MemberUpdateResponse | MemberUpdate updates the member configuration. |
| MemberList | MemberListRequest | MemberListResponse | MemberList lists all the members in the cluster. |



##### service `KV` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Range | RangeRequest | RangeResponse | Range gets the keys in the range from the store. |
| Put | PutRequest | PutResponse | Put puts the given key into the store. A put request increases the revision of the store, and generates one event in the event history. |
| DeleteRange | DeleteRangeRequest | DeleteRangeResponse | Delete deletes the given range from the store. A delete request increase the revision of the store, and generates one event in the event history. |
| Txn | TxnRequest | TxnResponse | Txn processes all the requests in one transaction. A txn request increases the revision of the store, and generates events with the same revision in the event history. It is not allowed to modify the same key several times within one txn. |
| Compact | CompactionRequest | CompactionResponse | Compact compacts the event history in etcd. User should compact the event history periodically, or it will grow infinitely. |



##### service `Lease` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| LeaseGrant | LeaseGrantRequest | LeaseGrantResponse | LeaseGrant creates a lease. A lease has a TTL. The lease will expire if the server does not receive a keepAlive within TTL from the lease holder. All keys attached to the lease will be expired and deleted if the lease expires. The key expiration generates an event in event history. |
| LeaseRevoke | LeaseRevokeRequest | LeaseRevokeResponse | LeaseRevoke revokes a lease. All the key attached to the lease will be expired and deleted. |
| LeaseKeepAlive | LeaseKeepAliveRequest | LeaseKeepAliveResponse | KeepAlive keeps the lease alive. |



##### service `Maintenance` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Alarm | AlarmRequest | AlarmResponse | Alarm activates, deactivates, and queries alarms regarding cluster health. |
| Status | StatusRequest | StatusResponse | Status gets the status of the member. |
| Defragment | DefragmentRequest | DefragmentResponse |  |
| Hash | HashRequest | HashResponse | Hash returns the hash of the local KV state for consistency checking purpose. This is designed for testing; do not use this in production when there are ongoing transactions. |
| Snapshot | SnapshotRequest | SnapshotResponse | Snapshot sends a snapshot of the entire backend |



##### service `Watch` (testdata/rpc.proto)

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Watch | WatchRequest | WatchResponse | Watch watches the events happening or happened. Both input and output are stream. One watch rpc can watch for multiple keys or prefixs and get a stream of events. The whole events history can be watched unless compacted. |



##### message `AlarmMember` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| memberID |  | uint64 | uint64 | long | int/long | uint64 |
| alarm |  | AlarmType | | | | |



##### message `AlarmRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| action |  | AlarmAction | | | | |
| memberID | MemberID is the member raising the alarm request | uint64 | uint64 | long | int/long | uint64 |
| alarm |  | AlarmType | | | | |



##### message `AlarmResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| alarms |  | (slice of) AlarmMember | | | | |



##### message `AuthDisableRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthDisableResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthEnableRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthEnableResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleAddRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| name |  | string | string | String | str/unicode | string |



##### message `AuthRoleAddResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleDeleteRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthRoleDeleteResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleGetRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthRoleGetResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleGrantRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| name |  | string | string | String | str/unicode | string |
| perm |  | authpb.Permission | | | | |



##### message `AuthRoleGrantResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleRevokeRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthRoleRevokeResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserAddRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| name |  | string | string | String | str/unicode | string |
| password |  | string | string | String | str/unicode | string |



##### message `AuthUserAddResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserChangePasswordRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| name |  | string | string | String | str/unicode | string |
| password |  | string | string | String | str/unicode | string |



##### message `AuthUserChangePasswordResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserDeleteRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| name |  | string | string | String | str/unicode | string |



##### message `AuthUserDeleteResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserGetRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthUserGetResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserGrantRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| user |  | string | string | String | str/unicode | string |
| role |  | string | string | String | str/unicode | string |



##### message `AuthUserGrantResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthUserRevokeRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthUserRevokeResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `AuthenticateRequest` (testdata/rpc.proto)

Empty field.



##### message `AuthenticateResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `CompactionRequest` (testdata/rpc.proto)

Compaction compacts the kv store upto a given revision. All superseded keys with a revision less than the compaction revision will be removed.

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| revision |  | int64 | int64 | long | int/long | int64 |
| physical | physical is set so the RPC will wait until the compaction is physically applied to the local database such that compacted entries are totally removed from the backing store. | bool | bool | boolean | boolean | bool |



##### message `CompactionResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `Compare` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| result |  | CompareResult | | | | |
| target |  | CompareTarget | | | | |
| key | key path | bytes | []byte | ByteString | str | string |
| target_union |  | oneof | | | | |
| version | version of the given key | int64 | int64 | long | int/long | int64 |
| create_revision | create revision of the given key | int64 | int64 | long | int/long | int64 |
| mod_revision | last modified revision of the given key | int64 | int64 | long | int/long | int64 |
| value | value of the given key | bytes | []byte | ByteString | str | string |



##### message `DefragmentRequest` (testdata/rpc.proto)

Empty field.



##### message `DefragmentResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `DeleteRangeRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| key | if the range_end is not given, the request deletes the key. | bytes | []byte | ByteString | str | string |
| range_end | if the range_end is given, it deletes the keys in range [key, range_end). | bytes | []byte | ByteString | str | string |



##### message `DeleteRangeResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| deleted | Deleted is the number of keys that got deleted. | int64 | int64 | long | int/long | int64 |



##### message `EmptyResponse` (testdata/raft_internal.proto)

Empty field.



##### message `HashRequest` (testdata/rpc.proto)

Empty field.



##### message `HashResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| hash |  | uint32 | uint32 | int | int/long | uint32 |



##### message `InternalRaftRequest` (testdata/raft_internal.proto)

An InternalRaftRequest is the union of all requests which can be sent via raft.

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | uint64 | uint64 | long | int/long | uint64 |
| v2 |  | Request | | | | |
| range |  | RangeRequest | | | | |
| put |  | PutRequest | | | | |
| delete_range |  | DeleteRangeRequest | | | | |
| txn |  | TxnRequest | | | | |
| compaction |  | CompactionRequest | | | | |
| lease_grant |  | LeaseGrantRequest | | | | |
| lease_revoke |  | LeaseRevokeRequest | | | | |
| auth_enable |  | AuthEnableRequest | | | | |
| auth_user_add |  | AuthUserAddRequest | | | | |
| auth_user_delete |  | AuthUserDeleteRequest | | | | |
| auth_user_change_password |  | AuthUserChangePasswordRequest | | | | |
| auth_user_grant |  | AuthUserGrantRequest | | | | |
| auth_role_add |  | AuthRoleAddRequest | | | | |
| auth_role_grant |  | AuthRoleGrantRequest | | | | |
| alarm |  | AlarmRequest | | | | |



##### message `LeaseGrantRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| TTL | advisory ttl in seconds | int64 | int64 | long | int/long | int64 |
| ID | requested ID to create; 0 lets lessor choose | int64 | int64 | long | int/long | int64 |



##### message `LeaseGrantResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| ID |  | int64 | int64 | long | int/long | int64 |
| TTL | server decided ttl in second | int64 | int64 | long | int/long | int64 |
| error |  | string | string | String | str/unicode | string |



##### message `LeaseKeepAliveRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | int64 | int64 | long | int/long | int64 |



##### message `LeaseKeepAliveResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| ID |  | int64 | int64 | long | int/long | int64 |
| TTL |  | int64 | int64 | long | int/long | int64 |



##### message `LeaseRevokeRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | int64 | int64 | long | int/long | int64 |



##### message `LeaseRevokeResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `Member` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | uint64 | uint64 | long | int/long | uint64 |
| name | If the member is not started, name will be an empty string. | string | string | String | str/unicode | string |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode | (slice of) string |
| clientURLs | If the member is not started, client_URLs will be an zero length string array. | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode | (slice of) string |



##### message `MemberAddRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode | (slice of) string |



##### message `MemberAddResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| member |  | Member | | | | |



##### message `MemberListRequest` (testdata/rpc.proto)

Empty field.



##### message `MemberListResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| members |  | (slice of) Member | | | | |



##### message `MemberRemoveRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | uint64 | uint64 | long | int/long | uint64 |



##### message `MemberRemoveResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `MemberUpdateRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | uint64 | uint64 | long | int/long | uint64 |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode | (slice of) string |



##### message `MemberUpdateResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `Metadata` (testdata/etcdserver.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| NodeID |  | uint64 | uint64 | long | int/long | uint64 |
| ClusterID |  | uint64 | uint64 | long | int/long | uint64 |



##### message `PutRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| key |  | bytes | []byte | ByteString | str | string |
| value |  | bytes | []byte | ByteString | str | string |
| lease |  | int64 | int64 | long | int/long | int64 |



##### message `PutResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |



##### message `RangeRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| key | if the range_end is not given, the request returns the key. | bytes | []byte | ByteString | str | string |
| range_end | if the range_end is given, it gets the keys in range [key, range_end) if range_end is nonempty, otherwise it returns all keys >= key. | bytes | []byte | ByteString | str | string |
| limit | limit the number of keys returned. | int64 | int64 | long | int/long | int64 |
| revision | range over the store at the given revision. if revision is less or equal to zero, range over the newest store. if the revision has been compacted, ErrCompaction will be returned in response. | int64 | int64 | long | int/long | int64 |
| sort_order | sort_order is the requested order for returned the results | SortOrder | | | | |
| sort_target | sort_target is the kv field to use for sorting | SortTarget | | | | |
| serializable | range request is linearizable by default. Linearizable requests has a higher latency and lower throughput than serializable request. To reduce latency, serializable can be set. If serializable is set, range request will be serializable, but not linearizable with other requests. Serializable range can be served locally without waiting for other nodes in the cluster. | bool | bool | boolean | boolean | bool |



##### message `RangeResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| kvs |  | (slice of) storagepb.KeyValue | | | | |
| more | more indicates if there are more keys to return in the requested range. | bool | bool | boolean | boolean | bool |



##### message `Request` (testdata/etcdserver.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| ID |  | uint64 | uint64 | long | int/long | uint64 |
| Method |  | string | string | String | str/unicode | string |
| Path |  | string | string | String | str/unicode | string |
| Val |  | string | string | String | str/unicode | string |
| Dir |  | bool | bool | boolean | boolean | bool |
| PrevValue |  | string | string | String | str/unicode | string |
| PrevIndex |  | uint64 | uint64 | long | int/long | uint64 |
| PrevExist |  | bool | bool | boolean | boolean | bool |
| Expiration |  | int64 | int64 | long | int/long | int64 |
| Wait |  | bool | bool | boolean | boolean | bool |
| Since |  | uint64 | uint64 | long | int/long | uint64 |
| Recursive |  | bool | bool | boolean | boolean | bool |
| Sorted |  | bool | bool | boolean | boolean | bool |
| Quorum |  | bool | bool | boolean | boolean | bool |
| Time |  | int64 | int64 | long | int/long | int64 |
| Stream |  | bool | bool | boolean | boolean | bool |
| Refresh |  | bool | bool | boolean | boolean | bool |



##### message `RequestUnion` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| request |  | oneof | | | | |
| request_range |  | RangeRequest | | | | |
| request_put |  | PutRequest | | | | |
| request_delete_range |  | DeleteRangeRequest | | | | |



##### message `ResponseHeader` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| cluster_id |  | uint64 | uint64 | long | int/long | uint64 |
| member_id |  | uint64 | uint64 | long | int/long | uint64 |
| revision | revision of the store when the request was applied. | int64 | int64 | long | int/long | int64 |
| raft_term | term of raft when the request was applied. | uint64 | uint64 | long | int/long | uint64 |



##### message `ResponseUnion` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| response |  | oneof | | | | |
| response_range |  | RangeResponse | | | | |
| response_put |  | PutResponse | | | | |
| response_delete_range |  | DeleteRangeResponse | | | | |



##### message `SnapshotRequest` (testdata/rpc.proto)

Empty field.



##### message `SnapshotResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header | header has the current store information. The first header in the snapshot stream indicates the point in time of the snapshot. | ResponseHeader | | | | |
| remaining_bytes | remaining_bytes is the number of blob bytes to be sent after this message | uint64 | uint64 | long | int/long | uint64 |
| blob | blob has the next chunk of the snapshot in the snapshot stream. | bytes | []byte | ByteString | str | string |



##### message `StatusRequest` (testdata/rpc.proto)

Empty field.



##### message `StatusResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| version |  | string | string | String | str/unicode | string |
| dbSize |  | int64 | int64 | long | int/long | int64 |
| leader |  | uint64 | uint64 | long | int/long | uint64 |
| raftIndex |  | uint64 | uint64 | long | int/long | uint64 |
| raftTerm |  | uint64 | uint64 | long | int/long | uint64 |



##### message `TxnRequest` (testdata/rpc.proto)

If the comparisons succeed, then the success requests will be processed in order, and the response will contain their respective responses in order. If the comparisons fail, then the failure requests will be processed in order, and the response will contain their respective responses in order. From google paxosdb paper: Our implementation hinges around a powerful primitive which we call MultiOp. All other database operations except for iteration are implemented as a single call to MultiOp. A MultiOp is applied atomically and consists of three components: 1. A list of tests called guard. Each test in guard checks a single entry in the database. It may check for the absence or presence of a value, or compare with a given value. Two different tests in the guard may apply to the same or different entries in the database. All tests in the guard are applied and MultiOp returns the results. If all tests are true, MultiOp executes t op (see item 2 below), otherwise it executes f op (see item 3 below). 2. A list of database operations called t op. Each operation in the list is either an insert, delete, or lookup operation, and applies to a single database entry. Two different operations in the list may apply to the same or different entries in the database. These operations are executed if guard evaluates to true. 3. A list of database operations called f op. Like t op, but executed if guard evaluates to false.

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| compare |  | (slice of) Compare | | | | |
| success |  | (slice of) RequestUnion | | | | |
| failure |  | (slice of) RequestUnion | | | | |



##### message `TxnResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| succeeded |  | bool | bool | boolean | boolean | bool |
| responses |  | (slice of) ResponseUnion | | | | |



##### message `WatchCancelRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| watch_id |  | int64 | int64 | long | int/long | int64 |



##### message `WatchCreateRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| key | the key to be watched | bytes | []byte | ByteString | str | string |
| range_end | if the range_end is given, keys in [key, range_end) are watched NOTE: only range_end == prefixEnd(key) is accepted now | bytes | []byte | ByteString | str | string |
| start_revision | start_revision is an optional revision (including) to watch from. No start_revision is "now". | int64 | int64 | long | int/long | int64 |
| progress_notify | if progress_notify is set, etcd server sends WatchResponse with empty events to the created watcher when there are no recent events. It is useful when clients want always to be able to recover a disconnected watcher from a recent known revision. etcdsever can decide how long it should send a notification based on current load. | bool | bool | boolean | boolean | bool |



##### message `WatchRequest` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| request_union |  | oneof | | | | |
| create_request |  | WatchCreateRequest | | | | |
| cancel_request |  | WatchCancelRequest | | | | |



##### message `WatchResponse` (testdata/rpc.proto)

| Field | Description | Type | Go | Java | Python | C++ |
| ----- | ----------- | ---- | --- | ---- | ------ | --- |
| header |  | ResponseHeader | | | | |
| watch_id | watch_id is the ID of the watching the response sent to. | int64 | int64 | long | int/long | int64 |
| created | If the response is for a create watch request, created is set to true. Client should record the watch_id and prepare for receiving events for that watching from the same stream. All events sent to the created watching will attach with the same watch_id. | bool | bool | boolean | boolean | bool |
| canceled | If the response is for a cancel watch request, cancel is set to true. No further events will be sent to the canceled watching. | bool | bool | boolean | boolean | bool |
| compact_revision | CompactRevision is set to the minimum index if a watching tries to watch at a compacted index.  This happens when creating a watching at a compacted revision or the watching cannot catch up with the progress of the KV.  Client should treat the watching as canceled and should not try to create any watching with same start_revision again. | int64 | int64 | long | int/long | int64 |
| events |  | (slice of) storagepb.Event | | | | |



