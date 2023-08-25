<script lang="ts">
    import {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';

    let form;
    export let country:CountryModel = new CountryModel("", "", "", "", "", false)

</script>

<form method="POST" action="?/update" use:enhance bind:this={form}>
    <input name="slug" type="hidden" bind:value={country.slug}>

        <input name="songName" type="hidden" bind:value={country.songName}>
        <input name="bandName" type="hidden" bind:value={country.bandName}>

    <!--    <input name="songName" type="text" bind:value={country.songName} on:change={() => form.requestSubmit()}>-->
    <!--    <input name="bandName" type="text" bind:value={country.bandName} on:change={() => form.requestSubmit()}>-->

    <label for="check-{country.slug}">
        <input type="checkbox" id="check-{country.slug}" name="participating" class="hidden" bind:checked={country.participating} on:change={() => form.requestSubmit()}>
        <span class="p-3 h-full w-full border-2 border-grey-4" class:border-amber-300={country.participating}>
            {country.flag} {country.name}
        </span>
    </label>
</form>