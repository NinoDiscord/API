# Contributing to Nino's API
> 🎀 **Thanks for thinking about contributing to Nino's API and making it better! This is a guide on how to contribute, and get on the contributors tab in the [about](https://nino.sh/about) page. :3**

## Prerequisites
Before you can start working on the API, I recommend some tools before getting started:

- [**Redis** v5+](https://redis.io) **~** This is the cache database for Nino, but we use a seperate database index for sessions.
- [**PostgreSQL** 9+](https://postgresql.org) **~** The main database for Nino, you must need to run the [bot](https://github.com/NinoDiscord/Nino) to create the database!
- [**Go** v1.16+](https://golang.org/) **~** Go 1.16+ is recommended to start contributing, though a version with Go Modules is supported also~
- [**GoLand**](https://www.jetbrains.com/go/) **~** Go IDE to contribute! This is optional if you prefer Visual Studio Code, it's a perfect Go IDE!

## 🎀 Contributing
1. [Fork](https://github.com/NinoDiscord/API/fork) the repository!
2. Clone the repository to your local machine (``git clone https://github.com/$USERNAME/API``)

###### NOTE: omit `$USERNAME` with your username on GitHub

3. Open the project in your IDE or run `go mod init` to initialize the project.
4. Run `go get` to fetch all dependencies
5. Code your heart out! If you plan to update the GraphQL schema, it is required to add your own resolvers in `graphql/resolvers/<type>.go`
6. [Submit](https://github.com/NinoDiscord/API/compare) your pull request.
