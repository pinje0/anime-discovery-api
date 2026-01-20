# Anime Discovery API

Go backend for anime discovery application using Jikan API.

## Features

- Fetches top anime list from Jikan API
- Retrieves anime details by ID
- Caching to respect Jikan rate limits
- Clean, simplified JSON responses
- Configurable via environment variables

## Getting Started

### Prerequisites

- Go 1.20+
- Internet connection (for Jikan API)

### Installation

1. Clone the repository
2. Navigate to this directory:
   ```bash
   cd anime-discovery-api
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Configuration

Create a `.env` file based on `.env.example`:

```env
SERVER_PORT=8080
JIKAN_API_URL=https://api.jikan.moe/v4
CACHE_TIMEOUT_MINUTES=10
```

### Running

```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /health` | Health check |
| `GET /anime/top?page=1` | Get top anime list |
| `GET /anime/:id` | Get anime details by ID |

### Response Examples

**Health Check:**
```json
{"status":"ok"}
```

**Top Anime:**
```json
{
  "data": [
    {
      "id": 59978,
      "title": "Sousou no Frieren 2nd Season",
      "title_english": "Frieren: Beyond Journey's End Season 2",
      "synopsis": "Second season of Sousou no Frieren.",
      "image_url": "https://cdn.myanimelist.net/images/anime/1921/154528.jpg",
      "score": 9.33,
      "episodes": 10,
      "rating": "PG-13 - Teens 13 or older",
      "year": 2026,
      "status": "Currently Airing",
      "genre": "Adventure"
    }
  ],
  "page": 1
}
```

**Anime by ID:**
```json
{
  "id": 1,
  "title": "Cowboy Bebop",
  "title_english": "Cowboy Bebop",
  "synopsis": "Crime is timeless...",
  "image_url": "https://cdn.myanimelist.net/images/anime/1/20450.jpg",
  "score": 8.75,
  "episodes": 26,
  "rating": "R - 17+ (violence & profanity)",
  "year": 1998,
  "status": "Finished Airing",
  "genre": "Action"
}
```

## Project Structure

```
├── .env              # Environment variables (not committed)
├── .env.example      # Environment template
├── .gitignore
├── README.md
├── config/
│   └── config.go     # Configuration loader
├── go.mod
├── go.sum
├── handlers/
│   └── anime.go      # HTTP handlers
├── main.go           # Entry point
├── models/
│   └── anime.go      # Data types
├── routes/
│   └── routes.go     # Route definitions
└── services/
    └── jikan.go      # Jikan API integration + caching
```

## Technologies

- [Gin](https://gin-gonic.com/) - Web framework
- [Jikan API](https://jikan.moe/) - Unofficial MyAnimeList API

## License

MIT
