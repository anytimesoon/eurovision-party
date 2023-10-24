import type {RequestHandler} from "@sveltejs/kit";
import {CountryModel} from "$lib/models/classes/country.model";
import {json} from "@sveltejs/kit";
import {countryGoEP} from "$lib/models/enums/endpoints.enum";

export const GET :RequestHandler = async ({fetch, params}): Promise<Response> => {
    const countryRes = await fetch(countryGoEP.FIND_ONE + params.country);
    let country:CountryModel = await countryRes.json()
    return json(country)
}