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
    echo "copy ${project}/app/admin/service/configs"
    cp "${project}/app/admin/service/configs/config.yaml" "$(pwd)/admin/service/configs/"
    cp "${project}/app/admin/service/configs/registry.yaml" "$(pwd)/admin/service/configs/"
fi

if [ "${server_name}" = "circle" ]
then
echo "copy ${project}/app/circle/service/configs"
cp "${project}/app/circle/service/configs/config.yaml" "$(pwd)/circle/service/configs/"
cp "${project}/app/circle/service/configs/registry.yaml" "$(pwd)/circle/service/configs/"
fi

if [ "${server_name}" = "circle_job" ]
then
echo "copy ${project}/app/circle/job/configs"
cp "${project}/app/circle/job/configs/config.yaml" "$(pwd)/circle/job/configs/"
cp "${project}/app/circle/job/configs/registry.yaml" "$(pwd)/circle/job/configs/"
fi

if [ "${server_name}" = "feed" ]
then
echo "copy ${project}/app/feed/service/configs"
cp "${project}/app/feed/service/configs/config.yaml" "$(pwd)/feed/service/configs/"
cp "${project}/app/feed/service/configs/registry.yaml" "$(pwd)/feed/service/configs/"
fi

if [ "${server_name}" = "feed_job" ]
then
echo "copy ${project}/app/feed/job/configs"
cp "${project}/app/feed/job/configs/config.yaml" "$(pwd)/feed/job/configs/"
cp "${project}/app/feed/job/configs/registry.yaml" "$(pwd)/feed/job/configs/"
fi

if [ "${server_name}" = "feed_interface" ]
then
echo "copy ${project}/app/feed/interface/configs"
cp "${project}/app/feed/interface/configs/config.yaml" "$(pwd)/feed/interface/configs/"
cp "${project}/app/feed/interface/configs/registry.yaml" "$(pwd)/feed/interface/configs/"
fi

if [ "${server_name}" = "mqtt" ]
then
echo "copy ${project}/app/mqtt/service/configs"
cp "${project}/app/mqtt/service/configs/config.yaml" "$(pwd)/mqtt/service/configs/"
cp "${project}/app/mqtt/service/configs/registry.yaml" "$(pwd)/mqtt/service/configs/"
fi

if [ "${server_name}" = "user" ]
then
echo "copy ${project}/app/user/service/configs"
cp "${project}/app/user/service/configs/config.yaml" "$(pwd)/user/service/configs/"
cp "${project}/app/user/service/configs/registry.yaml" "$(pwd)/user/service/configs/"
fi

if [ "${server_name}" = "shop" ]
then
echo "copy ${project}/app/shop/service/configs"
cp "${project}/app/shop/service/configs/config.yaml" "$(pwd)/shop/service/configs/"
cp "${project}/app/shop/service/configs/registry.yaml" "$(pwd)/shop/service/configs/"
fi