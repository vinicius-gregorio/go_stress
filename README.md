# Go Stress Testing CLI

A simple and efficient CLI tool for stress-testing URLs. This tool simulates a high volume of requests to a specified URL and generates a detailed report at the end, including metrics such as success rate, duration, and number of requests processed.

## Docker Usage

The stress-testing app is available as a Docker image, making it easy to run without installing Go locally.

### Pull and Run the Docker Image

You can run the image directly from Docker Hub without explicitly pulling it.

```bash
docker run --rm viniboyz/stresstest:latest --url="http://your_target_url"
```

### Parameters

The CLI supports the following parameters:

- `--url` (required): The target URL to stress-test.
- `--concurrency` (optional): The number of concurrent requests. Default is `1`.
- `--requests` (optional): The total number of requests to make. Default is `1`.

> **Note**: If you only pass the `--url` parameter, the app will default to 1 request with a concurrency of 1.

### Examples

1. **Basic Example (Single Request):**
   ```bash
   docker run --rm viniboyz/stresstest:latest --url="http://your_target_url"
   ```
   This will make a single request to `http://your_target_url` with 1 concurrency.

2. **Custom Concurrency and Request Count:**
   ```bash
   docker run --rm viniboyz/stresstest:latest --url="http://your_target_url" --concurrency=10 --requests=500
   ```
   This will make 500 requests to `http://your_target_url`, with up to 10 concurrent requests at a time.

3. **Testing a Local API (Using `host.docker.internal`):**
   When testing a local service, use `host.docker.internal` instead of `localhost` to direct Docker to your host machine:
   ```bash
   docker run --rm viniboyz/stresstest:latest --url="http://host.docker.internal:8080" --concurrency=5 --requests=1000
   ```

   > **Warning**: When targeting a service on your host machine, `localhost` won’t work from within Docker. Use `host.docker.internal` instead.

## Running Locally with Go

If you have Go installed, you can run the application directly using `main.go` or by setting it up with Go modules.

### 1. Running with `main.go`

To run the app directly from the source code, navigate to the project directory and use:

```bash
go run main.go --url="http://your_target_url" --concurrency=5 --requests=1000
```

This will compile and execute the code, allowing you to use all available parameters (`--url`, `--concurrency`, and `--requests`).

### 2. Running with Go Modules

If your project is set up with Go modules, you can install the dependencies and run it as follows:

1. Install the dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go --url="http://your_target_url" --concurrency=5 --requests=1000
   ```

This approach will automatically resolve and download any necessary dependencies listed in your `go.mod` file.

> **Note**: You can also build the application as an executable by running `go build -o stresstest main.go` and then executing `./stresstest` with the desired parameters.

### Sample Report

After the stress test completes, the app will display a report in the following format:

```plaintext
+--------------------------------------------------------+
|  Stress Testing URL: http://host.docker.internal:8080  |
+--------------------------------------------------------+
⠏ Processing request 999, status code: 200, duration: 4.078143ms
------------------------------------------
Stress test completed
Number of requests made: 1000
Time taken: 8.253720761s
Success rate: 100.00%
```

The report provides an overview of:
- Total requests made
- Total time taken for the test
- Success rate (percentage of requests with a successful response)

## License

This project is licensed under the MIT License.
