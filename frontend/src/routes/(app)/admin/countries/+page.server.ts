import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import type {Actions} from "./$types";
import {countryGoEP} from "$lib/models/enums/endpoints.enum";

export const actions : Actions = {
    update: async ({fetch, request}) => {
        const fd = await request.formData()
        const country:CountryModel = Object.fromEntries([...fd]) as CountryModel;
        const resProm = await fetch(countryGoEP.UPDATE, {
            method: "PUT",
            body: JSON.stringify(country, (key, value) => {
            if (key == "participating") {
                return value === "on"
            } else {
                return value
            }})
        })

        const res:ResponseModel<CountryModel> = await resProm.json()

        return {
            success: true,
            country: res.body,
            error: res.error
        }
    }
}