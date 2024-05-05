<script lang="ts">
    import { fade } from 'svelte/transition';
    import {errorStore} from "$lib/stores/error.store";
    import {onMount} from "svelte";

    let shouldDisplay = false

    onMount(() => {
        shouldDisplay = true
        setTimeout(closeToaster, 2000)
    })

    export const closeToaster = () => {
        shouldDisplay = false
        setTimeout(() => $errorStore = "", 400)
    }
</script>

{#if shouldDisplay}
<div transition:fade={{ delay: 250, duration: 300 }} class="fixed inset-0 z-30 overflow-y-auto h-full w-full z-40" on:mouseup={closeToaster}>
    <div class="top-20 max-w-screen-sm">
        <div class="bg-warning rounded py-3 px-1.5 mt-2 shadow-lg shadow-gray-800 w-[75%] mx-auto">
            <p class="text-center">
                {$errorStore}
            </p>
        </div>
    </div>
</div>
{/if}