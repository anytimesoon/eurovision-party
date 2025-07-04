import {writable} from "svelte/store";
import type {CommentModel} from "$lib/models/classes/comment.model";
import type {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
import {socketStore} from "$lib/stores/socket.store";
import {socketStateStore} from "$lib/stores/socketState.store";
import {browser} from "$app/environment";

export const commentQueue = newCommentQueue()

let currentQueue: Array<ChatMessageModel<CommentModel>>
commentQueue.subscribe( val => {
    browser && localStorage.setItem("commentQueue", JSON.stringify(val))
    currentQueue = val
})

let socketState:boolean
socketStateStore.subscribe(val => socketState = val)

let socket:WebSocket | undefined

function newCommentQueue() {
    const {subscribe, update} = writable(
        (browser && JSON.parse(
            localStorage.getItem("commentQueue") ||
            JSON.stringify(new Array<ChatMessageModel<CommentModel>>())
        )) || new Array<ChatMessageModel<CommentModel>>()
    )

    return {
        subscribe,
        update,
        restart: () => restart(),
        removeMessage: (messageId:string) => removeMessageHandler(messageId),
        removeFirstMessage: () => removeFirstMessageHandler(),
        addComment: (chatMessage:ChatMessageModel<CommentModel>) => addCommentHandler(chatMessage)
    }
}

function restart() {
    if(currentQueue.length > 0 && socketState) {
        send(currentQueue[0])
    }
}

function removeFirstMessageHandler() {
    commentQueue.update( queue => {
        queue.shift()
        return queue
    })
}

function removeMessageHandler(messageId:string) {
    commentQueue.update( queue => {
        return queue.filter( (message:ChatMessageModel<CommentModel>) => message.body.id != messageId)
    })

    if(currentQueue.length > 0 && socketState) {
        send(currentQueue[0])
    }
}

function addCommentHandler(chatMessage:ChatMessageModel<CommentModel>) {
    commentQueue.update( queue => [...queue, chatMessage])

    if (socketState) {
        send(currentQueue[0])
    }
}

// Effectively private. Do not call from outside the queue
function send(message: ChatMessageModel<CommentModel>) {
    if(socket === undefined) {
        connectToSocket()
    }

    if(socket !== undefined) {
        socket.send(JSON.stringify(message))
    }
}

function connectToSocket() {
    socketStore.subscribe(val => socket = val)
}