<script lang="ts">
    import type {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import Spinner from "$lib/components/Spinner.svelte";

    let form:HTMLFormElement
    let formState = formButtonState.ENABLED
    export let country:CountryModel


    $: borderColour = country.participating ? "border-primary" : "border-grey-400"
</script>

<li class="my-1.5 border-2 {borderColour} bg-canvas-secondary text-center w-full">
    <form method="POST" action="?/update" bind:this={form} class="m-3" use:enhance={() => {
        formState = formButtonState.SENDING

        return async ({ update }) => {
            await update()
            formState = formButtonState.ENABLED
        };
    }}>
        <input name="slug" type="hidden" bind:value={country.slug}>
        <input name="name" type="hidden" bind:value={country.name}>
        <input name="songName" type="hidden" bind:value={country.songName}>
        <input name="bandName" type="hidden" bind:value={country.bandName}>

        <label for="check-{country.slug}" class="cursor-pointer">
            <input type="checkbox" id="check-{country.slug}" name="participating" class="hidden" bind:checked={country.participating} on:change={() => form.requestSubmit()}>
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