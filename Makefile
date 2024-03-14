all: build demo

REPO_DIR := traefik
TRAEFIK_VERSION := v2.11.0

.PHONY: download_repo check_version

download_repo: $(REPO_DIR)
	@echo "Repository already exists."

check_version: $(REPO_DIR)
	@cd $(REPO_DIR) && git describe --tags | grep -q '^$(TRAEFIK_VERSION)$$' || (echo "Wrong version" && cd .. && $(MAKE) clean_repo download_repo)
	@echo "Version is correct."

clean_repo:
	@echo "Cleaning repository..."
	@rm -rf $(REPO_DIR)

$(REPO_DIR):
	git clone --depth=1 --branch=$(TRAEFIK_VERSION) https://github.com/traefik/traefik.git $(REPO_DIR)

build-ui:
	@cd $(REPO_DIR) && $(MAKE) clean-webui generate-webui && tar czvf webui.tar.gz ./webui/static/

build: download_repo check_version
	@cd $(REPO_DIR) && git reset --hard
	cp -rf ./pkg $(REPO_DIR)
	@cd $(REPO_DIR) && $(MAKE) binary-linux-amd64

demo:
	docker compose up
