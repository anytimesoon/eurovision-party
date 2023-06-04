import {countrySvelteEP} from "$lib/models/enums/endpoints.enum";
import {ResponseModel} from "$lib/models/classes/response.model";
import {CountryModel} from "$lib/models/classes/country.model";
import type {Actions, PageServerLoad} from "./$types";

export const actions : Actions = {
    update: async ({event, cookies, fetch, request}) => {
        const fd = await request.formData()
        const country:CountryModel = Object.fromEntries([...fd]) as CountryModel;
        const resProm = await fetch("http://localhost:8080/restricted/api/country/", {
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