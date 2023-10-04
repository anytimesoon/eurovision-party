<script lang="ts">
    import {countryStore, notParticipatingCountryStore, participatingCountryStore} from "$lib/stores/country.store";
    import type {ActionData} from "./$types";
    import CountryParticipationForm from "$lib/components/forms/CountryParticipationForm.svelte";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import {currentUser} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import AdminNav from "$lib/components/AdminNav.svelte";


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

<div class="h-full flex flex-col">
    <AdminNav page="countries"/>

    <h3 class="text-center">{$participatingCountryStore.length} countries selected</h3>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="grid grid-cols-2 gap-x-3">
                <div>
                    <ul>
                        {#each $notParticipatingCountryStore as country}
                            <li class="p-3 my-1.5 border-2 border-grey-400 bg-canvas-secondary text-center w-full">
                                <CountryParticipationForm country={country} />
                            </li>
                        {/each}
                    </ul>
                </div>

                <div>
                    <ul>
                        {#each $participatingCountryStore as country}
                            <li class="p-3 my-1.5 border-2 border-primary bg-canvas-secondary text-center w-full">
                                <CountryParticipationForm country={country} />
                            </li>
                        {/each}
                    </ul>
                </div>
            </div>
        </div>
    </div>

</div>



