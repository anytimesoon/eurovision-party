<script lang="ts">
    import { onMount } from "svelte";
    import { loop_guard } from "svelte/internal";

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

    onMount( () => promise = getCountries() );

    function updateParticipating(countryName: string, participatingStatus: boolean) {
        fetch("http://localhost:8080/country",{
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "name": countryName,
                "participating": participatingStatus
            })
        })
    }  
</script>

<h1>List of all Eurovision countries</h1>

<ul style="text-decoration: none;">
    {#await promise} 
        <p>hello</p>
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
                <button on:click="{() => updateParticipating(country.name, country.participating)}">toggle</button>
            </li>
        {/each}
    {:catch error}
        <p>{error.message}</p>
    {/await}
</ul>
