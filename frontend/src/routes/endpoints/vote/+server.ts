import {json, RequestHandler} from "@sveltejs/kit";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";
import {VoteModel} from "$lib/models/classes/vote.model";
import {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
import {ResponseModel} from "$lib/models/classes/response.model";

export const GET :RequestHandler = async ({fetch}): Promise<Response> => {
    const voteRes = await fetch(voteGoEP.RESULTS);
    let vote:VoteModel = await voteRes.json()
    return json(vote)
}

export const POST:RequestHandler = async ({fetch, request}):Promise<Response> => {
    const fd = await request.formData()
    const vote:VoteSingleModel = Object.fromEntries([...fd]) as VoteSingleModel;

    const resProm = await fetch(voteGoEP.UPDATE, {
        method: "PUT",
        body: JSON.stringify(vote,(key, value) => {
            if (key == "score") {
                return parseInt(value)
            }
            return value
        })
    })

    const res:ResponseModel<VoteModel> = await resProm.json()
    return json(res.body)
}