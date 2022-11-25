install:
	pip3 install -r requirements.txt

update_version:
	node ../bin/update_version.js requirements.txt 'bloock==(.*)' $(shell cat ../VERSION)

run:
	python3 setup.py install
	find samples/**/* -type f -name '*.py' | rush 'echo {}; python3 {}' 
