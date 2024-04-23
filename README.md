# dgo-commandler

The `dgo-commandler` library simplifies the creation and management of Discord slash commands using the DiscordGo package. It offers an easy-to-use framework that handles command registration, middleware integration, logging, and interaction with Discord API features.

## Features

- **Easy Registration**: Effortlessly register slash commands with simple function calls.
- **Middleware Support**: Incorporate middleware for preprocessing, logging, or any custom functionality.
- **Logging**: Integrated structured logging helps track command usage and errors efficiently.
- **Utility Functions**: Includes helper functions to manage user interactions and state within Discord environments.

## Getting Started

### Prerequisites

Before installing `dgo-commandler`, make sure you have the following:

- Go version 1.22.2 or higher
- DiscordGo library

### Installation

To install `dgo-commandler`, run the following command in your terminal:

```bash
go get github.com/CreativeUnicorns/dgo-commandler
```

This command retrieves the library and installs it in your Go workspace.

### Quick Start Guide

Here's how you can quickly set up a Discord bot using `dgo-commandler`:

1. **Create a Discord Bot** on the Discord developer portal and obtain your token.
2. **Import the library** into your Go project:

```go
import (
    "github.com/bwmarrin/discordgo"
    "github.com/CreativeUnicorns/dgo-commandler"
)
```

3. **Initialize the Discord session** and register commands:

```go
func main() {
    dg, err := discordgo.New("Bot " + "yourBotToken")
    if err != nil {
        log.Fatal("error creating Discord session,", err)
        return
    }

    // Open the Discord session.
    err = dg.Open()
    if err != nil {
        log.Fatal("error opening connection,", err)
        return
    }

    // Initialize command handler
    commandler.AddAndRegisterInteractionCommands(dg)

    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    select {}
}
```

**_Alternatively_**
You may split the handler registration out from the command registration.

```go
func main() {
    dg, err := discordgo.New("Bot " + "yourBotToken")
    if err != nil {
        log.Fatal("error creating Discord session,", err)
        return
    }

    // Initialize command handlers
    commandler.AddInteractionCommandHandlers(dg)

    // Open the Discord session.
    err = dg.Open()
    if err != nil {
        log.Fatal("error opening connection,", err)
        return
    }

    // Register commands
    commandler.RegisterInteractionCommands(dg)

    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    select {}
}
```

## Contributing

Contributions to `dgo-commandler` are welcome! For more information on how to contribute, please refer to the [contributing guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
