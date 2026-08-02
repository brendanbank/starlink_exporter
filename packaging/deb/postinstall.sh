#!/bin/sh
set -e

case "$1" in
    configure)
        # Create starlink-exporter user and group if they don't exist
        if ! getent group starlink-exporter >/dev/null 2>&1; then
            addgroup --system starlink-exporter
        fi

        if ! getent passwd starlink-exporter >/dev/null 2>&1; then
            adduser --system --ingroup starlink-exporter \
                --no-create-home --disabled-password \
                --gecos "Starlink Exporter Service" \
                starlink-exporter
        fi

        # Ensure binary has correct permissions
        chown root:root /usr/bin/starlink_exporter
        chmod 755 /usr/bin/starlink_exporter

        # Create state directory
        mkdir -p /var/lib/starlink-exporter
        chown starlink-exporter:starlink-exporter /var/lib/starlink-exporter
        chmod 755 /var/lib/starlink-exporter

        # Reload systemd
        systemctl daemon-reload || true

        # Enable the service, then activate the version just unpacked.
        # $2 is the previously configured version: empty on a fresh install,
        # set on an upgrade. On upgrade the service is usually already
        # running, and "start" is a no-op that leaves the old binary running
        # until something else restarts it. try-restart picks up the new
        # binary, while still leaving a deliberately stopped service stopped.
        systemctl enable starlink-exporter.service || true
        if [ -n "$2" ]; then
            systemctl try-restart starlink-exporter.service || true
        else
            systemctl start starlink-exporter.service || true
        fi

        echo "Starlink Exporter has been installed and started."
        echo "Service runs as user: starlink-exporter"
        echo "Check status: systemctl status starlink-exporter"
        echo "View logs: journalctl -u starlink-exporter -f"
        ;;
esac

exit 0
