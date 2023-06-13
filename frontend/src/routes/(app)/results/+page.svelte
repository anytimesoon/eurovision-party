<script lang="ts">
    import type {PageData} from './$types';
    import type {VoteModel} from "$lib/models/classes/vote.model";
    import {countryStore} from "$lib/stores/country.store";

    export let data:PageData
    let votes:VoteModel[]


    $: console.log(votes)
    $: votes = data.votes

    let sortBy = {col: "song", descending: true};

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
                    ? 1 * sortModifier
                    : 0;

        votes = votes.sort(sort);
    }
</script>

<table>
    <thead>
    <tr>
        <th>Country</th>
        <th on:click={sort("song")}>Song</th>
        <th on:click={sort("performance")}>Performance</th>
        <th on:click={sort("costume")}>Costume</th>
        <th on:click={sort("props")}>Props</th>
    </tr>
    </thead>
    <tbody>
    {#each votes as vote}
        <tr>
            <td>{$countryStore.find(c => c.slug === vote.countrySlug).name}</td>
            <td>{vote.song}</td>
            <td>{vote.performance}</td>
            <td>{vote.costume}</td>
            <td>{vote.props}</td>
        </tr>
    {/each}
    </tbody>
</table>