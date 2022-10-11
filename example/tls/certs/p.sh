etcdctl --endpoints=https://example-client.default.svc:2379 --cert-file=etcd-client.crt --key-file=etcd-client.key --ca-file=etcd-client-ca.crt set say1 hello1
etcdctl --endpoints=https://example-client.default.svc:2379 --cert-file=etcd-client.crt --key-file=etcd-client.key --ca-file=etcd-client-ca.crt set say hello

