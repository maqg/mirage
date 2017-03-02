#!/bin/sh

. ./global.sh

if [ -d /var/lib/mysql/$DB_NAME ]; then
	echo "$DB_NAME database exist, backup it ... ..."
	./dbbackup.sh
fi

create_db ()
{
	DBNAME=$1

$MYSQL -u$DB_USER -p$DB_PASSWD << EOF
	DROP DATABASE IF EXISTS $DBNAME;
	CREATE DATABASE $DBNAME DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
	USE $DBNAME;
EOF
	$MYSQL -u$DB_USER -p$DB_PASSWD $DBNAME < "mysql.sql"
	$MYSQL -u$DB_USER -p$DB_PASSWD $DBNAME < "view.sql"

	echo "initialing default tables...$DBNAME"
	$MYSQL -u$DB_USER -p$DB_PASSWD $DBNAME < "mysql.default"

	if [ -f "./mysql.default.auto" ]; then
		echo "initialing auto tables...$DBNAME"
		$MYSQL -u$DB_USER -p$DB_PASSWD $DBNAME < "mysql.default.auto"
	fi

	if [ -f "./mysql.default.example" ]; then
		echo "initialing example tables...$DBNAME"
		$MYSQL -u$DB_USER -p$DB_PASSWD $DBNAME < "./mysql.default.example"
	fi

	$MYSQL -u$DB_USER -p$DB_PASSWD -e "GRANT ALL ON *.* TO $DB_USER@'127.0.0.1' IDENTIFIED BY '$DB_PASSWD' WITH GRANT OPTION;"

	echo "DB $1 created ok"
}


create_db $DB_NAME

exit 0
