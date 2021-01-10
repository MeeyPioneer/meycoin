#!/usr/bin/env bash
echo "================= raft member join with backup test ===================="

BP_NAME=""

#rm BP*.toml
#./meycoinconf-gen.sh 10001 tmpl.toml 5
#clean.sh
#./inittest.sh
source test_common.sh

echo "clean all prev servers"
echo "kill_svr & clean 11004~11007"
kill_svr.sh
for i in  11004 11005 11006 11007; do
	echo "rm -rf $TEST_RAFT_INSTANCE/data/$i $TEST_RAFT_INSTANCE/BP$i.toml"
	rm -rf $TEST_RAFT_INSTANCE/data/$i $TEST_RAFT_INSTANCE/BP$i.toml
done

TEST_SKIP_GENESIS=0 make_node.sh
RUN_TEST_SCRIPT set_system_admin.sh

sleep 2

function backupJoin() {
	if ! [ $1 -lt 6 ] || ! [ $2 -lt 6 ]; then
		echo "Usage: $0 srcnodeNo(1<=no<=5) addnodeNo"
		echo "exam) $0 3 4"
		exit 100
	fi

	srcnodename=${nodenames[$1]}
	srcsvrport=${svrports[$srcnodename]}
	srcrpcport=${ports[$srcnodename]}

	addnodename=${nodenames[$2]}
	addsvrport=${svrports[$addnodename]}
	addrpcport=${ports[$addnodename]}

	echo "add $addsvrport with $srcsvrport data"
	
	echo ""
	echo "========= shutdown srcsvrport $srcsvrport   ========="
	kill_svr.sh $srcsvrport 

	echo ""
	echo "========= copy backup : cp -rf ./data/$srcsvrport ./data/$addsvrport ========="
	cp -rf ./data/$srcsvrport ./data/$addsvrport 

	run_svr.sh $srcsvrport

	checkSync 10001 $srcrpcport 180

	sleep 20

	echo ""
	echo "========= add $addnodename ========="
	add_member.sh $addnodename usebackup
	checkSyncWithLeader $addrpcport 60
}

backupJoin 3 4
backupJoin 3 5

echo "=========== rm member5 =========="
rm_member.sh meycoin5
rm BP11005*

echo "=========== rm member4 =========="
rm_member.sh meycoin4
rm BP11004*

echo ""
echo "========= check if reorg occured ========"
checkReorg
