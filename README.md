# Whispercleaner

Whispercleaner is a command-line tool designed to clean transcription output from the Whisper application. It removes timestamp information from each line, producing clean, readable text.

## Features

- Removes timestamp information from Whisper transcription output
- Processes input from stdin and writes to stdout, allowing easy integration into Unix-like pipelines
- Cross-platform support (Linux, Windows, macOS)
- Verbose logging option for detailed processing information

## Installation

### Prerequisites

- Go 1.16 or higher
- [just](https://github.com/casey/just) command runner (optional, but recommended for easier build management)

### Building from source

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/whispercleaner.git
   cd whispercleaner
   ```

2. Build the application:
   
   If you have `just` installed:
   ```
   just build
   ```
   
   Otherwise, use Go directly:
   ```
   go build -o whispercleaner
   ```

## Usage

Basic usage:
```
cat input_transcription.txt | ./whispercleaner > cleaned_output.txt
```

With verbose logging:
```
cat input_transcription.txt | ./whispercleaner --verbose > cleaned_output.txt 2>processing.log
```

For help:
```
./whispercleaner --help
```

## Example

### Input (Whisper transcription output)

```
[00:00.000 --> 00:08.520]  Welcome everyone to our quarterly meeting. Today, we'll be discussing the progress of Project Phoenix and our plans for the upcoming fiscal year.
[00:08.520 --> 00:15.760]  First, I'd like to call on Sarah from the development team to give us an update on the new features we've been working on.
[00:15.760 --> 00:24.040]  Thanks, John. We've made significant progress on the user interface redesign. Our latest user testing shows a 30% improvement in navigation efficiency.
[00:24.040 --> 00:32.280]  That's great news, Sarah. Mark, can you fill us in on how this aligns with our marketing strategy?
[00:32.280 --> 00:41.600]  Absolutely. With these improvements, we're in a strong position to target the enterprise market. We're planning a major campaign launch to coincide with the next release.
[00:41.600 --> 00:50.920]  Excellent. Now, let's move on to discussing our financial projections for the next quarter. Emily, would you like to take the floor?
```

### Output (Cleaned transcription)

```
Welcome everyone to our quarterly meeting. Today, we'll be discussing the progress of Project Phoenix and our plans for the upcoming fiscal year.
First, I'd like to call on Sarah from the development team to give us an update on the new features we've been working on.
Thanks, John. We've made significant progress on the user interface redesign. Our latest user testing shows a 30% improvement in navigation efficiency.
That's great news, Sarah. Mark, can you fill us in on how this aligns with our marketing strategy?
Absolutely. With these improvements, we're in a strong position to target the enterprise market. We're planning a major campaign launch to coincide with the next release.
Excellent. Now, let's move on to discussing our financial projections for the next quarter. Emily, would you like to take the floor?
```

## Development

This project uses `just` for managing build tasks. To see available commands:

```
just
```

Common tasks:
- `just build`: Build the application
- `just test`: Run tests
- `just build-all`: Build for all supported platforms
- `just clean`: Remove built executables

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
