#!/bin/bash

# Run tests and generate coverage profile
go test -coverprofile=coverage.out ./...

# Check if coverage file was created
if [ ! -f coverage.out ]; then
  echo "❌ Coverage file not found. Tests may have failed."
  exit 1
fi

# Show coverage summary
echo "📊 Coverage Summary:"
go tool cover -func=coverage.out

# Optional: Show total coverage only
TOTAL=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo "✅ Total Coverage: $TOTAL"

# Optional: Open HTML report in browser
read -p "Open detailed HTML report? (y/n): " OPEN
if [ "$OPEN" = "y" ]; then
  go tool cover -html=coverage.out -o coverage.html
  open coverage.html
fi
