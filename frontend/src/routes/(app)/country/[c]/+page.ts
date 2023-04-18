import {partCountryStore} from "$lib/stores/partCountry.store";
import type {CountryModel} from "$lib/models/classes/country.model";


/** @type {import('../../../index').PageLoad} */
// eslint-disable-next-line @typescript-eslint/no-unused-vars,@typescript-eslint/ban-ts-comment
// @ts-ignore
export const load = ({params}):CountryPage => {
    let countries = new Array<CountryModel>()
    partCountryStore.subscribe(val => countries = val)
    const country = countries.find( c => c.slug === params.c)
    return {
        c: country
    }
}

export class CountryPage{
    c!: CountryModel|undefined;
}