<script lang="ts">
    import { fade } from 'svelte/transition';
    import {errorStore} from "$lib/stores/error.store";
    import {onMount} from "svelte";

    let shouldDisplay = $state(false)

    onMount(() => {
        shouldDisplay = true
        setTimeout(closeToaster, 3000)
    })

    export const closeToaster = () => {
        shouldDisplay = false
        setTimeout(() => $errorStore = "", 400)
    }
</script>

{#if shouldDisplay}
<div transition:fade|global={{ delay: 250, duration: 300 }} class="fixed inset-0 overflow-y-auto h-full w-full z-50" onmouseup={closeToaster}>
    <div class="top-20">
        <div class="bg-warning rounded py-3 px-2 mt-2 shadow-lg shadow-gray-800 w-[75%] mx-auto">
            <p class="text-center">
                {$errorStore}
            </p>
        </div>
    </div>
</div>
{/if}