SHELL=/bin/bash

update_version:


run:
	(cd ./src/main/java/com/bloock/quickstart/samples && find ./**/* -type f -name '*.java') | rush 'bash run.sh {} {/} {%.}'
	(cd ./src/main/java/com/bloock/quickstart/samples && find ./**/* -type f -name '*.java') | rush 'java -jar ./target/{%.}.jar'
