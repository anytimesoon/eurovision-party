<script lang="ts">
    import {notParticipatingCountryStore, participatingCountryStore} from "$lib/stores/country.store";
    import CountryParticipationForm from "$lib/components/forms/CountryParticipationForm.svelte";
    import SettingsNav from "$lib/components/SettingsNav.svelte";
</script>

<div class="h-full flex flex-col">
    <SettingsNav page="countries"/>

    <h2 class="text-center">{$participatingCountryStore.length} countries selected</h2>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="grid grid-cols-2 gap-x-3">
                <div>
                    <ul>
                        <CountryParticipationForm countries={$notParticipatingCountryStore} isParticipating={false}/>
                    </ul>
                </div>

                <div>
                    <ul>
                        <CountryParticipationForm countries={$participatingCountryStore} isParticipating={true}/>
                    </ul>
                </div>
            </div>
        </div>
    </div>

</div>



