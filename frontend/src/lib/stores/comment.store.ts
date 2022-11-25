import { writable, readable } from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";

type State = {
  comments: Array<CommentModel>;
  error?: string;
};

export const commentStore = writable<State>({
  comments: []
});

export const connect = () => {
  const socket = new WebSocket("ws://localhost:8080/chat/connect");
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


export const timeline = commentStore.subscribe;





// const storeMessage = (message:string) => {
//     commentStore.set(`You: ${message}`);
// };

// export default {
//   subscribe: commentStore.subscribe,
//   storeMessage,
//   sendMessage,
// };