import { describe, it, expect, beforeEach, vi } from 'vitest'
import { get } from 'svelte/store'
import { userStore, currentUser, botId } from './user.store'
import { UserModel } from '$lib/models/classes/user.model'

vi.mock('$app/environment', () => ({
    browser: true
}))


/**
 * @vitest-environment jsdom
 */
describe('User Store', () => {
    const testUser = new UserModel(
        'test-id',
        'Test User',
        'test-user',
        'test-icon',
        0
    )
    beforeEach(() => {
        localStorage.clear()
        userStore.set(new Map<string, UserModel>())
        currentUser.set(UserModel.empty())
        botId.set('')
    })

    describe('userStore', () => {
        it('should initialize as empty Map', () => {
            const store = get(userStore)
            expect(store).toBeInstanceOf(Map)
            expect(store.size).toBe(0)
        })

        it('should allow adding new users', () => {
            userStore.update(store => {
                store.set(testUser.id, testUser)
                return store
            })

            const store = get(userStore)
            expect(store.get('test-id')).toEqual(testUser)
        })
    })

    describe('currentUserStore', () => {
        it('should initialize with empty user when no localStorage data', () => {
            const user = get(currentUser)
            expect(user).toEqual(UserModel.empty())
        })

        it('should load user from localStorage', () => {
            currentUser.set(testUser)

            const user = get(currentUser)
            expect(user).toEqual(testUser)

            const localStorageUser = UserModel.deserialize(JSON.parse(localStorage.getItem('currentUser')))
            expect(localStorageUser).toEqual(user)
        })

        it('should update localStorage when currentUser changes', () => {
            currentUser.set(testUser)
            const user = get(currentUser)
            expect(user).toEqual(testUser)

            testUser.name = 'New Name'
            currentUser.set(testUser)
            const updatedUser = get(currentUser)
            expect(updatedUser.name).toEqual(testUser.name)

            const localStorageUser = UserModel.deserialize(JSON.parse(localStorage.getItem('currentUser')))
            expect(localStorageUser).toEqual(updatedUser)
        })
    })

    describe('botId', () => {
        it('should initialize as empty string when no localStorage data', () => {
            const id = get(botId)
            expect(id).toBe('')
        })

        it('should load botId from localStorage', () => {
            botId.set('test-bot-id')

            const currentId = get(botId)
            expect(currentId).toEqual('test-bot-id')

            const id = localStorage.getItem('botUser')
            expect(id).toBe('test-bot-id')
        })

        it('should update localStorage when botId changes', () => {
            botId.set('test-bot-id')
            const id = get(botId)
            expect(id).toBe('test-bot-id')

            botId.set('new-bot-id')
            const newId = get(botId)
            expect(newId).toEqual('new-bot-id')

            const localStorageId = localStorage.getItem('botUser')
            expect(localStorageId).toEqual(newId)
        })
    })
})