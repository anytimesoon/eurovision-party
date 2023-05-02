<script lang="ts">
    import CountryVoteButtons from "$lib/components/CountryVoteButtons.svelte";

    /** @type {import('../../../index').PageData} */
    export let data;
</script>

{#key data.c}
    <h1>{data.c.flag} {data.c.name}</h1>

    <p>{data.c.songName} by {data.c.bandName}</p>

    <CountryVoteButtons country={data.c}/>
{/key}