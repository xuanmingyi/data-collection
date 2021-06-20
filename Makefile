.PHONY: api
# genreate api
api:
	find app -mindepth 2 -maxdepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) app'

.PHONY: wire
# generate wire
wire:
	find app -mindepth 2 -maxdepth 2 -type d -print

.PHONY: proto
# generate proto
proto:
	find app -mindepth 2 -maxdepth 2 -type d -print