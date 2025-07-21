import {writable} from "svelte/store";
import {CountryModel} from "$lib/models/classes/country.model";

export const currentlyVotingOn = writable<CountryModel>(
    new CountryModel(
        "",
        "",
        "",
        "",
        "",
        false
    )
);
