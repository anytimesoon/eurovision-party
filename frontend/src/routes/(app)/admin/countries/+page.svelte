<script lang="ts">
    import {countryStore, notParticipatingCountryStore, participatingCountryStore} from "$lib/stores/country.store";
    import { enhance } from '$app/forms';
    import type {ActionData} from "./$types";
    import CountryForm from "$lib/components/forms/CountryForm.svelte";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import {countryGoEP} from "$lib/models/enums/endpoints.enum";


    export let form:ActionData

    const updateCountry = (form) => {
        if (form != null) {
            $countryStore = $countryStore.map((country:CountryModel) => {
                if (country.slug === form.country.slug) {
                    return form.country
                }

                return country
            })
        }
    }

    $: updateCountry(form)
</script>

    <h1>List of all Eurovision countries</h1>

    <ul>
        {#each $notParticipatingCountryStore as country}
        <li>
            {country.flag}
            {country.name}

            is out of the running 😢
            <CountryForm country={country} />
        </li>
        {/each}
    </ul>

<div class="half-width">
            {#each $participatingCountryStore as country}
                <li>
                    {country.flag}
                    {country.name}

                    is in the running 🎉
                    <CountryForm country={country} />
                </li>
            {/each}
</div>

