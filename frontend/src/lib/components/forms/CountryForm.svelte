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

    <label for="check-{country.slug}"  class='flex w-fit items-center justify-evely text-grey-8 relative'>
        <input type="checkbox" id="check-{country.slug}" name="participating" class="peer hidden" bind:checked={country.participating} on:change={() => form.requestSubmit()}>
        <span class="relative z-10 inset-0 p-3 before:-z-10 before:content-[''] before:absolute before:inset-0 before:h-full before:w-full before:border-2 before:border-grey-4 peer-checked:before:bg-indigo-200 peer-checked:before:text-indigo-800 peer-checked:before:border-indigo-800 peer-checked:before:block">
            {country.flag} {country.name}
        </span>
    </label>
</form>