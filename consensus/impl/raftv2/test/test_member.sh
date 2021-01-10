#!/usr/bin/env bash
echo "================= raft member add/remove test ===================="
source test_common.sh

BP_NAME=""

#rm BP*.toml
#./meycoinconf-gen.sh 10001 tmpl.toml 5
#clean.sh
#./inittest.sh
rm BP11004* BP11005*

echo "kill_svr & clean 11004~11007"
kill_svr.sh
for i in  11004 11005 11006 11007; do
	rm -rf ./data/$i
	rm -rf ./BP$i.toml
done

TEST_SKIP_GENESIS=0 make_node.sh
RUN_TEST_SCRIPT set_system_admin.sh

date
echo ""
echo "========= add meycoin4 ========="
add_member.sh meycoin4
if [ "$?" != "0" ];then
	echo "fail to add"
	exit 100
fi
checkSync 10001 10004 120 

date
echo ""
echo "========= add meycoin5 ========="
add_member.sh meycoin5
if [ "$?" != "0" ];then
	echo "fail to add"
	exit 100
fi
checkSync 10001 10005 120 


date
echo ""
echo "========= add meycoin6 ========="
add_member.sh meycoin6
if [ "$?" != "0" ];then
	echo "fail to add"
	exit 100
fi
checkSync 10001 10006 120 result


date
echo ""
echo "========= add meycoin7 ========="
add_member.sh meycoin7
if [ "$?" != "0" ];then
	echo "fail to add"
	exit 100
fi
checkSync 10001 10007 120 

date
echo ""
echo "========= rm meycoin7 ========="
rm_member.sh meycoin7
if [ "$?" != "0" ];then
	echo "fail to rm"
	exit 100
fi
rm BP11007*
checkSync 10001 10006 60 

echo ""
echo "========= rm meycoin6 ========="
rm_member.sh meycoin6
if [ "$?" != "0" ];then
	echo "fail to rm"
	exit 100
fi
rm BP11006*
checkSync 10001 10005 60 

echo ""
echo "========= rm meycoin5 ========="
rm_member.sh meycoin5
if [ "$?" != "0" ];then
	echo "fail to rm"
	exit 100
fi
rm BP11005*
checkSync 10001 10004 60 

echo ""
echo "========= rm meycoin4 ========="
rm_member.sh meycoin4
if [ "$?" != "0" ];then
	echo "fail to rm"
	exit 100
fi
rm BP11004*
checkSync 10001 10003 60 

echo ""
echo "========= check if reorg occured ======="
checkReorg
