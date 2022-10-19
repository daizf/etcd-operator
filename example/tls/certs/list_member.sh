ETCDCTL_API=3 etcdctl --endpoints=https://example-client.default.svc:2379 \
    --cert=etcd-client.crt --key=etcd-client.key --cacert=etcd-client-ca.crt \
    member list -w table
