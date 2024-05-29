# iterate through services dir and run `make freeze` if it exists
freeze_refresh:
	@for dir in $(wildcard services/*); do \
		if [ -d "$$dir" ]; then \
			if [ -f "$$dir/makefile" ]; then \
				$(MAKE) -C $$dir freeze; \
			fi; \
		fi; \
	done