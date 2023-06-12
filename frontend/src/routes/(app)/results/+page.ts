import {voteSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {VoteModel} from "$lib/models/classes/vote.model";

export async function load({fetch}) {
    const voteRes = await fetch(voteSvelteEP.ALL);

    const votes:ResponseModel<VoteModel[]> = await voteRes.json()

    return {
        votes: votes.body
    }
}