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

    function sort() {
        sortByDescending = !sortByDescending

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortByDescending) ? -1 : 1;

        results = results.sort((a, b) =>
            (getScore(a, currentCategory) - getScore(b, currentCategory)) * sortModifier
        );
        console.log(results)
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

    <div class="flex">
        <div class="w-1/4 text-right pt-2 pr-3">
            <label for="id">Category</label>
        </div>
        <div class="w-1/2">
            <select class="w-full text-center py-2 capitalize"
                    name="id">
                {#each Object.values(voteCats) as category}
                    <option value={category} on:click={(e) => currentCategory = e.target.value}>
                        {category}
                    </option>
                {/each}
            </select>
        </div>
    </div>
</Modal>

<div class="h-full flex flex-col">

    <h1 class="text-center">Rankings</h1>

    <div class="py-3">
        <button class="absolute top-5 right-2 cursor-pointer py-2 px-2 rounded" on:click={openModal}>
            <span class="flex">Filter <Filter size="1.2em" class="pt-0.5 pl-0.5"/></span>
        </button>


    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="py-3">
                {#if results.length > 0}
                <table class="w-full border-spacing-y-3 border-collapse">
                    <thead>
                    <tr>
                        <th></th>
                        <th on:click={() => sort()} class="flex justify-center capitalize">
                            {currentCategory} <Sort size="1.4em" class="pl-1.5"/>
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
                {:else}
                    <div class="text-center">
                        <h1>🤷</h1>
                        <p>
                            Nothing to see yet. Come back when there are some votes!
                        </p>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>
