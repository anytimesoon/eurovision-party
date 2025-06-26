import { describe, it, expect, vi, beforeEach } from 'vitest'
import { results } from './results.store'
import { ResultModel } from '$lib/models/classes/result.model'
import { resultPageState } from '$lib/stores/resultPageState.store'
import * as genericFetch from '$lib/utils/genericFetch'
import { voteEP } from '$lib/models/enums/endpoints.enum'
import {voteCats} from "$lib/models/enums/categories.enum";
import {ResultPageStateModel} from "$lib/models/classes/resultPageState.model";

describe('results store', () => {
    let currentResults: ResultModel[]
    results.subscribe( val => currentResults = val)

    beforeEach(() => {
        results.set([])
        vi.clearAllMocks()
    })

    it('should sort results by category', () => {
        const testData = [
            generateRandomResult('france'),
            generateRandomResult('germany'),
            generateRandomResult('italy'),
            generateRandomResult('spain'),
        ]

        results.set(testData)

        results.sort(voteCats.SONG, false)
        expect(currentResults[0].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.SONG))
        expect(currentResults[1].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.SONG))
        expect(currentResults[2].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[3].getScore(voteCats.SONG))
        expect(currentResults[3].getScore(voteCats.SONG)).toBeLessThan(currentResults[0].getScore(voteCats.SONG))

        results.sort(voteCats.SONG, true)
        expect(currentResults[3].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.SONG))
        expect(currentResults[2].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.SONG))
        expect(currentResults[1].getScore(voteCats.SONG)).toBeGreaterThanOrEqual(currentResults[0].getScore(voteCats.SONG))
        expect(currentResults[0].getScore(voteCats.SONG)).toBeLessThan(currentResults[3].getScore(voteCats.SONG))

        results.sort(voteCats.TOTAL, false)
        expect(currentResults[0].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.TOTAL))
        expect(currentResults[1].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.TOTAL))
        expect(currentResults[2].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[3].getScore(voteCats.TOTAL))
        expect(currentResults[3].getScore(voteCats.TOTAL)).toBeLessThan(currentResults[0].getScore(voteCats.TOTAL))

        results.sort(voteCats.TOTAL, true)
        expect(currentResults[3].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.TOTAL))
        expect(currentResults[2].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.TOTAL))
        expect(currentResults[1].getScore(voteCats.TOTAL)).toBeGreaterThanOrEqual(currentResults[0].getScore(voteCats.TOTAL))
        expect(currentResults[0].getScore(voteCats.TOTAL)).toBeLessThan(currentResults[3].getScore(voteCats.TOTAL))

        results.sort(voteCats.PROPS, false)
        expect(currentResults[0].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.PROPS))
        expect(currentResults[1].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.PROPS))
        expect(currentResults[2].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[3].getScore(voteCats.PROPS))
        expect(currentResults[3].getScore(voteCats.PROPS)).toBeLessThan(currentResults[0].getScore(voteCats.PROPS))

        results.sort(voteCats.PROPS, true)
        expect(currentResults[3].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.PROPS))
        expect(currentResults[2].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.PROPS))
        expect(currentResults[1].getScore(voteCats.PROPS)).toBeGreaterThanOrEqual(currentResults[0].getScore(voteCats.PROPS))
        expect(currentResults[0].getScore(voteCats.PROPS)).toBeLessThan(currentResults[3].getScore(voteCats.PROPS))

        results.sort(voteCats.COSTUME, false)
        expect(currentResults[0].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.COSTUME))
        expect(currentResults[1].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.COSTUME))
        expect(currentResults[2].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[3].getScore(voteCats.COSTUME))
        expect(currentResults[3].getScore(voteCats.COSTUME)).toBeLessThan(currentResults[0].getScore(voteCats.COSTUME))

        results.sort(voteCats.COSTUME, true)
        expect(currentResults[3].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.COSTUME))
        expect(currentResults[2].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.COSTUME))
        expect(currentResults[1].getScore(voteCats.COSTUME)).toBeGreaterThanOrEqual(currentResults[0].getScore(voteCats.COSTUME))
        expect(currentResults[0].getScore(voteCats.COSTUME)).toBeLessThan(currentResults[3].getScore(voteCats.COSTUME))

        results.sort(voteCats.PERFORMANCE, false)
        expect(currentResults[0].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.PERFORMANCE))
        expect(currentResults[1].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.PERFORMANCE))
        expect(currentResults[2].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[3].getScore(voteCats.PERFORMANCE))
        expect(currentResults[3].getScore(voteCats.PERFORMANCE)).toBeLessThan(currentResults[0].getScore(voteCats.PERFORMANCE))

        results.sort(voteCats.PERFORMANCE, true)
        expect(currentResults[3].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[2].getScore(voteCats.PERFORMANCE))
        expect(currentResults[2].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[1].getScore(voteCats.PERFORMANCE))
        expect(currentResults[1].getScore(voteCats.PERFORMANCE)).toBeGreaterThanOrEqual(currentResults[0].getScore(voteCats.PERFORMANCE))
        expect(currentResults[0].getScore(voteCats.PERFORMANCE)).toBeLessThan(currentResults[3].getScore(voteCats.PERFORMANCE))
    })

    it('should check for no scores correctly', () => {
        const testData = []
        
        results.set(testData)
        expect(results.noScores()).toBe(true)
        
        const testDataWithScores = [
            generateRandomResult('france'),
            generateRandomResult('germany'),
        ]
        
        results.set(testDataWithScores)
        expect(results.noScores()).toBe(false)
    })

    it('should refresh results', async () => {
        const allResults = [
            generateRandomResult('france'),
            generateRandomResult('germany'),
            generateRandomResult('italy'),
            generateRandomResult('spain'),
        ]

        const userResults = [
            generateRandomResult('france'),
            generateRandomResult('germany'),
            generateRandomResult('italy'),
        ]
        
        let getSpy = vi.spyOn(genericFetch, 'get')
            .mockResolvedValue(allResults)

        vi.spyOn(resultPageState, 'hasUserSelected')
            .mockReturnValue(false)

        await results.refresh()

        expect(getSpy).toHaveBeenCalledWith(voteEP.RESULTS)
        expect(currentResults.length).toBe(4)

        getSpy = vi.spyOn(genericFetch, 'get')
            .mockResolvedValue(userResults)

        vi.spyOn(resultPageState, 'hasUserSelected')
            .mockReturnValue(true)
        resultPageState.set(new ResultPageStateModel('test-user-id', voteCats.TOTAL ))

        await results.refresh()
        expect(getSpy).toHaveBeenCalledWith(voteEP.RESULTS + 'test-user-id')
        expect(currentResults.length).toBe(3)
    })
})

const generateRandomResult = (country: string): ResultModel => {

    const costume = Math.floor(Math.random() * 11)
    const song = Math.floor(Math.random() * 11)
    const perf = Math.floor(Math.random() * 11)
    const props = Math.floor(Math.random() * 11)

    return new ResultModel(
        country,
        costume,
        song,
        perf,
        props,
        costume + song + perf + props
    )
}