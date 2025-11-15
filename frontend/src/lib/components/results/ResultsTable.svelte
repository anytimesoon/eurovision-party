<script lang="ts">
    import {results} from "$lib/stores/results.store";
    import {resultPageState} from "$lib/stores/resultPageState.store"
    import {countryStore} from "$lib/stores/country.store";
    import Sort from "svelte-material-icons/Sort.svelte";
    import ResetUserResultForm from "$lib/components/forms/ResetUserResultForm.svelte";
    import Filter from "svelte-material-icons/Filter.svelte";
    import SelectUserResultForm from "$lib/components/forms/SelectUserResultForm.svelte";
    import Modal from "$lib/components/Modal.svelte";
    import EmptyTotal from "$lib/components/results/EmptyTotal.svelte";
    import EmptyUser from "$lib/components/results/EmptyUser.svelte";
    import {flip} from "svelte/animate";

    let openModal:VoidFunction = $state()
    let closeModal:VoidFunction = $state()
    let isDefaultState:boolean = $state(resultPageState.isDefault())
    let hasScores:boolean = $state(results.hasScores())
    let hasUserSelected:boolean = $state(resultPageState.hasUserSelected())
    let filteredResults = $derived($results.filter(result => (result !== null || result !== undefined) && result.total !== 0))

    $effect(() => {
        if ($results) {
            hasScores = results.hasScores()
        }
        if ($resultPageState) {
            isDefaultState = resultPageState.isDefault()
            hasUserSelected = resultPageState.hasUserSelected()
        }
        console.log(filteredResults)
    })

    function sort() {
        $resultPageState.sortByDescending = !$resultPageState.sortByDescending
        results.sortResults()
    }

</script>

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <SelectUserResultForm showTitle={true}/>
</Modal>


{#if !hasScores && isDefaultState}
    <EmptyTotal />
{:else if !hasScores && hasUserSelected}
    <EmptyUser />
{:else}

    <table class="w-full border-spacing-y-3 border-collapse">
        <thead>
            <tr>
                <th>
                    <div class="flex align-end">
                        <button class="cursor-pointer py-2 px-2 rounded" onclick={openModal}>
                            <span class="flex"><Filter size="1.2em" class="pt-0.5 pl-0.5"/> Filter</span>
                        </button>
                        {#if !isDefaultState}
                            <div class="pl-3">
                                <ResetUserResultForm />
                            </div>
                        {/if}
                    </div>
                </th>
                <th onclick={() => sort()} class="capitalize cursor-pointer">
                    <div class="flex justify-center">
                        {$resultPageState.category} <Sort size="1.4em" class="pl-1.5"/>
                    </div>

                </th>
            </tr>
        </thead>
        <tbody>
        {#each filteredResults as result (result.countrySlug)}
            <tr animate:flip={{duration: 500}}>
                <td class="py-3 pl-3">
                    {$countryStore.find(c => c.slug === result.countrySlug).flag}
                    {$countryStore.find(c => c.slug === result.countrySlug).name}
                </td>
                <td class="text-center w-1/4">
                    {result.getScore($resultPageState.category)}
                </td>
            </tr>
        {/each}
        </tbody>
    </table>

{/if}