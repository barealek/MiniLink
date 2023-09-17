<!-- markdownlint-configure-file {
  "MD013": {
    "code_blocks": false,
    "tables": false
  },
  "MD033": false,
  "MD041": false
} -->

<div align="center">

# miniLink

⚡ Fast - 🪶 Lightweight - 📋 Open Source

miniLink is a fast, easy deployable URL shortener service. 

miniLink is written in Go, using the Fiber framework for easy path parameters.
The front end is written in Svelte and Tailwind.
It uses MongoDB to store the different shortened links. 
miniLink uses Cloudflared to expose the app

[Preview Service](#preview-service) 
•
[Installation](#installation) 
•
[Configuration](#configuration)

</div>

## Preview Service

If you would like to try the project without self-hosting it, head over to [minilink.imalek.me](https://minilink.imalek.me).


## Installation

miniLink can be self-hosted in x easy steps:

1. **Install Docker**

   The preferred way to run miniLink is through Docker Compose. You can install Docker from the [official Docker website](https://www.docker.com/).
   You can also just build the miniLink "backend" image and use it however you'd like.


2. **Clone down the repository**

   You want to download all the code from the repository. This can be done either by using the `git` CLI app, or with the "Download Zip" button.
   `git clone https://github.com/barealek/miniLink`


3. **Set up environment variables**

   miniLink requires a few environment variables to work. This includes the database connection string, (optionally) cloudflare tunnel


## Configuration

### Flags

When calling `zoxide init`, the following flags are available:

- `--cmd`
  - Changes the prefix of the `z` and `zi` commands.
  - `--cmd j` would change the commands to (`j`, `ji`).
  - `--cmd cd` would replace the `cd` command (doesn't work on Nushell / POSIX shells).
- `--hook <HOOK>`
  - Changes how often zoxide increments a directory's score:
    | Hook     | Description                       |
    | -------- | --------------------------------- |
    | `none`   | Never                             |
    | `prompt` | At every shell prompt             |
    | `pwd`    | Whenever the directory is changed |
- `--no-cmd`
  - Prevents zoxide from defining the `z` and `zi` commands.
  - These functions will still be available in your shell as `__zoxide_z` and
    `__zoxide_zi`, should you choose to redefine them.

### Environment variables

Environment variables[^2] can be used for configuration. They must be set before
`zoxide init` is called.
