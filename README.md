# Go Banking Application

This project is a comprehensive banking application developed in Go, following the principles of Hexagonal Architecture (also known as Ports and Adapters pattern). This architectural pattern allows the separation of concerns, making the codebase easier to maintain, more adaptable to changes, and highly testable.

In this application, the core business logic is isolated from external concerns through well-defined interfaces. These interfaces represent the 'ports', and the implementations that adapt the application to external factors (like databases or web services) are the 'adapters'. This means that the business logic doesn't know anything about who is invoking it or what will be done with the results.

The application provides modules for account management, customer relations, and transaction processing. Each module is designed to be independent and can be developed, tested, deployed, and scaled separately.

Whether you're a developer seeking to understand how banking systems work, or a student studying software architecture, this project provides a solid starting point. It's not just about the code - it's about learning and understanding the design principles and architecture that underpin effective banking software.

## Installation

Clone the repository and navigate into the directory:

```bash
git clone https://github.com/username/project.git
cd project
```

Install the dependencies:

```bash
go mod download
```

## Usage

Run the main file:

```bash
go run main.go
```

## Structure

The project is structured into several packages:

- `app`: Contains the application logic, including handlers and middleware.
- `domain`: Contains the domain models and repositories.
- `dto`: Contains data transfer objects.
- `mocks`: Contains mock objects for testing.
- `service`: Contains services that implement business logic.
- `resources`: Contains resources such as SQL scripts.

Authentication logic is separated into its own repository, and error handling and logging are handled by a shared library.

## Related Repositories

- [Authentication Service](https://github.com/Nishith-Savla/golang-banking-auth): This repository contains the authentication logic for the banking application.
- [Shared Library](https://github.com/Nishith-Savla/golang-banking-lib): This repository contains shared components such as error handling and logging.

## Testing

Run the tests with:

```bash
go test ./...
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
