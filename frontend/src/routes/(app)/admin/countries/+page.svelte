<script lang="ts">
    import {notParticipatingCountryStore, participatingCountryStore} from "$lib/stores/country.store";
    import CountryParticipationForm from "$lib/components/forms/CountryParticipationForm.svelte";
    import AdminNav from "$lib/components/AdminNav.svelte";
</script>

<div class="h-full flex flex-col">
    <AdminNav page="countries"/>

    <h2 class="text-center">{$participatingCountryStore.length} countries selected</h2>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="grid grid-cols-2 gap-x-3">
                <div>
                    <ul>
                        {#each $notParticipatingCountryStore as country}
                            <CountryParticipationForm country={country} />
                        {/each}
                    </ul>
                </div>

                <div>
                    <ul>
                        {#each $participatingCountryStore as country}
                            <CountryParticipationForm country={country} />
                        {/each}
                    </ul>
                </div>
            </div>
        </div>
    </div>

</div>



