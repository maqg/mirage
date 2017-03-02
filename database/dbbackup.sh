#!/bin/sh
# This is a ShellScript for auto DB backup when create octframe
# Jason
# 2015-10-27
 
. ./global.sh

FLAG=V5
BackupPath=/oct/Mysql_Backup
LogFile=$BackupPath/backup.log
NOW=$(date +%y%m%d%H%M%S)
NewFile=$BackupPath/$FLAG"_"$NOW.tgz
DumpFile=$BackupPath/$FLAG"_"$NOW.sql

mkdir -p $BackupPath

cd $BackupPath
ls -t | awk 'NR>99 {system("rm \"" $0 "\"")}'

echo "*********************************" >> $LogFile
echo "begin to backup $FLAG database $NOW" >> $LogFile

mysqldump -u$DB_USER -p$DB_PASSWD $DB_NAME -n -q -R > $DumpFile
if [ "$?" = 0 ]; then
	cd $BackupPath
	tar -czvf $NewFile $(basename $DumpFile) >> $LogFile 2>&1
	echo "[$NewFile]Backup Success!" >> $LogFile
	rm -rf $DumpFile
	echo "Backup octframe database Success $NOW $NewFile" >> $LogFile
	echo "*********************************" >> $LogFile
	exit 0
else
	echo "backup database failed $NOW"
	echo "*********************************" >> $LogFile
	exit 1
fi
