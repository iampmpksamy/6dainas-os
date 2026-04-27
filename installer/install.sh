#!/bin/bash

set -e

echo "🚀 Installing 6DAiNAS..."

VERSION=${1:-latest}
BASE_URL="https://github.com/iampmpksamy/6dainas-os/releases"

TMP_DIR="/tmp/6dainas"
INSTALL_DIR="/"

# Detect version

if [ "$VERSION" == "latest" ]; then
DOWNLOAD_URL="$BASE_URL/latest/download/6dainas-os.tar.gz"
else
DOWNLOAD_URL="$BASE_URL/download/$VERSION/6dainas-os-$VERSION.tar.gz"
fi

echo "📦 Downloading package..."
mkdir -p $TMP_DIR
curl -L $DOWNLOAD_URL -o $TMP_DIR/6dainas.tar.gz

echo "📂 Extracting..."
tar -xzf $TMP_DIR/6dainas.tar.gz -C $TMP_DIR

echo "📁 Installing files..."
cp -r $TMP_DIR/release/* $INSTALL_DIR

echo "⚙️ Setting permissions..."
chmod +x /usr/bin/6dainas-* || true

echo "🔄 Reloading systemd..."
systemctl daemon-reexec
systemctl daemon-reload

echo "🚀 Enabling services..."
systemctl enable 6dainas-gateway || true
systemctl enable 6dainas-message-bus || true
systemctl enable 6dainas-user-service || true
systemctl enable 6dainas-app-management || true
systemctl enable 6dainas-main || true

echo "▶️ Starting services..."
systemctl start 6dainas-gateway || true

echo "✅ 6DAiNAS Installed Successfully!"

echo "🌐 Access UI via: http://localhost"
