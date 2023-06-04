<script>
    import {currentUser} from "$lib/stores/user.store.ts";
    import {voteOptions} from "$lib/models/classes/voteOptions.model.ts";
    import { enhance } from '$app/forms';

    export let countrySlug = ""
    export let score = 0
    export let cat = ""

</script>

<form method="POST" action="?/vote" use:enhance>
    <input name="countrySlug" type="hidden" value={countrySlug}>
    <input name="cat" type="hidden" value={cat}>
    <input name="userId" type="hidden" value={$currentUser.id}>
    {#each voteOptions as { value, label }}
        <input type="radio" bind:group={score} value={value} name="score" on:click={(e) => {
            e.target.parentElement.requestSubmit()
        }}/>
        <label>{label}</label>
    {/each}
</form>