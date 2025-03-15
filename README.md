# 🏗️ JCL Generator v0.1

A **command-line tool** to generate **JCL (Job Control Language)** scripts dynamically based on user inputs. Designed for **mainframe job automation**, this tool simplifies the creation of structured JCL jobs without manually writing extensive scripts.

## 📌 Features
✅ **Fully optional parameters** – Only specified parameters are included in the final JCL.  
✅ **JCL syntax compliance** – Ensures generated JCL is valid and structured properly.  
✅ **Interactive Mode** – Prompts for missing values when `-interactive` or `-i` is used.  
✅ **Alias Flags** – Shortened versions of common flags for easier use.  
✅ **Cross-platform support** – Available for Windows, Linux, and macOS.  
✅ **Automated builds & releases** – Precompiled binaries available for easy download.  
✅ **Future roadmap** – Plans to support YAML input, job steps, and advanced system overrides in upcoming versions.  

## 🚀 Installation

### **🔹 Option 1: Download Precompiled Binary**
1. **Go to the [Releases](https://github.com/MederickBernier/JCLGenerator/releases) page.**
2. **Download the appropriate ZIP file** for your OS:
   - **Windows** → `jclgen-windows.zip`
   - **Linux** → `jclgen-linux.zip`
   - **macOS** → `jclgen-darwin.zip`
3. **Extract the ZIP** and move the binary to a directory in your `$PATH` (optional).

### **🔹 Option 2: Build from Source (Requires Go)**
```bash
git clone https://github.com/MederickBernier/JCLGenerator.git
cd JCLGenerator
go mod tidy
go build -o jclgen ./cmd/main.go
```


---

### **📌 Usage**
```markdown
## 🎯 Usage

Run the JCL Generator with command-line arguments:

```bash
jclgen -jobname=TESTJOB -class=A -msgclass=X -output=myjob
```
### **🔹 Interactive Mode (New in v0.2)**  
Run without flags to enter interactive mode:  
```bash
jclgen -interactive  
jclgen -i  
```
It will prompt for missing values:

Enter Job Name: TESTJOB  
Enter Class: A  
Enter MsgClass: X  
Enter Output Filename: myjob  

### **Example Output**
```
//TESTJOB JOB (ACCT),'Generated JCL'
// CLASS=A
// MSGCLASS=X

//STEP1 EXEC PGM=IEFBR14
```
---

### **🔹 Update Flags Table**  
📍 **Location**: Under **"## 🎯 Usage"**, in the flags table.  
📌 **Modification**: **Add alias flags and interactive flag**

| **Flag**            | **Alias** | **Description**                       | **Example**          |
|---------------------|----------|---------------------------------------|----------------------|
| -jobname         | -jn     | Job name (1-8 uppercase characters)  | -jobname=TESTJOB  |
| -class           | -c      | Job execution class (A-Z)            | -class=A          |
| -msgclass        | -mc     | Message output class (A-Z, 0-9)      | -msgclass=X       |
| -output          | -o      | Output JCL filename (without .jcl) | -output=myjob     |
| -interactive     | -i      | Enables interactive input mode       | -i                |


For advanced parameters like DD statements, JES2, system overrides, future versions will include YAML support.


---

### **📌 Development & Contribution**
```markdown
## 🛠️ Development & Contribution

### **🔹 Requirements**
- **Go 1.24 or higher**
- Git installed for cloning and versioning

### **🔹 Building the Project**
```bash
go build -o jclgen ./cmd/main.go
```

### **📌 Running Tests**
(Currently, validation functions are embedded but future versions will include unit tests.)

### **📌 Contributing**
Want to help? Open an issue, fork the repo, and submit a pull request!
Roadmap includes YAML support, interactive input mode, and JCL validation improvements.


---

### **📌 Versioning & Releases**
```markdown
## 🔄 Versioning & Releases

- **v0.1** – Core functionality, CLI-based JCL generation.  
- **v0.2** – Interactive CLI mode, alias flags for ease of use.  
- **Upcoming:**  
  - v0.3 – YAML import & job steps.  

To download a specific version, visit the [Releases](https://github.com/MederickBernier/JCLGenerator/releases) page.
```

## 📜 License

This project is licensed under the **MIT License** – you are free to modify, distribute, and use it.


## 🏛️ Credits & Acknowledgments

Developed by **Mederick Bernier**, designed to **simplify JCL script generation** for mainframe environments.  
For questions, feel free to reach out via GitHub issues.

🚀 Happy JCL scripting!
