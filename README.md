# Sweep 🧹

[![Go Report Card](https://goreportcard.com/badge/github.com/edwin0x/sweep)](https://goreportcard.com/report/github.com/edwin0x/sweep)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/edwin0x/sweep?status.svg)](https://godoc.org/github.com/edwin0x/sweep)

A fast, elegant system cleanup tool for macOS written in Go.

<p align="center">
  <img src="docs/assets/sweep-demo.gif" alt="Sweep Demo" width="600">
</p>

## ✨ Features

- 🚀 **Fast & Efficient**: Written in Go for maximum performance
- 🔍 **Smart Scanning**: Intelligent system analysis
- 🛡️ **Safe Operations**: Dry-run mode and backup options
- 🎨 **Beautiful CLI**: Clean and intuitive interface
- 🔄 **Real-time Monitoring**: Track system space usage
- ⚡ **Quick Clean**: Fast cleanup for common areas

## 🚀 Installation

### Using Go

```bash
go install github.com/edwin0x/sweep@latest
```

### From Source

```bash
# Clone the repository
git clone https://github.com/edwin0x/sweep.git

# Change to project directory
cd sweep

# Build the project
go build -o sweep cmd/sweep/main.go

# Optional: Install globally
sudo mv sweep /usr/local/bin/
```

## 📚 Usage

```bash
# Show help
sweep --help

# Basic cleanup
sudo sweep clean

# Preview cleanup (dry-run)
sudo sweep clean --dry-run

# Analyze system
sweep scan

# Monitor space usage
sweep monitor
```

### Commands

- `clean`: Remove unnecessary files and caches
- `scan`: Analyze disk usage
- `monitor`: Monitor system space usage
- `quick`: Perform quick cleanup

### Global Flags

- `--dry-run`: Preview changes without making them
- `--verbose, -v`: Show detailed output
- `--silent, -s`: Minimize output
- `--force, -f`: Skip confirmations

## 🛡️ Safety Features

- Dry-run mode to preview changes
- Skip system-critical files
- Confirmation prompts for major operations
- Detailed logging of all operations

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to all contributors who have helped with code, bug reports, and suggestions
- Built with [Cobra](https://github.com/spf13/cobra) and [Color](https://github.com/fatih/color)

## 📊 Project Status

This project is actively maintained and accepting contributions.