<script lang="ts" xmlns="http://www.w3.org/1999/html">
    import type {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';
    import FormButton from "$lib/components/forms/FormButton.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";

    let formState = formButtonState.ENABLED
    export let country:CountryModel

</script>

<form method="POST" action="?/update" use:enhance={() => {
        formState = formButtonState.SENDING

        return async ({ update }) => {
            await update()
            formState = formButtonState.ENABLED
        };
    }}>
    <input name="slug" type="hidden" bind:value={country.slug}>
    <input name="name" type="hidden" bind:value={country.name}>
    <input name="flag" type="hidden" bind:value={country.flag}>

    <input type="checkbox" name="participating" class="hidden" bind:checked={country.participating}>

    <input class="mb-3" id="{country.slug}-song" name="songName" type="text" bind:value={country.songName} placeholder="Song Title">

    <input class="mb-3" id="{country.slug}-act" name="bandName" type="text" bind:value={country.bandName} placeholder="Act Name">

    <FormButton state={formState}>
        <ContentSave size="1.4em" /> Save
    </FormButton>
</form>