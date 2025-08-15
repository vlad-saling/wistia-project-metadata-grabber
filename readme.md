# Wistia Project Metadata Grabber

This tool fetches metadata for Wistia projects using the Wistia API.

## How it works

- Retrieves metadata for all project ID listed in `input.txt`.
- Filter out the name and pairs it with ID
- Saves output to `output.csv`.

## Prerequisites

- Go 1.18+
- Wistia API token

## Usage

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/wistia-project-metadata-grabber.git
    cd wistia-project-metadata-grabber
    ```

2. **Provide your Wistia API token:**
    
    - Update `main.go` with token `bearerToken`

3. **Run the tool:**
    `go run main.go`


## License

MIT
