# ShellHit - AI-Powered Terminal Built with Go

🚀 **Revolutionize Your Command-Line Experience!**

Stop memorizing complex commands! **ShellHit** is your AI-powered terminal, written in **Go**, that understands natural language and executes the right commands for you. Whether setting up a Python virtual environment, cloning your github repos, or initializing a Node.js project, ShellHit has you covered. 🎯

## ✨ Features

- 🤖 **AI-driven command execution** – Describe what you need in plain English!
- 🔥 **Admin privileges support** – No more permission issues!
- ⚡ **Full terminal capabilities** – Supports common commands like `cd`, `ls`, `mkdir`, `rm`, `npm install`, `pip install`, and more.
- 🎯 **Google Gemini AI integration** – Translates natural language into valid Windows, macOS, and Linux commands.
- 🚀 **Fast & Seamless** – Automates tasks efficiently.
- 🎨 **Enhanced readability** – Beautifully formatted output.
- 📦 **One-click installer available** – Download from **GitHub Releases** and get started instantly!

---

## 🛠️ Installation

### Option 1: **Using the Installer** (Recommended)

1. **Download the latest installer** from [GitHub Releases](https://github.com/yourusername/shellhit/releases).
2. Run the installer and follow the setup instructions.
3. Launch ShellHit as an **Administrator** and start automating! 🚀

### Option 2: **Manual Installation**

1. **Clone the Repository**
   ```bash
   git clone https://github.com/saranb565/ShellHit.git
   cd shellhit
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Set Up API Key**
   - Create a `.env` file in the project root.
   - Add your **Google Gemini AI API key**:
     ```env
     API_KEY=your_google_gemini_api_key
     ```

4. **Run ShellHit**
   ```bash
   go run main.go
   ```

> **Note:** Some commands require **admin privileges**!

---

## 🚀 How It Works

### 🔹 Execute Basic Commands
```bash
ShellHit> ls
ShellHit> mkdir my_project
ShellHit> cd my_project
ShellHit> pwd
```

### 🔹 Use AI to Automate Tasks
```bash
ShellHit> Find and list all running processes
ShellHit> Create a directory 'test_project' and initialize npm
ShellHit> Show my current directory
ShellHit> Kill the process using port 8080
ShellHit> Clone a GitHub repository and navigate into it
ShellHit> Set up a Python virtual environment and install Flask
ShellHit> Install TensorFlow globally
ShellHit> Compile and execute a Java program
```
And much more....

✅ **Your AI-powered terminal converts requests into executable commands instantly!**

---

## 🏆 Why Choose ShellHit?

- **No need to memorize commands** – Just type naturally!
- **Optimized for speed** – Executes commands with minimal delay.
- **Beginner-friendly yet powerful** – Great for all experience levels.
- **Perfect for automation** – Boosts productivity.
- **Cross-platform support** – Works on **Windows, macOS, and Linux**.

---

## 🛠️ Troubleshooting

- **Getting API errors?**
  - Ensure your `.env` file contains a valid **API key**.
  - Check if your **Google Gemini AI quota** is exhausted.
- **Admin permission issues?**
  - Run the program as **Administrator**.
- **Command not recognized?**
  - AI may not have an equivalent command for your OS. Try rephrasing!
- **Installer issues?**
  - Download the latest version from **GitHub Releases** and ensure security software isn’t blocking it.

---

## 📜 License

This project is **open-source** and available under the **MIT License**.

---

## 🤝 Contributing

Want to enhance ShellHit? Fork the repo and submit a pull request! 🚀

---

## 🎯 Future Roadmap

- ✅ Support for **Linux & macOS**
- ✅ Custom AI model for improved command execution
- ✅ Enhanced **error handling and logging**
- ✅ Advanced scripting automation
- ✅ Cloud-based command execution for remote tasks
- ✅ Integration with CI/CD pipelines

---

### ❤️ Experience AI-powered efficiency with **ShellHit**! 🚀
### Built with ❤️ by Saran B

