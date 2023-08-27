# Commit - Automated Git Commit Messages with LLM 🤖📝
Artfully create commit messages that reflect the essence of your code changes 🎨🔍. Craftsmanship for your commits 👨‍🎨. Unleash the power of language models to automate your Git commit messages 🌐🤖. With Commit, save time ⏰ and generate meaningful commit messages based on your code changes 📜.

## Table of Contents 📚

- [Features 🌟](#features)
- [Prerequisites 📋](#prerequisites)
- [Installation 🛠️](#installation)
  - [Environment Variables 🌍](#environment-variables)
  - [Complete Install 📦](#complete-install)
- [Usage 🖱️](#usage)
- [Contributing 🤝](#contributing)
  - [How to Contribute 👷‍♂️](#how-to-contribute)
  - [Code of Conduct 📜](#code-of-conduct)
  - [Community 💬](#community)
- [License 📄](#license)
- [Acknowledgements 🙏](#acknowledgements)

## Features 🌟

- **Automated Commit Messages**: Leverages language models to craft meaningful commit messages 🤖💬.
- **Easy to Install**: One command installs everything you need 🚀.
- **GitHub CLI Integration**: Seamlessly integrates as a GitHub CLI extension 🔄.

## Prerequisites 📋

- Go 1.16+ 🟢
- Git 2.30+ 📦
- GitHub CLI 2.0+ 🔗

## Installation 🛠️

### Environment Variables 🌍

Before running AutoCommit, it's advisable to set a few environment variables 🔑:

- `GPT_API_KEY`: The API key for the GPT-4 model (🚨 **Required**).
- `LLM_MODEL`: Specify a different language model 🔄 (Optional; Default: `gpt-4.5-turbo`).
- `FINE_TUNE_PARAMS`: Additional parameters for fine-tuning the model output ⚙️ (Optional; Default: `{}`).

Add these environment variables by appending them to your `.bashrc`, `.zshrc`, or other shell configuration files 📄:

\```bash
export GPT_API_KEY=your-api-key-here
export LLM_MODEL=gpt-4.5-turbo
export FINE_TUNE_PARAMS='{"temperature": 0.7}'
\```

Or, you can set them inline before running the AutoCommit command 🖱️:

\```bash
GPT_API_KEY=your-api-key-here LLM_MODEL=gpt-4.5-turbo FINE_TUNE_PARAMS='{"temperature": 0.7}' git auto-commit
\```

### Complete Install 📦

For an end-to-end installation experience, execute 👇:

\```bash
bash <(curl -s https://raw.githubusercontent.com/ghcli/commit/main/install.sh)
\```

This comprehensive script accomplishes the following 📋:

1. Downloads the latest `generateCommitMessage` binary ⬇️.
2. Makes the binary executable 🏃.
3. Sets up a Git alias: `auto-commit` 🏷️.
4. Installs the GitHub CLI extension for AutoCommit 🔄.

## Usage 🖱️

### Native Git 🌐

To auto-generate a commit message, type ⌨️:

\```bash
git auto-commit
\```

### GitHub CLI Extension 🔗

For the same functionality through GitHub CLI, execute 🤖:

\```bash
gh commit
\```

Both commands invoke a Git diff, pass the changes to GPT-4, and craft a commit message based on the model's output 💬🎉.

## Contributing 🤝

### How to Contribute 👷‍♂️

1. Fork the repository 🍴.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`) 🌳.
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`) 📝.
4. Push to the branch (`git push origin feature/AmazingFeature`) ⬆️.
5. Open a pull request 🤲.

### Code of Conduct 📜

Please read the `CODE_OF_CONDUCT.md` for guidelines on community behavior 👥.

### Community 💬

See community discussions, and follow the project board for current and upcoming features 📅.

## License 📄

MIT License. For more information, please refer to the `LICENSE` file in the repo 📑.

## Acknowledgements 🙏

- Thanks to OpenAI for providing the models 🌐.
- All the contributors who made this project possible 👨‍👩‍👧‍👦.
