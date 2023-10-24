import {countrySvelteEP, voteSvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import {VoteModel} from "$lib/models/classes/vote.model";

export const prerender = 'auto'

export async function load({params, fetch}):PageLoad {
    let res = await fetch(countrySvelteEP.FIND_ONE + params.country)
    const countryRes:ResponseModel<CountryModel> = await res.json()

    res = await fetch(voteSvelteEP.BY_COUNTRY_AND_USER + params.country)
    const voteRes:ResponseModel<VoteModel> = await res.json()
    const vote = new VoteModel().deserialize(voteRes.body)

    return {
        country: countryRes.body,
        vote
    }
}