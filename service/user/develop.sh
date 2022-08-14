#!/bin/bash

if [ $1 == "run" ] ; then

    docker stop app-mysql consul |   xargs echo "停止 "
    docker start app-mysql consul | xargs echo "运行 "
    kratos run

fi

