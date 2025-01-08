<script lang="ts">
    import {CountryModel} from "$lib/models/classes/country.model";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Spinner from "$lib/components/Spinner.svelte";
    import {countryEP} from "$lib/models/enums/endpoints.enum";
    import {put} from "$lib/utils/genericFetch";
    import {countryStore} from "$lib/stores/country.store";

    let formState = formButtonState.ENABLED
    export let country:CountryModel

    const submit = async (country: CountryModel) => {
        formState = formButtonState.SENDING
        const updatedCountry = await put<CountryModel, CountryModel>(countryEP.UPDATE, country)
            .then(res => CountryModel.deserialize(res))
        $countryStore[$countryStore.findIndex(c => c.slug === updatedCountry.slug)] = updatedCountry
        formState = formButtonState.ENABLED
    }

    $: borderColour = country.participating ? "border-primary" : "border-grey-400"
</script>

<li class="my-1.5 border-2 {borderColour} bg-canvas-secondary text-center w-full">
    <form class="m-3" >
        <input name="slug" type="hidden" bind:value={country.slug}>
        <input name="name" type="hidden" bind:value={country.name}>
        <input name="flag" type="hidden" bind:value={country.flag}>
        <input name="songName" type="hidden" bind:value={country.songName}>
        <input name="bandName" type="hidden" bind:value={country.bandName}>

        <label for="check-{country.slug}" class="cursor-pointer">
            <input type="checkbox" id="check-{country.slug}" name="participating" class="hidden" bind:checked={country.participating} on:change={() => submit(country)}>
            <span>
                {#if formState === formButtonState.SENDING}
                    <Spinner size={"s"} thickness={"s"} isBlock={false}/>
                {:else}
                    {country.flag}
                {/if}
                    {country.name}
            </span>
        </label>
    </form>
</li>