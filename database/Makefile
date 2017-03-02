all::
	python3 ./createExample.py > ./mysql.default.example

init::
	python3 ./createExample.py > ./mysql.default.example
	./init_db.sh
	./dump.sh

clean::
	rm -rf *.xml
	rm -rf *.example
	rm -rf *.auto
