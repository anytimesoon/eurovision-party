import { writable } from "svelte/store";
import type { CountryModel } from "$lib/models/classes/country.model";
  
export const countryStore = writable<CountryModel[]>([]);
