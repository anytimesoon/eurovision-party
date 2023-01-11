<script lang="ts">
    import type {CountryModel} from "$lib/models/classes/country.model";
	import { countryStore } from "$lib/stores/country.store";
    import { countryEP } from "$lib/models/enums/endpoints.enum"
	import { onMount } from "svelte";
    import {sendGet} from '$lib/helpers/sender.helper';
    import {
        updateCountryBand,
        updateCountryParticipation,
        updateCountrySong
    } from "$lib/helpers/country.helper.js";

    onMount( () => {
        sendGet<CountryModel[]>(countryEP.ALL).then( data => $countryStore = data.body);
    })
</script>

<div>
    <h1>List of all Eurovision countries</h1>

    <ul id="full-list" style="text-decoration: none;">

        {#if $countryStore.length > 0}
        {#each $countryStore as country}
            {#if !country.participating }
        <li>
            {country.flag}
            {country.name}

            is out of the running 😢

            <button on:click="{() => updateCountryParticipation(country)}">toggle</button>
        </li>
            {/if}
        {/each}
        {/if}
    </ul>
</div>
<div>
        {#if $countryStore.length > 0}
            {#each $countryStore as country}
                {#if country.participating }
                <li>
                    {country.flag}
                    {country.name}

                    is in the running 🎉

                    <button on:click="{() => updateCountryParticipation(country)}">toggle</button>

                    <label for="bandName">Band Name</label>
                    <input value={country.bandName} id="bandName" on:focusout={(e) => updateCountryBand(country, e.target.value)}>

                    <label for="songName">Song Name</label>
                    <input value={country.songName} id="songName" on:focusout={(e) => updateCountrySong(country, e.target.value)}>
                </li>
                {/if}
            {/each}
        {/if}
</div>

