<script lang="ts">
    import Radio from "$lib/components/RadioButton.svelte";
    import CountryVoteButtons from "$lib/components/CountryVoteButtons.svelte";

    /** @type {import('../../../index').PageData} */
    export let data;
</script>

<h1>{data.c.flag} {data.c.name}</h1>

<p>{data.c.songName} by {data.c.bandName}</p>

{#key data.c}
    <CountryVoteButtons country={data.c}/>
{/key}