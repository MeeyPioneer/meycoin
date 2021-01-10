#!/usr/bin/env bash
source test_common.sh

if [ "$1" = "" ] ; then
	echo "use:rm_member.sh meycoin1~meycoin3"
	exit 100
fi


rmnode=$1


# get leader
myleader=
getleader myleader
echo "myleader=$myleader"


getLeaderPort leaderport
prevCnt=$(getClusterTotal $leaderport)


raftID=""
getRaftID $leaderport $rmnode raftID

# get leader port

echo "leader=$myleader, port=$leaderport, raftId=$raftID"

#echo "meycoincli -p $leaderport cluster remove --nodeid $raftID"
#meycoincli -p $leaderport cluster remove --nodeid $raftID

walletFile="$TEST_RAFT_INSTANCE/genesis_wallet.txt"
ADMIN=
getAdminUnlocked $leaderport $walletFile ADMIN

rmJson="$(makeRemoveMemberJson $raftID)"

echo "meycoincli -p "$leaderport" contract call --governance "$ADMIN" meycoin.enterprise changeCluster "$rmJson" --keystore . --password 1234"
meycoincli -p $leaderport contract call --governance $ADMIN meycoin.enterprise changeCluster "$rmJson" --keystore . --password 1234
echo "remove Done" 

# check if total count is decremented
reqCnt=$((prevCnt-1))
echo "reqClusterTotal=$reqCnt"
if [ "$myleader" = "$rmnode" ];then
	leaderport=10001
fi
waitClusterTotal $reqCnt 100
if [ $? -ne 1 ]; then
	echo "remove failed"
	exit 100
fi
