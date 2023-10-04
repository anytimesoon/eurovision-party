<script lang="ts" xmlns="http://www.w3.org/1999/html">
    import {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';

    let form;
    export let country:CountryModel = new CountryModel("", "", "", "", "", false)

</script>

<form method="POST" action="?/update" use:enhance bind:this={form}>
    <input name="slug" type="hidden" bind:value={country.slug}>
    <input type="checkbox" name="participating" class="hidden" bind:checked={country.participating}>

    <input class="mb-3" id="{country.slug}-song" name="songName" type="text" bind:value={country.songName} placeholder="Song Title">

    <input class="mb-3" id="{country.slug}-act" name="bandName" type="text" bind:value={country.bandName} placeholder="Act Name">

    <button class="mb-3">Save</button>
</form>