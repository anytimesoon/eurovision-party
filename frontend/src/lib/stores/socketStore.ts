import {writable} from "svelte/store";
import {chatEP} from "$lib/models/enums/endpoints.enum";
import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
import {CommentModel} from "$lib/models/classes/comment.model";
import {commentStore} from "$lib/stores/comment.store";

export const socketStore = socket()

function socket() {
    let ws = connectToSocket()
    const {subscribe, set, update} = writable(ws)
    return {
        subscribe,
        send: (message:string) => ws.send(message),
        reconnect: () => ws = connectToSocket()
    }
}

function connectToSocket(){
    let socket = new WebSocket(chatEP)

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉")
    };

    socket.onclose = function (e) {
        console.log("Connection stopped. Attempting to reconnect")
        if(socket.readyState === 3) {
            setTimeout(connectToSocket(), 1000)
        }
    };

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c: string) => {
            const chatMessage = JSON.parse(c)
            switch (chatMessage.category) {
                case chatMsgCat.COMMENT:
                    let comment: CommentModel = chatMessage.body
                    comment.createdAt = new Date(chatMessage.body.createdAt)
                    commentStore.update(comments => {
                        return [comment, ...comments]
                    });
                    break
                case chatMsgCat.COMMENT_ARRAY:
                    let commentModels: CommentModel[] = chatMessage.body

                    for (let i = 0; i < commentModels.length; i++) {
                        commentModels[i].createdAt = new Date(commentModels[i].createdAt)
                        commentStore.update(comments => {
                            return [commentModels[i], ...comments]
                        });
                    }
                    break
                default:
                    console.log("bad message: " + c)
            }


        })
    }
    return socket
}