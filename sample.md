### testdata


##### service `Auth`

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



##### service `Cluster`

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| MemberAdd | MemberAddRequest | MemberAddResponse | MemberAdd adds a member into the cluster. |
| MemberRemove | MemberRemoveRequest | MemberRemoveResponse | MemberRemove removes an existing member from the cluster. |
| MemberUpdate | MemberUpdateRequest | MemberUpdateResponse | MemberUpdate updates the member configuration. |
| MemberList | MemberListRequest | MemberListResponse | MemberList lists all the members in the cluster. |



##### service `KV`

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Range | RangeRequest | RangeResponse | Range gets the keys in the range from the store. |
| Put | PutRequest | PutResponse | Put puts the given key into the store. A put request increases the revision of the store, and generates one event in the event history. |
| DeleteRange | DeleteRangeRequest | DeleteRangeResponse | Delete deletes the given range from the store. A delete request increase the revision of the store, and generates one event in the event history. |
| Txn | TxnRequest | TxnResponse | Txn processes all the requests in one transaction. A txn request increases the revision of the store, and generates events with the same revision in the event history. It is not allowed to modify the same key several times within one txn. |
| Compact | CompactionRequest | CompactionResponse | Compact compacts the event history in etcd. User should compact the event history periodically, or it will grow infinitely. |



##### service `Lease`

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| LeaseGrant | LeaseGrantRequest | LeaseGrantResponse | LeaseGrant creates a lease. A lease has a TTL. The lease will expire if the server does not receive a keepAlive within TTL from the lease holder. All keys attached to the lease will be expired and deleted if the lease expires. The key expiration generates an event in event history. |
| LeaseRevoke | LeaseRevokeRequest | LeaseRevokeResponse | LeaseRevoke revokes a lease. All the key attached to the lease will be expired and deleted. |
| LeaseKeepAlive | LeaseKeepAliveRequest | LeaseKeepAliveResponse | KeepAlive keeps the lease alive. |



##### service `Maintenance`

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Alarm | AlarmRequest | AlarmResponse | Alarm activates, deactivates, and queries alarms regarding cluster health. |
| Status | StatusRequest | StatusResponse | Status gets the status of the member. |
| Defragment | DefragmentRequest | DefragmentResponse |  |
| Hash | HashRequest | HashResponse | Hash returns the hash of the local KV state for consistency checking purpose. This is designed for testing; do not use this in production when there are ongoing transactions. |
| Snapshot | SnapshotRequest | SnapshotResponse | Snapshot sends a snapshot of the entire backend |



##### service `Watch`

| Method | Request Type | Response Type | Description |
| ------ | ------------ | ------------- | ----------- |
| Watch | WatchRequest | WatchResponse | Watch watches the events happening or happened. Both input and output are stream. One watch rpc can watch for multiple keys or prefixs and get a stream of events. The whole events history can be watched unless compacted. |



##### message `AlarmMember`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| memberID |  | uint64 | uint64 | uint64 | long | int/long |
| alarm |  | AlarmType | | | | |



##### message `AlarmRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| action |  | AlarmAction | | | | |
| memberID | MemberID is the member raising the alarm request | uint64 | uint64 | uint64 | long | int/long |
| alarm |  | AlarmType | | | | |



##### message `AlarmResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| alarms |  | (slice of) AlarmMember | | | | |



##### message `AuthDisableRequest`

Empty field.



##### message `AuthDisableResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthEnableRequest`

Empty field.



##### message `AuthEnableResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleAddRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| name |  | string | string | string | String | str/unicode |



##### message `AuthRoleAddResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleDeleteRequest`

Empty field.



##### message `AuthRoleDeleteResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleGetRequest`

Empty field.



##### message `AuthRoleGetResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleGrantRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| name |  | string | string | string | String | str/unicode |
| perm |  | authpb.Permission | | | | |



##### message `AuthRoleGrantResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthRoleRevokeRequest`

Empty field.



##### message `AuthRoleRevokeResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserAddRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| name |  | string | string | string | String | str/unicode |
| password |  | string | string | string | String | str/unicode |



##### message `AuthUserAddResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserChangePasswordRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| name |  | string | string | string | String | str/unicode |
| password |  | string | string | string | String | str/unicode |



##### message `AuthUserChangePasswordResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserDeleteRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| name |  | string | string | string | String | str/unicode |



##### message `AuthUserDeleteResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserGetRequest`

Empty field.



##### message `AuthUserGetResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserGrantRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| user |  | string | string | string | String | str/unicode |
| role |  | string | string | string | String | str/unicode |



##### message `AuthUserGrantResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthUserRevokeRequest`

Empty field.



##### message `AuthUserRevokeResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `AuthenticateRequest`

Empty field.



##### message `AuthenticateResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `CompactionRequest`

Compaction compacts the kv store upto a given revision. All superseded keys with a revision less than the compaction revision will be removed.

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| revision |  | int64 | int64 | int64 | long | int/long |
| physical | physical is set so the RPC will wait until the compaction is physically applied to the local database such that compacted entries are totally removed from the backing store. | bool | bool | bool | boolean | boolean |



##### message `CompactionResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `Compare`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| result |  | CompareResult | | | | |
| target |  | CompareTarget | | | | |
| key | key path | bytes | []byte | string | ByteString | str |
| target_union |  | oneof | | | | |
| version | version of the given key | int64 | int64 | int64 | long | int/long |
| create_revision | create revision of the given key | int64 | int64 | int64 | long | int/long |
| mod_revision | last modified revision of the given key | int64 | int64 | int64 | long | int/long |
| value | value of the given key | bytes | []byte | string | ByteString | str |



##### message `DefragmentRequest`

Empty field.



##### message `DefragmentResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `DeleteRangeRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| key | if the range_end is not given, the request deletes the key. | bytes | []byte | string | ByteString | str |
| range_end | if the range_end is given, it deletes the keys in range [key, range_end). | bytes | []byte | string | ByteString | str |



##### message `DeleteRangeResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| deleted | Deleted is the number of keys that got deleted. | int64 | int64 | int64 | long | int/long |



##### message `EmptyResponse`

Empty field.



##### message `HashRequest`

Empty field.



##### message `HashResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| hash |  | uint32 | uint32 | uint32 | int | int/long |



##### message `InternalRaftRequest`

An InternalRaftRequest is the union of all requests which can be sent via raft.

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | uint64 | uint64 | uint64 | long | int/long |
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



##### message `LeaseGrantRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| TTL | advisory ttl in seconds | int64 | int64 | int64 | long | int/long |
| ID | requested ID to create; 0 lets lessor choose | int64 | int64 | int64 | long | int/long |



##### message `LeaseGrantResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| ID |  | int64 | int64 | int64 | long | int/long |
| TTL | server decided ttl in second | int64 | int64 | int64 | long | int/long |
| error |  | string | string | string | String | str/unicode |



##### message `LeaseKeepAliveRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | int64 | int64 | int64 | long | int/long |



##### message `LeaseKeepAliveResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| ID |  | int64 | int64 | int64 | long | int/long |
| TTL |  | int64 | int64 | int64 | long | int/long |



##### message `LeaseRevokeRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | int64 | int64 | int64 | long | int/long |



##### message `LeaseRevokeResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `Member`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | uint64 | uint64 | uint64 | long | int/long |
| name | If the member is not started, name will be an empty string. | string | string | string | String | str/unicode |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode |
| clientURLs | If the member is not started, client_URLs will be an zero length string array. | (slice of) string | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode |



##### message `MemberAddRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode |



##### message `MemberAddResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| member |  | Member | | | | |



##### message `MemberListRequest`

Empty field.



##### message `MemberListResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| members |  | (slice of) Member | | | | |



##### message `MemberRemoveRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | uint64 | uint64 | uint64 | long | int/long |



##### message `MemberRemoveResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `MemberUpdateRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | uint64 | uint64 | uint64 | long | int/long |
| peerURLs |  | (slice of) string | (slice of) string | (slice of) string | (slice of) String | (slice of) str/unicode |



##### message `MemberUpdateResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `Metadata`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| NodeID |  | uint64 | uint64 | uint64 | long | int/long |
| ClusterID |  | uint64 | uint64 | uint64 | long | int/long |



##### message `PutRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| key |  | bytes | []byte | string | ByteString | str |
| value |  | bytes | []byte | string | ByteString | str |
| lease |  | int64 | int64 | int64 | long | int/long |



##### message `PutResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |



##### message `RangeRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| key | if the range_end is not given, the request returns the key. | bytes | []byte | string | ByteString | str |
| range_end | if the range_end is given, it gets the keys in range [key, range_end) if range_end is nonempty, otherwise it returns all keys >= key. | bytes | []byte | string | ByteString | str |
| limit | limit the number of keys returned. | int64 | int64 | int64 | long | int/long |
| revision | range over the store at the given revision. if revision is less or equal to zero, range over the newest store. if the revision has been compacted, ErrCompaction will be returned in response. | int64 | int64 | int64 | long | int/long |
| sort_order | sort_order is the requested order for returned the results | SortOrder | | | | |
| sort_target | sort_target is the kv field to use for sorting | SortTarget | | | | |
| serializable | range request is linearizable by default. Linearizable requests has a higher latency and lower throughput than serializable request. To reduce latency, serializable can be set. If serializable is set, range request will be serializable, but not linearizable with other requests. Serializable range can be served locally without waiting for other nodes in the cluster. | bool | bool | bool | boolean | boolean |



##### message `RangeResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| kvs |  | (slice of) storagepb.KeyValue | | | | |
| more | more indicates if there are more keys to return in the requested range. | bool | bool | bool | boolean | boolean |



##### message `Request`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| ID |  | uint64 | uint64 | uint64 | long | int/long |
| Method |  | string | string | string | String | str/unicode |
| Path |  | string | string | string | String | str/unicode |
| Val |  | string | string | string | String | str/unicode |
| Dir |  | bool | bool | bool | boolean | boolean |
| PrevValue |  | string | string | string | String | str/unicode |
| PrevIndex |  | uint64 | uint64 | uint64 | long | int/long |
| PrevExist |  | bool | bool | bool | boolean | boolean |
| Expiration |  | int64 | int64 | int64 | long | int/long |
| Wait |  | bool | bool | bool | boolean | boolean |
| Since |  | uint64 | uint64 | uint64 | long | int/long |
| Recursive |  | bool | bool | bool | boolean | boolean |
| Sorted |  | bool | bool | bool | boolean | boolean |
| Quorum |  | bool | bool | bool | boolean | boolean |
| Time |  | int64 | int64 | int64 | long | int/long |
| Stream |  | bool | bool | bool | boolean | boolean |
| Refresh |  | bool | bool | bool | boolean | boolean |



##### message `RequestUnion`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| request |  | oneof | | | | |
| request_range |  | RangeRequest | | | | |
| request_put |  | PutRequest | | | | |
| request_delete_range |  | DeleteRangeRequest | | | | |



##### message `ResponseHeader`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| cluster_id |  | uint64 | uint64 | uint64 | long | int/long |
| member_id |  | uint64 | uint64 | uint64 | long | int/long |
| revision | revision of the store when the request was applied. | int64 | int64 | int64 | long | int/long |
| raft_term | term of raft when the request was applied. | uint64 | uint64 | uint64 | long | int/long |



##### message `ResponseUnion`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| response |  | oneof | | | | |
| response_range |  | RangeResponse | | | | |
| response_put |  | PutResponse | | | | |
| response_delete_range |  | DeleteRangeResponse | | | | |



##### message `SnapshotRequest`

Empty field.



##### message `SnapshotResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header | header has the current store information. The first header in the snapshot stream indicates the point in time of the snapshot. | ResponseHeader | | | | |
| remaining_bytes | remaining_bytes is the number of blob bytes to be sent after this message | uint64 | uint64 | uint64 | long | int/long |
| blob | blob has the next chunk of the snapshot in the snapshot stream. | bytes | []byte | string | ByteString | str |



##### message `StatusRequest`

Empty field.



##### message `StatusResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| version |  | string | string | string | String | str/unicode |
| dbSize |  | int64 | int64 | int64 | long | int/long |
| leader |  | uint64 | uint64 | uint64 | long | int/long |
| raftIndex |  | uint64 | uint64 | uint64 | long | int/long |
| raftTerm |  | uint64 | uint64 | uint64 | long | int/long |



##### message `TxnRequest`

If the comparisons succeed, then the success requests will be processed in order, and the response will contain their respective responses in order. If the comparisons fail, then the failure requests will be processed in order, and the response will contain their respective responses in order. From google paxosdb paper: Our implementation hinges around a powerful primitive which we call MultiOp. All other database operations except for iteration are implemented as a single call to MultiOp. A MultiOp is applied atomically and consists of three components: 1. A list of tests called guard. Each test in guard checks a single entry in the database. It may check for the absence or presence of a value, or compare with a given value. Two different tests in the guard may apply to the same or different entries in the database. All tests in the guard are applied and MultiOp returns the results. If all tests are true, MultiOp executes t op (see item 2 below), otherwise it executes f op (see item 3 below). 2. A list of database operations called t op. Each operation in the list is either an insert, delete, or lookup operation, and applies to a single database entry. Two different operations in the list may apply to the same or different entries in the database. These operations are executed if guard evaluates to true. 3. A list of database operations called f op. Like t op, but executed if guard evaluates to false.

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| compare |  | (slice of) Compare | | | | |
| success |  | (slice of) RequestUnion | | | | |
| failure |  | (slice of) RequestUnion | | | | |



##### message `TxnResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| succeeded |  | bool | bool | bool | boolean | boolean |
| responses |  | (slice of) ResponseUnion | | | | |



##### message `WatchCancelRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| watch_id |  | int64 | int64 | int64 | long | int/long |



##### message `WatchCreateRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| key | the key to be watched | bytes | []byte | string | ByteString | str |
| range_end | if the range_end is given, keys in [key, range_end) are watched NOTE: only range_end == prefixEnd(key) is accepted now | bytes | []byte | string | ByteString | str |
| start_revision | start_revision is an optional revision (including) to watch from. No start_revision is "now". | int64 | int64 | int64 | long | int/long |
| progress_notify | if progress_notify is set, etcd server sends WatchResponse with empty events to the created watcher when there are no recent events. It is useful when clients want always to be able to recover a disconnected watcher from a recent known revision. etcdsever can decide how long it should send a notification based on current load. | bool | bool | bool | boolean | boolean |



##### message `WatchRequest`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| request_union |  | oneof | | | | |
| create_request |  | WatchCreateRequest | | | | |
| cancel_request |  | WatchCancelRequest | | | | |



##### message `WatchResponse`

| Field | Description | Type | Go | C++ | Java | Python |
| ----- | ----------- | ---- | --- | --- | ---- | ------ |
| header |  | ResponseHeader | | | | |
| watch_id | watch_id is the ID of the watching the response sent to. | int64 | int64 | int64 | long | int/long |
| created | If the response is for a create watch request, created is set to true. Client should record the watch_id and prepare for receiving events for that watching from the same stream. All events sent to the created watching will attach with the same watch_id. | bool | bool | bool | boolean | boolean |
| canceled | If the response is for a cancel watch request, cancel is set to true. No further events will be sent to the canceled watching. | bool | bool | bool | boolean | boolean |
| compact_revision | CompactRevision is set to the minimum index if a watching tries to watch at a compacted index.  This happens when creating a watching at a compacted revision or the watching cannot catch up with the progress of the KV.  Client should treat the watching as canceled and should not try to create any watching with same start_revision again. | int64 | int64 | int64 | long | int/long |
| events |  | (slice of) storagepb.Event | | | | |



