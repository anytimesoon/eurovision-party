<script lang="ts">
    import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
    import CountryActForm from "$lib/components/forms/CountryActForm.svelte";
    import type {ActionData} from "./$types";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import AdminNav from "$lib/components/AdminNav.svelte";

    export let form:ActionData

    const updateCountry = (form:ActionData) => {
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

<div class="h-full flex flex-col">
    <AdminNav page="acts" />

    <h2 class="text-center">Acts</h2>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            {#each $participatingCountryStore as country}

                <div class="p-3">
                    <div class="p-3 border-2 border-secondary rounded">
                        <h3>{country.flag} {country.name}</h3>
                        <CountryActForm country={country}/>
                    </div>
                </div>
            {/each}
        </div>
    </div>
</div>