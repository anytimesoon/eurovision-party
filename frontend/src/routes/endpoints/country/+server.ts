import type {RequestHandler} from "@sveltejs/kit";
import type {CountryModel} from "$lib/models/classes/country.model";
import {countryGoEP} from "$lib/models/enums/endpoints.enum";
import {json} from "@sveltejs/kit";

export const GET :RequestHandler = async ({fetch, cookies}): Promise<Response> => {
    const countryRes:Response = await fetch(countryGoEP.ALL)

    let countries:Array<CountryModel> = await countryRes.json()
    return json(countries)
}