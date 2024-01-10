# Shai - AI Terminal Assistant

Shai (pronounced `shy`) is a terminal-based ChatGPT-powered assistant written in Go. 
Standing for ShellAI, it enables users to e employ ChatGPT to generate shell commands 
based on user free-text input. 


## Demo

![Shai Demo](https://github.com/vonhraban/shai/blob/master/demo.gif)

## Configuration

### Environment Variables

- **SHAI_OPENAI_API_KEY**: Set this variable to your ChatGPT API key.
- **SHAI_DEBUG**: Set to any value for debug mode. When in Debug Mode, 
raw ChatGPT responses are output for debugging purposes. Set to 0 to explicitly disable
- **DUMMY_API**: Set to any value to bypass querying the actual API and receive a dummy response. 
Useful for saving tokens during manual testing or development. Set to 0 to explicitly disable

Environment variables can be set using either the `export` shell built-in or a `.env` file. 
To set variables via export, use:

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

## Usage

To use Shai, run the following command in your terminal:

```bash
shai <request string>
```

or for interactive query input 

```bash
shai
```

**Note:** The ChatGPT API returns pure commands, so it's recommended to avoid overly ambitious queries.

## Setup Wizard

**Note:** the setup wizard is currently a placeholder and not fully implemented.

A setup wizard is there so simplify the configuration process. To invoke the setup wizard, run:

```bash
shai setup
```

## Known issues

At the moment, there is no human-friendly error handling of the ChatGPT API non-200 responses or
connectivity issues. Be aware of potential issues, and comprehensive error handling is planned for future updates. 

Currently, Shai supports ChatGPT 3.5 only, but adding support for different versions as a configuration 
option is planned in the nearest future.

The generated commands are tailored for zsh. Support for other shells is planned.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/vonhraban/shai.git
```

2. Navigate to the project directory:

```bash
cd shai
```

3. Build the project:

```bash
go build .
```

4. Set up the necessary environment variables or create a `.env` file.

5. Run Shai:

```bash
./shai what is my current external ip address
```

6. Optionally, symlink the binary to `/usr/local/bin` (or other folder in $PATH)
   
```bash
ln -s $(pwd)/shai /usr/local/bin/shai
```

## Running Unit Tests

To run unit tests, navigate to the project directory and execute:

```bash
go test ./...
```