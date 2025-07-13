# Installation Guide

## Quick Install with go install

If you have Go installed and want to install directly from GitHub:

```bash
go install github.com/mph-llm-experiments/denote-contacts@latest
```

## Build from Source

### Prerequisites
- Go 1.21 or later

### Steps

1. Clone the repository:
```bash
git clone https://github.com/mph-llm-experiments/denote-contacts.git
cd denote-contacts
```

2. Build the binary:
```bash
go build -o denote-contacts .
```

3. Install to your PATH:
```bash
# Option 1: Copy to /usr/local/bin (requires sudo)
sudo cp denote-contacts /usr/local/bin/

# Option 2: Copy to ~/bin (create if needed)
mkdir -p ~/bin
cp denote-contacts ~/bin/
# Add ~/bin to your PATH if not already there
```

4. Set up configuration:
```bash
./setup-config.sh
# Edit ~/.config/denote-contacts/config.toml to set your notes directory
```

## Daily Use

### Running the Application
```bash
denote-contacts
```

### Configuration
The app looks for configuration in this order:
1. Environment variable: `DENOTE_CONTACTS_DIR`
2. Config file: `~/.config/denote-contacts/config.toml`
3. Default: `~/Documents/denote`

### First Run
1. Create your notes directory if it doesn't exist
2. Either set `DENOTE_CONTACTS_DIR` or edit the config file
3. Run `denote-contacts`

## Updating

To update to the latest version:

```bash
# If installed with go install
go install github.com/mph-llm-experiments/denote-contacts@latest

# If built from source
cd denote-contacts
git pull
go build -o denote-contacts .
# Then copy to your installation location
```

## Troubleshooting

### Binary not found
Make sure the installation directory is in your PATH:
```bash
echo $PATH
```

### Config not loading
Check that config file exists:
```bash
ls -la ~/.config/denote-contacts/config.toml
```

### Notes directory not found
Verify the path in your config:
```bash
cat ~/.config/denote-contacts/config.toml
```