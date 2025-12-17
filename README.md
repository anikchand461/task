# Task CLI 📝

A simple, fast, cross-platform **command-line task manager** written in Go.

This tool works as a **terminal command** (`task`) and does **NOT require Go** to be installed for normal users.

## Use CLI for task mangement
![SCR-20251217-qvaz](https://github.com/user-attachments/assets/68ae944d-a0fa-4853-94bf-364b54513c37)
![SCR-20251217-qvfa](https://github.com/user-attachments/assets/bf36b908-4d7f-4c25-8100-ccfd7b17c629)

---

## 🚀 Setup & Installation (No Go Required)

Below are step-by-step setup instructions for **each operating system**.

---

## 🟢 macOS Setup

### 1️⃣ Download the binary

Go to the **GitHub Releases** page:

👉 [https://github.com/anikchand461/task/releases](https://github.com/anikchand461/task/releases)

Download the correct file:

* **Apple Silicon (M1 / M2 / M3)** → `task-mac`
* **Intel Mac** → `task-mac-intel`

---

### 2️⃣ Make the binary executable

Open Terminal and run:

```bash
cd ~/Downloads
chmod +x task-mac
```

(Use `task-mac-intel` if you downloaded the Intel version)

---

### 3️⃣ Move it to PATH (make it a command)

```bash
sudo mv task-mac /usr/local/bin/task
```

Now the command `task` is available system-wide.

---

### 4️⃣ macOS security fix (IMPORTANT)

macOS may block the binary because it is not notarized.

If the command is killed immediately, run:

```bash
sudo xattr -rd com.apple.quarantine /usr/local/bin/task
```

---

### 5️⃣ Verify installation

```bash
task
```

---

## 🟦 Windows Setup

### 1️⃣ Download the binary

From the Releases page, download:

```
task.exe
```

---

### 2️⃣ Move to a folder

Example:

```
C:\task
```

---

### 3️⃣ Add folder to PATH

1. Search **Environment Variables**
2. Open **Edit system environment variables**
3. Click **Environment Variables**
4. Select **Path** → Edit → New
5. Add:

   ```
   C:\task
   ```
6. Save and restart terminal

---

### 4️⃣ Verify installation

Open Command Prompt or PowerShell:

```powershell
task
```

---

## 🟠 Linux Setup

### 1️⃣ Download the binary

From the Releases page, download:

```
task-linux
```

---

### 2️⃣ Make it executable

```bash
chmod +x task-linux
```

---

### 3️⃣ Move to PATH

```bash
sudo mv task-linux /usr/local/bin/task
```

---

### 4️⃣ Verify installation

```bash
task
```

---

## 📌 Usage Guide

### Add a task

```bash
task add "Buy groceries"
```

---

### List all tasks

```bash
task list
```

Output example:

```
1. [ ] Buy groceries
2. [x] Learn Go
```

---

### Mark a single task as done

```bash
task done 1
```

---

### Mark **all** tasks as done

```bash
task done all
```

---

### Clear (delete) all tasks

```bash
task clear
```

⚠️ This deletes **all tasks for the current user only**.

---

### Help

```bash
task help
task -h
task --help
```

Displays all available commands and usage information.

---

## 📂 Data Storage

Tasks are stored **per user** at:

```text
~/.task/tasks.txt
```

* Each user has a separate task list
* Data persists across updates
* Removing the binary does **not** delete tasks
* Use `task clear` to remove all stored tasks

---

## 🗑️ Uninstall

### macOS / Linux

Remove the binary:

```bash
sudo rm /usr/local/bin/task
```

(Optional) Remove all task data:

```bash
rm -rf ~/.task
```

---

### Windows

1. Delete `task.exe` from its folder
2. (Optional) Remove task data:

```powershell
rmdir /s %USERPROFILE%\.task
```

---

## 🛠 Development (For Contributors)

If you want to modify the source code:

```bash
git clone https://github.com/anikchand461/task.git
cd task
go run ./cmd/task list
```

Build locally:

```bash
go build -o task ./cmd/task
```

---

## 📄 License

MIT License

---

## 🙌 Author

**Anik Chand**
Built with ❤️ using Go.

