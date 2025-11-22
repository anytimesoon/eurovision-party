import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest'
import { socketStore } from '$lib/stores/socket.store'
import { commentStore } from '$lib/stores/comment.store'
import { userStore } from '$lib/stores/user.store'
import { socketStateStore } from '$lib/stores/socketState.store'
import { chatMsgCat } from '$lib/models/enums/chatMsgCat'
import { CommentModel } from '$lib/models/classes/comment.model'
import { sessionStore } from '$lib/stores/session.store'
import WS from 'vitest-websocket-mock'
import { ChatMessageModel } from "$lib/models/classes/chatMessage.model";
import { UserModel } from "$lib/models/classes/user.model";
import { authLvl } from "$lib/models/enums/authLvl.enum";

/**
 * @vitest-environment jsdom
 */
describe.sequential('Socket Store', () => {
    const session = 'test-session'
    let wsURL = 'ws://localhost:9999/'
    let mockServer: WS
    let currentComments: CommentModel[] = []
    let currentState: boolean = false

    const wsConnectionMessage = new ChatMessageModel(
        chatMsgCat.LATEST_COMMENT,
        ""
    )

    beforeEach(async () => {
        vi.clearAllMocks()
        WS.clean()
        mockServer = new WS(wsURL+session, {jsonProtocol: true})

        sessionStore.set(session)
        commentStore.set(new Map<string, CommentModel>)
        commentStore.subscribe(val => currentComments = Array.from(val.values()))
        socketStateStore.subscribe(val => currentState = val)

        socketStore.connect(wsURL)
        await mockServer.connected
    })

    afterEach(() => {
        vi.restoreAllMocks()
        socketStore.set(null)
        socketStateStore.set(false)
        commentStore.set(new Map<string, CommentModel>)
        mockServer.close()
        WS.clean()
    })

    it('should handle socket open event correctly', () => {
        expect(currentState).toBe(true)
    })


    it('should handle new comment message', async () => {
        // WS has already connected. Reading the initial message
        await expect(mockServer).toReceiveMessage(wsConnectionMessage)

        const testComment = {
            text: 'Hello World',
            userId: 'user1',
            id: 'comment1',
            replyToComment: null,
            createdAt: new Date().toISOString(),
            fileName: ""
        }

        const chatMessage = {
            category: chatMsgCat.COMMENT,
            body: testComment
        }
        // Doesn't work in a test suite. Need to spend time figuring out why
        // const payload = JSON.stringify(chatMessage)
        //
        // socketStore.send(payload)
        // await expect(mockServer).toReceiveMessage(chatMessage)

        mockServer.send(chatMessage)

        commentStore.subscribe(val => currentComments = Array.from(val.values()))

        expect(currentComments).toHaveLength(1)
        expect(currentComments[0].text).toBe(testComment.text)
        expect(currentComments[0].id).toBe(testComment.id)
        expect(currentComments[0].userId).toBe(testComment.userId)
        expect(currentComments[0].createdAt.toISOString()).toBe(testComment.createdAt)
        expect(currentComments[0].replyToComment).toBe(undefined)
        expect(currentComments[0].fileName).toBe(testComment.fileName)
    })

    it('should handle user update message', () => {
        const updateMessage = {
            category: chatMsgCat.UPDATE_USER,
            body: {
                updatedUser: {
                    id: 'user1',
                    name: 'Test User Updated',
                    icon: 'test-icon',
                    slug: 'test-user',
                    authLvl: authLvl.USER
                },
                comment: {
                    id: 'comment1',
                    text: 'User has been updated',
                    userId: 'user1',
                    replyTo: null,
                    createdAt: new Date().toISOString(),
                    fileName: "",
                    authLvl: authLvl.BOT
                }
            }
        }
        userStore.set(new Map<string, UserModel>([
            ['user1', new UserModel(
                'user1',
                'Test User',
                'test-user',
                'test-icon',
                authLvl.USER,
                [],
                'other-user',
                true
            )],
        ]))

        let users: Map<string, UserModel>
        userStore.subscribe(val => users = val)

        mockServer.send(updateMessage)

        expect(users.get('user1').name).toBe('Test User Updated')

        expect(currentComments).toHaveLength(1)
        expect(currentComments[0].text).toBe('User has been updated')
    })

    it('should handle socket close and attempt reconnection', async () => {
        const spy = vi.spyOn(window, 'setTimeout');
        expect(currentState).toBe(true)

        mockServer.close()
        expect(currentState).toBe(false)

        mockServer = new WS(wsURL+session, {jsonProtocol: true})

        // @ts-ignore
        spy.mock.calls.forEach(([cb, , ...args]) => cb(...args));
        await mockServer.connected

        await expect(mockServer).toReceiveMessage(wsConnectionMessage)
        expect(mockServer).toHaveReceivedMessages([wsConnectionMessage])

        expect(currentState).toBe(true)
        spy.mockRestore()
    })

    it('should handle comment array message', async () => {
        const comments = [
            {
                text: 'Test 1',
                userId: 'user1',
                id: 'comment1',
                replyTo: null,
                createdAt: new Date().toISOString()
            },
            {
                text: 'Test 2',
                userId: 'user2',
                id: 'comment2',
                replyTo: null,
                createdAt: new Date().toISOString()
            }
        ]

        const chatMessage = {
            category: chatMsgCat.COMMENT_ARRAY,
            body: comments
        }

        await expect(mockServer).toReceiveMessage(wsConnectionMessage)
        mockServer.send(chatMessage)

        expect(currentComments[0].text).toBe('Test 1')
        expect(currentComments[1].text).toBe('Test 2')
    })
})