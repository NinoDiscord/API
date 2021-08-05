<div align="center">
    <h3>Nino API</h3>
    <blockquote><strong>Source code for Nino's API (located at <a href="https://api.nino.sh">api.nino.sh</a>), made with Go and GraphQL 💜</strong></blockquote>
</div>

<div align='center'>
  <img alt='api.nino.sh Workflow' src='https://img.shields.io/github/workflow/status/NinoDiscord/API/Deploy%20to%20Kubernetes/master?style=flat-square' />
</div>

<hr />

## Why?
I decided to move Nino's API to a different repository so no overhead of requests are correlated to the bot.
Since Node.js is single-threaded, I decided to not run an API and the bot at the same time and having an API implemented
in the bot *could be* "optional" but, I decided not to, which was the original plan on the rewrite (v0 -> v1).

## Support
Need support related to Nino or any microservices under the organization? Join in the **Noelware** Discord server in #support under the **Nino** category:

[![discord embed owo](https://discord.com/api/v8/guilds/824066105102303232/widget.png?style=banner3)](https://discord.gg/ATmjFH9kMH)

## Self-hosting
Before we run this API, I recommend not to. Since builds are most likely unstable, it's not really recommended to Nino and private instances
don't need this type of functionality...

### Prerequisites
Before running your instance of **API**, you will need the following tools before starting:

- [**PostgreSQL** v10+](https://postgresql.org) - Used for retrieving data from the bot -> **API**. 
- [**Git** v2.31+](https://git-scm.com/) - Useful for retrieving new updates with **API**.
- [**Go** v1.16+](https://go.dev) - Language build-tools to run **API**

## Building
```sh
$ git clone https://github.com/NinoDiscord/API.git && cd API
$ go get && go build
$ ./api # or run api.exe in Command Prompt
```

## License
**API** is released under the **GPL-3.0** License.
