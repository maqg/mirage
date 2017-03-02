all::
	python ./createExample.py > ./mysql.default.example

init::
	python ./createExample.py > ./mysql.default.example
	./init_db.sh
	./dump.sh

clean::
	rm -rf *.xml
	rm -rf *.example
	rm -rf *.auto
