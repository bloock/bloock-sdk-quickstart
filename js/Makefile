update_version:
	node ../bin/update_version.js package.json '"@bloock/sdk": "(.*)"' $(shell cat ../VERSION)

run:
	find src/samples/**/* -type f -name '*.ts' | rush 'echo {}; npx ts-node {}' 
