<script lang="ts">
    import { onMount } from "svelte";
    import { loop_guard } from "svelte/internal";
    import type {CountryModel} from "src/lib/models/classes/country.model"

    async function getCountries(){
        const response = await fetch('http://localhost:8080/');
        countries = await response.json();

        return countries
    }

    let countries: CountryModel[];
    let promise = getCountries();

    onMount( () => promise = getCountries() );

    
</script>

<h1>List of all Eurovision countries</h1>

<ul style="text-decoration: none;">
    {#await promise} 
        <p>loading</p>
    {:then countries}
        {#each countries as country}
            <li>
                {country.flag} 
                {country.name} 
                {#if country.participating } 
                    is in the running 🎉
                {:else} 
                    is out of the running 😢
                {/if}
                <button on:click="{() => country.update()}">toggle</button>
            </li>
        {/each}
    {:catch error}
        <p>{error.message}</p>
    {/await}
</ul>
