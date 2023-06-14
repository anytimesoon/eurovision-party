<script lang="ts">
    import type {PageData, ActionData} from './$types';
    import {countryStore} from "$lib/stores/country.store";
    import type {ResultModel} from "$lib/models/classes/result.model";
    import { enhance } from '$app/forms';
    import {userStore} from "$lib/stores/user.store";

    export let data:PageData
    export let form:ActionData
    let results:ResultModel[]
    let sortBy = {col: "total", descending: true}

    $userStore = data.users
    results = data.results
    let userArray = [...$userStore]

    $: sort = (column) => {

        if (sortBy.col == column) {
            sortBy.descending = !sortBy.descending
        } else {
            sortBy.col = column
            sortBy.descending = true
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

    const updateResults = (form) => {
        if(form !== null) {
            results = form.results
        }
    }

    $: updateResults(form)
</script>

<form method="POST" action="?/getUserResults" use:enhance>
    <select value={form?.selection ?? ""} name="id" on:change={(e) => {e.target.parentElement.requestSubmit()}}>
        <option value="">Main Results</option>
        {#each userArray as userInfo}
            <option value={userInfo[0]}>{userInfo[1].name}</option>
        {/each}
    </select>
</form>

<table>
    <thead>
    <tr>
        <th>Country</th>
        <th on:click={sort("song")}>Song</th>
        <th on:click={sort("performance")}>Performance</th>
        <th on:click={sort("costume")}>Costume</th>
        <th on:click={sort("props")}>Props</th>
        <th on:click={sort("total")}>Total</th>
    </tr>
    </thead>
    <tbody>
    {#each results as result}
        {#if result.total !== 0}
            <tr>
                <td>{$countryStore.find(c => c.slug === result.countrySlug).name}</td>
                <td>{result.song}</td>
                <td>{result.performance}</td>
                <td>{result.costume}</td>
                <td>{result.props}</td>
                <td>{result.total}</td>
            </tr>
        {/if}
    {/each}
    </tbody>
</table>