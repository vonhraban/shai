# Shai - Terminal ChatGPT Assistant

Shai is a terminal-based ChatGPT-powered assistant written in Go. It enables users to interact with ChatGPT and generate commands based on user input. The project supports three configuration options through environment variables, allowing users to customize their experience.

## Demo

![Shai Demo](https://giphy.com/embed/c4ANrJvHDlhDEmIPKN)

## Configuration

### Environment Variables

- **SHAI_OPENAI_API_KEY**: Set this variable to your ChatGPT API key.
- **SHAI_DEBUG**: Set to 1 for debug mode (print raw ChatGPT responses), and 0 to disable debug mode.
- **DUMMY_API**: Set to any value to bypass querying the actual API and receive a dummy response. Useful for saving tokens during manual testing or development.

Environment variables can be set using either the `export` shell built-in or a `.env` file. To set variables via export, use:

```bash
export SHAI_OPENAI_API_KEY=your_api_key
export SHAI_DEBUG=1
export DUMMY_API=1
```

Alternatively, create a `.env` file in the project root and add the following:

```dotenv
SHAI_OPENAI_API_KEY=your_api_key
SHAI_DEBUG=1
DUMMY_API=1
```

**Note:** At the moment, there is no human-friendly error handling with the ChatGPT API. Be aware of potential issues, and comprehensive error handling is planned for future updates. Currently, Shai supports ChatGPT 3.5 only, but adding support for different versions as a configuration option is planned in the nearest future.

## Usage

To use Shai, run the following command in your terminal:

```bash
shai <request string>
```

**Note:** The ChatGPT API returns commands, so it's recommended to avoid overly ambitious queries.

## Setup Wizard

A setup wizard is under development to simplify the configuration process. To invoke the setup wizard, run:

```bash
shai setup
```

Please note that the setup wizard is currently a placeholder and not fully implemented.

## Running Unit Tests

To run unit tests, navigate to the project directory and execute:

```bash
go test ./...
```

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your_username/shai.git
```

2. Navigate to the project directory:

```bash
cd shai
```

3. Build the project:

```bash
go build
```

4. Set up the necessary environment variables or create a `.env` file.

5. Run Shai:

```bash
./shai "Hello, Shai! What can you do for me?"
```

Feel free to explore and contribute to the project. Happy chatting with Shai!