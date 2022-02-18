#!/bin/bash

trap 'cleanup $LINENO' ERR

# shellcheck disable=SC2120
cleanup() {
    MINIO_VERSION=dev docker-compose \
                 -f "buildscripts/tiering-tests/compose.yml" \
                 down -v
}

add_aliases() {
    # add hot tier alias
    for i in $(seq 1 4); do
        echo "... attempting to add hot tier alias $i"
        until (mc alias set hot http://127.0.0.1:9000 minioadmin minioadmin); do
            echo "...waiting... for 5secs" && sleep 5
        done
    done

    # add warm tier aliaas
    for i in $(seq 1 4); do
        echo "... attempting to add warm tier alias $i"
        until (mc alias set warm http://127.0.0.1:9011 minioadmin minioadmin); do
            echo "...waiting... for 5secs" && sleep 5
        done
    done

}

main() {
    sudo apt install curl -y
    export GOPATH=/tmp/gopath
    export PATH=${PATH}:${GOPATH}/bin

    go install github.com/minio/mc@latest

    TAG=minio/minio:dev make docker

    MINIO_VERSION=dev docker-compose \
                 -f "buildscripts/tiering-tests/compose.yml" \
                 up -d

    add_aliases

    # set aggressive scanner settings
    mc admin config set hot scanner delay=1 max_wait=100ms cycle=10s

    # create a bucket on hot tier
    mc mb hot/mybucket

    # create a bucket on warm tier to transition objects from hot tier
    mc mb warm/tierbucket

    # add a warm tier
    mc admin tier add s3 hot WARM --endpoint http://warmtier:9000 --access-key minioadmin --secret-key minioadmin --bucket tierbucket --prefix prefix

    # add ILM rules to transition objects from hot/mybucket to warm/tierbucket
    mc ilm add --transition-days 0 --storage-class WARM hot/mybucket

    mc cp /etc/hosts hot/mybucket/obj-1

    for i in $(seq 1 4); do
        echo "... checking if object has transitioned $i"
        remote_tier=$(mc stat hot/mybucket/obj-1 --json | jq '.metadata."X-Amz-Storage-Class"')
        if "$remote_tier" == "WARM"
        then
            break
        fi
        echo "... waiting for 2 secs" && sleep 2
    done

    # restore obj-1
    mc ilm restore hot/mybucket/obj-1

    obj_count=$(mc ls hot/mybucket/obj-1 | wc -l)
    if $obj_count != 1
    then
        echo "expected 1 but got ${obj_count}"
        exit 1;
    fi

    cleanup
}


main "$@"
