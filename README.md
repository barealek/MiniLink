<!-- markdownlint-configure-file {
  "MD013": {
    "code_blocks": false,
    "tables": false
  },
  "MD033": false,
  "MD041": false
} -->

<div align="center">

# miniLink🔗

⚡ Fast - 🪶 Lightweight - 📋 Open Source

miniLink is a fast, easily deployable URL shortener service. 

miniLink is written in Go, using the Fiber framework for easy path parameters.
The front end is written in Svelte and Tailwind.
It uses MongoDB to store the different shortened links. 

[Preview Service](#preview-service) 
•
[Installation](#installation) 

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

   miniLink requires a few environment variables to work. This includes the database connection string, and (optionally) cloudflare tunnel token.

4. **Run the project**
   Lastly, you can run `docker compose up` to start the project. The first time it will take a minute or so to build.


Note: miniLink uses Cloudflared to expose the app - this isn't required, but it was the easiest way for me to deploy it.


If you'd like to, feel free to check out some of my other projects. <3
