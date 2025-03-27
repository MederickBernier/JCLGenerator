# 🏗️ JCL Generator v0.3

A **command-line tool** to generate **JCL (Job Control Language)** scripts dynamically based on user input or configuration files. Designed for **mainframe job automation**, this tool simplifies JCL creation for developers, architects, and operations engineers.

---

## 📌 Features

✅ **Fully optional parameters** – Only defined parameters are rendered in the final JCL  
✅ **JCL syntax compliance** – Ensures generated jobs are valid for mainframe execution  
✅ **Interactive Mode** – Prompt-based input when `-interactive` or `-i` is used  
✅ **YAML Config Support (v0.3)** – Load parameters from `config.yaml` via `--from-yaml`  
✅ **Short & long flag aliases** – Flexible and ergonomic CLI interface  
✅ **Cross-platform** – Windows, Linux, macOS supported  
✅ **Easy builds** – Build with Go or download from GitHub releases

---

## 🚀 Installation

### 🔹 Option 1: Download Precompiled Binary

1. Go to the [Releases](https://github.com/MederickBernier/JCLGenerator/releases)
2. Download the binary for your platform:
   - Windows → `jclgen-windows.zip`
   - Linux → `jclgen-linux.zip`
   - macOS → `jclgen-darwin.zip`
3. Extract and (optionally) move to a directory in your `$PATH`

### 🔹 Option 2: Build from Source

```bash
git clone https://github.com/MederickBernier/JCLGenerator.git
cd JCLGenerator
go mod tidy
go build -o jclgen ./cmd/main.go
```

---

## 🎯 Usage

### 🔸 Basic CLI Mode

```bash
jclgen -jobname=TESTJOB -class=A -msgclass=X -output=myjob
```

### 🔸 Interactive Mode

```bash
jclgen -interactive
# or
jclgen -i
```

You'll be prompted to enter values like:

```text
Enter Job Name: TESTJOB
Enter Class: A
Enter MsgClass: X
Enter Output Filename: myjob
```

### 🔸 YAML Input (New in v0.3)

Prepare a `config.yaml` like:

```yaml
job_name: TEST
account_number: ACCT123
job_description: "Generated from YAML"
class: A
msg_class: X
enable_comments: true
exec_pgm: IEFBR14
ds_name: MY.DATA.SET
disp: (NEW,CATLG,DELETE)
steps:
  - step_name: STEP1
    exec_pgm: IEFBR14
```
See [`examples/config.yaml`](examples/config.yaml) for a sample config.

Then run:

```bash
jclgen --from-yaml=config.yaml
# or
jclgen -fy=config.yaml
```

---

## 🧾 Flags Table

| **Flag**         | **Alias** | **Description**                        | **Example**                |
|------------------|-----------|----------------------------------------|----------------------------|
| `--jobname`      | `-j`      | JCL Job name (1-8 chars)               | `--jobname=MYJOB`          |
| `--class`        | `-c`      | Execution class (A-Z)                  | `--class=A`                |
| `--msgclass`     | `-m`      | Message class                          | `--msgclass=X`             |
| `--output`       | `-o`      | Output filename (.jcl extension added) | `--output=myjob`           |
| `--interactive`  | `-i`      | Launch in interactive mode             | `--interactive`            |
| `--with-comments`| `-wc`     | Toggle verbose comments in output      | `--with-comments=false`    |
| `--from-yaml`    | `-fy`     | Load configuration from YAML file      | `--from-yaml=config.yaml`  |

---

## 🛠️ Development & Contribution

### Requirements
- Go 1.21+  
- Git

### Building Locally

```bash
go mod tidy
go build -o jclgen ./cmd/main.go
```

### Roadmap

| Version | Features |
|---------|----------|
| **v0.1** | Base CLI support |
| **v0.2** | Interactive mode + alias flags |
| ✅ **v0.3** | YAML input support + job steps scaffolding |
| 🚧 **v0.4** | Fully interactive wizard mode (planned) |
| 🚧 **v0.5** | GUI/Browser interface (experimental idea) |

---

## 🧪 Testing

Unit tests will be added in future versions for validation logic and YAML parsing. For now, test by running various flag combinations or YAML input files.

---

## 📜 License

MIT License – Free to use, modify, and distribute.

---

## 🏛️ Credits

Developed by **Mederick Bernier** to support legacy modernization, JCL automation, and structured mainframe workflows.

---

> ⭐ Star the repo if this tool helps you!
