#!/bin/bash
set -e

# Build Debian packages for multiple architectures
# This script builds .deb packages for AMD64, ARM64, and ARM (armhf)

VERSION="${VERSION:-1.0.0}"
PACKAGE_NAME="starlink-exporter"
BUILD_DIR="$(pwd)/build"
DEB_DIR="${BUILD_DIR}/deb"

echo "Building Debian packages for ${PACKAGE_NAME} version ${VERSION}"

# Clean previous builds
rm -rf "${BUILD_DIR}"
mkdir -p "${DEB_DIR}"

# Function to build package for specific architecture
build_deb() {
    local ARCH=$1
    local GOARCH=$2
    local GOARM=$3

    echo "Building .deb package for ${ARCH}..."

    # Build the binary
    echo "  - Building binary for ${ARCH}..."
    if [ "${ARCH}" = "armhf" ]; then
        CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} GOARM=${GOARM} go build \
            -ldflags="-s -w -extldflags '-static'" \
            -trimpath \
            -o starlink_exporter \
            ./cmd/starlink_exporter
    else
        CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build \
            -ldflags="-s -w -extldflags '-static'" \
            -trimpath \
            -o starlink_exporter \
            ./cmd/starlink_exporter
    fi

    # Create package directory structure
    local PKG_DIR="${BUILD_DIR}/${PACKAGE_NAME}_${VERSION}_${ARCH}"
    mkdir -p "${PKG_DIR}/DEBIAN"
    mkdir -p "${PKG_DIR}/usr/bin"
    mkdir -p "${PKG_DIR}/lib/systemd/system"
    mkdir -p "${PKG_DIR}/usr/share/doc/${PACKAGE_NAME}"

    # Copy files
    cp starlink_exporter "${PKG_DIR}/usr/bin/"
    chmod 755 "${PKG_DIR}/usr/bin/starlink_exporter"

    cp starlink-exporter.service "${PKG_DIR}/lib/systemd/system/"
    chmod 644 "${PKG_DIR}/lib/systemd/system/starlink-exporter.service"

    cp LICENSE "${PKG_DIR}/usr/share/doc/${PACKAGE_NAME}/"
    cp README.md "${PKG_DIR}/usr/share/doc/${PACKAGE_NAME}/"

    # Create control file
    cat > "${PKG_DIR}/DEBIAN/control" <<EOF
Package: ${PACKAGE_NAME}
Version: ${VERSION}
Section: net
Priority: optional
Architecture: ${ARCH}
Maintainer: Brendan Bank <brendan@brendanbank.nl>
Homepage: https://github.com/brendanbank/starlink_exporter
Description: Prometheus exporter for Starlink metrics
 A Prometheus exporter that collects metrics from Starlink dishes
 including signal quality, obstruction data, alignment statistics,
 and power consumption.
EOF

    # Create postinst script
    cat > "${PKG_DIR}/DEBIAN/postinst" <<'EOF'
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

        # Enable the service (but don't start it automatically)
        systemctl enable starlink-exporter.service || true

        echo "Starlink Exporter has been installed."
        echo "Service runs as user: starlink-exporter"
        echo "To start the service: systemctl start starlink-exporter"
        echo "To view logs: journalctl -u starlink-exporter -f"
        ;;
esac

exit 0
EOF
    chmod 755 "${PKG_DIR}/DEBIAN/postinst"

    # Create prerm script
    cat > "${PKG_DIR}/DEBIAN/prerm" <<'EOF'
#!/bin/sh
set -e

case "$1" in
    remove|deconfigure)
        systemctl stop starlink-exporter.service || true
        systemctl disable starlink-exporter.service || true
        ;;
esac

exit 0
EOF
    chmod 755 "${PKG_DIR}/DEBIAN/prerm"

    # Create postrm script
    cat > "${PKG_DIR}/DEBIAN/postrm" <<'EOF'
#!/bin/sh
set -e

case "$1" in
    purge)
        # Remove state directory
        rm -rf /var/lib/starlink-exporter || true

        # Remove user and group on purge
        if getent passwd starlink-exporter >/dev/null 2>&1; then
            deluser --quiet starlink-exporter || true
        fi

        if getent group starlink-exporter >/dev/null 2>&1; then
            delgroup --quiet starlink-exporter || true
        fi
        ;;

    remove|upgrade|failed-upgrade|abort-install|abort-upgrade|disappear)
        # Do nothing on remove/upgrade - keep user for logs access
        ;;
esac

exit 0
EOF
    chmod 755 "${PKG_DIR}/DEBIAN/postrm"

    # Build the package
    echo "  - Creating .deb package..."
    dpkg-deb --build "${PKG_DIR}" "${DEB_DIR}/${PACKAGE_NAME}_${VERSION}_${ARCH}.deb"

    # Clean up binary
    rm -f starlink_exporter

    echo "  ✓ Created ${PACKAGE_NAME}_${VERSION}_${ARCH}.deb"
}

# Build for each architecture
build_deb "amd64" "amd64" ""
build_deb "arm64" "arm64" ""
build_deb "armhf" "arm" "7"

echo ""
echo "All packages built successfully in ${DEB_DIR}:"
ls -lh "${DEB_DIR}"/*.deb

echo ""
echo "To install a package locally:"
echo "  sudo dpkg -i ${DEB_DIR}/${PACKAGE_NAME}_${VERSION}_<arch>.deb"
