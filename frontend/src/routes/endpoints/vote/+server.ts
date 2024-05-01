import {json, type RequestHandler} from "@sveltejs/kit";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";
import type {VoteModel} from "$lib/models/classes/vote.model";
import type {ResponseModel} from "$lib/models/classes/response.model";

export const GET :RequestHandler = async ({fetch}): Promise<Response> => {
    const voteRes = await fetch(voteGoEP.RESULTS);
    let vote:VoteModel = await voteRes.json()
    return json(vote)
}

export const PUT :RequestHandler = async ({fetch, request}): Promise<Response> => {
    let text = await request.text()

    const resProm = await fetch(voteGoEP.UPDATE, {
        method: "PUT",
        body: text
    })

    const res:ResponseModel<VoteModel> = await resProm.json()
    return json(res)
}