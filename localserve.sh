#!/usr/bin/env bash

PHOTOPRISM_DEBUG=true
PHOTOPRISM_SERVER_MODE=true
PHOTOPRISM_ASSETS_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets
PHOTOPRISM_CACHE_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/cache
PHOTOPRISM_RESOURCES_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/resources
PHOTOPRISM_CONFIG_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/config
PHOTOPRISM_IMPORT_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/photos/import
PHOTOPRISM_EXPORT_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/photos/export
PHOTOPRISM_ORIGINALS_PATH=~/Documents/Arbeit/Projects/PhpStorm/PhotoPrism/assets/photos/originals
PHOTOPRISM_DATABASE_DRIVER=mysql
PHOTOPRISM_DATABASE_DSN="photoprism:photoprism@tcp(localhost:4001)/photoprism"
PHOTOPRISM_HTTP_HOST=0.0.0.0
PHOTOPRISM_HTTP_PORT=2342
PHOTOPRISM_SQL_HOST=0.0.0.0
PHOTOPRISM_SQL_PORT=4001
PHOTOPRISM_SQL_PASSWORD=photoprism
TF_CPP_MIN_LOG_LEVEL=0

go run cmd/photoprism/photoprism.go migrate
