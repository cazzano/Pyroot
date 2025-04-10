# 🌲 PyRoot
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**PyRoot** is a powerful command-line tool designed to simplify Python project management and execution.

## ✨ Features
- **Project Initialization** - Set up new Python projects with proper structure
- **Virtual Environment Management** - Create and manage Python virtual environments
- **Build Automation** - Package your Python projects into standalone executables
- **Execution Control** - Run Python files directly with support for arguments
- **Flexible Usage** - Target specific files or modules for execution

## 📦 Build && Installation From Source
```bash
# Clone the repository
git clone https://github.com/cazzano/pyroot.git

# Navigate to the directory
cd pyroot/stable/go/src/

# Install the tool
go build && mv src pyroot && sudo mv pyroot /usr/bin && echo "You Installed It Hah !!!"
```
## Installation From Release

```bash
wget https://github.com/cazzano/pyroot/releases/download/vpn/pyroot_vr-1.0_x86_64_arch.zip;

unzip pyroot_vr-1.0_x86_64_arch.zip && sudo mv pyroot /usr/bin/ && rm pyroot_vr-1.0_x86_64_arch.zip && echo "Yeah You Installed It!!!!";
```


## 🚀 Usage
PyRoot provides several commands to simplify your Python development workflow:

### Initialize an existing project with requirements.txt
```bash
pyroot init
```
This will create a requirements.txt file in the current directory.

### Manage virtual environments
```bash
pyroot vnv
```
This will set up virtual environment commands (pyvnv and pypip) in your shell.

### Create a new project
```bash
pyroot new my-project
```
This will create a project structure:
```
my-project/
├── src/
├── target/
└── requirements.txt
```

### Build your project
```bash
pyroot build
```
Packages your Python project into an executable binary in the `target/release/` directory using PyInstaller.

### Run Python files
Run Python files in the current directory:
```bash
pyroot run
```

### Run a specific file or module
```bash
pyroot run --1 ./path/to/file.py
```
```bash
pyroot run --1 specific_module.py
```

### Display version information
```bash
pyroot --v
```

### Display help message
```bash
pyroot --h
```

## 📝 Examples
### Example 1: Quick Start Project
```bash
# Create a new project
pyroot new my-awesome-app

# Navigate to the source directory
cd my-awesome-app/src

# Run the project
pyroot run

# Build the project
pyroot build

# Your compiled binary will be in my-awesome-app/target/release/
```

### Example 2: Virtual Environment Management
```bash
# Set up virtual environment commands
pyroot vnv

# Use the virtual environment Python
pyvnv -m pip install requests

# Install packages with the virtual environment pip
pypip install numpy pandas
```

### Example 3: Working with Specific Files
```bash
# Run a specific file
pyroot run --1 main.py

# Run a specific module
pyroot run --1 utils.py
```

## 🛠️ Command Reference
| Command | Description |
|---------|-------------|
| `init` | Initialize the existing project with requirements.txt |
| `vnv` | Initialize commands to manage your Python virtual environment |
| `new` | Initialize a new project structure |
| `build` | Build the project to target/release folder using PyInstaller |
| `run` | Run Python files in the current directory |
| `run --1` | Run a specific file or module |
| `--v` | Display version information |
| `--h` | Display help message |

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support
For support, please open an issue in the GitHub repository or contact the maintainers.
