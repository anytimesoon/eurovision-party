<script lang="ts">
    import type {PageData, ActionData} from './$types';
    import {countryStore} from "$lib/stores/country.store";
    import type {ResultModel} from "$lib/models/classes/result.model";
    import { enhance } from '$app/forms';
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import Sort from "svelte-material-icons/Sort.svelte";
    import Filter from "svelte-material-icons/Filter.svelte";
    import {voteCats} from "$lib/models/enums/categories.enum";
    import Modal from "$lib/components/Modal.svelte";

    export let data:PageData
    export let form:ActionData
    let results:ResultModel[]
    let sortByDescending = true
    let currentCategory = voteCats.TOTAL
    let userSelectForm:HTMLFormElement
    let openModal:VoidFunction
    let closeModal:VoidFunction

    results = data.results
    let userArray = [...new Map(Object.entries(data.users))]

    function noScores():boolean {
        const filtered = results.filter((res) => res.total > 0)
        return filtered.length === 0
    }

    function sort() {
        sortByDescending = !sortByDescending

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortByDescending) ? -1 : 1;

        results = results.sort((a, b) =>
            (getScore(a, currentCategory) - getScore(b, currentCategory)) * sortModifier
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
    }

    $: if (currentCategory) {
        sortByDescending = false
        sort()
    }
</script>

<Modal bind:openModal={openModal} bind:closeModal={closeModal}>
    <h3 class="text-center">Filters</h3>

    <div class="pb-3">
        <form method="POST" action="?/getUserResults" use:enhance bind:this={userSelectForm}>
            <div class="flex">
                <div class="w-1/4 text-right pt-2 pr-3">
                    <label for="id">Person</label>
                </div>
                <div class="w-1/2">
                    <select class="w-full text-center py-2"
                            value={form?.selection ?? ""}
                            name="id"
                            on:change={() => userSelectForm.requestSubmit()}>
                        <option value="">Everyone</option>
                        {#each userArray as userInfo}
                            {#if userInfo[1].authLvl !== authLvl.BOT}
                                <option value={userInfo[0]}>{userInfo[1].name}</option>
                            {/if}
                        {/each}
                    </select>
                </div>
            </div>
        </form>
    </div>

    <div class="pb-3">
        <div class="flex">
            <div class="w-1/4 text-right pt-2 pr-3">
                <label for="id">Category</label>
            </div>
            <div class="w-1/2">
                <select class="w-full text-center py-2 capitalize" name="id">
                    <option value="total" on:click={(e) => currentCategory = e.target.value}>
                        Total
                    </option>

                    <option value="song" on:click={(e) => currentCategory = e.target.value}>
                        Song
                    </option>
                    <option value="performance" on:click={(e) => currentCategory = e.target.value}>
                        Performance
                    </option>
                    <option value="costume" on:click={(e) => currentCategory = e.target.value}>
                        Costume
                    </option>
                    <option value="props" on:click={(e) => currentCategory = e.target.value}>
                        Props
                    </option>
                </select>
            </div>
        </div>
    </div>

</Modal>

<div class="h-full flex flex-col">

    <h1 class="text-center">Rankings</h1>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="py-3">

                {#if noScores()}
                    <div class="text-center">
                        <h1 class="font-sans">🤷</h1>
                        <p>
                            Nothing to see yet. Come back when there are some votes!
                        </p>
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
                                </div>
                            </th>
                            <th on:click={() => sort()} class="capitalize">
                                <div class="flex justify-center">
                                    {currentCategory} <Sort size="1.4em" class="pl-1.5"/>
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
                                        {getScore(result, currentCategory)}
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
