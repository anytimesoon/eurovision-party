import {ResultModel} from "$lib/models/classes/result.model";
import {writable} from "svelte/store";
import {resultPageState} from "$lib/stores/resultPageState.store";
import {get} from "$lib/utils/genericFetch";
import {voteEP} from "$lib/models/enums/endpoints.enum";
import type {ResultPageStateModel} from "$lib/models/classes/resultPageState.model";
import type { IResultModel } from "$lib/models/interfaces/iresultmodel.interface";

export const results = newResultsStore()

let currentResults: ResultModel[]
results.subscribe(val => currentResults = val)

let currentStatus: ResultPageStateModel
resultPageState.subscribe(val => currentStatus = val)

function newResultsStore() {
    const {subscribe, update, set} = writable(new Array<ResultModel>())

    return {
        subscribe,
        update,
        set,
        refresh: async () => refresh(),
        noScores: (): boolean => noScores(),
    }
}

async function refresh() {
    if (!resultPageState.hasUserSelected()) {
        const newResults = await get(voteEP.RESULTS)
        results.set(newResults.map((result: IResultModel) => ResultModel.deserialize(result)))
    } else {
        const newResults = await get(voteEP.RESULTS + currentStatus.userId)
        results.set(newResults.map((result: IResultModel) => ResultModel.deserialize(result)))
    }
}

function noScores():boolean {
    const filtered = currentResults.filter((res) => res.total > 0)
    return filtered.length === 0
}