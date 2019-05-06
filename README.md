# Try Out VS Cloud Environments: Go

This is a sample project that lets you try out VS Cloud Environments in a few easy steps.
   
## Things to try

This sample has been cloned into your VS Cloud Environment. You're able to work with it like you would any local code.

Some things to try:

1. **Install language extensions**
  - Open `server.go`
  - Install the recommended Go extension
  - View the list of `Cloud Environment - Installed` extensions
    > Note: The Go extension has been installed in the remote environment
  - Reload VS Code
    > Note: VS Code automatically reconnects to your cloud environment

2. **Edit code:**
   - Open `server.go`
   - Try adding some code and check out the IntelliSense

3. **Terminal:** 
  - Press ctrl+shift+[backtick] to open the terminal
  - From the terminal run `go version`
    > Note: go is not installed on the local machine, yet you're able to use it! 
  - Type other Linux commands (`uname`, `ls`, etc.) to interact with the underlying environment
    > Note: The local machine is Windows, yet you're able to issue Linux commands! 

5. **Build, Run, and Debug:**
  - Open `server.go`
  - Change the message to "Hello {your name} from Go!"
  - Add a breakpoint (e.g. on line 22)
  - Press F5 to launch the app in the container and install all tools at the prompt
  - Once the breakpoint is hit, try hovering over variables, examining locals, and more.
  - Continue, then open a local browser and go to `http://localhost:9000` and note you can't connect to the server

6. **Forward port:**
  - Stop debugging and remove the breakpoint.
  - Open `server.go`
  - Change the server port to 5000. (`portNumber := "5000"`)
  - Press F5< to launch the app in the container.
  - Run the **Cloud Environments: Forward Port from Environment...** command
    - Enter port 5000
    - Name it anything you'd like
    - View the forwarded ports in the Environment Details panel
  - Open http://localhost:5000/ and show the booth attendant the running site

7. **Finish**
  - Complete steps 5-7 on the Challenge Worksheet