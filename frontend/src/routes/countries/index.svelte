<script lang="ts">
    import type {CountryModel} from "$lib/models/classes/country.model";
	import { countryStore } from "$lib/stores/country.store";
    import { countryEP } from "$lib/models/enums/endpoints.enum"
	import { onMount } from "svelte";
    import { sendGet } from '$lib/modules/sender';

    onMount( () => {
        sendGet<CountryModel[]>(countryEP.ALL).then( data => $countryStore = data.body);
    })
</script>

<div style="width: 50%">
    <h1>List of all Eurovision countries</h1>

    <ul id="full-list" style="text-decoration: none;">

        {#if $countryStore.length > 0}
        {#each $countryStore as country}
        <li>
            {country.flag}
            {country.name}
            {#if country.participating }
            is in the running 🎉
            {:else}
            is out of the running 😢
            {/if}
<!--            <button on:click="{() => sendPost()}">toggle</button>-->
        </li>
        {/each}
        {/if}
    </ul>

    <ul id="active-list" style="text-decoration: none;">

    </ul>
</div>

