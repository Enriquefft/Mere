#!/bin/bash
set -e

# Mere Installation Script
# This script downloads and installs the Mere CLI tool

VERSION="${VERSION:-latest}"
REPO="Enriquefft/Mere"
GITHUB_BASE="https://github.com/${REPO}/releases"

# Detect NixOS early so it can influence defaults
IS_NIXOS=false
if [ -f /etc/os-release ] && grep -q "NixOS" /etc/os-release; then
    IS_NIXOS=true
fi

# On NixOS, default to ~/.local/bin (no sudo, writable, usually in PATH)
if [ "$IS_NIXOS" = true ]; then
    INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
else
    INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
fi

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print colored output
info() {
    echo -e "${GREEN}[INFO]${NC} $1" >&2
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1" >&2
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

# Detect OS
detect_os() {
    local os
    case "$(uname -s)" in
        Linux*)     os=linux;;
        Darwin*)    os=darwin;;
        MINGW*|MSYS*|CYGWIN*) os=windows;;
        *)          error "Unsupported OS: $(uname -s)"; exit 1;;
    esac
    echo "$os"
}

# Detect architecture
detect_arch() {
    local arch
    case "$(uname -m)" in
        x86_64|amd64)    arch=amd64;;
        arm64|aarch64)    arch=arm64;;
        i386|i686)       arch=386;;
        armv6l|armv7l)   arch=arm;;
        *)               error "Unsupported architecture: $(uname -m)"; exit 1;;
    esac
    echo "$arch"
}

# Get latest version from GitHub
get_latest_version() {
    if [ "$VERSION" != "latest" ]; then
        echo "$VERSION"
        return
    fi

    info "Fetching latest version from GitHub..."
    local latest_version
    latest_version=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$latest_version" ]; then
        error "Failed to fetch latest version"
        exit 1
    fi

    echo "$latest_version"
}

# Download the binary
download_binary() {
    local version="$1"
    local os="$2"
    local arch="$3"
    local filename="mere-${os}-${arch}.tar.gz"
    local download_url="${GITHUB_BASE}/download/${version}/${filename}"

    info "Downloading Mere ${version} for ${os}/${arch}..."
    info "From: ${download_url}"

    # Create temp directory for download
    local tmp_dir=$(mktemp -d)
    local tmp_archive="${tmp_dir}/${filename}"
    local tmp_file="${tmp_dir}/mere"

    # Download archive
    if ! curl -fsSL "$download_url" -o "$tmp_archive"; then
        error "Failed to download Mere from ${download_url}"
        exit 1
    fi

    # Extract binary from archive
    tar -xzf "$tmp_archive" -C "$tmp_dir" mere

    # Make binary executable
    chmod +x "$tmp_file"

    echo "$tmp_file"
}

# Install the binary
install_binary() {
    local source="$1"
    local target="${INSTALL_DIR}/mere"

    # Check if install directory exists, create if needed
    if [ ! -d "$INSTALL_DIR" ]; then
        warn "Install directory ${INSTALL_DIR} does not exist. Creating it..."
        if mkdir -p "$INSTALL_DIR" 2>/dev/null; then
            :
        else
            sudo mkdir -p "$INSTALL_DIR"
        fi
    fi

    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        info "Installing Mere to ${target}..."
        cp "$source" "$target"
        chmod +x "$target"
    else
        info "Installing Mere to ${target} (requires sudo)..."
        sudo cp "$source" "$target"
        sudo chmod +x "$target"
    fi

    # Clean up temp directory
    rm -rf "$(dirname "$source")"
}

# Check PATH
check_path() {
    if ! echo "$PATH" | grep -q "${INSTALL_DIR}"; then
        warn "${INSTALL_DIR} is not in your PATH"
        warn "Add the following to your shell configuration:"
        warn "  export PATH=\"\$PATH:${INSTALL_DIR}\""
    fi
}

# Detect user's shell and validate config
detect_shell_config() {
    local shell_name=$(basename "$SHELL")
    local shell_config=""

    case "$shell_name" in
        zsh)
            shell_config="$HOME/.zshrc"
            ;;
        bash)
            shell_config="$HOME/.bashrc"
            ;;
        fish)
            shell_config="$HOME/.config/fish/config.fish"
            ;;
        ksh)
            shell_config="$HOME/.kshrc"
            ;;
        *)
            warn "Unknown shell: $shell_name"
            return 1
            ;;
    esac

    # Check if config exists
    if [ ! -f "$shell_config" ]; then
        warn "Shell config not found at $shell_config"
        return 1
    fi

    echo "$shell_config"
}

# Main installation
main() {
    echo ""
    echo "Mere Installation Script"
    echo "========================"
    echo ""

    # Detect OS and architecture
    local os=$(detect_os)
    local arch=$(detect_arch)
    info "Detected OS: ${os}"
    info "Detected Architecture: ${arch}"

    if [ "$IS_NIXOS" = true ]; then
        warn "NixOS detected. Installing to ${INSTALL_DIR}."
        warn "If the binary fails to run, glibc path mismatch may be the cause."
        warn "In that case, consider using: nix run github:${REPO}"
    fi

    echo ""

    # Get version
    local version
    version=$(get_latest_version) || exit 1
    info "Installing version: ${version}"
    echo ""

    # Download binary
    local binary_path
    binary_path=$(download_binary "$version" "$os" "$arch") || exit 1

    # Install binary
    install_binary "$binary_path"

    echo ""
    info "Mere installed successfully!"
    echo ""
    info "Version: $(mere version 2>/dev/null || echo 'unknown')"
    echo ""

    # Check PATH
    check_path

    echo ""
    info "Next steps:"
    echo "  1. Reload your shell configuration:"

    shell_config=$(detect_shell_config)
    if [ $? -eq 0 ]; then
        echo "     source $shell_config"
    else
        echo "     • Create your shell config file, or"
        echo "     • Run: /usr/local/bin/mere init (with full path)"
    fi

    echo "  2. Run: mere init"
    echo "  3. Use /module to create your first deep module"
    echo ""
}

# Run main function
main
