#!/bin/bash
version=v.1.21.14
if [ $# -gt 0 ]; then
  version=$1
fi

./hack/build/operator/build
./hack/build/backup-operator/build
./hack/build/restore-operator/build

docker build -t daizf/etcd-operator:$version -f ./hack/build/Dockerfile  .
docker push daizf/etcd-operator:$version

kubectl set image deploy etcd-operator etcd-operator=daizf/etcd-operator:$version
