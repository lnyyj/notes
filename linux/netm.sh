#!/bin/bash
##作用：监控eth0端口的流量
##可以将eth0换为eth1等其他端口
##作者：Barlow
##2012-12-10

while true 
do 
#取初始流量值
R1=`cat /sys/class/net/eth0/statistics/rx_bytes`
T1=`cat /sys/class/net/eth0/statistics/tx_bytes`
#
#暂停10秒后再次取值
sleep 1
R2=`cat /sys/class/net/eth0/statistics/rx_bytes`
T2=`cat /sys/class/net/eth0/statistics/tx_bytes`
#
#计算1秒内平均流量值，以kb/s为单位
TBPS=`expr $T2 - $T1`
RBPS=`expr $R2 - $R1`

TKBPS=`expr $TBPS / 1024`
RKBPS=`expr $RBPS / 1024`
echo "上传速率 eth0: $TKBPS kb/s 下载速率 eth0: $RKBPS kb/s at $(date +%Y%m%d%H:%M:%S)"
done
