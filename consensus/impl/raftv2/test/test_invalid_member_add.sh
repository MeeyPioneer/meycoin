#!/usr/bin/env bash
echo "================= raft invalid member add test ===================="

BP_NAME=""

#rm BP*.toml
#./meycoinconf-gen.sh 10001 tmpl.toml 5
#clean.sh
#./inittest.sh
source test_common.sh

echo "kill_svr"
kill_svr.sh
rm -rf $TEST_RAFT_INSTANCE_DATA
rm $TEST_RAFT_INSTANCE/BP*

echo ""
echo "========= join invalid config member meycoin4 ========="
pushd $TEST_RAFT_INSTANCE/config
do_sed.sh BP11004.toml meycoin4 meycoinxxx =
popd

TEST_SKIP_GENESIS=0 make_node.sh 
RUN_TEST_SCRIPT set_system_admin.sh

add_member.sh meycoin4
if [ $? -ne 0 ];then
	echo "Adding of invalid config member must succeed"
	exit 100
fi

sleep 40
existProcess 10004
if [ "$?" = "1" ]; then
	echo "error! process must be killed."
	exit 100
fi

pushd $TEST_RAFT_INSTANCE/config
do_sed.sh BP11004.toml meycoinxxx meycoin4 =
popd

echo ""
echo "========= rm meycoin4 ========="
rm_member.sh meycoin4
rm BP11004*
checkSync 10001 10003 20 

echo ""
echo "========= check if reorg occured ======="
checkReorg
