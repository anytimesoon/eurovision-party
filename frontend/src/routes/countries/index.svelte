<script lang="ts" context="module">
    import {getCountries} from "$lib/stores/country.store";
    import {updateCountry} from "$lib/stores/country.store";
    getCountries();
</script>

<script lang="ts">
    import type {CountryModel} from "$lib/models/classes/country.model"
	import { countryStore } from "$lib/stores/country.store";

    let countries:CountryModel[];
    countryStore.subscribe(val => {
            countries = val;
        });
</script>

<div style="width: 50%">
    <h1>List of all Eurovision countries</h1>

    <ul id="full-list" style="text-decoration: none;">
        {#await countries} 
            <p>loading</p>
        {:then countries}
            {#if countries.length > 0}
                {#each countries as country}
                    <li>
                        {country.flag} 
                        {country.name} 
                        {#if country.participating } 
                            is in the running 🎉
                        {:else} 
                            is out of the running 😢
                        {/if}
                        <button on:click="{() => updateCountry(country)}">toggle</button>
                    </li>
                {/each}
            {/if}
        {:catch error}
            <p>{error.message}</p>
        {/await}
    </ul>
</div>

<div style="width: 50%">
    <ul id="active-list" style="text-decoration: none;">

    </ul>
</div>

