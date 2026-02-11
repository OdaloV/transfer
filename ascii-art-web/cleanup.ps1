# cleanup.ps1 - COMPLETE GARBAGE COLLECTION
Write-Host "Running Docker garbage collection..."

# STOP AND REMOVE YOUR CONTAINER (ADD THIS HERE)
docker stop ascii-container 2>$null
docker rm ascii-container 2>$null

# REMOVE YOUR IMAGE (ADD THIS HERE)
docker rmi ascii-art-web 2>$null

# NOW RUN GARBAGE COLLECTION ON ALL UNUSED OBJECTS
docker system prune -f --volumes

Write-Host "Cleanup complete!"
docker system df