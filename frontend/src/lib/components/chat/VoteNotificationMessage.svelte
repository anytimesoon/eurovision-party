<script lang="ts">
    import {currentlyVotingOn} from "$lib/stores/currentlyVotingOn.store";
</script>

<div class="rounded my-5 mx-2 p-3 text-center border border-secondary bg-canvas-primary">
    <p class="text-[3em]">{$currentlyVotingOn.flag} 🤖 {$currentlyVotingOn.flag}</p>

    <p class="mb-3">People are voting for {$currentlyVotingOn.name}!</p>

    <p><a href={`/country/${$currentlyVotingOn.slug}`} class="rounded bg-primary p-3 block w-40 mx-auto">🌟 Vote now! 🌟</a></p>
</div>