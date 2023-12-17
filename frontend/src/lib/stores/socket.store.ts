import {writable} from "svelte/store";
import {chatEP} from "$lib/models/enums/endpoints.enum";
import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
import {CommentModel} from "$lib/models/classes/comment.model";
import {commentStore} from "$lib/stores/comment.store";
import type {UpdateMessageModel} from "$lib/models/classes/updateMessage.model";
import {botId, userStore} from "$lib/stores/user.store";

export const socketStore = socket()
let botUserId:string
botId.subscribe(val => botUserId = val)


function socket() {
    let ws = connectToSocket()
    const {subscribe} = writable(ws)
    return {
        subscribe,
        send: (message:string) => ws.send(message),
        reconnect: () => ws = connectToSocket()
    }
}

function connectToSocket(){
    let socket = new WebSocket(chatEP)

    socket.onerror = function (error){
        console.log("Connection was lost. " + error)
        if(socket.readyState == WebSocket.CLOSED) {
            setTimeout(connectToSocket(), 1000)
        }
    }

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉")
    };

    socket.onclose = function () {
        console.log("Connection stopped. Attempting to reconnect")
        if(socket.readyState == WebSocket.CLOSED) {
            setTimeout(connectToSocket(), 1000)
        }
    };

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c: string) => {
            const chatMessage = JSON.parse(c)
            switch (chatMessage.category) {
                case chatMsgCat.COMMENT:
                    addNewComment(chatMessage.body)
                    break
                case chatMsgCat.COMMENT_ARRAY:
                    let commentModels: CommentModel[] = chatMessage.body

                    for (let i = 0; i < commentModels.length; i++) {
                        addNewComment(commentModels[i])
                    }
                    break
                case chatMsgCat.UPDATE_USER:
                    let updateMessage:UpdateMessageModel = chatMessage.body
                    console.log(updateMessage)
                    // user needs to be updated before message gets published
                    userStore.update(users => {
                        users[updateMessage.updatedUser.id] = updateMessage.updatedUser
                        return users
                    })
                    addNewComment(updateMessage.comment)
                    break
                default:
                    console.log("bad message: " + c)
            }


        })
    }
    return socket
}

function addNewComment(comment:CommentModel){
    comment.createdAt = new Date(comment.createdAt)
    commentStore.update(comments => {
        const first = comments[0]

        // if (first === undefined) {
        //     return [comment, ...comments]
        // }

        if (first && first.createdAt.getDay() != comment.createdAt.getDay() && first.userId !== botUserId) {
            let date = new Date()
            const botComment = new CommentModel(
                `${comment.createdAt.getDate()}/${comment.createdAt.getMonth()}/${comment.createdAt.getFullYear()}`,
                botUserId,
                null,
                date
                )
            return [botComment, ...comments]
        }

        if (first && first.userId === comment.userId) {
            comment.isCompact = true
        }
        return [comment, ...comments]
    });
}