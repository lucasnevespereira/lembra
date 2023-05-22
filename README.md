# Lembra CLI

Lembra is a command-line reminder application that allows you to create and manage reminders for various tasks and
events.

## Installation

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
Run the application:

```bash
make run title="Hello Again" message="you have another meeting" sound="default" time="20h10"
```

## Usage
The application supports the following commands:

- `create`: Create a new reminder.
- `update`: Update an existing reminder.
- `delete`: Delete a reminder by its ID.
- `list`: List all existing reminders.

For detailed information about each command and its options, use the `--help` flag. 
For example:
```bash
./bin/lembra create --help
```

## Configuration
The application uses a SQLite database to store reminders. The database file (`reminders.db`) will be created automatically when you run the application.


## Contributing
Contributions to Lembra are welcome! If you find a bug or have a suggestion for improvement, please open an issue or submit a pull request.


## License
This project is licensed under the MIT License.