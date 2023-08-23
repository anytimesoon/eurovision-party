<script lang="ts">
    import {countryStore, notParticipatingCountryStore, participatingCountryStore} from "$lib/stores/country.store";
    import type {ActionData} from "./$types";
    import CountryForm from "$lib/components/forms/CountryForm.svelte";
    import type {CountryModel} from "$lib/models/classes/country.model";


    export let form:ActionData

    const updateCountry = (form) => {
        if (form != null) {
            $countryStore = $countryStore.map((country:CountryModel) => {
                if (country.slug === form.country.slug) {
                    return form.country
                }

                return country
            })
        }
    }

    $: updateCountry(form)
</script>



<div class="flex flex-col h-full">
    <h1>List of all Eurovision countries</h1>
    <div class="flex-grow flex-auto h-0 p-4 rounded mb-3">
        <div class="grid grid-cols-1">
            <div class="col-end-1">
                <ul>
                    {#each $notParticipatingCountryStore as country}
                        <li class="p-1">
                            <CountryForm country={country} />
                        </li>
                    {/each}
                </ul>
            </div>

            <div class="col-end-2">
                <ul>
                    {#each $participatingCountryStore as country}
                        <li class="p-1 text-right">
                            <CountryForm country={country} />
                        </li>
                    {/each}
                </ul>
            </div>
        </div>



    </div>
</div>

