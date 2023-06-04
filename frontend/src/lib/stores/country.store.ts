import {derived, writable} from "svelte/store";
import type {CountryModel} from "$lib/models/classes/country.model";

const defaultCountryStore = new Array<CountryModel>()

export const countryStore = writable<CountryModel[]>(defaultCountryStore);

export const participatingCountryStore = derived(countryStore, $countryStore => {
    return $countryStore.filter(country => country.participating )
})

export const notParticipatingCountryStore = derived(countryStore, $countryStore => {
    return $countryStore.filter(country => !country.participating )
})