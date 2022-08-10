# Faces Rarity Backend
## Summary
- This backend is based on the Polymorphs V2 rarity backend. Some functionalities are cut-off.
- Everytime new faces(s) is/are claimed 
- Also adds some concurrency and database optimizations
## Deployment
- Currently, the backend runs as a process in a GCloud Virtual Machine:
  - `go run main.go`
- Notes:
  - When deploying for `production`, take into account all constants in `constants/metadata.go`
  - This backend only queries the contract ~ 15 seconds for specific events, and updates the Mongo DB accordingly
  - The mongo collections themselves are queried using the following gcloud function: `https://github.com/UniverseXYZ/Polymorph-Rarity-Cloud`