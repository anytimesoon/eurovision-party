<script lang="ts" context="module">
    import type { Load } from '@sveltejs/kit';

    export const load : Load = async ({ fetch }) => ({
        props: {
            countries: await fetch("http://localhost:8080/api/country").then((res) => res.json() as Promise<CountryModel[]>),
        }
     })
</script>

<script lang="ts">
    import type {CountryModel} from "src/lib/models/classes/country.model"

    export let countries: CountryModel[] = [];
</script>

<h1>List of all Eurovision countries</h1>

<ul style="text-decoration: none;">
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
                    <button on:click="{() => country.update()}">toggle</button>
                </li>
            {/each}
        {/if}
    {:catch error}
        <p>{error.message}</p>
    {/await}
</ul>
