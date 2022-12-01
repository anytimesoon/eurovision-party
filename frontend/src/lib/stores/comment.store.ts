import { writable } from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";
import type { TokenModel } from "$lib/models/classes/token.model";
import { tokenStore } from "./token.store";

type State = {
  comments: Array<CommentModel>;
  error?: string;
};

export const commentStore = writable<State>({
  comments: []
});

let token:TokenModel;
tokenStore.subscribe((val) => {
  token = val
})

export const connectToChat = () => {
  const socket = new WebSocket("ws://localhost:8080/restricted/chat/connect/" + token);
  if (!socket) {
    // Store an error in our state.  The function will be
    // called with the current state;  this only adds the
    // error.
    commentStore.update((s: State) => { return {...s, error: "Unable to connect" }});
    return;
  }

  // Connection opened
  socket.addEventListener("open", () => {
    console.log("You're connected. Welcome to the party!!!🎉");
    // TODO: Set up ping/pong, etc.
  });

  // Listen for messages
  socket.addEventListener("message", function (event) {
    const data:CommentModel = JSON.parse(event.data);
    commentStore.update((s: State) => ({ ...s, comments: s.comments.concat(data) }));
  });

  // Send message
  socket.addEventListener("close", (_event: any) => {
    console.log("The connection has been closed. Goodbye!");
  });

  return socket
}
