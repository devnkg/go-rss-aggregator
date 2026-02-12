#!/bin/bash

# RSS Aggregator - Quick Start Script

echo "🚀 RSS Aggregator Setup"
echo "========================"
echo ""

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL is not installed. Please install PostgreSQL first."
    exit 1
fi

echo "✅ PostgreSQL found"
echo ""

# Check if database exists
echo "📦 Setting up database..."
if psql -U postgres -lqt | cut -d \| -f 1 | grep -qw rss_aggregator; then
    echo "⚠️  Database 'rss_aggregator' already exists"
else
    echo "Creating database 'rss_aggregator'..."
    psql -U postgres -c "CREATE DATABASE rss_aggregator;" 2>/dev/null
    if [ $? -eq 0 ]; then
        echo "✅ Database created successfully"
    else
        echo "❌ Failed to create database"
        exit 1
    fi
fi

echo ""
echo "📚 Installing Go dependencies..."
go mod download
go mod tidy

echo ""
echo "✅ Setup complete!"
echo ""
echo "🎯 To start the server:"
echo "   go run main.go json.go models.go db.go"
echo ""
echo "📖 For testing guide, see: TESTING.md"
echo ""
