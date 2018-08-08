#!/bin/bash

CUR_DIR=$(cd `dirname $0`; pwd)
cd ${CUR_DIR}

mkdir -p "${CUR_DIR}/local_data/logs"

CONTROL_LOG="${CUR_DIR}/local_data/logs/control.log"
>${CONTROL_LOG}
exec 1>>${CONTROL_LOG} 2>&1

export PS4='+[`basename ${BASH_SOURCE[0]}`:$LINENO ${FUNCNAME[0]} \D{%F %T} $$]'
set -x

PROG_NAME=$0
ACTION=$1
BIN_NAME="gd_log_tailer"
LOG_DIR="${CUR_DIR}/local_data/logs"
DAEMON_LOG="daemon.log"

function usage
{
    echo "Usage: $PROG_NAME {start|stop|restart|deploy|status|kill}"
    exit 1;
}

function start
{
    ./${BIN_NAME} > ${LOG_DIR}/${DAEMON_LOG} 2>&1 &
    echo $?
}

function kill
{
    ps aux | grep gd_log_tailer | grep -v grep | awk '{print $2}' | xargs -i kill -9 {}
}

function restart
{
    kill
    sleep 1
    start
}

function status
{
    ps aux | grep "./gd_log_tailer" | grep -v grep > /dev/null
    return $?
}

function upgrade
{
    restart
}

if [ $# -lt 1 ]; then
    usage
fi

case "$ACTION" in
    start)
        kill
        start
    ;;
    stop)
        kill
    ;;
    restart)
        restart
    ;;
    kill)
        kill
    ;;
    status)
        status
    ;;
    upgrade)
        upgrade
    ;;
    *)
        usage
    ;;
esac
