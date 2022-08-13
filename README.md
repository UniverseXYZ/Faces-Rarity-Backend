# Faces Rarity Backend
## Summary
- This backend is based on the Polymorphs' V2 rarity backend. Some functionalities are cut-off.
- Extended to process events on Polygon chain
- Also adds some concurrency and database optimizations
## Deployment
- The process runs in a docker container on GCE
  - Build the docker image
    - ```bash
      docker build --platform linux/amd64  .
      ```
- run the image as a container
- Since some double-inserting problems were occurring, MongoDB's collections should be configured with unique primary key `tokenID`
## Notes:
  - This backend only queries the contract ~ 15 seconds for specific events, and updates the Mongo DB accordingly
  - The mongo collections themselves are queried using the following gcloud function: `https://github.com/UniverseXYZ/Polymorph-Rarity-Cloud`