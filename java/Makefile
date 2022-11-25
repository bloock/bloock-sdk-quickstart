SHELL=/bin/bash

update_version:
	node ../bin/update_version.js pom.xml '<artifactId>bloock-sdk</artifactId>[\S\s]*?<version>(.*?)</version>' $(shell cat ../VERSION)

run:
	(cd ./src/main/java/com/bloock/quickstart/samples && find ./**/* -type f -name '*.java') | rush 'bash run.sh {} {/} {%.}'
	(cd ./src/main/java/com/bloock/quickstart/samples && find ./**/* -type f -name '*.java') | rush 'java -jar ./target/{%.}.jar'
