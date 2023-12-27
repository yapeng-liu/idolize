#!/bin/sh
pwd=$(pwd)
echo "create $(pwd)/admin/service/bin"
echo "create $(pwd)/admin/service/configs"
mkdir -p "$(pwd)/admin/service/bin" -p $(pwd)/admin/service/configs
echo "create $(pwd)/circle/job/bin"
echo "create $(pwd)/circle/job/configs"
echo "create $(pwd)/circle/service/bin"
echo "create $(pwd)/circle/service/configs"
mkdir -p $(pwd)/circle/job/bin -p $(pwd)/circle/job/configs -p $(pwd)/circle/service/bin -p $(pwd)/circle/service/configs
echo "create $(pwd)/feed/interface/bin"
echo "create $(pwd)/feed/interface/configs"
echo "create $(pwd)/feed/job/bin"
echo "create $(pwd)/feed/job/configs"
echo "create $(pwd)/feed/service/bin"
echo "create $(pwd)/feed/service/configs"
mkdir -p $(pwd)/feed/interface/bin -p $(pwd)/feed/interface/configs -p $(pwd)/feed/job/bin -p $(pwd)/feed/job/configs -p $(pwd)/feed/service/bin -p $(pwd)/feed/service/configs
echo "create $(pwd)/mqtt/service/bin"
echo "create $(pwd)/mqtt/service/configs"
mkdir -p $(pwd)/mqtt/service/bin -p $(pwd)/mqtt/service/configs
echo "create $(pwd)/user/service/bin"
echo "create $(pwd)/user/service/configs"
mkdir -p $(pwd)/user/service/bin -p $(pwd)/user/service/configs
echo "create $(pwd)/shop/service/bin"
echo "create $(pwd)/shop/service/configs"
mkdir -p $(pwd)/shop/service/bin -p $(pwd)/shop/service/configs
