all::
	go build -o bin/mirage src/main.go


prepare::
	go run tools/funcmap/create_funcmap.go > src/api/funcmap.go	


clean::
	rm -rf bin/*
	rm -rf mirage


fmt::
	go fmt .
