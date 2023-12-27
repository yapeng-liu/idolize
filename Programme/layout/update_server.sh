#!/bin/sh

total=$#
if [ ${total} -lt 1 ]
then
    echo "please use ./update_server.sh user"
    exit 0
fi

project="/home/robo/2golangProject/daoke_server"
server_name=$1
pwd=$(pwd)

if [ "${server_name}" = "admin" ]
then
    echo "remove ${project}/app/admin/service/bin/server"
    rm "$(pwd)/admin/service/bin/server"
    echo "copy ${project}/app/admin/service/bin/server"
    cp "${project}/app/admin/service/bin/server" "$(pwd)/admin/service/bin/"
fi

if [ "${server_name}" = "circle" ]
then
    echo "remove ${project}/app/circle/service/bin/server"
    rm "$(pwd)/circle/service/bin/server"
    echo "copy ${project}/app/circle/service/bin/server"
    cp "${project}/app/circle/service/bin/server" "$(pwd)/circle/service/bin/"
fi

if [ "${server_name}" = "circle_job" ]
then
    echo "remove ${project}/app/circle/job/bin/server"
    rm "$(pwd)/circle/job/bin/server"
    echo "copy ${project}/app/circle/job/bin/server"
    cp "${project}/app/circle/job/bin/server" "$(pwd)/circle/job/bin/"
fi

if [ "${server_name}" = "feed" ]
then
    echo "remove ${project}/app/feed/service/bin/server"
    rm "$(pwd)/feed/service/bin/server"
    echo "copy ${project}/app/feed/service/bin/server"
    cp "${project}/app/feed/service/bin/server" "$(pwd)/feed/service/bin/"
fi

if [ "${server_name}" = "feed_job" ]
then
    echo "remove ${project}/app/feed/job/bin/server"
    rm "$(pwd)/feed/job/bin/server"
    echo "copy ${project}/app/feed/job/bin/server"
    cp "${project}/app/feed/job/bin/server" "$(pwd)/feed/job/bin/"
fi

if [ "${server_name}" = "feed_interface" ]
then
    echo "remove ${project}/app/feed/interface/bin/server"
    rm "$(pwd)/feed/interface/bin/server"
    echo "copy ${project}/app/feed/interface/bin/server"
    cp "${project}/app/feed/interface/bin/server" "$(pwd)/feed/interface/bin/"
fi

if [ "${server_name}" = "mqtt" ]
then
    echo "remove ${project}/app/mqtt/service/bin/server"
    rm "$(pwd)/mqtt/service/bin/server"
    echo "copy ${project}/app/mqtt/service/bin/server"
    cp "${project}/app/mqtt/service/bin/server" "$(pwd)/mqtt/service/bin/"
fi

if [ "${server_name}" = "user" ]
then
    echo "remove ${project}/app/user/service/bin/server"
    rm "$(pwd)/user/service/bin/server"
    echo "copy ${project}/app/user/service/bin/server"
    cp "${project}/app/user/service/bin/server" "$(pwd)/user/service/bin/"
fi

if [ "${server_name}" = "shop" ]
then
    echo "remove ${project}/app/shop/service/bin/server"
    rm "$(pwd)/shop/service/bin/server"
    echo "copy ${project}/app/shop/service/bin/server"
    cp "${project}/app/shop/service/bin/server" "$(pwd)/shop/service/bin/"
fi
