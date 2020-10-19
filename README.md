# Git Lit 💡

A proof-of-concept project to interact with Tuya smart-lights from Github Actions.

## Behind the scenes

The Tuya smart-light can be controlled remotely using Tuya's REST API and supplementary GoLang library.

Based on different Github events, workflows are triggered which run the `main.go` and control the smart-lights

 ## Running locally

```bash
export P4L_DEVICE_ID=00000000000000 
export P4L_ACCESS_ID=00000000000000
export P4L_ACCESS_SECRET=00000000000000     

# Turn on the lights, wait for 10 seconds and then turn them off
go run main.go

# Toggles the device
go run main.go -toggle=true
```