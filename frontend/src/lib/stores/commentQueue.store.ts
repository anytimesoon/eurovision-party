import {type Writable, writable} from "svelte/store";
import type {CommentModel} from "$lib/models/classes/comment.model";
import type {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
import {socketStore} from "$lib/stores/socket.store";
import {socketStateStore} from "$lib/stores/socketState.store";

export const commentQueue = newCommentQueue()

let currentQueue: Array<ChatMessageModel<CommentModel>>
commentQueue.subscribe( val => currentQueue = val)

let socketState:boolean
socketStateStore.subscribe(val => socketState = val)

function newCommentQueue(): Writable<Array<ChatMessageModel<CommentModel>>> {
    const {subscribe, update} = writable(new Array<ChatMessageModel<CommentModel>>())
    return {
        subscribe,
        update,
        restart: () => restart(),
        removeMessage: (messageId:string) => removeMessageHandler(messageId),
        addComment: (chatMessage:ChatMessageModel<CommentModel>) => addCommentHandler(chatMessage)
    }
}

function restart() {
    if(currentQueue.length > 0 && socketState) {
        send(currentQueue[0])
    }
}

function removeMessageHandler(messageId:string) {
    commentQueue.update( queue => {
        return queue.filter( message => message.body.id != messageId)
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

function send(message) {
    socketStore.send(JSON.stringify(message))
}