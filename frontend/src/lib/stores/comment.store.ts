import { writable } from "svelte/store";
import type { CommentModel } from "$lib/models/classes/comment.model";
import type { TokenModel } from "$lib/models/classes/token.model";
import { tokenStore } from "./token.store";
import {chatEP} from "$lib/models/enums/endpoints.enum";

const comments = new Array<CommentModel>()

export const commentStore = writable<CommentModel[]>(comments);

let auth:TokenModel;
tokenStore.subscribe((val) => {
  auth = val
})

export function connectToChat():WebSocket {
  const socket = new WebSocket(chatEP, [auth.token, "chat"]);
  if (!socket) {
    // Store an error in our state.  The function will be
    // called with the current state;  this only adds the
    // error.
    commentStore.update(comments => { return {...comments}});
    return socket;
  }

  // Connection opened
  socket.addEventListener("open", () => {
    console.log("You're connected. Welcome to the party!!!🎉");
    // TODO: Set up ping/pong, etc.
  });

  // Listen for messages
  socket.addEventListener("message", function (event) {
    const data:CommentModel = JSON.parse(event.data);
    commentStore.update(comments => {
      comments.push(data)
      return comments
    } );
  });

  // Send message
  socket.addEventListener("close", (_event: Event) => {
    //TODO gracefully close websocket
    console.log("The connection has been closed. Goodbye!" + _event);
  });

  return socket
}
