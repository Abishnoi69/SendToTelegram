# Telegram Messaging API

This is a simple Go-based API that allows users to send messages to a Telegram chat by making HTTP requests. The server exposes a single endpoint, `/send`, which accepts a `POST` request with a JSON payload. This payload should contain the `chat_id` and `message` fields to specify the Telegram chat and message content.


## Why?
I created this API for my own use to easily send messages to Telegram channels or groups from various applications. By setting up this server, I can programmatically communicate with any chat using my Telegram bot, making it convenient to send notifications, alerts, or updates.

## Features

- **Send Messages**: Easily send messages to Telegram by making a `POST` request.
- **HTML Formatting**: Messages support HTML formatting.
- **No Web Preview**: Link previews are disabled in messages to prevent clutter.

## Prerequisites

- **Go** : https://golang.org
- A **Telegram bot token** from [BotFather](https://core.telegram.org/bots#botfather)

## Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/AshokShau/SendToTelegram.git
   cd SendToTelegram
   ```

2. **Add your Telegram bot token**:

    Add your Telegram bot token as an environment variable named `TOKEN`.

3. **Run the server**:

   ```bash
   go run main.go
   ```

   The server will start on port `3000` by default.

## Usage

### Endpoint

- **URL**: `/send`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Payload**:

  ```json
  {
    "chat_id": "CHAT_ID",
    "message": "Hello from API!"
  }
  ```

### Example Request

You can use `curl` to test the endpoint:

```bash
curl -X POST http://localhost:3000/send -H "Content-Type: application/json" -d '{"chat_id": "CHAT_ID", "message": "Hello from API!"}'
```

### Example Response

- **Success**: Returns `200 OK` with a message indicating success.
- **Failure**: Returns `400 Bad Request` or `500 Internal Server Error` with error details.

## Error Handling

The API handles and responds with appropriate error messages:
- **200 OK**: If the message was sent successfully.
- **400 Bad Request**: For invalid JSON payloads.
- **404 Not Found**: If the endpoint is not found.
- **405 Method Not Allowed**: If the request method is not `POST`.
- **500 Internal Server Error**: If the message could not be sent to Telegram.

## Security

This example does not include authentication or rate limiting. If you intend to make this publicly accessible, consider implementing authentication (such as API keys or OAuth) and rate limiting to prevent abuse.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

