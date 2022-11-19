SHELL=/bin/bash

update_version:


run:
	rm -rf build
	find ./src/samples/**/* -type f -name '*.java' | rush 'javac -d build -cp ./target/bloock-1.0.0-dependencies.jar:. -sourcepath src {} src/utils/*.java'
	find ./src/samples/**/* -type f -name '*.java' | rush ' \
		echo {}; \
		java -cp ./target/bloock-1.0.0-dependencies.jar:build {%.} \
	'
