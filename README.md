# Eurovision Party
A privately hosted eurovision voting and chat. Easily set up a remote eurovision party with your friends using this server. It will provide a live chat, as well as voting and rank the acts and display charts. It will be simple enough for a person with no technical knowledge to set up and host.

# Running locally
This project is still a work in progress, but the back end is currently in a working state. It requres a mysql server running on your machine. Once that is set up, run from the root directory:```go run .```

The front end uses SvelteKit, and still requires a lot of work, but some views are currently working. This can be run by navigating to the `frontend` directory and running ```npm run dev```

You should now be able to log in using the signin link in the go logs.
