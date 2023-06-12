import {json, RequestHandler} from "@sveltejs/kit";
import {voteGoEP} from "$lib/models/enums/endpoints.enum";
import {VoteModel} from "$lib/models/classes/vote.model";

export const GET :RequestHandler = async ({fetch}): Promise<Response> => {
    const voteRes = await fetch(voteGoEP.ALL);
    let vote:VoteModel = await voteRes.json()
    return json(vote)
}