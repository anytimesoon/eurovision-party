<script lang="ts">
    import {countryStore, participatingCountryStore} from "$lib/stores/country.store";
    import CountryActForm from "$lib/components/forms/CountryActForm.svelte";
    import type {ActionData} from "./$types";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import {currentUser} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
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

{#if $currentUser.authLvl === authLvl.ADMIN }
    <AdminNav />
{/if}

<div class="rounded mb-3 overflow-auto max-h-[calc(100vh-10em)]">
    <h1 class="text-center">Acts</h1>

    {#each $participatingCountryStore as country}

        <div class="py-3 border-b-2 border-b-amber-500">
            <h3>{country.flag} {country.name}</h3>
            <CountryActForm country={country}/>
        </div>
    {/each}
</div>