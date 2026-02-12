# RSS Aggregator API - Setup & Testing Guide

## Database Setup

### Prerequisites
- PostgreSQL installed and running
- Port 5432 available

### Create Database
```bash
# Connect to PostgreSQL
psql -U postgres

# Create the database
CREATE DATABASE rss_aggregator;

# Exit psql
\q
```

Alternative: Use the connection string in `.env`
```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/rss_aggregator?sslmode=disable
```

## Running the Application

```bash
# Install dependencies
go get -u gorm.io/gorm gorm.io/driver/postgres

# Run the application
go run main.go json.go models.go db.go

# Server will start on http://localhost:8080
```

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### Feeds Endpoints

#### Create a Feed (POST)
```bash
curl -X POST http://localhost:8080/v1/feeds \
  -H "Content-Type: application/json" \
  -d '{
    "name": "TechCrunch",
    "url": "https://techcrunch.com/feed/",
    "user_id": 1
  }'
```

#### Get All Feeds (GET)
```bash
curl http://localhost:8080/v1/feeds
```

#### Get Single Feed (GET)
```bash
curl http://localhost:8080/v1/feeds/1
```

#### Update Feed (PUT)
```bash
curl -X PUT http://localhost:8080/v1/feeds/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "TechCrunch Updated",
    "url": "https://techcrunch.com/feed/"
  }'
```

#### Delete Feed (DELETE)
```bash
curl -X DELETE http://localhost:8080/v1/feeds/1
```

### Articles Endpoints

#### Get All Articles (GET)
```bash
curl http://localhost:8080/v1/articles
```

#### Get Articles by Feed (GET)
```bash
curl http://localhost:8080/v1/feeds/1/articles
```

## Using Postman

1. Import these requests into Postman
2. Create environment variables:
   - `base_url`: http://localhost:8080

3. Create requests for each endpoint above

## Using REST Client Extension (VS Code)

Create a `requests.http` file:
```http
### Health Check
GET http://localhost:8080/health

### Create Feed
POST http://localhost:8080/v1/feeds
Content-Type: application/json

{
  "name": "TechCrunch",
  "url": "https://techcrunch.com/feed/",
  "user_id": 1
}

### Get All Feeds
GET http://localhost:8080/v1/feeds

### Get Single Feed
GET http://localhost:8080/v1/feeds/1

### Update Feed
PUT http://localhost:8080/v1/feeds/1
Content-Type: application/json

{
  "name": "TechCrunch Updated"
}

### Delete Feed
DELETE http://localhost:8080/v1/feeds/1

### Get All Articles
GET http://localhost:8080/v1/articles

### Get Feed Articles
GET http://localhost:8080/v1/feeds/1/articles
```

## Database Tables

The application automatically creates these tables:

### users
- id (Primary Key)
- name
- api_key (Unique)
- created_at
- updated_at

### feeds
- id (Primary Key)
- name
- url
- user_id (Foreign Key)
- created_at
- updated_at

### articles
- id (Primary Key)
- title
- description
- url
- published_at
- feed_id (Foreign Key)
- created_at

## Troubleshooting

### Database Connection Error
```
error: database "rss_aggregator" does not exist
```
**Solution**: Create the database first using the SQL commands above

### Port Already in Use
Change the PORT in `.env`:
```
PORT=3000
```

### GORM Migration Issues
Clear the database and restart:
```sql
DROP DATABASE rss_aggregator;
CREATE DATABASE rss_aggregator;
```

## Next Steps

1. Add RSS feed parser to fetch articles automatically
2. Implement background job to periodically update feeds
3. Add authentication with API keys
4. Add pagination to articles endpoint
5. Write unit tests

