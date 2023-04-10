import {writable} from "svelte/store";
import type { CountryModel } from "$lib/models/classes/country.model";

const defaultCountryList:CountryModel[] = new Array<CountryModel>();
export const countryStore = writable<CountryModel[]>(defaultCountryList);