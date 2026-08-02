#!/bin/sh
set -e

case "$1" in
    remove|deconfigure)
        systemctl stop starlink-exporter.service || true
        systemctl disable starlink-exporter.service || true
        ;;
esac

exit 0
