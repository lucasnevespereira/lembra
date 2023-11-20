# Lembra CLI

Lembra is a command-line reminder application that allows you to create and manage reminders for various tasks and
events.

<p>
    <a href="https://github.com/lucasnevespereira/lembra/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/lucasnevespereira/lembra.svg?logo=github&style=flat-square"></a>
    <a href="https://github.com/lucasnevespereira/lembra/actions/workflows/release.yml"><img alt="GitHub release" src="https://github.com/lucasnevespereira/lembra/actions/workflows/release.yml/badge.svg"></a> 
</p>

## Installation

### Install with Homebrew

```
brew tap lucasnevespereira/tools
```

```
brew install lembra
```

or
```
brew install lucasnevespereira/tools/lembra
```

Update lembra with

```shell
brew upgrade lucasnevespereira/tools/lembra

```

### Install manually
Clone the repository:

```bash
git clone https://github.com/lucasnevespereira/lembra.git
```

Navigate to the project directory:

```bash
cd lembra
```

Build the application:

```bash
make build
```

## Usage

The application supports the following commands:

- `create`: Create a new reminder.
- `update`: Update an existing reminder.
- `delete`: Delete a reminder by its ID.
- `list`: List all existing reminders.
- `listen`: Start the reminder listener daemon.
- `stop`: Stop the reminder listener daemon.
- `logs`: Read the logs of the reminder listener daemon process.

For detailed information about each command and its options, use the `--help` flag.

### Create a Reminder

To create a new reminder, use the create command. Specify the title, message, and time for the reminder.

```bash
lembra create --title "Meeting" --message "You have a meeting at 2 PM" --time "14:00"
```

### Update a Reminder

To update an existing reminder, use the update command. Specify the ID of the reminder and provide the updated
information.

```bash
lembra update --id 123 --title "Updated Meeting" --message "The meeting time has changed to 3 PM"
```

### Delete a Reminder

To delete a reminder, use the delete command. Specify the ID of the reminder to be deleted.

```bash
lembra delete --id 123
```

### List Reminders

To list all existing reminders, use the list command.

```bash
lembra list
```

### Start Reminder Listener Daemon

To start the reminder listener daemon, use the listen command. The daemon will continuously check for reminders and
display notifications when the time matches.

```bash
lembra listen
```

### Stop Reminder Listener Daemon

To stop the reminder listener daemon, use the stop command.

```bash
lembra stop
```

### Read Daemon Logs

To read the logs of the reminder listener daemon process, use the logs command.

```bash
lembra logs
```

## Configuration

The application uses a SQLite database to store reminders. The database file (`reminders.db`) will be created
automatically when you run the application.

## Contributing

Contributions to Lembra are welcome! If you find a bug or have a suggestion for improvement, please open an issue or
submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).