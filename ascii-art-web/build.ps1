# build.ps1 - Build image and run container
Write-Host "Building Docker image..."
docker build -t ascii-art-web .
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed!"
    exit 1
}

Write-Host "Cleaning up old container..."
docker stop ascii-container 2>$null
docker rm ascii-container 2>$null

Write-Host "Starting new container..."
docker run -d -p 8080:8080 --name ascii-container ascii-art-web

Write-Host "Container running at http://localhost:8080"
docker ps --filter "name=ascii-container"