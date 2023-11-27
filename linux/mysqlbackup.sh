#!/usr/bin/env bash
#存放目录
BackupDir=~/mysqlbackup
#数据库库名
DataBaseName=thss
#日期命名
DateTag=`date +%Y%m%d`
#sql脚本名字
sqltag=$DataBaseName'_'$DateTag'.'sql
#压缩文件名字
tartag=$sqltag'.'tar'.'gz
#备份
mysqldump -uyjy -php51yYnbQqjBAWS1 --databases $DataBaseName > $BackupDir/$sqltag
#进行压缩并删除原文件
# shellcheck disable=SC2164
cd $BackupDir
tar -czf  $tartag $sqltag
rm -rf $sqltag
#定时清除文件，以访长期堆积占用磁盘空间(删除5天以前带有tar.gz文件)
find $BackupDir -mtime +5 -name '*.tar.gz' -exec rm -rf {} \;