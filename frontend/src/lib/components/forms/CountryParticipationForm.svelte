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

    <label for="check-{country.slug}">
        <input type="checkbox" id="check-{country.slug}" name="participating" class="hidden" bind:checked={country.participating} on:change={() => form.requestSubmit()}>
        <span>
            {country.flag} {country.name}
        </span>
    </label>
</form>