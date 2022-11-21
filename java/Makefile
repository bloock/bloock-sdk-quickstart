SHELL=/bin/bash

update_version:


run:
	rm -rf build
	find ./src/samples/**/* -type f -name '*.java' | rush 'javac -d build -cp ./target/bloock-sdk-2.0.0-beta.23.jar:. -sourcepath src {} src/utils/*.java'
	find ./src/samples/**/* -type f -name '*.java' | rush ' \
		echo {}; \
		java -cp ./target/bloock-sdk-2.0.0-beta.23.jar:build {%.} \
	'
