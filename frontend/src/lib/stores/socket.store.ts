import {writable} from "svelte/store";
import {chatMsgCat} from "$lib/models/enums/chatMsgCat";
import {CommentModel} from "$lib/models/classes/comment.model";
import {commentStore} from "$lib/stores/comment.store";
import type {UpdateMessageModel} from "$lib/models/classes/updateMessage.model";
import {botId, userStore} from "$lib/stores/user.store";
import {commentQueue} from "$lib/stores/commentQueue.store";
import {socketStateStore} from "$lib/stores/socketState.store";
import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
import {ChatMessageModel} from "$lib/models/classes/chatMessage.model";
import {errorStore} from "$lib/stores/error.store";
import type {UserModel} from "$lib/models/classes/user.model";
import type {IComment} from "$lib/models/interfaces/icomment.interface";
import {v4 as uuid} from 'uuid';
import {sessionStore} from "$lib/stores/session.store";
import type {VoteTracker} from "$lib/models/classes/voteNotification.model";
import {currentlyVotingOn} from "$lib/stores/currentlyVotingOn.store";
import {CountryModel} from "$lib/models/classes/country.model";
import type {CommentReactionModel} from "$lib/models/classes/commentReaction.model";
import {currentUser} from "$lib/stores/user.store";


let session:string
sessionStore.subscribe(val => session = val)

let botUserId:string
botId.subscribe(val => botUserId = val)

let commentLog: Array<CommentModel>
commentStore.subscribe(val => commentLog = val.values().toArray())

let thisUser: UserModel
currentUser.subscribe(val => thisUser = val)

export const socketStore = socket()
let timeoutDuration = 1000 // in milis

let cachedUrl: string

function socket() {
    let ws: WebSocket | null = null
    const {subscribe, update, set} = writable(ws)
    return {
        subscribe,
        update,
        set,
        send: (message: string) => ws.send(message),
        connect: (url: string)=> {
            if (!ws) {
                cachedUrl = url
                ws = connectToSocket(cachedUrl)
                set(ws)
            }
        }
    }
}

function connectToSocket(url: string){
    let socket = new WebSocket(url + session)

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
        setTimeout(() => socketStore.set(connectToSocket(cachedUrl)), timeoutDuration)
    }

    socket.onmessage = function (event) {
        const split = event.data.split("\n")
        split.map((c: string) => {
            if (c.length === 0 ) {
                // do nothing
            } else {
                const chatMessage = JSON.parse(c)
                switch (chatMessage.category) {
                    case chatMsgCat.COMMENT:
                        addNewComment(chatMessage.body)
                        commentQueue.removeMessage(chatMessage.body.id)
                        break
                    case chatMsgCat.COMMENT_ARRAY:
                        let commentModels: CommentModel[] = chatMessage.body
                        for (let i = 0; i < commentModels.length; i++) {
                            commentQueue.removeMessage(commentModels[i].id)
                            addNewComment(commentModels[i])
                        }
                        break
                    case chatMsgCat.UPDATE_USER:
                        let updateMessage:UpdateMessageModel = chatMessage.body
                        // a user needs to be updated before a message gets published
                        userStore.update(users => {
                            const user = users.get(updateMessage.updatedUser.id)
                            if (user.icon === updateMessage.updatedUser.icon) {
                                updateMessage.updatedUser.icon += `?${Date.now()}`
                            }
                            users.set(user.id, updateMessage.updatedUser)
                            return users
                        })
                        addNewComment(updateMessage.comment)
                        break
                    case chatMsgCat.NEW_USER:
                        let newUser:UserModel = chatMessage.body
                        userStore.update(users => {
                            users[newUser.id] = newUser
                            return users
                        })
                        break
                    case chatMsgCat.ERROR:
                        const error = chatMessage.body
                        console.log(error)
                        errorStore.set(error)
                        commentQueue.removeFirstMessage()
                        break
                    case chatMsgCat.VOTE_NOTIFICATION:
                        const voteNotification: VoteTracker = chatMessage.body
                        currentlyVotingOn.set(CountryModel.deserialize(voteNotification.country))
                        addNewComment(CommentModel.deserialize(voteNotification.comment))
                        break
                    case chatMsgCat.UPDATE_COMMENT:
                        const updateComment: CommentReactionModel = chatMessage.body
                        if (updateComment.userId !== thisUser.id) {
                            commentStore.update(val => {
                                const comment = val.get(updateComment.commentId)
                                if (comment) {
                                    comment.addOrRemoveReaction(updateComment.userId, updateComment.reaction)
                                    val.set(comment.id, comment)
                                }
                                return val
                            })
                        }
                        break
                    default:
                        errorStore.set("Oops... something just went very wrong. Please stay seated while the performances continue.")
                }
            }


        })
    }
    return socket
}

function addNewComment(commentObject:IComment){
    const comment = CommentModel.deserialize(commentObject)
    comment.createdAt = new Date(comment.createdAt)
    let latestComment = commentLog[commentLog.length - 1]

    if (latestComment && latestComment.createdAt.getDay() != comment.createdAt.getDay() && latestComment.userId !== botUserId) {
        commentStore.update( comments => {
            const date = new Date()
            const commentId = uuid()
            const botComment = new CommentModel(
                `${comment.createdAt.getDate()}/${comment.createdAt.getMonth() + 1}/${comment.createdAt.getFullYear()}`,
                botUserId,
                commentId,
                null,
                date
            )
            comments.set(commentId, botComment)
            return comments
        })
    }

    commentStore.update( comments => {
        if (latestComment && latestComment.userId === comment.userId) {
            comment.isCompact = true
        }

        comments.set(comment.id, comment)
        return comments
    })
}


