<script lang="ts">
    import type {PageData, ActionData} from './$types';
    import {countryStore} from "$lib/stores/country.store";
    import type {ResultModel} from "$lib/models/classes/result.model";
    import { enhance } from '$app/forms';
    import {authLvl} from "$lib/models/enums/authLvl.enum";

    export let data:PageData
    export let form:ActionData
    let results:ResultModel[]
    let sortBy = {col: "total", descending: true, icon: "fa-sort-down"}

    results = data.results
    let userArray = [...new Map(Object.entries(data.users))]

    function sort(column:string) {

        if (sortBy.col === column) {
            sortBy.descending = !sortBy.descending
            sortBy.descending ? sortBy.icon = "fa-sort-down" : sortBy.icon = "fa-sort-up"
        } else {
            sortBy.col = column
            sortBy.descending = true
            sortBy.icon = "fa-sort-down"
        }

        // Modifier to sorting function for ascending or descending
        let sortModifier = (sortBy.descending) ? -1 : 1;

        let sort = (a, b) =>
            (a[column] < b[column])
                ? -1 * sortModifier
                : (a[column] > b[column])
                    ? sortModifier
                    : 0;

        results = results.sort(sort);
    }

    $: if (form) {
        results = form.results
    }

</script>

<div class="h-full flex flex-col">

    <h1 class="text-center">Rankings</h1>

    <div class="py-3">
        <form method="POST" action="?/getUserResults" use:enhance>
            <select class="w-full text-center py-3" value={form?.selection ?? ""} name="id" on:change={(e) => {e.target.parentElement.requestSubmit()}}>
                <option value="">Main Results</option>
                {#each userArray as userInfo}
                    {#if userInfo[1].authLvl !== authLvl.BOT}
                        <option value={userInfo[0]}>{userInfo[1].name}</option>
                    {/if}
                {/each}
            </select>
        </form>
    </div>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="py-3">
                <table class="w-full border-spacing-y-3 border-collapse">
                    <thead>
                    <tr>
                        <th></th>
                        <th on:click={() => sort("total")}>Total ⬇</th>
                    </tr>
                    </thead>
                    <tbody>
                    {#each results as result}
                        {#if result.total !== 0}
                            <tr>
                                <td class="py-3">
                                    {$countryStore.find(c => c.slug === result.countrySlug).flag}
                                    {$countryStore.find(c => c.slug === result.countrySlug).name}
                                </td>
<!--                                <td>{result.song}</td>-->
<!--                                <td>{result.performance}</td>-->
<!--                                <td>{result.costume}</td>-->
<!--                                <td>{result.props}</td>-->
                                <td class="text-center">{result.total}</td>
                            </tr>
                        {/if}
                    {/each}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
