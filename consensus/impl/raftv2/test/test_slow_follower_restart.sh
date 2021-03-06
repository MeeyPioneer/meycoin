#!/usr/bin/env bash
source test_common.sh

echo ""
echo "============================== raft server slow follower restart test ============================"

BP_NAME=""

#rm BP*.toml
#./meycoinconf-gen.sh 10001 tmpl.toml 5
#clean.sh
#./inittest.sh

echo ""
echo "======== make initial server ========="
make_node.sh 

checkSync 10001 10002 60
checkSync 10001 10003 60

sleep 10

kill_svr.sh 11003

echo "run meycoin3(11003). this node is slower than other nodeds."
DEBUG_CHAIN_OTHER_SLEEP=15000 run_svr.sh 11003
#DEBUG_CHAIN_OTHER_SLEEP=2000 run_svr.sh 11003

checkSyncRunning 10001 10003 100

kill_svr.sh 11003
run_svr.sh 11003

checkSync	10001 10003 300
echo "------------ success--------------"


