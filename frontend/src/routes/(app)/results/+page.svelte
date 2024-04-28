<script lang="ts">
    import type {PageData, ActionData} from './$types';
    import {countryStore} from "$lib/stores/country.store";
    import type {ResultModel} from "$lib/models/classes/result.model";
    import Sort from "svelte-material-icons/Sort.svelte";
    import Filter from "svelte-material-icons/Filter.svelte";
    import {voteCats} from "$lib/models/enums/categories.enum";
    import Modal from "$lib/components/Modal.svelte";
    import {resultPageState} from "$lib/stores/resultPageState.store";
    import SelectUserResultForm from "$lib/components/forms/SelectUserResultForm.svelte";
    import ResetUserResultForm from "$lib/components/forms/ResetUserResultForm.svelte";
    import {userStore} from "$lib/stores/user.store";

    export let data:PageData
    export let form:ActionData
    let results:ResultModel[]
    let sortByDescending = true
    let openModal:VoidFunction
    let closeModal:VoidFunction

    results = data.results

    function noScores():boolean {
        const filtered = results.filter((res) => res.total > 0)
        return filtered.length === 0
    }

    function sort() {
        sortByDescending = !sortByDescending

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortByDescending) ? -1 : 1;

        results = results.sort((a, b) =>
            (getScore(a, $resultPageState.category) - getScore(b, $resultPageState.category)) * sortModifier
        );

    }

    function getScore(res: ResultModel, cat: string): number {
        switch (cat) {
            case voteCats.COSTUME:
                return res.costume
            case voteCats.SONG:
                return res.song
            case voteCats.PERFORMANCE:
                return res.performance
            case voteCats.PROPS:
                return res.props
            default:
                return res.total
        }
    }

    $: if (form) {
        results = form.results
        $resultPageState.userId = form.selection
    }

    $: if ($resultPageState.category) {
        sortByDescending = false
        sort()
    }
</script>

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <SelectUserResultForm showTitle={true}/>
</Modal>

<div class="h-full flex flex-col">

    <h1 class="text-center">Rankings</h1>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="py-3">

                {#if noScores() && $resultPageState.userId === "" && $resultPageState.category === voteCats.TOTAL}
                    <div class="text-center">
                        <h1 class="font-sans">🤷</h1>
                        <p>
                            Nothing to see yet. Come back when there are some votes!
                        </p>
                    </div>
                {:else if noScores() && $resultPageState.userId !== "" && $resultPageState.category }
                    <div class="text-center">
                        <h1 class="font-sans">🤷</h1>
                        <p>
                            Looks like {$userStore[$resultPageState.userId].name} hasn't voted yet.
                        </p>
                        <SelectUserResultForm />
                        <ResetUserResultForm />
                    </div>
                {:else}

                    <table class="w-full border-spacing-y-3 border-collapse">
                        <thead>
                        <tr>
                            <th>
                                <div class="flex align-end">
                                    <button class="cursor-pointer py-2 px-2 rounded" on:click={openModal}>
                                        <span class="flex"><Filter size="1.2em" class="pt-0.5 pl-0.5"/> Filter</span>
                                    </button>
                                    {#if $resultPageState.category !== voteCats.TOTAL || $resultPageState.userId !== ""}
                                        <div class="pl-3">
                                            <ResetUserResultForm />
                                        </div>
                                    {/if}
                                </div>
                            </th>
                            <th on:click={() => sort()} class="capitalize">
                                <div class="flex justify-center">
                                    {$resultPageState.category} <Sort size="1.4em" class="pl-1.5"/>
                                </div>

                            </th>
                        </tr>
                        </thead>
                        <tbody>
                        {#each results as result}
                            {#if result.total !== 0}
                                <tr>
                                    <td class="py-3 pl-3">
                                        {$countryStore.find(c => c.slug === result.countrySlug).flag}
                                        {$countryStore.find(c => c.slug === result.countrySlug).name}
                                    </td>
                                    <td class="text-center w-1/4">
                                        {getScore(result, $resultPageState.category)}
                                    </td>
                                </tr>
                            {/if}
                        {/each}
                        </tbody>
                    </table>

                {/if}
            </div>
        </div>
    </div>
</div>
