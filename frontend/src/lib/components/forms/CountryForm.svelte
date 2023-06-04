<script lang="ts">
    import {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';

    let form;
    export let country:CountryModel = new CountryModel("", "", "", "", "", false)

</script>

<form method="POST" action="?/update" use:enhance bind:this={form}>
    <input name="slug" type="hidden" bind:value={country.slug}>
    {#if !country.participating}
        <input name="songName" type="hidden" bind:value={country.songName}>
        <input name="bandName" type="hidden" bind:value={country.bandName}>
    {:else }
        <input name="songName" type="text" bind:value={country.songName} on:change={() => form.requestSubmit()}>
        <input name="bandName" type="text" bind:value={country.bandName} on:change={() => form.requestSubmit()}>
    {/if}
    <input name="participating" type="checkbox" bind:checked={country.participating} on:change={() => form.requestSubmit()}>
</form>