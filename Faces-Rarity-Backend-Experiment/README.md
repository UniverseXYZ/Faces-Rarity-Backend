# Faces Rarity Backend
## Summary
- This backend is based on the Polymorphs V2 rarity backend. Some functionalities are cut-off.
- Also adds some concurrency and database optimizations
## Deployment
- Build the docker image
  - ```bash
    docker build --platform linux/amd64  .
    ```
- run the image as a container
## Notes:
  - This backend only queries the contract ~ 15 seconds for specific events, and updates the Mongo DB accordingly
  - The mongo collections themselves are queried using the following gcloud function: `https://github.com/UniverseXYZ/Polymorph-Rarity-Cloud`