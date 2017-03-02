#!/usr/bin/python
# -*- coding: utf-8 -*-

import json	
import time

import sys 
now = int(time.time() * 1000)

session_itme1 = {
	"ID": "00000000000000000000000000000000",
	"S_User": "admin",
	"S_Cookie": {
		"role": 7,
		"id": "097ff6191a7b4623a8d912c2adb62784",
		"name": "admin"
	},
	"S_CreateTime": now,
	"S_ExpireTime": now + 315360000000
}

DEFAULT_ITEMS = [
	#("tb_session", session_itme1)
]


def createSql(table, obj):
	front = "INSERT INTO %s (" % table
	end = " VALUES ("
	for (k, v) in list(obj.items()):
		front += "%s," % k
		if (type(v) == type(1)):
			end += "%ld," % v
		elif (type(v) == type("str")):
			end += "'%s'," % v
		else:
			end += "'%s'," % json.dumps(v, ensure_ascii=False)

	return front[:-1] + ")" + end[:-1] + ");\n"

if __name__ == '__main__':

	sql = ""

	for (table, item) in DEFAULT_ITEMS:
		sql += createSql(table, item)

	print(sql)
