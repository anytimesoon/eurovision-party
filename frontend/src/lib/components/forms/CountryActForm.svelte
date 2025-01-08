<script lang="ts" xmlns="http://www.w3.org/1999/html">
    import {CountryModel} from "$lib/models/classes/country.model";
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {countryEP} from "$lib/models/enums/endpoints.enum";
    import {countryStore} from "$lib/stores/country.store";
    import {put} from "$lib/utils/genericFetch";

    let formState = formButtonState.ENABLED
    export let country:CountryModel
    let form:HTMLFormElement

    const submit = async () => {
        formState = formButtonState.SENDING
        const fd = new FormData(form)
        const newCountry = new CountryModel(
            country.name,
            country.slug,
            fd.get("bandName") as string,
            fd.get("songName") as string,
            country.flag,
            country.participating
        )

        const updatedCountry = await put<CountryModel, CountryModel>(countryEP.UPDATE, newCountry)
            .then(res => CountryModel.deserialize(res))
        $countryStore[$countryStore.findIndex(c => c.slug === updatedCountry.slug)] = updatedCountry

        formState = formButtonState.ENABLED
    }

</script>

<form bind:this={form} on:submit|preventDefault={submit}>
    <input class="mb-3" id="{country.slug}-song" name="songName" type="text" bind:value={country.songName} placeholder="Song Title">

    <input class="mb-3" id="{country.slug}-act" name="bandName" type="text" bind:value={country.bandName} placeholder="Act Name">

    <FormButton state={formState}>
        <ContentSave size="1.4em" /> Save
    </FormButton>
</form>