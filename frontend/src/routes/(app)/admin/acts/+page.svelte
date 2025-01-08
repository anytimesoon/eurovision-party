<script lang="ts">
    import {participatingCountryStore} from "$lib/stores/country.store";
    import CountryActForm from "$lib/components/forms/CountryActForm.svelte";
    import AdminNav from "$lib/components/AdminNav.svelte";
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