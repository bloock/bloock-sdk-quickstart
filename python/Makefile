update_version:


run:
	python3 setup.py install
	find samples/**/* -type f -name '*.py' | rush 'echo {}; python3 {}' 
