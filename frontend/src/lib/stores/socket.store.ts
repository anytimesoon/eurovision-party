import {writable} from "svelte/store";
import {chatEP} from "$lib/models/enums/endpoints.enum";
import {commentStore} from "$lib/stores/comment.store";
import type {TokenModel} from "$lib/models/classes/token.model";
import {tokenStore} from "$lib/stores/token.store";
import type {IComment} from "$lib/models/interfaces/icomment.interface";

let auth:TokenModel
tokenStore.subscribe(val => auth = val)

export const socketStore = connect()

function connect(){
        const {subscribe, set} = writable({})
        const socket = connectToSocket()

        return {
            subscribe,
            set,
            send: (message:string) => socket.send(message)
        }
}

function connectToSocket():WebSocket{
    const token = auth.token || "";
    const socket = new WebSocket(chatEP, [token, "chat"]);

    let timeout = 250;

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉");
        timeout = 250;
    };

    socket.onmessage = function (event) {
        const data: IComment = JSON.parse(event.data);
        commentStore.update(comments => {
            return [...comments, data]
        });
    };

    socket.onclose = function (e) {
        console.log('Socket is closed. Reconnect will be reattempted in ' + timeout + "milliseconds. " + e.reason);
        setTimeout(connectToSocket, Math.min(10000, timeout += timeout));
    };

    return socket
}