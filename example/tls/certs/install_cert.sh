#!/bin/bash
kubectl delete secret etcd-server-tls
kubectl delete secret etcd-peer-tls
kubectl delete secret etcd-client-tls
kubectl create secret generic etcd-server-tls --from-file=server-ca.crt --from-file=server.crt --from-file=server.key
kubectl create secret generic etcd-peer-tls --from-file=peer-ca.crt --from-file=peer.crt --from-file=peer.key
kubectl create secret generic etcd-client-tls --from-file=etcd-client-ca.crt --from-file=etcd-client.crt --from-file=etcd-client.key
