update_version:
	$(MAKE) -C go update_version
	$(MAKE) -C java update_version
	$(MAKE) -C js update_version
	$(MAKE) -C python update_version
	