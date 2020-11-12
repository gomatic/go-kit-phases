# can override:
# - PROJECT_NAME
# - BRANCH_NAME
# - BUILD_NUMBER
# - CANDIDATE_TAG
# - GIT_BRANCH
# - GIT_COMMIT
# - GIT_COMMIT_LONG

export TZ := UTC

THIS := $(abspath $(firstword $(MAKEFILE_LIST)))
THISD := $(patsubst %/,%,$(dir $(THIS)))
THISR := $(patsubst $(ROOT)/%,%,$(THIS))
MAKER := $(MAKE) -f $(THISR)

PROJECT_NAME ?= $(notdir $(PWD))

GIT_HOST := github.com
GIT_ORG := gomatic
GIT_REPO := $(PROJECT_NAME)
GIT_BRANCH ?= $(shell git --no-pager rev-parse --abbrev-ref HEAD)
GIT_COMMIT ?= $(shell git --no-pager rev-parse --short=12 HEAD)
GIT_COMMIT_LONG ?= $(shell git --no-pager rev-parse HEAD)
GIT_DATETIME := $(shell git --no-pager log -1 --format='%ad' --abbrev=12 --date=format:'%Y%m%dT%H%M%SUTC' HEAD 2>/dev/null || date '+%Y%m%dT%H%M%S%Z')
BRANCH_NAME ?= $(GIT_BRANCH)
BUILD_NUMBER ?= 9999999
