#!/bin/bash
set -e

# Publish Debian packages to GitHub Packages
# Requires: gh CLI (GitHub CLI) authenticated with repo permissions

VERSION="${VERSION:-1.0.0}"
PACKAGE_NAME="starlink-exporter"
DEB_DIR="$(pwd)/build/deb"
REPO="brendanbank/starlink_exporter"

echo "Publishing Debian packages to GitHub Packages..."

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI (gh) is not installed."
    echo "Install it from: https://cli.github.com/"
    exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
    echo "Error: Not authenticated with GitHub CLI."
    echo "Run: gh auth login"
    exit 1
fi

# Check if packages exist
if [ ! -d "${DEB_DIR}" ]; then
    echo "Error: No packages found in ${DEB_DIR}"
    echo "Run 'make build-deb' first to build packages."
    exit 1
fi

# Create a release if it doesn't exist
echo "Creating GitHub release v${VERSION}..."
gh release create "v${VERSION}" \
    --repo "${REPO}" \
    --title "Release v${VERSION}" \
    --notes "Debian packages for starlink-exporter v${VERSION}

## Installation

### Ubuntu/Debian AMD64
\`\`\`bash
wget https://github.com/${REPO}/releases/download/v${VERSION}/${PACKAGE_NAME}_${VERSION}_amd64.deb
sudo dpkg -i ${PACKAGE_NAME}_${VERSION}_amd64.deb
\`\`\`

### Ubuntu/Debian ARM64
\`\`\`bash
wget https://github.com/${REPO}/releases/download/v${VERSION}/${PACKAGE_NAME}_${VERSION}_arm64.deb
sudo dpkg -i ${PACKAGE_NAME}_${VERSION}_arm64.deb
\`\`\`

### Ubuntu/Debian ARMHF (Raspberry Pi)
\`\`\`bash
wget https://github.com/${REPO}/releases/download/v${VERSION}/${PACKAGE_NAME}_${VERSION}_armhf.deb
sudo dpkg -i ${PACKAGE_NAME}_${VERSION}_armhf.deb
\`\`\`

## After Installation

Start the service:
\`\`\`bash
sudo systemctl start starlink-exporter
sudo systemctl status starlink-exporter
\`\`\`

View logs:
\`\`\`bash
sudo journalctl -u starlink-exporter -f
\`\`\`
" \
    2>/dev/null || echo "Release v${VERSION} already exists, uploading packages..."

# Upload each package
for deb in "${DEB_DIR}"/*.deb; do
    if [ -f "$deb" ]; then
        echo "Uploading $(basename "$deb")..."
        gh release upload "v${VERSION}" "$deb" \
            --repo "${REPO}" \
            --clobber
        echo "  ✓ Uploaded $(basename "$deb")"
    fi
done

echo ""
echo "All packages published successfully!"
echo "View release at: https://github.com/${REPO}/releases/tag/v${VERSION}"
echo ""
echo "Users can now install with:"
echo "  wget https://github.com/${REPO}/releases/download/v${VERSION}/${PACKAGE_NAME}_${VERSION}_<arch>.deb"
echo "  sudo dpkg -i ${PACKAGE_NAME}_${VERSION}_<arch>.deb"
