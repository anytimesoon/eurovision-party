import {writable} from "svelte/store";
import {chatEP} from "$lib/models/enums/endpoints.enum";
import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
import {CommentModel} from "$lib/models/classes/comment.model";
import {commentStore} from "$lib/stores/comment.store";
import type {UpdateMessageModel} from "$lib/models/classes/updateMessage.model";
import {botId, currentUser, userStore} from "$lib/stores/user.store";
import {commentQueue} from "$lib/stores/commentQueue.store";
import {loginURI} from "$lib/stores/loginURI.store";
import {socketStateStore} from "$lib/stores/socketState.store";
import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";

let loginEP:string
loginURI.subscribe(val => loginEP = val)

let currentUserID:string
currentUser.subscribe(val => currentUserID = val.id)

let botUserId:string
botId.subscribe(val => botUserId = val)

let commentLog: Array<CommentModel>
commentStore.subscribe(val => commentLog = val)

let retryCount:number
socketRetryCount.subscribe(val => retryCount = val)

export const socketStore = socket()
let timeoutDuration = 1000 // in milis


function socket() {
    let ws = connectToSocket()
    const {subscribe, update} = writable(ws)
    return {
        subscribe,
        update,
        send: (message:string) => ws.send(message),
        reconnect: () => ws = connectToSocket()
    }
}

function connectToSocket(){
    let socket = new WebSocket(chatEP + loginEP)

    socket.onopen = function () {
        console.log("You're connected. Welcome to the party!!!🎉")
        socketRetryCount.reset()
        const latestCommentId: string = commentLog.length > 0 ? commentLog[0].id : ""
        const latestCommentMessage = new ChatMessageModel(
            chatMsgCat.LATEST_COMMENT,
            latestCommentId
        )
        socket.send(JSON.stringify(latestCommentMessage))

        socketStateStore.isReady(true)
        commentQueue.restart()
    }

    socket.onerror = function () {
        console.log("Websocket error")
    }

    socket.onclose = function () {
        console.log("Connection stopped. Attempting to reconnect")
        socketRetryCount.increment()
        socketStateStore.isReady(false)
        setTimeout(() => socketStore.reconnect(), timeoutDuration * retryCount)
    }

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c: string) => {
            const chatMessage = JSON.parse(c)
            switch (chatMessage.category) {
                case chatMsgCat.COMMENT:
                    socketStateStore.isReady(true)
                    addNewComment(chatMessage.body)
                    commentQueue.removeMessage(chatMessage.body.id)
                    break
                case chatMsgCat.COMMENT_ARRAY:
                    let commentModels: CommentModel[] = chatMessage.body
                    for (let i = 0; i < commentModels.length; i++) {
                        commentQueue.removeMessage(commentModels[i].id)
                        addNewComment(commentModels[i])
                    }
                    socketStateStore.isReady(true)
                    break
                case chatMsgCat.UPDATE_USER:
                    socketStateStore.isReady(true)
                    let updateMessage:UpdateMessageModel = chatMessage.body
                    // setTimeout(()=>{}, 500)
                    // user needs to be updated before message gets published
                    userStore.update(users => {
                        if (users[updateMessage.updatedUser.id].icon === updateMessage.updatedUser.icon) {
                            updateMessage.updatedUser.icon += `?${Date.now()}`
                        }
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
    let latestComment = commentLog[0]

    if (latestComment && latestComment.createdAt.getDay() != comment.createdAt.getDay() && latestComment.userId !== botUserId) {
        commentStore.update( comments => {
            let date = new Date()
            const botComment = new CommentModel(
                `${comment.createdAt.getDate()}/${comment.createdAt.getMonth() + 1}/${comment.createdAt.getFullYear()}`,
                botUserId,
                null,
                date
            )
            return [botComment, ...comments]
        })
    }

    commentStore.update( comments => {
        if (latestComment && latestComment.userId === comment.userId) {
            comment.isCompact = true
        }

        return [comment, ...comments]
    })
}


