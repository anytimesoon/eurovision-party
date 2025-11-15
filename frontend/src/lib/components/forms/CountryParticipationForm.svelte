<script lang="ts">
    import {CountryModel} from "$lib/models/classes/country.model";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Spinner from "$lib/components/Spinner.svelte";
    import {countryEP} from "$lib/models/enums/endpoints.enum";
    import {put} from "$lib/utils/genericFetch";
    import {countryStore} from "$lib/stores/country.store";
    import {crossfade} from "$lib/utils/crossfade";
    import {flip} from "svelte/animate";

    let formState = $state(formButtonState.ENABLED)
    const [send, receive] = crossfade
    interface Props {
        countries: CountryModel[]
        isParticipating: boolean
    }

    let { countries, isParticipating }: Props = $props();

    const submit = async (country: CountryModel) => {
        formState = formButtonState.SENDING
        const updatedCountry = await put<CountryModel, CountryModel>(countryEP.UPDATE, country)
            .then(res => CountryModel.deserialize(res))
        $countryStore[$countryStore.findIndex(c => c.slug === updatedCountry.slug)] = updatedCountry
        formState = formButtonState.ENABLED
    }

    let borderColour = $derived(isParticipating ? "border-primary" : "border-grey-400")
</script>

{#each countries as country (country.slug)}
    <li class="my-1.5 border-2 {borderColour} bg-canvas-secondary text-center w-full"
        animate:flip={{duration: 300}}
        in:receive={{key: country.slug}}
        out:send={{key: country.slug}}>
        <form class="m-3" >
            <input name="slug" type="hidden" bind:value={country.slug}>
            <input name="name" type="hidden" bind:value={country.name}>
            <input name="flag" type="hidden" bind:value={country.flag}>
            <input name="songName" type="hidden" bind:value={country.songName}>
            <input name="bandName" type="hidden" bind:value={country.bandName}>

            <label for="check-{country.slug}" class="cursor-pointer">
                <input type="checkbox" id="check-{country.slug}" name="participating" class="hidden" bind:checked={country.participating} onchange={() => submit(country)}>
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
{/each}