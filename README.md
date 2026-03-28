# Katara

An anime tracker backend API built with Go, featuring a GraphQL interface, session-based authentication, and automatic anime data synchronization from AniList.

## Stack

- **Go** + **Gin** — language and web framework
- **GraphQL** — API layer via [gqlgen](https://github.com/99designs/gqlgen)
- **MongoDB** — primary database (anime catalog, users, lists)
- **Redis** — session storage
- **AniList API** — external anime data source
- **Hexagonal Architecture** — clean separation of domain, adapters, and infrastructure
- **Air** — hot reload for local development

## Features

- User registration and login with session-based auth (httpOnly cookie)
- Protected routes via auth middleware
- Full anime catalog synced from AniList API (~20,000+ entries)
- Personal anime list management (add, update, delete entries)
- Filter and sort anime by status, format, season, year, genre, popularity, and more
- Background sync worker that refreshes data every 24 hours
- CORS configured for local frontend development

## Project Structure

```
cmd/               — application entrypoint
graph/             — GraphQL resolvers, mappers, generated code
  model/           — generated GraphQL models
internal/
  adapters/
    anilist/       — AniList API client
    anime_repo/    — MongoDB anime repository
    user_repo/     — MongoDB user repository
    list_repo/     — MongoDB list repository
    redis/         — Redis session repository
    documents/     — MongoDB document types and mappers
    schema/        — GraphQL schema files (.graphqls)
  domain/
    anime/         — anime entity, repository interface, service
    user/          — user entity, repository interface, service
    list/          — list entity, repository interface, service
    session/       — session repository interface
  config/          — environment config loader
  database/        — MongoDB and Redis connection setup
  worker/          — background AniList sync worker
middleware/        — auth middleware (session validation)
```

## GraphQL API

### Queries

| Query | Auth | Description |
|-------|------|-------------|
| `getAnimeByID(aniListID)` | ✅ | Get anime by AniList ID |
| `getAnimeWithFilters(animeFilter)` | ✅ | Filter and sort anime catalog |
| `me` | ✅ | Get current authenticated user |
| `getAllLists` | ✅ | Get user's anime list |
| `getListByStatus(status)` | ✅ | Get list filtered by status |

### Mutations

| Mutation | Auth | Description |
|----------|------|-------------|
| `register(registerInput)` | — | Create a new account |
| `login(loginInput)` | — | Login and receive session cookie |
| `logout` | ✅ | Invalidate session |
| `updateEmail` | ✅ | Update email address |
| `updatePassword` | ✅ | Change password |
| `updateAvatar` | ✅ | Update avatar URLs |
| `updateBannerImage` | ✅ | Update banner image |
| `updateAbout` | ✅ | Update bio/about text |
| `addList(input)` | ✅ | Add anime to personal list |
| `updateList(aniListID, input)` | ✅ | Update list entry |
| `deleteList(aniListID)` | ✅ | Remove anime from list |

## Getting Started

### Prerequisites

- Go 1.21+
- MongoDB
- Redis
- [Air](https://github.com/air-verse/air) (optional, for hot reload)

### Setup

1. Clone the repository and install dependencies:

```bash
git clone https://github.com/terrakuta/Katara.git
cd Katara
go mod tidy
```

2. Copy the example environment file and fill in your values:

```bash
cp .env.example .env
```

3. Start the server:

```bash
# with hot reload
air

# or directly
go run ./cmd
```

4. Open GraphQL Playground at `http://localhost:8080/playground`

## Environment Variables

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=katara
REDIS_ADDR=localhost:6379
PORT=8080
SYNC_ENABLED=false
```

Set `SYNC_ENABLED=true` to start the background AniList sync worker on startup. The worker fetches all anime from AniList and stores them in MongoDB, then repeats every 24 hours.

## Authentication

Authentication uses session-based cookies. On successful login, a `session_id` is set as an `httpOnly` cookie. Subsequent requests to protected endpoints require this cookie to be present. Sessions are stored in Redis with a 1-hour TTL.

To access protected endpoints from the playground, login first via the `login` mutation, then include the cookie in subsequent requests.
