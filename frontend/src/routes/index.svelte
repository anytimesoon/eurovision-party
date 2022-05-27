<script lang="ts">
    import { onMount } from "svelte";

    type Country = {
        id: string
        name: string
        bandName: string
        songName: string
        flag: string
        participating: boolean
    }

    async function getCountries(){
        const response = await fetch('http://localhost:8080/');
        countries = await response.json();

        return countries
    }

    let countries: Country[];
    let promise = getCountries();

    onMount(async function() {
        promise = getCountries();
    });

    
    
</script>

<h1>List of all Eurovision countries</h1>

<ol>
    {#await promise} 
        <p>hello</p>
    {:then countries}
        {#each countries as country}
            <li>{country.flag} {country.name} {#if country.participating } is in the final 🎉 {:else} is out of the running 😢 {/if}</li>
        {/each}
    {:catch error}
        <p>{error.message}</p>
    {/await}
</ol>
